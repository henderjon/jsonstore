/*
Package json-store reads/writes data to json flat files

Given a top level directory, `Get`, `Put`, and `Del` json encoded flat files of
a given name.
*/
package jsonstore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	// "errors"
)

// default dir/file permission
const PERM = 0755

// var MISS = errors.New("File Not Found")

// Bucket holds the name of the directory to which the files are written and
// has all the attached methods
type Bucket struct {
	prefix string
}

// NewConnection creates our root data dir
func Open(dir string) (*Bucket, error) {

	if len(dir) == 0 {
		return &Bucket{""}, nil
	}

	if dir[:len(dir)-1] == "/" {
		dir = dir[:len(dir)-2]
	}

	if err := os.MkdirAll(dir, PERM); err != nil {
		return nil, err
	}

	return &Bucket{dir}, nil
}

// Get retrieves the contents of the given file and unmarshals it to the given interface.
// `os.IsNotExist` ought to be used to discover if a `Get` couldn't find a file
func (b *Bucket) Get(key string, v interface{}) error {
	contents, err := ioutil.ReadFile(b.mkkey(key))
	// if os.IsNotExist(err) {
	// 	return MISS
	// }

	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, v)
	if err != nil {
		return err
	}
	return nil
}

// Put marshals and writes the contents of the given interface to the given file
func (b *Bucket) Put(key string, v interface{}) error {
	contents, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.mkkey(key), contents, PERM)
	if err != nil {
		return err
	}
	return nil
}

// Del deletes the given file
func (b *Bucket) Del(key string) error {
	err := os.Remove(b.mkkey(key))
	if err != nil {
		return err
	}
	return nil
}

// DellAll deletes the top level data dir
func (b *Bucket) DelAll() error {
	err := os.RemoveAll(b.prefix)
	if err != nil {
		return err
	}
	return nil
}

// mkkey creates the full path to the file given the prefix and the key
func (b *Bucket) mkkey(key string) string {
	if len(b.prefix) > 0 {
		b.prefix += "/"
	}
	return b.prefix + key
}
