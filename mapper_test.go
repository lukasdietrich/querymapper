package querymapper

import (
	"net/url"
	"testing"
)

type Struct1 struct {
	KeyA string
	KeyB bool
	KeyC uint
	KeyD int
	KeyE float32
}

type Struct2 struct {
	KeyA string `query:"custom"`
}

type Struct3 struct {
	KeyA Struct1
}

func parseValues(u string) url.Values {
	a, _ := url.Parse(u)
	return a.Query()
}

func TestCorrectValue(t *testing.T) {
	tests := map[string]Struct1{
		"http://a.b?keya=test&keyb=true&keyc=13&keyd=-37&keye=0.1337": Struct1{
			"test", true, 13, -37, 0.1337,
		},
		"http://a.b?keya=example&keyb=false&keyc=0&keyd=234&keye=234.23": Struct1{
			"example", false, 0, 234, 234.23,
		},
	}

	for u, correct := range tests {
		var actual Struct1
		err := MapQuery(parseValues(u), &actual)

		if err != nil {
			t.Error(err)
		}

		if actual != correct {
			t.Error("structs did not match")
		}
	}
}

func TestMissingValue(t *testing.T) {
	tests := []string{
		"http://a.b?keya=test&keyc=13&keyd=-37&keye=0.1337",
		"http://a.b?keya=example&keyc=false&keyc=0&keyd=234&keye=234.23",
	}

	for _, u := range tests {
		err := MapQuery(parseValues(u), &Struct1{})

		if err == nil {
			t.Error("did not error on missing value")
		}
	}
}

func TestCustomKey(t *testing.T) {
	tests := map[string]Struct2{
		"http://a.b?custom=test": Struct2{"test"},
		"http://a.b?custom=":     Struct2{},
	}

	for u, correct := range tests {
		var actual Struct2
		err := MapQuery(parseValues(u), &actual)

		if err != nil {
			t.Error(err)
		}

		if actual != correct {
			t.Error("structs did not match")
		}
	}
}

func TestUnsupportedType(t *testing.T) {
	err := MapQuery(parseValues("http://a.b?keya="), &Struct3{})
	if err == nil {
		t.Error("did not error on unsupported type")
	}
}
