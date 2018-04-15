package main

import (
	"crypto/rand"
	"errors"

	"github.com/boltdb/bolt"
	"magi.systems/math/mshuf"
)

var errSequenceDrain = errors.New("sequence drain")

var keyMatrix = []byte("matrix")

func newID(db *bolt.DB, name string) (ret uint64, err error) {
	err = db.Update(func(tx *bolt.Tx) (err error) {
		// bucket
		var bkt *bolt.Bucket
		if bkt, err = tx.CreateBucketIfNotExists([]byte(name)); err != nil {
			return
		}
		// matrix
		var m mshuf.Matrix
		b := bkt.Get(keyMatrix)
		if len(b) == mshuf.MatrixLength {
			m = mshuf.Matrix(b)
		} else {
			m = mshuf.NewMatrix()
			for i := 0; i < mshuf.MatrixSize; i++ {
				m.RandomRowAt(rand.Reader, i)
			}
			bkt.Put(keyMatrix, m)
		}
		// new sequence
		var seq uint64
		if seq, err = bkt.NextSequence(); err != nil {
			return
		}
		// check sequence exceeded
		if seq >= shardSize {
			err = errSequenceDrain
			return
		}
		// shuffle
		ret = m.Shuffle(seq)
		// mask to shard
		ret = ret&shardMask + shardPrefix
		return
	})
	return
}
