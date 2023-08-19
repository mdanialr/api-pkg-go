package helper

// Ptr return pointer to given t.
func Ptr[T comparable](t T) *T {
	return &t
}

// Def return the value of given pointer and return the default value of that
// type instead if nil.
func Def[T comparable](t *T) (n T) {
	if t == nil {
		return n
	}
	return *t
}
