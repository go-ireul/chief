package main

import (
	"crypto/rand"

	"ireul.com/bolt"
	"ireul.com/mshuf"
)

func newID(db *bolt.DB, name string) (ret uint64, err error) {
	err = db.Update(func(tx *bolt.Tx) (err error) {
		// bucket
		var bkt *bolt.Bucket
		if bkt, err = tx.CreateBucketIfNotExists([]byte(name)); err != nil {
			return
		}
		// matrix
		var m mshuf.Matrix
		b := bkt.Get([]byte("matrix"))
		if len(b) == mshuf.MatrixLength {
			m = mshuf.Matrix(b)
		} else {
			m = mshuf.NewMatrix()
			for i := 0; i < mshuf.MatrixSize; i++ {
				m.RandomRowAt(rand.Reader, i)
			}
			bkt.Put([]byte("matrix"), m)
		}
		// shuffle new sequence
		var seq uint64
		if seq, err = bkt.NextSequence(); err != nil {
			return
		}
		ret = m.Shuffle(seq)
		return
	})
	return
}
