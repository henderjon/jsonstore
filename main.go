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
)

// default dir/file permission
const PERM = 0755

// holds the name of the directory to which the files are written
type connection struct {
	dir string
}

// NewConnection creates our root data dir
func NewConnection(dir string) (*connection, error) {
	if dir[:len(dir)-1] == "/" {
		dir = dir[:len(dir)-2]
	}

	if err := os.MkdirAll(dir, PERM); err != nil {
		return nil, err
	}

	return &connection{dir}, nil
}

// Get retrieves the contents of the given file and unmarshals it to the given interface
func (c *connection) Get(key string, v interface{}) error {
	contents, err := ioutil.ReadFile(mkkey(c.dir, key))
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
func (c *connection) Put(key string, v interface{}) error {
	contents, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(mkkey(c.dir, key), contents, PERM)
	if err != nil {
		return err
	}
	return nil
}

// Del deletes the given file
func (c *connection) Del(key string) error {
	err := os.Remove(mkkey(c.dir, key))
	if err != nil {
		return err
	}
	return nil
}

// DellAll deletes the top level data dir
func (c *connection) DelAll() error {
	err := os.RemoveAll(c.dir)
	if err != nil {
		return err
	}
	return nil
}

// mkkey creates the full path to the file given the prefix and the key
func mkkey(prefix, key string) string {
	return prefix + "/" + key
}
