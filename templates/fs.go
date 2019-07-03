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

var _emptymoduleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x6e\xc2\x30\x10\x44\xef\xfb\x15\x53\x4e\xd0\x4a\xf8\x23\xa0\x07\x0e\x2d\x17\x7e\xc0\xb1\x37\xc6\x6d\xec\x75\x9d\x75\xa5\x08\xf1\xef\x15\x24\xa5\xbd\xed\xcc\x6a\xde\x8c\x31\x3b\xf1\x8c\xc0\x99\xab\x55\xf6\xe8\x26\xf4\x83\x4d\x31\x07\xd9\x26\x36\xa1\xda\x72\xfe\x1a\xb0\x3f\xe2\xfd\x78\xc2\xeb\xfe\x70\x7a\x22\x63\x5e\xba\x16\x07\x8f\xe5\x4b\x54\xac\xfb\xb4\x81\xff\x8c\x98\x8a\x54\xc5\xea\x3f\xcb\xdf\x8e\x15\x91\x31\xd0\x73\x1c\xc1\xa9\xe8\x84\x24\xbe\x0d\x8c\x38\xa2\x8d\x73\x7f\x90\xc7\x1e\xa8\xa0\x54\xf9\x8e\x9e\x61\x71\x07\xfc\x06\x3a\x76\xb6\x8d\xbc\xc8\x6d\x90\x1b\x22\x8b\xc2\x49\x2a\x71\x60\x0f\xdf\x6a\xcc\xe1\xc1\x22\x9d\x0a\xe3\x6d\x4e\x8f\x5a\x9b\xd3\xcb\x95\xa8\x6f\xd9\x61\x3d\xdb\x1b\xec\x24\xf7\x31\xb4\xca\xeb\x98\x3f\xd8\xa9\x54\x3c\xdf\x6b\xb7\x87\x45\x6f\x70\x21\xba\xd2\x4f\x00\x00\x00\xff\xff\x68\x0a\x88\x30\x39\x01\x00\x00")

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

	info := bindataFileInfo{name: "emptymodule.go.tpl", size: 313, mode: os.FileMode(438), modTime: time.Unix(1562163068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _moduleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xc1\x4e\x02\x31\x10\x86\xcf\x3b\x4f\x31\xee\xa9\x45\xd3\x9e\xf1\x28\x70\xe0\x20\x18\xe5\x05\xba\xed\x50\xaa\xdd\x76\xe9\xb6\x20\x21\xfb\xee\x66\x05\x8c\x31\xf1\xd2\x74\xfe\x3f\xf3\x7d\x23\x25\xce\xa2\x21\xb4\x14\x28\xa9\x4c\x06\x9b\x13\x6e\xbd\x6a\x5d\xb0\x51\xb4\x24\x6d\x52\xdd\x6e\xef\x1f\x70\xbe\xc6\xd5\x7a\x83\x8b\xf9\x72\x23\x40\xca\xfb\xa6\x38\x6f\xf0\xee\xda\x03\x74\x4a\x7f\x28\x4b\xf8\x13\xb8\xb6\x8b\x29\x23\x83\xaa\xfe\x0d\x34\xe3\xa7\x86\xaa\xb6\x2e\xef\x4a\x23\x74\x6c\xe5\x74\x6a\xa8\x77\x36\xf4\xd2\xee\xbd\xa5\x70\xb3\xd6\xc0\x01\xf2\xa9\x23\x7c\x8e\xa6\x78\xc2\x3e\xa7\xa2\xf3\x79\x00\xd8\x96\xa0\x91\x4d\x2e\x39\xc7\x59\x0c\x5b\x67\x4b\x22\xe6\xc2\x3b\xe9\x1c\x13\x4e\xbe\x4d\x62\x79\x9d\x39\x9e\xa1\xba\x95\xe2\xc9\x05\xc3\x02\x1d\xd9\xd5\x24\x16\x9f\xa4\x4b\x56\x8d\xa7\x37\xbd\xa3\x56\x71\x2e\x36\xf1\x25\xc5\x83\x33\x94\xd8\x68\x63\x29\xc6\x8c\x93\xf1\x7d\xa5\x3e\xfa\x03\x25\x8e\xff\xad\x8f\xb2\x2a\x51\x2e\x29\xe0\x8a\x8e\x7f\x6b\x76\xb9\xf7\x7c\x03\xf5\x8f\x38\x72\x07\x0e\xd5\xc0\x61\x80\xaf\x00\x00\x00\xff\xff\x2e\xa0\x1c\x6c\x97\x01\x00\x00")

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

	info := bindataFileInfo{name: "module.go.tpl", size: 407, mode: os.FileMode(438), modTime: time.Unix(1562071640, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resolverGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4e\x03\x31\x0c\x45\xf7\x3e\xc5\x67\x37\xb0\xe9\x09\x38\x02\x0b\x7a\x01\x14\x82\x67\xc6\x10\xec\xd4\x71\x5a\x55\x55\xef\x8e\xa8\xda\x42\x2b\xcd\xca\x96\xbf\xde\xb3\x7e\x4d\xf9\x2b\x4d\x8c\xc9\x53\x9d\x37\x85\x68\xb5\x42\xcc\xd2\x30\x4a\x61\x48\x43\x42\x8b\xe4\x21\x3a\xa1\x9a\x68\x60\x34\x47\xcc\x8c\xea\xf6\xc9\x39\xd0\x2a\x67\x19\x25\xc3\xb9\x59\xd9\xb2\xb7\x5f\x87\x04\x76\x52\x0a\xd4\x02\xef\x0c\xe7\x89\x95\x3d\x05\x7f\x3c\x10\xc5\xbe\x32\xdc\x2c\xd6\x67\x04\x68\xe1\x3d\xc7\xe1\x48\xb4\x4d\x8e\x37\x5c\x92\xb5\x59\xe0\x19\xca\xbb\xe1\x3f\xf0\x48\x34\x76\xcd\x18\x9e\x6e\xae\x78\xed\xec\xfb\xe1\x3c\xaf\xf6\x03\x39\x47\x77\x85\x4a\xa1\xe3\x02\xfa\xd2\x23\x85\x98\x0e\x7f\xeb\xa2\xe0\x54\x60\x73\xf3\xe3\x5a\xe0\x94\x7d\xdf\x1b\x2e\xf1\x4f\x00\x00\x00\xff\xff\xd6\xae\x57\x56\x70\x01\x00\x00")

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

	info := bindataFileInfo{name: "resolver.go.tpl", size: 368, mode: os.FileMode(438), modTime: time.Unix(1562071634, 0)}
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

	info := bindataFileInfo{name: "schema.graphql", size: 89, mode: os.FileMode(438), modTime: time.Unix(1562071634, 0)}
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
