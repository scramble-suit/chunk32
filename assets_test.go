// Code generated by go-bindata.
// sources:
// assets-test/test1
// assets-test/test2
// assets-test/test3
// assets-test/test4
// assets-test/test5
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _test1 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x92\x31\x8e\xd4\x4c\x10\x85\xf3\xff\x14\x2f\xfc\x91\xbc\x23\x40\xe2\x00\x04\x0b\x09\x11\x02\x11\xd7\xb8\x9f\xed\xd2\xb4\xab\x9a\xee\xf2\x5a\x43\xc4\x35\xb8\x1e\x27\x41\x6d\xc3\x6a\x13\x27\x76\x55\x7d\xef\x7b\x7e\xd4\x79\x09\xdc\x29\xb5\x41\x66\x1f\x10\x0b\xf1\xd1\x11\x94\x15\x6a\x51\x3d\x6d\x23\x13\x66\x57\x6b\x21\x39\xe3\xff\x7d\xd1\x71\x41\x66\x42\x38\x66\xc7\xcc\x78\x05\xb1\x84\x5d\x63\x81\xc6\xb1\x22\x71\xa4\x45\x95\xac\x3f\x98\x06\x7c\xfd\xfc\xe9\x21\xeb\x8d\xd0\xb5\x78\x0d\x14\x89\xa5\x21\x16\x89\x7e\x2c\xf1\x89\xd9\x0b\x3b\x43\x25\x26\x59\x35\xab\xd4\x73\x61\x78\x92\xfb\x05\xef\xa7\x60\xc5\x4e\x54\x66\x4a\x7b\x89\x34\xc0\x8d\xf0\xe9\x38\x3c\x69\x6d\x81\xef\x1b\x5b\xa8\x5b\x43\xa1\x97\x4c\x48\xbb\x31\x61\x97\x86\xc5\xf7\x0e\xae\x36\x7a\x2d\x5e\x25\x88\x27\xd6\xa6\x6e\x50\x9b\xbc\xae\xd2\x07\x2f\xf8\x46\x48\x5a\x35\xa2\xcf\x11\x49\x93\xfd\xfe\xf9\x2b\x70\x33\xdf\x2f\xf8\xe0\x15\x82\xec\x36\x23\x74\xe5\xd0\x3f\xb9\x32\x2b\x9f\xba\x97\x9e\xab\xc3\x94\xea\xd7\xcc\xb5\xb3\x15\x19\x6f\x32\x3f\xdf\x52\x9b\xb1\xfb\x96\x13\xae\x7d\xb0\x05\x9a\xe7\x3e\x7b\xbd\x43\x0c\x92\xd2\x83\x1b\xc2\x3d\x0f\xa7\x5b\x82\x36\xfa\x56\x65\x66\xfa\x17\x2a\x1c\x63\x65\x4f\xe0\xc6\x0b\xbe\x9c\xd5\x8d\xbe\xae\x9b\x69\xdc\xff\xbe\x4c\x58\xc5\xee\xc7\xae\x76\x2a\x4d\x3a\x4d\xac\xb4\x80\x94\x52\x5d\xc6\x85\xed\x82\x47\x19\x97\xc3\xe4\xc2\x5c\x98\xb0\x35\xf4\xbe\xaf\x8c\x6e\x7e\xb3\xc4\xda\xa2\xb3\xbc\x48\x36\xe0\xba\x45\x67\x5e\x35\x3d\xbc\x7d\xfd\xe6\x5d\xef\xbf\x5b\x1e\x33\xa5\x3e\x8b\xa8\xc4\xde\x1f\x76\xb8\xf7\x13\xa8\x79\xde\x8e\x8e\x0e\xd7\x46\xa6\xf3\x97\x92\xe4\x25\x20\x68\x6a\x73\xe6\x00\x9f\x26\x1d\x55\xf2\x91\xe0\xf2\xdf\x9f\x00\x00\x00\xff\xff\xd0\xd6\xf7\xe5\xb2\x02\x00\x00")

func test1Bytes() ([]byte, error) {
	return bindataRead(
		_test1,
		"test1",
	)
}

func test1() (*asset, error) {
	bytes, err := test1Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test1", size: 690, mode: os.FileMode(436), modTime: time.Unix(1524023998, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test2 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x4b\x8e\x13\x31\x10\xdd\x73\x8a\x77\x80\x51\xc4\xc0\x8e\x1d\x1b\xa2\x59\x22\x40\xac\x2b\xee\xea\xb8\x34\xb6\xcb\x72\x55\x27\x93\x5d\xae\x31\x12\x5c\x2e\x27\x41\x76\x4f\x98\xf0\x91\x7a\x55\x7e\x5d\xef\x53\xef\x6b\x64\x64\x35\x87\xe4\xaa\xcd\xa9\x38\x0a\x1f\x31\x33\xf9\xd2\x18\x3a\x63\xab\xb8\xc7\x91\x0c\x45\x1d\x84\x44\x65\xbf\xd0\x9e\xaf\x90\x0d\x1e\x7c\x3c\x77\xdc\xe5\xfc\x6c\xe0\x5c\x23\x99\x18\xb4\x60\x47\xe1\xf1\x48\x6d\x32\x04\xcd\x95\x5c\x76\x92\xc4\x4f\x1b\x7c\x2b\x2e\x09\x1e\xc9\x51\x55\x8a\xe3\xc8\x97\xf3\xf3\x04\x31\x5b\x78\x82\x39\xed\x12\xa3\x71\x62\x32\x86\x15\xaa\x16\xd5\x0d\x54\x6b\xd3\x27\xc9\xe4\x9c\x4e\xc8\x5a\x3c\xa6\xd3\x1d\x98\x42\xc4\x51\x3c\xc2\x64\x5f\x64\x96\xd0\x9d\x48\xb9\xb2\x26\x46\x88\x54\xf6\x6c\x1b\x7c\x67\xe8\xce\xb8\x1d\x3a\xcf\x0d\x9a\x42\xe0\xc4\x8d\x5c\xb4\x40\xfa\xe7\xdc\xd8\x1c\x54\x26\xd0\xa4\x75\x7d\xc8\x99\x27\x59\xe9\x69\x76\x6e\xf0\xf8\x2a\xf4\x25\xaf\x41\xb2\xe3\x24\x7c\xe0\xd5\x64\x07\xd5\xa6\x59\x56\xd0\x1f\x69\x20\xd3\xc4\x98\xf8\xc0\x49\x2b\x37\xc3\xcc\x9c\x90\x97\x10\x91\xb5\x71\x07\xcf\xfd\x36\x2f\x89\x9c\xa4\xec\x7b\xb6\x5b\xc5\xac\xad\x6f\x9d\x96\x30\xb4\x2d\xc6\x43\xac\x18\x08\x8f\x7c\x42\x63\x32\x2d\xab\x82\xad\xf6\x79\xd5\xba\x24\x6a\x70\x9d\xe8\xb4\xc1\x17\x29\x81\xf1\xee\xed\xfd\xfb\xa1\x70\xab\xf8\xf4\xf1\x33\x22\x19\xb8\x04\x5d\x1a\xed\x79\x42\xa5\xf0\xd8\x2f\x7e\xa3\xd0\xb5\xf3\x1e\x64\xea\xf6\x58\x1a\xf4\x38\xe8\x9b\x5d\xaf\x90\xa5\xd3\xf0\x53\xe5\xe0\x23\x52\xfb\xc7\xf7\x48\x29\x50\xea\x3d\x10\x1b\x02\xd6\x16\xfe\x15\x4f\x5b\x12\x7f\xc0\xe5\xfc\xe3\x61\x06\x15\x68\x7a\x95\x34\x4e\x33\x1a\x7b\x9d\x44\x1a\x91\x33\x8c\xf2\xef\x7d\x95\x3c\xde\x8d\xe9\x2d\x34\x2f\xe6\xd8\xf1\x7f\x4a\x9a\x78\xb5\xd1\xff\xb8\xa1\xdb\x5c\xce\x3f\xdf\xfc\x0a\x00\x00\xff\xff\xe6\x17\x36\x61\x32\x03\x00\x00")

func test2Bytes() ([]byte, error) {
	return bindataRead(
		_test2,
		"test2",
	)
}

func test2() (*asset, error) {
	bytes, err := test2Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test2", size: 818, mode: os.FileMode(436), modTime: time.Unix(1524024240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test3 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x76\xdb\x30\x0c\x44\xf7\x39\xc5\x1c\xc0\x95\x13\xa7\xdd\xe4\x02\x79\xd9\xb7\x07\x80\x48\x48\x62\x4d\x02\x2a\x01\xc9\xd5\xed\xfb\x68\xbb\x8e\xd3\x66\xa7\x27\x60\x66\x80\x0f\xbe\x49\xe4\x99\x25\xb2\x78\xde\x76\x30\x2e\x24\x9e\x02\x56\xae\x96\x54\x92\x8c\x98\xc8\xd0\x73\xd0\xc2\xf0\x89\x11\x19\x03\x05\x57\x98\x93\x44\xaa\x11\x83\x56\x44\xb6\x50\x53\xdf\xfa\x4d\x07\x3f\x51\xe5\xbf\x1e\x86\x24\x28\x24\x1b\x32\xc9\xb8\xd0\xc8\x08\x5a\xca\x22\xc9\x13\xdb\x0e\x49\x42\x5e\x62\x53\x36\xfb\x57\xbd\x55\xb7\x0e\x3f\xec\xec\xf8\xff\x54\x3b\x64\x72\xae\xef\x19\x2d\x90\x7f\xcf\x1c\x9c\x23\x5c\xd1\x33\x7a\x0a\xc7\x13\xd5\x68\x5f\x82\x96\x99\x3c\xf5\x99\x71\x4a\x3e\x81\xa9\xe6\x74\xa7\xde\xa1\x5f\x1c\x2a\x79\x3b\xd7\x93\x80\xd0\x92\x33\xa3\xd0\x4f\xbd\x35\xbe\x60\x7d\xea\x0e\xdd\x33\xca\x62\xde\x22\xfe\x35\x3e\x97\x9f\x40\x12\xdb\xe7\x53\xf7\xed\xe2\xbc\x1e\xba\xe7\xee\x2b\x84\x39\x42\xf4\x53\x65\xe3\xa3\x03\x7c\x52\xe3\xee\xe1\xe1\x6d\xc0\x89\x41\x51\x67\xff\xf4\x28\x8d\xf9\xab\x62\xa6\x70\xa4\xb1\x51\x24\x43\x51\xf3\xf6\x33\xf2\xca\x59\x67\xae\x76\x25\xb2\x6b\x64\xe5\x8c\x37\x95\x59\xab\xdf\xd2\x53\x4e\xbe\xa1\x2e\x99\x51\xf9\xd7\x92\x2a\x1b\x7c\x22\x47\x4c\xc3\xc0\x95\xc5\x3f\x02\xb0\xcb\xe6\x8b\xf1\x5d\xc7\xd5\x73\x26\x9f\xac\xc3\xf7\x29\x19\xb4\x37\xae\x2b\x79\x52\x41\xe6\x88\xc5\xda\x4d\x6e\x8b\x5c\x15\xf7\xe7\x4c\x82\xd3\x94\xc2\xf4\x9e\x64\x4e\xd5\xdb\xae\x74\x06\xf8\xd8\x3d\x5e\x9f\xca\xe5\x1d\x7e\x98\xab\xc9\xef\xd6\x6b\xa3\xbc\xa0\x6c\xfb\x76\xcb\x71\xbf\x1e\xf6\xb6\xf4\xfb\xf9\x38\x76\x0f\x7f\x02\x00\x00\xff\xff\xaa\xd2\x87\x46\xf1\x02\x00\x00")

func test3Bytes() ([]byte, error) {
	return bindataRead(
		_test3,
		"test3",
	)
}

func test3() (*asset, error) {
	bytes, err := test3Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test3", size: 753, mode: os.FileMode(436), modTime: time.Unix(1524024335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test4 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x56\x4b\x92\xe3\x36\x0c\xdd\xcf\x29\xde\x01\x3c\xae\x9a\x2c\x7b\x97\x4f\xd5\x54\xaf\x93\xaa\xac\x21\x0a\x96\x50\x43\x11\x6a\x82\xb4\x46\x59\xe5\x1a\xb9\x5e\x4e\x92\x02\x69\xbb\xdb\xd3\xee\x2c\x2d\x82\x7c\x00\xde\x7b\x80\x7f\xc6\xce\x94\x41\x93\xe2\x19\x56\xb2\xa6\x29\xee\x18\x38\x0a\x9f\x79\x44\x99\xa9\x60\x9b\xb9\xcc\x9c\x51\x14\x92\x42\xac\x23\xe3\xcc\xd9\x44\x13\x52\x5d\x06\xce\x06\x49\x90\x65\xd5\x5c\xb0\x52\x99\x0d\x1b\x19\x22\xe5\x89\xe3\x0e\xc2\x42\xa5\x70\x86\x9e\x50\xc8\x0a\x1f\x40\x69\xc4\x73\x0b\xb2\x6f\xbc\x16\x09\x14\x3b\xd4\x4c\x67\x49\x13\xca\xcc\x4b\x3b\x5e\x29\x17\x09\x35\x52\x8e\x3b\x38\xf2\x44\xa9\x1c\xf1\x4b\x2d\x1e\x82\x91\x83\xb4\x34\x4a\xcd\xc9\xa0\xfe\x59\x31\xf0\x2b\x64\xd2\x72\x83\xc5\x50\xdb\x8f\xa8\x93\x84\xa7\x6b\xbe\x41\x97\x95\x8a\x0c\x12\xa5\xec\x2d\x31\xe3\x85\x52\x91\x70\x2d\xb2\x25\xa4\x53\xef\x41\xe6\x97\x2a\x99\x5f\x83\x2e\xcf\xbc\xc6\x1e\xf1\xe7\xcc\x09\xcf\xc8\x4c\x51\xfe\x6a\x4d\x14\x3b\xb4\x84\x1b\x34\x45\x24\x0e\x6c\xe6\x80\x56\xf3\x9a\xc5\x78\xc4\xc2\xc7\x4f\x9f\x7a\x53\x28\x9a\xbe\x39\x29\x7a\x7d\xab\x37\xc9\x13\x61\x88\x81\x60\x1c\x34\x8d\x07\x48\x1a\x79\xe5\x34\x72\x2a\x37\x90\xac\xb5\xb0\x5f\xfe\x38\xd7\x27\x4c\x99\xc6\x4a\x11\x41\x47\x46\xe6\x95\x24\x43\x73\xef\xfb\xf5\x73\x5d\x3d\x8a\xed\x88\xe7\x04\xea\xbc\x62\xcd\x3a\x65\x5a\x0e\x90\xf2\xef\xdf\xff\x18\x6a\x6a\x39\x9a\xe3\x14\x05\x7f\x5f\x39\x14\x50\x8c\x58\x29\x7c\xa3\x89\x9b\x48\xbc\x09\x97\x9b\x1e\x55\xd7\x91\x0a\xe3\x94\x75\xc1\xf9\x8b\x7f\x39\xff\xe4\x14\xd1\x1b\xe2\x71\xad\x2c\xec\xe8\xc5\xc3\x68\x61\x14\x59\xd8\x33\xb2\xc2\xe4\x1d\x28\x58\xaa\x15\x67\x7f\x55\x33\x19\x22\xe3\xa4\x19\xa6\x0b\x37\x09\xdc\x23\x7f\x63\x5e\x51\xcd\xb9\x3d\x7f\xc1\x36\x4b\x64\x68\x63\xd8\x91\xcd\x85\x78\x2b\x7c\xec\x89\xdd\x74\x77\x57\x46\x2b\x7e\xa8\x12\xc7\x2e\xeb\x1f\x8f\x4e\x92\x28\x62\x90\x44\x79\x3f\xf4\x14\xaf\x26\x1a\xb4\xcc\x8e\xee\xd7\x7a\xdd\x5d\xd4\xd7\x72\x8f\xf8\x2a\xaf\x76\xb8\x15\xfe\xc6\x68\xd8\xb4\xc6\x11\x91\xa9\xe5\x18\x34\x9d\xaa\x53\x7b\xc0\x59\x34\x52\xf1\xcb\x5b\xb3\x30\x63\x91\x69\x2e\x08\xce\x88\x3f\x75\x79\xa5\x26\x79\xa9\x9c\xd8\x0c\xb9\x46\x7e\xc2\x28\xa7\x13\x67\xd7\xd1\x8d\xb7\x96\x74\x6b\xc8\xeb\xe1\x5b\xb7\x1f\xf1\xc7\xcc\xd0\x14\x77\x6c\xb4\x7b\x22\x2d\xf8\xa1\x86\x0e\xef\x81\x0f\x1f\xba\x4e\xcc\x1f\xa3\x51\xd7\xf2\x3f\x22\x06\x19\x36\x8e\xd1\xed\x53\xfc\x8e\x9e\x10\xb4\x66\x7b\xa3\x04\x1f\x0b\xce\x11\x6c\xb7\xc2\x8b\x75\x23\x55\xe3\x87\xb8\x9b\x94\xd9\x87\xc9\xc7\x98\x87\x3e\x4c\xbc\xe4\x61\xc7\xd4\x69\xaa\x2b\x58\x6e\x1a\x7a\x57\xba\xfb\xea\x5d\xf1\x47\xfc\x4a\x79\x52\x37\x8a\x6e\xf6\xc1\xc5\x3b\x88\x77\x4f\x3c\x81\xfc\x98\xef\x46\x30\x02\xa5\x1f\x39\x5b\x98\x3c\xf7\x66\xc4\xb7\x34\xbb\xde\x9b\xe9\xba\xb3\x5b\xa3\x8e\xf8\x8d\x57\x70\xb2\x9a\xdd\xb9\xef\xc4\x72\x97\xd2\xc3\xac\x9f\x7e\x34\xff\x59\xa3\xaf\x14\x49\xf7\x48\x5d\x5e\x27\x49\xa3\xcf\x33\x49\x53\x64\xd0\x94\x99\xc7\xcf\x75\xd5\x74\x5b\x35\x2d\xc3\x5e\xe7\xab\x45\x0e\xc8\x24\x76\xf1\xc8\x85\xee\x3e\xc9\x1b\xc1\x77\xa3\xca\xb0\x49\x8c\x3e\x20\x6a\x6a\xc8\x34\x44\xbe\xf6\x5f\x0c\xb9\x19\xa4\x2d\x39\x13\x73\x76\x3f\x92\xf0\x36\x4b\x98\x41\x99\x11\xb2\x5c\xb6\x97\x76\xb0\xcf\x16\x28\x32\x4c\x4f\x65\xf3\x80\x91\xcf\x1c\x75\x5d\xd8\xf7\x96\xb7\x54\x0c\xfc\x52\x29\xc6\xfd\x11\xe0\x23\x7d\xe8\xb2\x46\xfe\xee\x62\x6d\x34\x7d\xd5\x36\x57\x42\xcd\x8d\xbe\x33\xa7\x51\xb3\x77\xc0\xea\xda\x37\x1a\xa5\x8b\xfd\x1f\xb8\xfc\xb2\x9d\xca\xcc\x7b\x93\x47\xdf\x4b\x99\xad\xc6\x36\x2d\xd6\xac\x43\x74\x83\x34\xed\x0c\xcc\x09\x2f\x55\x0a\x23\xcc\x14\x23\xa7\xc9\x83\x7c\xae\xb6\xe1\x75\xa9\xce\xff\x01\xb4\xd9\xa7\x1a\x9b\x65\x6b\x1a\x39\x5b\xa1\xd4\x74\x14\x64\xf4\x5b\x03\x97\xcd\xdf\x7b\xac\x71\xbf\xff\x60\x26\xf5\x75\x6b\x58\x33\x8f\x12\x2e\x03\xcd\xbb\xdf\xb8\x9c\x6b\x2e\x58\x34\x37\x83\xbb\x36\x50\xd7\x23\x7e\xff\x78\x54\x44\x2e\x86\x6a\xa0\xb3\x4a\x9f\xd5\x61\x56\x09\xdc\xe0\xdb\x4e\x68\x75\x49\xdf\x2a\xc7\x4f\xff\x05\x00\x00\xff\xff\x30\x0f\x6f\x6f\x1f\x09\x00\x00")

func test4Bytes() ([]byte, error) {
	return bindataRead(
		_test4,
		"test4",
	)
}

func test4() (*asset, error) {
	bytes, err := test4Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test4", size: 2335, mode: os.FileMode(436), modTime: time.Unix(1524029078, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test5 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x54\xcb\x6d\xf4\x46\x0c\xbe\xbb\x8a\xaf\x80\xfd\xb7\x81\x9c\xf2\x40\x02\x9f\x02\x04\x7f\x03\xd4\x0c\xb5\x62\xcc\x99\x51\x86\x94\xd6\xba\xb9\x8d\x00\x49\x73\xae\x24\xa0\xe4\xb5\xbd\x88\x2f\xba\x48\xe4\xf7\xa4\x1e\x71\x25\x03\xa9\x35\xd8\xd2\xe7\x2e\xc6\x19\xde\x90\xc5\x52\x5b\xb9\x63\x6a\x57\x94\x25\x4d\x90\x32\xb7\xee\x48\xad\xcc\xe4\x32\x88\x8a\x6f\x30\x29\xb3\xca\x28\x6c\x58\xb9\x9b\xb4\x0a\x63\xe5\xe4\xd2\xea\x09\xd7\x49\x62\xd0\xe0\x13\x63\xee\x6d\x50\x2e\x68\x23\x32\x27\xc9\x52\x2f\x6f\x1f\xcc\x94\x9e\xe8\xc2\xb7\x0d\x16\xf8\x8b\x31\xc6\xd6\x41\xb8\xc8\xca\x15\xc3\x22\x9a\xcf\xf8\x3e\x31\x52\xab\xe6\x9d\xa4\xba\xc5\xb2\x9f\xa9\x5f\x1a\xa8\x66\xfc\xc2\x33\x0a\x3d\xf1\xff\xa9\x80\xff\x5a\x64\x25\xe5\xea\xb1\xdc\x9a\xae\x01\xff\x53\x6b\xca\x54\x61\xe4\x62\xa3\xd0\x21\xea\x84\xc2\x54\xe3\xbd\x38\x12\x55\x0c\xfb\xc6\x0d\xfc\x3c\x73\x35\x59\x79\x37\x88\x9d\x7b\x91\xca\xb8\x4e\xec\x13\x07\xd5\x95\x54\xf2\x3b\x7a\x6a\x75\x94\xcb\xd2\xe9\x60\x10\x2a\xf8\x59\xcc\xed\x8c\x1f\x6b\x0e\x4f\x6a\x3c\x3a\xa3\xd0\x16\x20\x85\xea\xf6\xb6\xe3\x6e\xd6\x4e\xb8\x8a\x4f\xa8\x0d\x49\x99\x3a\x52\x17\xe7\x2e\xb4\x3b\x94\xa6\xd6\x2c\xd8\x86\xc9\xaf\x2f\xff\x0c\x6c\xfe\xfa\xf2\x2f\x5a\xe5\x33\xfe\x60\xdd\xe2\x5d\xab\x5f\xe7\x17\xf2\xa4\x9a\x33\x65\x28\x3b\x7e\x3b\x9c\x27\x78\x97\x55\x48\x4f\x50\xa9\x4c\xfd\x9b\x4b\x61\x90\x5e\x5a\x17\x9f\x4a\x18\x30\xca\x21\x02\x01\xae\x8c\xc0\xbd\xe7\x7d\x2b\x00\xe9\x95\x36\x7b\x17\xff\x7d\x12\xfb\x58\x75\xfb\xe8\x11\x89\x54\x51\xa4\x4a\x21\xfd\xaa\x4d\x52\xe1\x4b\xaf\x60\x95\x22\x95\x9c\x8f\x5e\x55\xe6\xbc\x1b\x61\x3c\x53\x27\x67\x68\x4b\x4f\x7b\x21\x0a\x55\x19\x83\xd6\x28\xca\x76\xc6\xa3\xa3\xf3\xac\x94\x8e\xd1\x72\xd8\x4a\x6f\x0a\x4e\xb0\xe9\x70\xe8\x73\x6e\x31\x7a\x02\x67\x71\xce\xc8\xd2\x39\xb9\x6e\x18\x36\x0c\xcd\x27\x64\x5e\x59\xdb\xcc\xdd\x76\x40\x6f\x4d\xed\x04\x9f\xc8\x61\x2e\xaa\xb0\x65\x0e\xdb\x2d\x80\x7b\xcb\x4b\x92\x21\xbc\x8a\x36\xdb\xf9\xe1\xe1\xf7\xa5\xef\xb5\xea\xc2\x35\xf1\xc1\x27\x6a\x9c\xb9\x1c\x2d\xbf\xa9\x94\x32\x53\xf2\xe8\xfb\x5d\x80\x67\xfc\xda\x54\xdb\xf5\x96\xbf\x46\x90\x77\x47\xc1\xd4\x55\xb8\xc3\x36\x73\x2e\x51\x25\x46\x66\x93\x4b\xe5\xe3\x62\xbc\xed\x27\x86\x65\xfe\xba\x22\x64\x98\xa9\xef\xd0\x94\xdb\xec\x01\x65\x5c\xa8\xba\xa4\x5b\x4c\x52\x2f\x67\x3c\x22\xb7\xfa\xfa\xf2\xb7\x63\x60\x15\x5e\xf9\x80\x4a\x92\xe3\x97\x12\xa1\x67\x56\x19\x38\x44\xe9\xf6\x43\xbc\xfd\x73\x89\x70\x76\x05\xfb\x37\xcd\x18\x6d\x3f\xa6\x37\xba\xc7\xc5\x8f\xd2\xcd\xbf\x4d\xbb\x9c\x0f\xb3\xda\x88\x65\x2f\x7e\xa8\x98\x58\x67\xce\x58\x0c\x03\xbb\x73\xc7\x52\x33\x77\xf3\x63\x86\xf6\xd0\xde\x7f\x65\x21\x50\xf9\x39\xd4\x89\x21\x75\xa6\x08\x77\xd8\x30\xc7\x49\xfb\x2e\x51\xea\xcd\x06\xe5\x9b\x31\x33\xf9\x64\x71\x54\xab\xac\x37\xc7\xbf\xf4\xac\x2f\x11\xf2\x06\xa9\x7e\x84\xfe\xd9\xb3\xb7\x89\x0f\xeb\xee\x1b\x4d\xfe\x89\xdf\x69\x4f\x74\xc7\x6a\xa0\x83\xfd\xfe\xcf\x7d\xb7\xe8\xfc\xf0\x5f\x00\x00\x00\xff\xff\x63\x5a\xb8\x76\xc3\x05\x00\x00")

func test5Bytes() ([]byte, error) {
	return bindataRead(
		_test5,
		"test5",
	)
}

func test5() (*asset, error) {
	bytes, err := test5Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test5", size: 1475, mode: os.FileMode(436), modTime: time.Unix(1524029089, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"test1": test1,
	"test2": test2,
	"test3": test3,
	"test4": test4,
	"test5": test5,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"test1": &bintree{test1, map[string]*bintree{}},
	"test2": &bintree{test2, map[string]*bintree{}},
	"test3": &bintree{test3, map[string]*bintree{}},
	"test4": &bintree{test4, map[string]*bintree{}},
	"test5": &bintree{test5, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

