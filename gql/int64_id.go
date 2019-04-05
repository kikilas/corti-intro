package gql

//Int64ID resolver
import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

const encodeHex = "0123456789ABCDEF"
const int64Base = 10

func MarshalInt64ID(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		writeQuotedString(w, strconv.FormatInt(i, int64Base))
	})
}

func UnmarshalInt64ID(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, int64Base, 64)
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case json.Number:
		return strconv.ParseInt(string(v), int64Base, 64)
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

func writeQuotedString(w io.Writer, s string) {
	start := 0
	io.WriteString(w, `"`)

	for i, c := range s {
		if c < 0x20 || c == '\\' || c == '"' {
			io.WriteString(w, s[start:i])

			switch c {
			case '\t':
				io.WriteString(w, `\t`)
			case '\r':
				io.WriteString(w, `\r`)
			case '\n':
				io.WriteString(w, `\n`)
			case '\\':
				io.WriteString(w, `\\`)
			case '"':
				io.WriteString(w, `\"`)
			default:
				io.WriteString(w, `\u00`)
				w.Write([]byte{encodeHex[c>>4], encodeHex[c&0xf]})
			}

			start = i + 1
		}
	}

	io.WriteString(w, s[start:])
	io.WriteString(w, `"`)
}
