package jsonstore

import "testing"
import "os"

// import "fmt"
// import "log"

func TestConnect_trailing(t *testing.T) {
	_, err := Open("./test/")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConnect_no_trailing(t *testing.T) {
	_, err := Open("")
	if err != nil {
		t.Error(err.Error())
	}
}

type TestPayload struct {
	Dig string
}

func TestFlow(t *testing.T) {
	c, err := Open("./test")
	if err != nil {
		t.Error(err.Error())
	}

	data := TestPayload{"this"}

	// test Put
	err = c.Put("test.json", data)
	if err != nil {
		t.Error(err.Error())
	}

	r := &TestPayload{}

	// test Get
	err = c.Get("test.json", r)
	if err != nil {
		t.Error(err.Error())
	}

	if r.Dig != "this" {
		t.Error("failed to unmarshal")
	}

	// test Del
	err = c.Del("test.json")
	if err != nil {
		t.Error(err.Error())
	}

	r = &TestPayload{}

	// test Get w/ a miss
	err = c.Get("test.json", r)
	if os.IsExist(err) {
		t.Error(err.Error())
	}

	err = c.DelAll()
	if err != nil {
		t.Error(err.Error())
	}

}
