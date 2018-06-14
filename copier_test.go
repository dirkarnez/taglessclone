package taglessclone

import (
"testing"
)

type TopSecret struct {
	value int
}

type A struct {
	ASecret string
	ATopSecret TopSecret
}

type B struct {
	BSecret string
	BTopSecret TopSecret
}

func TestSetField_1(t *testing.T) {
	a := A{ASecret: "Hello", ATopSecret: TopSecret{ value: 123}}
	b := B{}

	err := SetField(&b.BSecret, a.ASecret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if b.BSecret != a.ASecret && b.BSecret != "Hello" {
		t.Errorf("copy fail")
	}
}

func TestSetField_2(t *testing.T) {
	a := A{ASecret: "Hello Again", ATopSecret: TopSecret{ value: 456}}
	b := B{}

	err := SetField(&b.BTopSecret, a.ATopSecret)
	if err != nil {
		t.Errorf(err.Error())
	}

	if b.BTopSecret.value != a.ATopSecret.value && b.BTopSecret.value != 456 {
		t.Errorf("copy fail")
	}
}



