package gointf

// Storer is an interface for DB connection
type Storer interface {
	Lock()
	Unlock()
	Flush() error
	
	NewBucket(bucket string) error
	DelBucket(bucket string) error
	
	Get(bucket, key []byte) ([]byte, error)
	Put(bucket, key, val []byte) error
	Del(bucket, key string) error

	// Do will get/update/date using func([]byte)([]byte,error). If returned []byte is nil,
	// it means delete the key.
	Do(bucket, key []byte, func(val []byte)([]byte, error))([]byte, error)
	
	// Do iter will take keyPrefix, and based on that iterate.
	// When the func returns nil for []byte, it will delete the key.
	DoIter(bucket, keyPrefix []byte, func(key, val []byte)([]byte, error)) error
}
