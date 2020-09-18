// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package templates generated by go-bindata.// sources:
// templates/emptymodule.go.tpl
// templates/module.go.tpl
// templates/schema.graphql
package templates

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _emptymoduleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x4d\x6e\xeb\x20\x14\x46\xe7\xac\xe2\x7b\x19\x25\xaf\x92\xbd\x88\xa4\x83\x0c\xda\x4c\xb2\x01\x0c\xd7\x98\xd6\x70\x29\x5c\x2a\x59\x51\xf6\x5e\xf9\xa7\x49\x67\x7c\xc0\x39\x3a\x6d\x8b\x23\x5b\x82\xa3\x48\x59\x0b\x59\x74\x13\xfa\x51\x07\x1f\x1d\x37\x81\x5a\x97\x75\x1a\xbe\x46\x9c\x2e\x78\xbf\x5c\xf1\x7a\x3a\x5f\xff\xa9\xb6\xc5\x4b\x57\xfd\x68\xb1\x3d\x2b\x95\xb4\xf9\xd4\x8e\x9e\x17\x3e\x24\xce\x82\xdd\x5f\x99\x9d\x0f\x3b\x35\xf3\x32\xf8\x02\x0a\x49\x26\x04\xb6\x75\x24\xf8\x82\x5a\xd6\x00\xc7\x8f\x20\x08\x23\x65\xfe\xf6\x96\xa0\xb1\x08\x7e\x81\x8e\x8c\xae\x85\xb6\xd9\x38\x9e\x15\x91\x05\x86\x43\xf2\x23\x59\xd8\x9a\x7d\x74\x0f\x97\x92\x29\x11\xde\x56\xba\x48\xae\x46\x6e\xf7\xa5\xe6\xc8\xb1\xf7\xae\x66\x82\xad\x21\x4c\x08\x24\x03\x5b\xd5\xd7\x68\xb0\x5f\x81\xc3\xf3\xd3\xde\xc7\x0f\x32\xc2\x19\xff\x97\xa0\xe6\xbc\xed\x03\x6e\x77\xf5\x13\x00\x00\xff\xff\x37\x32\xf2\x7f\x53\x01\x00\x00")

func emptymoduleGoTplBytes() ([]byte, error) {
	return bindataRead(
		_emptymoduleGoTpl,
		"emptymodule.go.tpl",
	)
}

func emptymoduleGoTpl() (*asset, error) {
	bytes, err := emptymoduleGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "emptymodule.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _moduleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x4f\x6f\xdb\x38\x13\xc6\xcf\xd2\xa7\x98\x57\xc0\xbb\x21\x13\x97\xaa\xb1\xa7\xba\x9b\x4b\xd3\x1c\x7c\x88\x53\xb4\xce\x5e\x1c\xa3\xa0\xa5\x11\xcd\x9a\x22\x15\x92\x72\x1a\x18\xfa\xee\x0b\xd2\x52\xec\xa8\x28\xb0\xd8\x8b\x25\xcf\xbf\xe7\xf7\x8c\x28\xe5\x39\xdc\x98\x12\x41\xa0\x46\xcb\x3d\x96\xb0\x79\x81\x4a\xf1\x5a\x6a\x61\x58\x8d\xb9\xb0\xbc\xd9\x3e\xa9\x09\x7c\xbe\x87\xc5\xfd\x12\x6e\x3f\xcf\x97\x2c\xcd\x73\xb8\xda\xb4\x52\x95\xf0\xbf\xbe\x20\x4d\x1b\x5e\xec\xb8\x40\x78\x0d\xc8\xba\x31\xd6\x03\x49\x93\xac\x30\xda\xe3\x4f\x9f\xa5\x49\x86\xba\x30\xa5\xd4\x22\xff\xe1\x8c\x0e\x01\x65\x44\xb8\x58\xac\x14\x16\x3e\x4b\xd3\x24\x3b\x07\x08\xb5\x26\x14\x08\xe9\xb7\xed\x86\x15\xa6\xce\x3f\x7c\x28\xd1\x49\xa1\x5d\x2e\x9e\x94\x40\x3d\x50\x8e\xca\x5c\x53\x4d\xff\xcc\x0b\xb3\xb1\x3c\x4b\x69\x1a\xa8\xef\x4c\xd9\x2a\x04\xe9\x80\x6b\xe0\xad\x37\x27\xe3\x51\x68\x28\xf0\x06\x36\x52\x97\xe0\xb7\x08\xd6\x18\x0f\x16\x9d\x51\x7b\xb4\xa9\x7f\x69\x70\xa8\x72\xde\xb6\x85\x3f\x74\x71\xf4\x8d\xd1\x95\x14\xad\x45\x70\xe8\x5d\xec\xec\xb1\xd8\xed\x4f\x2c\x5a\xcf\x37\x0a\xbf\x15\x5b\xac\x79\x9c\x2d\xb5\x80\xbd\xe4\xc0\xa1\xb1\x66\x2f\x4b\xb4\x13\x68\xb8\x73\x21\x2e\x75\xec\x2f\x8c\xb5\x58\xf8\x11\x41\xd5\xea\x02\xc8\xe5\x91\x81\x9e\x74\x89\xd4\x3f\xb0\xf0\xc6\xc2\x65\xf4\xc2\xe6\xfd\x7f\x0a\x87\x34\x19\x92\xec\x93\xd4\x25\xd1\xf8\x4c\x7e\x47\x47\x29\x5b\x9a\x2f\x3d\x13\x09\x6a\x24\x02\x5c\x86\xdf\xaf\x3d\x05\xfd\xbd\xb9\x43\x9a\x24\x16\x7d\x6b\x35\x2c\xf0\x79\x9c\x26\x47\xde\xc3\x30\xc8\xcd\xa2\xbd\x8e\xa6\x49\x47\xd3\x11\xe7\x5d\xab\xbc\x8c\xb0\xf1\x29\xb2\x1b\x53\xd7\x5c\x97\xff\x92\xf0\xf2\x4d\xd3\x39\xd7\x1f\x6f\x32\x21\x91\x3c\x38\x9c\x41\x26\x9e\x14\x14\x5c\x29\x07\x1c\xc2\x7d\x98\xed\xa5\xd1\xa0\xe4\x0e\xe1\x42\x18\xb0\xad\x86\x9a\x4b\xcd\x84\x89\x15\x0f\x0e\x2d\x5b\x9a\xd2\x38\xb8\x38\x3c\x66\x0b\x5e\xe3\x63\x36\x7b\xcc\x42\x1c\xf8\x63\xd6\x5d\x40\x36\x09\x02\x5f\x5b\x3d\x8b\xf3\x48\x51\x97\x23\xb6\x09\x70\x2b\x1c\xac\xd6\xce\x5b\xa9\x45\x7c\x60\x49\x92\xc8\x0a\x14\x6a\x12\x72\x14\xae\xaf\xe1\x7d\x1f\x4f\x2a\x63\x61\x07\xb3\x6b\xb0\x5c\x0b\x84\x52\x86\x63\x12\x77\x30\xb4\x26\x89\x32\x82\x7d\xb1\x52\x7b\xa5\xc9\x8e\x1e\x83\xdd\xf1\x72\xdc\x42\xbc\xef\x00\x95\x43\x18\x2b\x4d\x87\x31\xfb\xa8\x72\x7c\x37\xd9\xdf\x5c\xb5\x78\x5f\x91\x73\xbd\x55\xe8\x59\xbd\x5f\xaf\x7b\x89\x57\xd9\x8a\x64\xff\x77\xd9\x04\xf6\x6c\xf9\xd2\x20\xa1\xec\x5b\xf4\x46\x28\xfd\x15\x22\x4d\xff\x8b\x96\xd4\xa1\xa1\xe6\x3b\x24\xab\xf5\x9b\xbe\x33\xd5\x45\x5b\xcf\x35\xa1\xef\xa6\xc7\x9e\xb0\x39\x19\xda\xa6\x1f\x41\xc2\x5f\xbf\xd4\x7d\x04\x79\x75\x35\x98\x97\x7a\x25\xdf\x4d\xd7\x70\xa2\x5a\xe0\x33\x79\x6d\x99\x6b\x22\xe9\xd8\xb6\xd2\x24\x7c\xd9\xd8\x83\xae\xb9\x75\x5b\xae\xc8\x6a\xbd\x79\xf1\x18\x77\xbb\x92\x6b\x3a\x81\x7e\x2c\x9b\x6b\x8f\xb6\xe2\x05\x12\x3a\x8c\x39\x29\x0e\x45\xb7\x0a\x6b\x42\xd3\xd3\xd3\xb3\xe8\x02\xff\x9e\xdd\x70\xa5\x08\x6f\x1a\xd4\xe5\xd8\xff\x61\xbc\xc5\xfe\x03\xcc\x3e\xf1\x62\x27\xac\x69\x75\x49\x28\xed\x02\x0a\x63\x8c\x9e\x56\xf3\x7d\x02\xf6\x74\xae\x82\xd4\x70\x0e\x26\xf0\x3d\x24\xa2\xb7\xbb\xde\x99\x7d\x63\x61\xb4\x08\x72\x3c\xcb\x64\x4f\xcf\xe8\xbb\xf0\x2a\x74\xf1\x55\xef\xd2\x7f\x02\x00\x00\xff\xff\x73\xfe\xc7\x98\x7d\x06\x00\x00")

func moduleGoTplBytes() ([]byte, error) {
	return bindataRead(
		_moduleGoTpl,
		"module.go.tpl",
	)
}

func moduleGoTpl() (*asset, error) {
	bytes, err := moduleGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "module.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\x54\xa8\x56\x48\xcb\x49\xcc\xcd\xcc\x4b\xcf\xb7\x52\x08\x2e\x29\xca\xcc\x4b\x57\xa8\xe5\x02\xcb\xfb\x96\x96\x24\x96\x64\xe6\xe7\x61\x57\x52\x9c\x9c\x98\x93\x58\xa4\x10\x92\x99\x9b\x0a\x63\xfb\x26\x16\x00\x02\x00\x00\xff\xff\x60\x5b\x70\x13\x59\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"emptymodule.go.tpl": emptymoduleGoTpl,
	"module.go.tpl":      moduleGoTpl,
	"schema.graphql":     schemaGraphql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"emptymodule.go.tpl": &bintree{emptymoduleGoTpl, map[string]*bintree{}},
	"module.go.tpl":      &bintree{moduleGoTpl, map[string]*bintree{}},
	"schema.graphql":     &bintree{schemaGraphql, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
