// Code generated by go-bindata.
// sources:
// assets/errors.csv
// DO NOT EDIT!

package errors

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

var _assetsErrorsCsv = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x8e\xb1\x4e\x33\x31\x10\x84\xfb\x3c\xc5\x76\x7f\x73\x8a\xfd\x4b\x49\x43\x87\x10\x0d\x05\x20\x85\x17\x70\xec\x41\x67\xb8\xdb\x35\xeb\xf5\x25\x79\x7b\x74\x3e\x51\x50\x51\xee\xe8\x9b\xf9\xf6\xe8\xbd\xf7\xc3\xd1\xfb\x61\x34\x2b\x77\xce\xe1\x1a\xe6\x32\x61\x1f\x65\x76\x97\xfc\x99\x1d\x54\x45\xab\xdb\xc0\x13\x74\x81\xd2\xe3\x9a\x0d\xf7\x4c\x8d\x13\x0c\x3a\x67\x46\xa2\x4e\xd2\x18\x4a\xc1\x7a\x0a\x93\x8d\xa0\xad\xb2\x3b\x78\xef\xff\x0f\x87\xbf\x45\x1b\xf8\x10\x98\xc5\xa8\x04\xad\xa0\xb3\xa4\x1b\x65\x36\x21\x5c\x0b\xa2\x21\xd1\xd3\xe9\xe5\x99\xe4\xfc\x81\x68\xc3\xdb\x08\x52\x7c\x35\x54\xa3\x24\xa8\xfc\xcf\x28\x8a\x2a\x6a\x11\x4e\x64\xd2\xff\xa8\xa6\x2d\x5a\x53\x10\x03\x09\x3d\xaf\x32\x2d\xa0\x9b\x34\xfd\x19\xd8\xd3\xeb\x84\x50\xd7\xc1\x25\xe3\xd2\x9b\x5d\x2f\xef\xbf\xb9\xdd\x77\x00\x00\x00\xff\xff\x5d\x21\x36\xa6\x3a\x01\x00\x00")

func assetsErrorsCsvBytes() ([]byte, error) {
	return bindataRead(
		_assetsErrorsCsv,
		"assets/errors.csv",
	)
}

func assetsErrorsCsv() (*asset, error) {
	bytes, err := assetsErrorsCsvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/errors.csv", size: 314, mode: os.FileMode(438), modTime: time.Unix(1482920414, 0)}
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
	"assets/errors.csv": assetsErrorsCsv,
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
	"assets": &bintree{nil, map[string]*bintree{
		"errors.csv": &bintree{assetsErrorsCsv, map[string]*bintree{}},
	}},
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

