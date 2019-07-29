package graphql

import (
	"math/big"
	"testing"
)

func TestMarshaller(t *testing.T) {
	i := new(big.Float).SetFloat64(1)

	v, err := UnmarshalFloat("1")
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("string unmatch")
	}

	v, err = UnmarshalFloat(1)
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("int unmatch")
	}

	v, err = UnmarshalFloat(int64(1))
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("int64 unmatch")
	}

	v, err = UnmarshalFloat(float64(1.0))
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("float64 unmatch")
	}

	_, err = UnmarshalFloat("test")
	if err == nil || err.Error() != "syntax error scanning number" {
		t.Error("invalid float error fails")
	}
}
