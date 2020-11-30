package graphql

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalFloat for graphql Float scalars to be compatible with big.Float
func MarshalFloat(f big.Float) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, f.String())
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

// MarshalDate marshals time.Time inf form of YYYY-MM-DD
func MarshalDate(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(t.Format("2006-01-02")))
	})
}

// UnmarshalDate unmarshals YYYY-MM-DD to time.Time
func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse("2006-01-02", tmpStr)
	}
	return time.Time{}, errors.New("date should be a string")
}
