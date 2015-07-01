package validator

import (
	"testing"
	"fmt"
)

type Human struct {
	Name		string		`validator:"regex(value:name, code:1001, message:name只能为英文字母)"`
	Age			int
}

type Student struct {
	Human
}

func Test_Human(t *testing.T) {
	AddRegex("name", "^[a-zA-Z]+$")

	var v = NewValidator()

	var h1 = Student{}
	h1.Name = "name1"
	h1.Age = 17

	v.AddValidator("Age", "gte", "18", "1002", "age必须大于等于18")

	fmt.Println(v.Validate(h1), "  ", v.Error())

	fmt.Println(v.Errors())

}