package main

import (
	"fmt"
	"testing"

	"ireul.com/bolt"
)

func Test_newID(t *testing.T) {
	db, err := bolt.Open("test.db", 0660, nil)
	if err != nil {
		panic(err)
	}
	id, err := newID(db, "test")
	if err != nil {
		panic(err)
	}
	fmt.Println("NewID", id)
}
