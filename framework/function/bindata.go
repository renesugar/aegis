// Code generated by go-bindata.
// sources:
// example_aegis
// example_main
// DO NOT EDIT!

package function

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

var _example_aegis = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4f\x6f\xd4\x30\x10\xc5\xef\xf9\x14\x4f\xe5\x4a\x8b\xca\x9f\x8b\x6f\xa9\x04\x12\x12\x42\x1c\xb8\x70\xaa\x26\xf6\xcb\xae\x85\x63\x5b\x1e\xbb\xe9\x7e\x7b\xe4\x64\x77\x11\x20\x7a\xca\xc8\xef\xcd\x9b\xf9\x4d\x24\x67\x33\x00\x51\x16\x1a\x7c\x7c\x96\x25\x07\x62\xe4\xc1\x2b\xc6\x9c\x07\xe0\x27\x99\x1f\x9a\x0f\xee\x93\x0f\x54\x83\x5a\x1a\x07\x59\xb5\x77\x15\x1e\x7c\x8a\x06\x4d\x6f\x29\x5a\x6f\xef\x87\x20\xcb\xe4\xa4\x6b\x8e\x6a\x8b\xcf\x75\x33\x8c\x11\xfc\x23\x5b\xb6\xec\xb9\x45\xdb\x0d\x5f\xb7\xe9\xd2\x95\xc7\xb3\x6f\x00\x5e\x61\x84\xe3\x2c\x2d\x54\xf0\x99\xb6\x75\x2b\x4a\x0a\xc4\xea\x43\xc0\x44\xd8\x42\xa9\x74\x98\x53\xc1\x29\xb5\xd7\x98\x5a\xed\x05\xac\x44\x48\x58\xe5\xa4\x68\x4a\x08\x34\xd3\xfa\xd9\x5b\xa4\xc8\xbb\x2d\xbc\x07\x19\x48\x89\xa6\xd3\x78\x59\x8c\xb9\x7f\xfb\xee\xfd\x07\xd3\x85\x37\xfb\x32\x3b\xce\xe3\x65\xcf\x41\xb2\xff\xef\xb5\xbe\x7d\x7e\x89\x7a\x53\x3b\xd1\x4d\x2e\xc9\xdd\x40\xab\x1c\xfe\xe5\x90\x56\xd3\x22\xd5\x5b\x09\xe1\x74\x87\x1f\x67\x12\xc7\xd9\x47\xee\x3d\x8a\x23\x0b\x21\x8a\x95\x21\xf4\xef\x1e\xf5\x24\xc5\xcb\x14\xa8\x3b\xdd\xf7\x23\x95\x7f\x4b\x90\x42\x64\x51\xa5\x83\x8f\xa8\x47\x82\x4f\x8c\x15\x0b\x55\xbb\xb3\xa6\x7e\xbd\x82\x2f\x1b\xf7\xf5\xff\xec\x91\xfb\x78\xb3\xd5\x40\xc7\xb8\xd4\x97\x8b\xf4\xb7\xeb\xd3\x75\xea\x6f\x17\x30\xa7\x64\xf0\x20\xe5\x57\x00\x00\x00\xff\xff\xbc\x4c\xe7\x68\x77\x02\x00\x00")

func example_aegisBytes() ([]byte, error) {
	return bindataRead(
		_example_aegis,
		"example_aegis",
	)
}

func example_aegis() (*asset, error) {
	bytes, err := example_aegisBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "example_aegis", size: 631, mode: os.FileMode(436), modTime: time.Unix(1510187626, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _example_main = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x93\x41\x4f\xdc\x3a\x10\xc7\xcf\xeb\x4f\x31\x2f\x97\x97\x3c\x85\x04\x3d\x71\x42\xe2\xf0\x84\x78\xd0\x8a\x52\x04\xb4\x3d\x54\x55\x35\x9b\xcc\x26\x16\x8e\x27\x8c\xc7\xdd\x45\x68\xbf\x7b\xe5\x64\x0b\x14\xd1\x6b\x7b\x88\xc6\x8a\x67\xc6\xff\xff\xcf\xe3\x11\x9b\x5b\xec\x08\x06\xb4\xde\x18\x3b\x8c\x2c\x0a\xb9\x59\x64\x0d\x7b\xa5\x8d\x66\x66\x91\x39\xee\x52\xf0\xa4\x75\x14\x97\x96\x9d\xd5\x3e\x2e\xab\x86\x87\x1a\xd7\x21\x7d\x7b\x0e\x87\x65\x8b\x7b\x1d\xd7\xf3\xea\xa9\x1e\xa9\xb3\x01\x9e\xd7\xe8\x80\x16\x85\x95\xeb\x69\xaf\x5e\x09\x0e\xb4\x66\xb9\xcd\x4c\x61\xcc\x2a\xfa\x66\xd2\x93\x17\xf0\x60\x16\x75\x0d\x67\xe8\x5b\x47\x80\x1e\xfe\xbb\x7c\x73\x8a\x4a\x6b\xbc\xbf\x14\xde\xdc\x5f\xd1\x5d\xa4\xa0\x40\xdf\xc8\x2b\xac\xad\xf6\x80\xf0\xe1\xea\x1c\x84\xee\x28\x06\x85\x11\xb5\x87\x2b\x8e\x4a\x62\x16\x32\x45\x38\x3c\x82\xe9\xd8\xea\x82\xd6\xf3\x56\xbe\x42\xe7\x6e\x7a\xe1\xd8\xf5\xc5\x8f\xbc\x6a\x3e\x35\xcf\x4e\x4f\x6e\xb2\x12\xb2\x3a\x2b\x41\x98\xb5\x84\x9e\x9c\xe3\x77\xb6\x6d\x1d\xad\x51\xe8\xa9\xe2\xdc\x06\x25\x9f\x17\x66\x6b\x4c\x5d\xc3\xb3\xae\xd0\x4f\xcd\x02\xa0\xbf\x9f\x45\x69\x8f\x0a\x0d\x47\xd7\xfa\xbf\x15\x96\xe9\x06\xb4\xe9\xa9\x05\x65\x40\xcf\xda\x93\xec\x8a\x64\x26\xf2\xac\x5b\xde\xe8\x06\x76\x80\xab\xe3\x39\x96\xc9\x33\xfc\x33\x3b\x7b\x1d\x53\x4a\x09\xbf\x4c\x09\x23\xfb\x40\x25\x8c\x28\x38\x04\x88\xe2\xaa\x8f\xe8\x22\x85\x02\x48\x84\x25\xdd\x85\x50\xa8\xae\x15\x35\x86\x63\x6e\x09\x8e\xe0\x60\xff\x20\xfd\xd5\x28\x1e\xbc\x75\x3b\xe3\x09\x13\xd8\x30\xeb\xb7\xbe\x83\xd3\x93\x9b\x04\x10\xac\x07\xed\x6d\x80\x06\x03\xcd\xae\x52\xea\x1f\xb3\xe3\x9a\x12\xbe\xa6\x79\xf8\x69\x64\xab\xff\x85\x87\x9d\x8c\x24\xad\x98\x7d\xbf\xbd\x7e\x7f\x91\xff\xbb\xbf\x5f\xc2\x80\xe3\xe7\xa0\x62\x7d\xf7\xc5\x7a\x25\x59\x61\x43\x0f\xdb\x87\x6c\x9a\xc2\xec\x30\x29\x2f\xe1\xf1\x01\x1d\x82\x6b\xb6\xc5\x2b\x94\x5e\x8c\x51\x02\x86\x10\xec\x30\x3a\x02\xda\xe0\x14\x79\x05\xc3\x63\xc6\x4c\xec\x45\xd9\xef\x86\xb7\x64\x76\x13\x3b\xee\xaa\x4b\xb1\x5e\x9d\xcf\xb3\xb3\xa4\x09\x8e\x1d\xc7\xf6\x53\x1a\xe3\xbf\xb2\x27\xc3\x2a\x91\xcc\xf6\x7b\x00\x00\x00\xff\xff\xa3\xc6\x37\xe9\x68\x04\x00\x00")

func example_mainBytes() ([]byte, error) {
	return bindataRead(
		_example_main,
		"example_main",
	)
}

func example_main() (*asset, error) {
	bytes, err := example_mainBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "example_main", size: 1128, mode: os.FileMode(436), modTime: time.Unix(1522126316, 0)}
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
	"example_aegis": example_aegis,
	"example_main":  example_main,
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
	"example_aegis": &bintree{example_aegis, map[string]*bintree{}},
	"example_main":  &bintree{example_main, map[string]*bintree{}},
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
