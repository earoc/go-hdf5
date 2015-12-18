package hdf5

// #cgo LDFLAGS: -lhdf5 -lhdf5_hl
// #cgo linux CFLAGS: -I/usr/include/hdf5/serial/
// #cgo linux LDFLAGS: -L/usr/lib/x86_64-linux-gnu/hdf5/serial/
// #cgo darwin CFLAGS: -I/usr/local/include
// #cgo darwin LDFLAGS: -L/usr/local/lib
// #include "hdf5.h"
import "C"
