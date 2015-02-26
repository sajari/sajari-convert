package justext

import (
		"bytes"
		"compress/gzip"
		"io"
		"reflect"
		"unsafe"
)

var _DefaultTemplate = ""+
"\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\x3d\x6e\xc4\x20"+
"\x10\x46\x6b\xfb\x14\xc8\x4a\x8d\xfb\x08\x53\x24\x45\x92\x26\x8a"+
"\x14\x5f\x60\x6c\x88\x41\xc2\x80\x80\x22\x2b\xc4\xdd\x17\xc3\x7a"+
"\xff\xb4\x5b\x81\xde\x3c\xe6\x1b\x86\x88\xb0\x2a\xda\x92\xc9\xb0"+
"\x03\x6d\x63\x7c\xd1\xe6\xcd\x48\xc5\x9d\x55\x10\x38\x7a\x1d\x10"+
"\xfe\xbe\x26\x29\x65\xc9\x81\x5e\x38\xc2\x3f\xe0\x60\x71\x60\x85"+
"\xcf\xb4\x89\x51\xfe\xa1\x2f\xff\x61\x0c\x43\xf8\x5d\x81\x2f\xb4"+
"\x62\xfc\xc9\x81\x49\xbd\x14\xd2\x10\x41\x63\x1c\x9d\x5c\x7f\x2d"+
"\xcc\xb9\xcf\xc8\xff\x43\x4a\xa4\x17\xb4\xf8\x5c\x79\x7e\x12\xed"+
"\x63\xd1\xd2\x5c\xac\xae\x66\x35\xfb\xfc\xa8\xe4\x69\x13\xd0\xed"+
"\x4f\xf6\x86\x68\xde\x26\x1b\xba\xe9\x52\xea\x9e\x85\xdc\x25\x94"+
"\xcb\x7e\x92\xbe\x6e\x2c\x4f\xbd\x2d\xf0\x18\x00\x00\xff\xff\x2c"+
"\xc5\xf5\x5d\x47\x01\x00\x00"

// DefaultTemplate returns the binary data for a given file.
func DefaultTemplate() []byte {
		// This bit of black magic ensures we do not get
		// unneccesary memcpy's and can read directly from
		// the .rodata section.
		var empty [0]byte
		sx := (*reflect.StringHeader)(unsafe.Pointer(&_DefaultTemplate))
		b := empty[:]
		bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		bx.Data = sx.Data
		bx.Len = len(_DefaultTemplate)
		bx.Cap = bx.Len
		
		gz, err := gzip.NewReader(bytes.NewBuffer(b))

		if err != nil {
			panic("Decompression failed: " + err.Error())
		}

		var buf bytes.Buffer
		io.Copy(&buf, gz)
		gz.Close()

		return buf.Bytes()
}