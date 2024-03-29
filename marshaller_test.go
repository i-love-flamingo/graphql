package graphql_test

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"
	"time"

	"flamingo.me/graphql"
)

func TestMarshalFloats(t *testing.T) {
	t.Parallel()

	i := new(big.Float).SetFloat64(1)

	v, err := graphql.UnmarshalFloat("1")
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("string unmatch")
	}

	v, err = graphql.UnmarshalFloat(1)
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("int unmatch")
	}

	v, err = graphql.UnmarshalFloat(int64(1))
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("int64 unmatch")
	}

	v, err = graphql.UnmarshalFloat(float64(1.0))
	if err != nil {
		t.Error(err)
	}

	if v.Cmp(i) != 0 {
		t.Error("float64 unmatch")
	}

	_, err = graphql.UnmarshalFloat("test")
	if err == nil {
		t.Error("invalid float error fails")
	}
}

func TestMarshalDates(t *testing.T) {
	t.Parallel()

	writer := graphql.MarshalDate(time.Time{})
	b := bytes.NewBufferString("")

	writer.MarshalGQL(b)

	if b.String() != "null" {
		t.Error("zero date should be null")
	}

	now := time.Now()
	writer = graphql.MarshalDate(now)
	b = bytes.NewBufferString("")

	writer.MarshalGQL(b)

	if b.String() != fmt.Sprintf("%q", now.Format("2006-01-02")) {
		t.Error("date should be marshalled to format YYYY-MM-DD")
	}

	date, err := graphql.UnmarshalDate("2011-08-12")
	if err != nil {
		t.Fatal("Unmarshal of correct date string should work")
	}

	expectedDate, _ := time.Parse("2006-01-02", "2011-08-12")
	if !date.Equal(expectedDate) {
		t.Error("Umarshalled date is wrong")
	}

	_, err = graphql.UnmarshalDate("foobar")
	if err == nil {
		t.Error("Unmarshal of invalid date string should lead to an error")
	}

	_, err = graphql.UnmarshalDate(42)
	if err == nil {
		t.Error("Unmarshal of invalid date type should lead to an error")
	}

	date, err = graphql.UnmarshalDate("")
	if err != nil {
		t.Fatal("Unmarshal of empty string should work")
	}

	if !date.Equal(time.Time{}) {
		t.Fatal("Date should be empty time")
	}
}
