package store_bbolt

import (
	"bytes"
	"errors"
	"go.etcd.io/bbolt"
	"time"
)

var ERR_BUCKET_NOT_EXIST = errors.New("bucket not exist")
var ERR_KEY_NOT_EXIST = errors.New("key not exist")
var ERR_KEY_ALREADY_EXISTS = errors.New("key already exists")
var ERR_CANNOT_GET_BUCKET = errors.New("cannot get bucket")

func NewBoltDB(filename string) (*boltStore, error) {
	b, err := bbolt.Open(filename, 0666, &bbolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}
	s := &boltStore{db: b}

	return s, nil
}

// Default database for storer will be badger
type boltStore struct {
	db *bbolt.DB
}

func (s *boltStore) Lock() {
	return
}

func (s *boltStore) Unlock() {
	return
}

func (s *boltStore) Flush() error {
	return nil
}

func (s *boltStore) NewBucket(bucket []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
			return err
		}
		return nil
	})
}

func (s *boltStore) DelBucket(bucket []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket(bucket)
	})
}

func (s *boltStore) Get(bucket, key []byte) ([]byte, error) {
	var out []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		if b := tx.Bucket(bucket); b != nil {
			out = b.Get(key)
			if out != nil {
				return nil
			}
			return ERR_KEY_NOT_EXIST
		}
		return ERR_BUCKET_NOT_EXIST
	})
	return out, err
}

func (s *boltStore) Put(bucket, key, data []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		// If bucket exists, use it to put.
		b := tx.Bucket(bucket)
		var err error
		if b == nil {
			b, err = tx.CreateBucketIfNotExists(bucket)
			if err != nil {
				return err
			}
		}
		return b.Put(key, data)
	})
}

func (s *boltStore) Del(bucket, key []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if b := tx.Bucket(bucket); b != nil {
			return b.Delete(key)
		}
		return ERR_BUCKET_NOT_EXIST
	})
}

func (s *boltStore) Iter(bucket, keyPrefix []byte, f func(key []byte, val []byte) error) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		if b := tx.Bucket(bucket); b != nil {
			c := b.Cursor()
			for k, v := c.Seek(keyPrefix); k != nil && bytes.HasPrefix(k, keyPrefix); k, v = c.Next() {
				if err := f(k, v); err != nil {
					return err
				}
			}
			return nil
		}
		return ERR_BUCKET_NOT_EXIST
	})
}

func (s *boltStore) Do(bucket, key []byte, f func(val []byte) ([]byte, error)) ([]byte, error) {
	var newVal []byte
	var err error
	err = s.db.Update(func(tx *bbolt.Tx) error {
		if b := tx.Bucket(bucket); b != nil {
			val := b.Get(key)
			newVal, err = f(val)
			if err != nil {
				return err
			}
			// delete if newVal is nil
			if val != nil && newVal == nil {
				return b.Delete(key)
			}
			return b.Put(key, newVal)
		}
		return ERR_BUCKET_NOT_EXIST
	})
	return newVal, err
}

func (s *boltStore) DoIter(bucket, keyPrefix []byte, f func(key, val []byte) ([]byte, error)) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if b := tx.Bucket(bucket); b != nil {
			c := b.Cursor()
			for k, v := c.Seek(keyPrefix); k != nil && bytes.HasPrefix(k, keyPrefix); k, v = c.Next() {
				val, err := f(k, v)
				if err != nil {
					return err
				}
				if val == nil {
					if err = b.Delete(k); err != nil {
						return err
					}
					continue
				}
				if bytes.Compare(val, v) != 0 {
					b.Put(k, val)
				}
			}
			return nil
		}
		return ERR_BUCKET_NOT_EXIST
	})
}
