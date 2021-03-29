package gointf

// Storer is an interface for DB connection
// If any value is nil, it means, the key doesn't exist.
// error will be returned if DB has problem running the command.
type Storer interface {
	// Lock will lock the mutex if supported
	Lock()
	// Unlock will unlock the mutex if supported
	Unlock()
	// Flush will flush the file if supported
	Flush() error

	// NewBucket will create a bucket if not exist. If already exists, it will
	// return an error.
	NewBucket(bucket string) error
	// DelBucket will remove bucket. If not found, it will return error.
	DelBucket(bucket string) error

	// Get will get values; if nil returns, there's no item.
	Get(bucket, key []byte) ([]byte, error)
	// Put will add values, it will overwrite if exists.
	Put(bucket, key, val []byte) error
	// Del will remove the item if exist.
	Del(bucket, key string) error

	// Do will get/update/date using func([]byte)([]byte,error).
	// If returned []byte is nil, it means delete the key.
	Do(bucket, key []byte, f func(val []byte)([]byte, error))([]byte, error)
	
	// Do iter will take keyPrefix, and based on that iterate.
	// When the func returns nil for []byte, it will delete the key.
	DoIter(bucket, keyPrefix []byte, f func(key, val []byte)([]byte, error)) error
}
