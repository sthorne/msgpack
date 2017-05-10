package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"reflect"
	"unicode"

	"github.com/pkg/errors"
)

// This utility generates the various encodeMapXXX methods, which
// uses raw map[string]XXX types instead of reflect for maximum
// efficiency
func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func ucfirst(s string) string {
	res := make([]rune, len(s))
	for i, r := range s {
		if i == 0 {
			res[i] = unicode.ToUpper(r)
		} else {
			res[i] = r
		}
	}
	return string(res)
}

func _main() error {
	types := []reflect.Kind{
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String,
	}

	var buf bytes.Buffer
	buf.WriteString("package msgpack")
	buf.WriteString("\n\n// Auto-generated by internal/cmd/genmap/genmap.go. DO NOT EDIT!")
	buf.WriteString("\n\nimport \"github.com/pkg/errors\"")
	for _, typ := range types {
		fmt.Fprintf(&buf, "\n\nfunc (e *Encoder) encodeMap%s(in interface{}) error {", ucfirst(typ.String()))
		fmt.Fprintf(&buf, "\nfor k, v := range in.(map[string]%s) {", typ)
		buf.WriteString("\nif err := e.EncodeString(k); err != nil {")
		buf.WriteString("\nreturn errors.Wrap(err, `failed to encode key`)")
		buf.WriteString("\n}")
		buf.WriteString("\nif err := e.Encode(v); err != nil {")
		buf.WriteString("\nreturn errors.Wrapf(err, `failed to encode value for %s`, k)")
		buf.WriteString("\n}")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return errors.Wrap(err, `failed to format`)
	}

	var fn string
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] != "-" {
			fn = os.Args[i]
			break
		}
	}

	dst, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, `failed to open file`)
	}
	defer dst.Close()

	dst.Write(formatted)
	return nil
}
