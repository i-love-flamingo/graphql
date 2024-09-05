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

var ErrUnexpectedValue = errors.New("unexpected value")

// MarshalFloat for graphql Float scalars to be compatible with big.Float
func MarshalFloat(f big.Float) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, f.String())
	})
}

// UnmarshalFloat for graphql Float scalars to be compatible with big.Float
func UnmarshalFloat(v interface{}) (big.Float, error) {
	switch floatValue := v.(type) {
	case string:
		return parseString(floatValue)
	case int:
		return *big.NewFloat(float64(floatValue)), nil
	case int64:
		return *big.NewFloat(float64(floatValue)), nil
	case float64:
		return *big.NewFloat(floatValue), nil
	case json.Number:
		return parseString(string(floatValue))
	default:
		return big.Float{}, fmt.Errorf("%w: %T is not a float", ErrUnexpectedValue, floatValue)
	}
}

//nolint:mnd // options for float parsing are clear
func parseString(floatValue string) (big.Float, error) {
	f, _, err := big.ParseFloat(floatValue, 10, 64, big.ToNearestEven)
	if f == nil {
		f = &big.Float{}
	}

	if err != nil {
		return *f, fmt.Errorf("can not parse string %q to float: %w", floatValue, err)
	}

	return *f, nil
}

// MarshalDate marshals time.Time in form of YYYY-MM-DD
func MarshalDate(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(t.Format("2006-01-02")))
	})
}

// UnmarshalDate unmarshalls YYYY-MM-DD to time.Time
func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		if len(tmpStr) == 0 {
			return time.Time{}, nil
		}

		parse, err := time.Parse(time.DateOnly, tmpStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("date must be in format %q: %w", time.DateOnly, err)
		}

		return parse, nil
	}

	return time.Time{}, fmt.Errorf("%w: date should be a string", ErrUnexpectedValue)
}
