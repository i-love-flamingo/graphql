package graphql

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalFloat for graphql Float scalars to be compatible with big.Float
func MarshalFloat(f big.Float) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, f.String())
	})
}

// UnmarshalFloat for graphql Float scalars to be compatible with big.Float
func UnmarshalFloat(v interface{}) (big.Float, error) {
	switch v := v.(type) {
	case string:
		f, _, err := big.ParseFloat(v, 10, 64, big.ToNearestEven)
		if f == nil {
			return big.Float{}, err
		}
		return *f, err
	case int:
		return *big.NewFloat(float64(v)), nil
	case int64:
		return *big.NewFloat(float64(v)), nil
	case float64:
		return *big.NewFloat(v), nil
	case json.Number:
		f, _, err := big.ParseFloat(string(v), 10, 64, big.ToNearestEven)
		if f == nil {
			return big.Float{}, err
		}
		return *f, err
	default:
		return big.Float{}, fmt.Errorf("%T is not an float", v)
	}
}
