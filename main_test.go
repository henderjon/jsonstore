package jsonstore

import "testing"
import "os"

// import "fmt"
// import "log"

func TestConnect(t *testing.T) {
	_, err := Open("./test/store/depth/")
	if err != nil {
		t.Error(err.Error())
	}
}

type TestPayload struct {
	Dig string
}

func TestPutGet(t *testing.T) {
	c, err := Open("./test")
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

func TestPutGet2(t *testing.T) {
	c, err := Open("")
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

func TestGetMiss(t *testing.T) {
	c, err := Open("")
	if err != nil {
		t.Error(err.Error())
	}

	r := &TestPayload{}
	err = c.Get("fnf.json", r)
	// this should miss
	if os.IsExist(err) {
		t.Error(err.Error())
	}

	err = c.DelAll()
	if err != nil {
		t.Error(err.Error())
	}

}
