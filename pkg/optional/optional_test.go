package optional

import (
	"errors"
	"fmt"
	"testing"
)

type Customer struct {
	Id   int64
	Name *NamePair
	Age  int
}

type NamePair struct {
	FirstName, LastName string
}

func TestOptional(t *testing.T) {
	c := &Customer{
		Id:   1,
		Name: nil,
		Age:  1,
	}
	// --------- not nil ----------------
	o := OfNullable(c)
	// IfPresent
	activeIfPresent := 0
	o.IfPresent(func(v *Customer) {
		activeIfPresent = 1
		if v == nil {
			t.Errorf("v is nil")
		}
		if v.Id != 1 {
			t.Errorf("expect v.Id = %v, but got %v", 1, v.Id)
		}
	})
	if activeIfPresent == 0 {
		t.Errorf("not active IfPresent")
	}

	// ======== nil =================
	c = nil
	o = OfNullable(c)

	// OrElse
	if o.OrElse(&Customer{Age: 100}).Age != 100 {
		t.Errorf("OrElse asset failed")
	}

	// OrElseGet
	if o.OrElseGet(func() *Customer {
		return &Customer{Age: 100}
	}).Age != 100 {
		t.Errorf("OrElseGet asset failed")
	}

	// OrErr
	e := fmt.Errorf("test")
	if !errors.Is(e, o.OrErr(e)) {
		t.Errorf("OrErr asset failed")
	}

	// IfNone
	activeIfNone := 0
	o.IfNone(func() {
		activeIfNone = 1
	})
	if activeIfNone == 0 {
		t.Errorf("IfNone asset failed")
	}

}
