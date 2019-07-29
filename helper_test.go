package graphql

import (
	"testing"
)

type TestModel struct{}

func (*TestModel) TestModelMethod() {}

func TestHelperModels(t *testing.T) {
	res := ModelMap{
		"TestModel": ModelMapEntry{
			Type: new(TestModel),
			Fields: map[string]string{
				"testMethod": "TestModelMethod",
			},
		},
	}.Models()

	if _, ok := res["TestModel"]; !ok {
		t.Error("Test Model is not implemented")
	}

	TestModelValue := res["TestModel"]

	if TestModelValue.Model.Has("flamingo.me/graphql.TestModel") == false {
		t.Error("Test Model not mapped correctly")
	}

	if _, ok := TestModelValue.Fields["testMethod"]; !ok {
		t.Error("Test Model method does not exists")
	}

	TestModelFieldsTestMethod := TestModelValue.Fields["testMethod"]
	if TestModelFieldsTestMethod.FieldName != "TestModelMethod" {
		t.Error("Test Model method field name is not matching")
	}
}

func TestHelperModelsFlat(t *testing.T) {
	res := ModelMap{
		"TestModel": TestModel{},
	}.Models()

	if _, ok := res["TestModel"]; !ok {
		t.Error("Test Model is not implemented")
	}

	TestModelValue := res["TestModel"]

	if TestModelValue.Model.Has("flamingo.me/graphql.TestModel") == false {
		t.Error("Test Model not mapped correctly")
	}

	if TestModelValue.Fields != nil {
		t.Error("Test Model Fields are not empty")
	}
}
