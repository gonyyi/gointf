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

func NewBoltDB(filename string) (*boltStore, error) {
	b, err := bbolt.Open(filename, 0666, &bbolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}
	s := &boltStore{db: b}

	return s, nil
}

//  gointf.Storer interface
//		CreateBucket(bucket string) error
//		DeleteBucket(bucket string) error
//		Flush() error
//		Put(bucket, key string, data []byte) error
//		Get(buckt, key string) ([]byte, error)
//		Delete(bucket, key string) error
//		Iterate(bucket, prefix string, fn func(key string, value []byte)) error

// Default database for storer will be badger
type boltStore struct {
	db  *bbolt.DB
}

func (s *boltStore) CreateBucket(bucket string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
			return err
		}
		return nil
	})
}

func (s *boltStore) DeleteBucket(bucket string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(bucket))
	})
}

func (s *boltStore) Put(bucket, key string, data []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			return b.Put([]byte(key), data)
		}
		return ERR_BUCKET_NOT_EXIST
	})
}

func (s *boltStore) Get(bucket, key string) ([]byte, error) {
	var out []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			out = b.Get([]byte(key))
			if out == nil {
				return ERR_KEY_NOT_EXIST
			}
			return nil
		}
		return ERR_BUCKET_NOT_EXIST
	})
	return out, err
}

func (s *boltStore) Delete(bucket, key string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		if b := tx.Bucket([]byte(bucket)); b != nil {
			return b.Delete([]byte(key))
		}
		return ERR_BUCKET_NOT_EXIST
	})
}

func (s *boltStore) getBuckets() ([]string, error) {
	var out []string
	err := s.db.View(func(tx *bbolt.Tx) error {
		tx.ForEach(func(name []byte, b *bbolt.Bucket)error {
			if b != nil {
				out = append(out, string(name))
			}
			return nil
		})
		return nil
	})
	return out, err
}

func (s *boltStore) Iterate(bucket, prefix string, fn func(key string, value []byte)) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		// check if bucket exists,
		prefix := []byte(prefix)
		if b := tx.Bucket([]byte(bucket)); b != nil {
			c := b.Cursor()
			for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, _ = c.Next() {
				fn(string(k), v)
			}
			return nil
		}
		return ERR_BUCKET_NOT_EXIST
	})
}

func (s *boltStore) Flush() error {
	return nil
}
