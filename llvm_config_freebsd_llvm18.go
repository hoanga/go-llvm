//go:build !byollvm && freebsd && llvm18

package llvm

// #cgo CPPFLAGS: -I/usr/local/llvm18/include -I/usr/local/llvm18/include/llvm-c -D_GNU_SOURCE -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS
// #cgo CXXFLAGS: -std=c++17
// #cgo LDFLAGS: -L/usr/local/llvm18/lib -lLLVM
import "C"

type run_build_sh int
