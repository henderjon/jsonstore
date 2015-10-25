package jsonstore

import "testing"
// import "fmt"

func TestConnect(t *testing.T){
	_, err := NewConnection("./test/store/depth/")
	if err != nil {
		t.Error(err.Error())
	}
}

type TestPayload struct {
	Dig string
}

func TestPutGet(t *testing.T){
	c, err := NewConnection("./test/store/depth/")
	if err != nil {
		t.Error(err.Error())
	}

	data := TestPayload{"this"}

	err = c.Put("test.json", data)
	if err != nil {
		t.Error(err.Error())
	}

	r := &TestPayload{}
	err = c.Get("test.json", r)
	if err != nil {
		t.Error(err.Error())
	}

	if r.Dig != "this" {
		t.Error("failed to unmarshal")
	}

	err = c.Del("test.json")
	if err != nil {
		t.Error(err.Error())
	}

	err = c.DelAll()
	if err != nil {
		t.Error(err.Error())
	}

}
