package graphql

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"
	"time"
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
	if err == nil {
		t.Error("invalid float error fails")
	}

	writer := MarshalDate(time.Time{})
	b := bytes.NewBufferString("")

	writer.MarshalGQL(b)
	if b.String() != "null" {
		t.Error("zero date should be null")
	}

	now := time.Now()
	writer = MarshalDate(now)
	b = bytes.NewBufferString("")

	writer.MarshalGQL(b)
	if b.String() != fmt.Sprintf("%q", now.Format("2006-01-02")) {
		t.Error("date should be marshalled to format YYYY-MM-DD")
	}

	date, err := UnmarshalDate("2011-08-12")
	if err != nil {
		t.Fatal("Unmarshal of correct date string should work")
	}

	expectedDate, _ := time.Parse("2006-01-02", "2011-08-12")
	if !date.Equal(expectedDate) {
		t.Error("Umarshalled date is wrong")
	}

	_, err = UnmarshalDate("foobar")
	if err == nil {
		t.Error("Unmarshal of invalid date string should lead to an error")
	}

	_, err = UnmarshalDate(42)
	if err == nil {
		t.Error("Unmarshal of invalid date type should lead to an error")
	}
}
