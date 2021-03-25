package gointf

// Storer is an interface for DB connection
type Storer interface {
	NewBucket(bucket string) error
	DeleteBucket(bucket string) error
	Flush() error
	
	Put(bucket, key string, data []byte) error
	Get(buckt, key string) ([]byte, error)
	Delete(bucket, key string) error
	Iterate(bucket, prefix string, fn func(key string, value []byte)) error
}
