package cache

// Cache represents an in-memory cache, regardless of cache policy.
// Cache variants two queues (for least frequently as well as recently used) and ARC cache can be incorporated
type Cache interface {
	Put(key []byte, value []byte)
	Get(key []byte) []byte
	Len() int
	Cap() int
	Clear()
}
