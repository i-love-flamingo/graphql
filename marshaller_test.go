package graphql

import (
	"math/big"
	"testing"
)

func TestMarshaller(t *testing.T) {
	i := new(big.Float).SetFloat64(1)

	v, e := UnmarshalFloat("1")
	if e != nil {
		t.Error(e)
	}

	if v.Cmp(i) != 0 {
		t.Error("string unmatch")
	}

	v, e = UnmarshalFloat(1)
	if e != nil {
		t.Error(e)
	}

	if v.Cmp(i) != 0 {
		t.Error("int unmatch")
	}

	v, e = UnmarshalFloat(int64(1))
	if e != nil {
		t.Error(e)
	}

	if v.Cmp(i) != 0 {
		t.Error("int64 unmatch")
	}

	v, e = UnmarshalFloat(float64(1.0))
	if e != nil {
		t.Error(e)
	}

	if v.Cmp(i) != 0 {
		t.Error("float64 unmatch")
	}

	_, e = UnmarshalFloat("test")
	if e == nil || e.Error() != "syntax error scanning number" {
		t.Error("invalid float error fails")
	}
}
