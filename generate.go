package protobuf

import (
	"io"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

const protoTemplate = `[[range .Types]]
message [[.Name]] {[[range .|Fields]]
  [[.|TypeName]] [[.|FieldName]] = [[.ID]][[.|Options]];[[end]]
}
[[end]]
`

var fixName = regexp.MustCompile(`((?:ID)|(?:[A-Z][a-z_0-9]+)|([\w\d]+))`)

func fieldName(f ProtoField) string {
	if f.Name != "" {
		return f.Name
	}
	parts := fixName.FindAllString(f.Field.Name, -1)
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
	}
	return strings.Join(parts, "_")
}

func typeIndirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func typeName(f ProtoField) string {
	t := f.Field.Type
	if t.Kind() == reflect.Slice {
		if t.Elem().Kind() == reflect.Uint8 {
			return fieldPrefix(f, TagNone) + "bytes"
		}
		return "repeated " + innerTypeName(typeIndirect(t.Elem()))
	}
	if t.Kind() == reflect.Ptr {
		return fieldPrefix(f, TagOptional) + innerTypeName(t.Elem())
	}
	return fieldPrefix(f, TagNone) + innerTypeName(t)
}

func fieldPrefix(f ProtoField, def TagPrefix) string {
	opt := def
	if def == TagNone {
		opt = f.Prefix
	}
	switch opt {
	case TagOptional:
		return "optional "
	case TagRequired:
		return "required "
	default:
		if f.Field.Type.Kind() == reflect.Ptr {
			return "optional "
		}
		return "required "
	}
}

func innerTypeName(t reflect.Type) string {
	if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 {
		return "bytes"
	}
	if t.PkgPath() == "time" {
		if t.Name() == "Time" {
			return "sfixed64"
		}
		if t.Name() == "Duration" {
			return "sint64"
		}
	}
	switch t.Name() {
	case "Ufixed32":
		return "fixed32"
	case "Ufixed64":
		return "ufixed64"
	case "Sfixed32":
		return "sfixed32"
	case "Sfixed64":
		return "sfixed64"
	}

	switch t.Kind() {
	case reflect.Float64:
		return "double"
	case reflect.Float32:
		return "float"
	case reflect.Int32:
		return "sint32"
	case reflect.Int, reflect.Int64:
		return "sint64"
	case reflect.Bool:
		return "bool"
	case reflect.Uint32:
		return "uint32"
	case reflect.Uint, reflect.Uint64:
		return "uint64"
	case reflect.String:
		return "string"

	case reflect.Struct:
		return t.Name()
	default:
		panic("unsupported type " + t.Name())
	}
}

func options(f ProtoField) string {
	if f.Field.Type.Kind() == reflect.Slice {
		switch f.Field.Type.Elem().Kind() {
		case reflect.Bool,
			reflect.Int32, reflect.Int64,
			reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return " [packed=true]"
		}
	}
	return ""
}

type FieldNamer func(ProtoField) string

type reflectedTypes []reflect.Type

func (r reflectedTypes) Len() int           { return len(r) }
func (r reflectedTypes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r reflectedTypes) Less(i, j int) bool { return r[i].Name() < r[j].Name() }

// GenerateProtobufDefinition generates a .proto file from a list of structs via reflection.
// fieldNamer is a function that maps ProtoField types to generated protobuf field names.
func GenerateProtobufDefinition(w io.Writer, types []interface{}, fieldNamer FieldNamer) error {
	rt := reflectedTypes{}
	for _, t := range types {
		rt = append(rt, reflect.Indirect(reflect.ValueOf(t)).Type())
	}
	sort.Sort(rt)
	if fieldNamer == nil {
		fieldNamer = fieldName
	}
	t := template.Must(template.New("protobuf").Funcs(template.FuncMap{
		"Fields":    ProtoFields,
		"FieldName": fieldNamer,
		"TypeName":  typeName,
		"Options":   options,
	}).Delims("[[", "]]").Parse(protoTemplate))
	return t.Execute(w, map[string]interface{}{
		"Types": rt,
		"Ptr":   reflect.Ptr,
		"Slice": reflect.Slice,
	})
}
