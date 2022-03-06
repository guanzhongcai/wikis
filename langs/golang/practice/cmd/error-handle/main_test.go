package main

import (
	"io/ioutil"
	"testing"
)

func TestNotFoundError_Error(t *testing.T) {
	e := NotFoundError{File: "hello.txt"}
	t.Log(e.Error())

	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString("hello"); err != nil {
		t.Error(err)
	}
}
