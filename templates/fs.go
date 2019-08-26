// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/emptymodule.go.tpl
// templates/module.go.tpl
// templates/resolver.go.tpl
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

// Mode return file modify time
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

	info := bindataFileInfo{name: "emptymodule.go.tpl", size: 339, mode: os.FileMode(420), modTime: time.Unix(1566801235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _moduleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcd\x8e\xe2\x30\x10\x84\xcf\xf1\x53\xf4\xe6\x64\xb3\x28\x39\xb3\xc7\x05\x0e\x1c\x16\x56\xbb\xbc\x80\x63\x37\x8e\x67\x12\x3b\xf8\x07\x06\x45\x79\xf7\x51\x4c\x32\x8c\x90\xb8\x44\x71\x77\x75\xd5\x57\x65\x09\x6b\x2b\x11\x14\x1a\x74\x3c\xa0\x84\xea\x06\xa7\x86\xb7\xda\x28\x5b\xb4\x58\x2a\xc7\xbb\xfa\xdc\x2c\x61\x73\x80\xfd\xe1\x08\xdb\xcd\xee\x58\x90\xb2\x84\x9f\x55\xd4\x8d\x84\x1f\x93\x80\x90\x8e\x8b\x77\xae\x10\xbe\x06\xba\xed\xac\x0b\x40\x49\x96\x7f\x77\x94\xe3\x4f\x4e\xb2\x5c\xe9\x50\xc7\xaa\x10\xb6\x2d\x57\x2b\x89\x5e\x2b\xe3\x4b\x75\x6e\x14\x9a\x39\x36\x27\x8c\x8c\x61\x7f\xac\x8c\x0d\x82\xf6\xc0\x0d\xf0\x18\xec\x83\x37\xd9\xcd\x82\x60\xa1\xd2\x46\x42\xa8\x11\x9c\xb5\x01\x1c\x7a\xdb\x5c\xd0\x91\x70\xeb\x70\x56\xf9\xe0\xa2\x08\xfd\x90\xac\xd7\xd6\x9c\xb4\x8a\x0e\xc1\x63\xf0\xe9\x72\x0a\x2f\xb6\x1f\x28\x62\xe0\x55\x83\xff\x45\x8d\x2d\x4f\xde\xda\x28\xb8\x68\x0e\x1c\x3a\x67\x2f\x5a\xa2\x5b\x42\xc7\xbd\x1f\xe7\xda\xa4\x7b\x61\x9d\x43\x11\x9e\x08\x4e\xd1\x08\xa0\x8b\x3b\x03\x7b\xe4\x52\x6d\xde\x50\x04\xeb\x60\x91\xba\x14\xbb\xe9\xcd\xa0\x27\xd9\xbc\x2c\x7e\x6b\x23\xa9\xc1\x2b\x7d\x45\xc7\x58\x71\xb4\x7f\x27\x26\x3a\xa6\xd1\x04\xb0\x18\xbf\xff\x26\x0a\xf6\xba\x5c\x4f\xb2\xcc\x61\x88\xce\xc0\x1e\xaf\xcf\x6b\x7a\xe7\xed\x67\x23\xff\x2b\xd5\x1b\x18\xc9\x06\x46\x06\xf2\x19\x00\x00\xff\xff\xe1\xff\x42\xd5\x49\x02\x00\x00")

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

	info := bindataFileInfo{name: "module.go.tpl", size: 585, mode: os.FileMode(420), modTime: time.Unix(1566801320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resolverGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x52\xc1\xae\x9b\x30\x10\x3c\x3f\x7f\xc5\xbc\x77\xe2\xd1\x0a\xbe\x80\x6b\xa5\x1e\x7a\x48\x7e\xa0\x72\xcc\x42\xdc\x82\xed\xac\x97\x44\x51\xd5\x7f\xaf\x40\x0e\x29\x81\x36\x97\xaa\x27\xd0\xee\xec\xec\x8c\x77\xca\x12\x1f\x0e\x83\xed\x6a\xbc\xb6\xac\xc3\xf1\xd4\x29\x15\xb4\xf9\xae\x5b\xc2\x5c\xb0\x7d\xf0\x2c\x78\x6b\x3a\xdd\x5b\xd7\xfa\xa2\xa7\x32\x35\xdf\x94\x2a\x4b\xc8\xd1\x46\x34\xb6\x23\xd8\x08\x8d\x28\x9a\xc5\xba\x16\xc1\x5b\x27\x68\x3c\x43\x8e\x84\xc0\xfe\x1b\x19\x41\x0c\x64\x6c\x63\x0d\x98\xa2\xef\xce\xc4\x71\xe4\xb0\x82\x8b\xed\x3a\x38\x2f\x38\x10\x98\x5a\x72\xc4\x5a\xa8\x7e\x55\x4a\xae\x81\xc0\xde\xcb\x3e\x8d\x20\x0a\x0f\x46\xf0\x43\xbd\x9c\x06\xe2\xeb\x5c\x07\x90\x2f\x2a\xea\xa5\x1f\x44\x8b\xf5\x6e\xc6\xe4\x8f\x15\xf5\x73\xb2\x61\x9d\x10\x37\xda\x10\xda\x41\x73\xad\xce\x9a\xf1\x15\x37\xd0\xde\x7b\x41\x05\x47\x97\xec\x77\x25\xef\xd3\xe8\x67\x37\x59\x1b\x1b\xb3\x2d\xd5\x0c\xce\x20\x63\xe4\x0b\x7c\xc2\x66\x4b\xdd\x4b\xd1\x1f\xf1\x5c\xf4\xfb\x68\x9e\x8b\x25\x4d\x85\x07\xf3\x5c\xac\x98\xaa\x15\x79\xf2\xbf\x1b\x47\xd1\x92\xc8\x9f\xc5\x4f\x98\x2c\x7d\x67\xca\x51\x09\xc9\xc0\x0e\x0f\x82\x12\xf3\x97\xb4\xf0\x09\xf9\x0d\x96\xdd\x7f\x37\x57\x6c\xe9\x9f\x22\xb2\x7c\x8c\x7b\x46\xf2\x14\xd7\xe2\x53\x8a\xf0\x6e\x43\x64\xba\x61\x4d\x81\x5c\x4d\xce\x58\x8a\x77\xa1\x0b\xe6\xf9\x86\xcd\x16\x1d\xfe\xbe\x2d\xdd\x6d\xb3\x87\x0a\x9b\x94\xb3\xc1\xd5\x31\xff\xa5\xc7\x75\xc2\xfe\xbf\xcd\x5f\x01\x00\x00\xff\xff\x0b\x4c\xe9\xbb\x91\x04\x00\x00")

func resolverGoTplBytes() ([]byte, error) {
	return bindataRead(
		_resolverGoTpl,
		"resolver.go.tpl",
	)
}

func resolverGoTpl() (*asset, error) {
	bytes, err := resolverGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resolver.go.tpl", size: 1169, mode: os.FileMode(420), modTime: time.Unix(1566801566, 0)}
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

	info := bindataFileInfo{name: "schema.graphql", size: 89, mode: os.FileMode(420), modTime: time.Unix(1563362778, 0)}
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
	"emptymodule.go.tpl": emptymoduleGoTpl,
	"module.go.tpl":      moduleGoTpl,
	"resolver.go.tpl":    resolverGoTpl,
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
	"emptymodule.go.tpl": &bintree{emptymoduleGoTpl, map[string]*bintree{}},
	"module.go.tpl":      &bintree{moduleGoTpl, map[string]*bintree{}},
	"resolver.go.tpl":    &bintree{resolverGoTpl, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
