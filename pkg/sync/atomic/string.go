package string

// AtomicString is a wrapper type for string that supports atomic operations
type AtomicString struct {
	ptr unsafe.Pointer
}

// NewAtomicString creates a new AtomicString with the given value
func NewAtomicString(s string) *AtomicString {
	return &AtomicString{
		ptr: unsafe.Pointer(&s),
	}
}

// Load returns the current value of the AtomicString
func (as *AtomicString) Load() string {
	return *(*string)(atomic.LoadPointer(&as.ptr))
}

// Store sets the value of the AtomicString to s
func (as *AtomicString) Store(s string) {
	atomic.StorePointer(&as.ptr, unsafe.Pointer(&s))
}

// CompareAndSwap compares the current value of the AtomicString with old and if they are equal, sets the value to new and returns true. Otherwise, it returns false.
func (as *AtomicString) CompareAndSwap(old, new string) bool {
	return atomic.CompareAndSwapPointer(&as.ptr, unsafe.Pointer(&old), unsafe.Pointer(&new))
}

// Swap sets the value of the AtomicString to new and returns the old value
func (as *AtomicString) Swap(new string) string {
	return *(*string)(atomic.SwapPointer(&as.ptr, unsafe.Pointer(&new)))
}
