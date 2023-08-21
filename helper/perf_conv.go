package helper

import "unsafe"

type (
	goSlice struct {
		Ptr unsafe.Pointer
		Len int
		Cap int
	}
	goString struct {
		Ptr unsafe.Pointer
		Len int
	}
)

// BytesToStr zero allocation bytes to string conversion. Because default
// string([]byte) take allocation so this helper may help with performance
// improvement but should be used in caution and appropriately.
//
// Caution: Make sure the given b bytes will never change.
//
// Refs:
//   - https://github.com/golang/go/issues/19367
//   - https://github.com/golang/go/issues/25484
//   - https://github.com/bytedance/sonic/blob/1a770c051a6201efd3aced2663f4e5d30b8513c0/internal/rt/fastmem.go#L40
//
//go:nosplit
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StrToBytes zero allocation string to bytes conversion. Because default
// []byte("string") take allocation so this helper may help with performance
// improvement but should be used in caution and appropriately.
//
// Caution: Make sure the given s string will never change.
//
// Refs:
//   - https://github.com/golang/go/issues/19367
//   - https://github.com/golang/go/issues/25484
//   - https://github.com/bytedance/sonic/blob/1a770c051a6201efd3aced2663f4e5d30b8513c0/internal/rt/fastmem.go#L40
//
//go:nosplit
func StrToBytes(s string) (b []byte) {
	(*goSlice)(unsafe.Pointer(&b)).Cap = (*goString)(unsafe.Pointer(&s)).Len
	(*goSlice)(unsafe.Pointer(&b)).Len = (*goString)(unsafe.Pointer(&s)).Len
	(*goSlice)(unsafe.Pointer(&b)).Ptr = (*goString)(unsafe.Pointer(&s)).Ptr
	return
}
