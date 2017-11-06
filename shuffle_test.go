package main

import (
	"testing"

	"ireul.com/bolt"
)

func Test_newID(t *testing.T) {
	db, err := bolt.Open("test.db", 0660, nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		_, err := newID(db, "test")
		if err != nil {
			panic(err)
		}
	}
}
