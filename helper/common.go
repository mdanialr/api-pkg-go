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

// Default return the default value while it's present and t is blank.
func Default[T comparable](t T, def ...T) (tt T) {
	// if given t is not empty just return it back
	if t != tt {
		return t
	}
	// return the given default if present
	if len(def) > 0 {
		return def[0]
	}
	return
}
