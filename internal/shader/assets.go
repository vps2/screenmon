// Code generated for package shader by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../assets/shader.vert
// ../../assets/shader.frag
package shader

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

var _shaderVert = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x41\x4e\xc2\x40\x14\x86\xf7\x93\xcc\x1d\xfe\xc4\x0d\xa0\xb4\x45\xdc\x11\xee\xe0\x0d\x4c\x03\x05\x26\x29\x1d\x52\x4a\x15\x8d\x09\xe2\xc6\xa4\x1e\xa6\x21\x12\x1b\x4a\x7b\x86\x7f\x6e\x64\xda\x22\x0b\x63\x8c\xb3\x7b\xef\x9f\xef\xbd\xf7\x5d\xc4\x5e\xb8\x54\x3a\x40\xbf\xef\x60\xa4\x43\x0f\x90\x02\x52\xf8\xee\x5a\xaf\x22\xb4\x7c\x3d\x72\xa3\x2a\x1f\xc2\x69\x43\x05\x88\xbd\x51\x1f\x0b\xbd\x54\x55\x77\x20\x85\x14\xb6\x6d\xde\x98\x99\x0d\x33\x16\x4c\xc1\x0c\xdc\x99\xc4\xbc\xb0\x34\xdb\xaa\x2e\x79\xa8\x02\x29\x56\x81\x9a\xe8\x70\x8e\x89\xaf\xdd\x08\xf7\x6a\x1c\xcd\x06\x3f\xbb\x33\x4f\x4d\x67\x51\x3d\x37\xd6\x6a\x8c\xb9\xab\x82\x56\x1b\x4f\x52\x00\x80\x6d\xb3\x60\x69\x36\x3c\x32\x65\xce\x8c\x9f\xe6\x95\x7b\x1e\xc1\x03\xcb\x3a\xf8\x68\xae\x30\x5b\x93\x9c\x37\x83\x3b\x30\xaf\xab\x94\xb9\x79\x67\x61\x12\xee\xa1\x17\x5e\x30\xf5\xd1\x4c\x9e\xfa\x77\xb7\x27\x29\x0c\x2b\xc9\x9b\xd6\xb7\xa4\xf5\x80\x0e\xae\x2d\x07\x76\x73\x33\xba\xe8\x59\xce\xd5\x09\xfc\xe5\x9d\xc1\x35\x3a\xe8\x36\x64\xe3\x85\xcb\xff\xa2\x8f\x7f\x7c\xea\x59\x4e\x7b\x20\xc5\xf3\x57\x00\x00\x00\xff\xff\xfb\x9c\x9c\x95\xbc\x01\x00\x00")

func shaderVertBytes() ([]byte, error) {
	return bindataRead(
		_shaderVert,
		"shader.vert",
	)
}

func shaderVert() (*asset, error) {
	bytes, err := shaderVertBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shader.vert", size: 444, mode: os.FileMode(438), modTime: time.Unix(1603110870, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _shaderFrag = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x50\xcd\x6a\x83\x40\x10\xbe\x0b\xbe\xc3\x07\x05\xd9\xa5\x35\xa8\xe9\x4d\xeb\xa5\xd0\xd7\x28\x56\x57\xb3\x60\x1c\x98\x18\x59\x49\xf2\xee\x25\x6e\x88\x76\x45\xe8\x1e\xe6\xf0\xcd\xf7\xb7\xf3\x32\x28\x3e\x69\xea\xb0\xdf\x47\x28\x89\x95\xef\xf9\xde\xb9\xd3\x35\xf1\x11\x75\x4b\x45\x0f\x13\xa7\x2e\x34\xae\x21\x93\xac\x59\x6b\xe8\x87\xb8\x52\xbc\x82\x0f\x4a\x37\x87\x3e\xbd\x67\xd3\xb9\xc7\xa0\xca\x77\x94\xd4\x12\x4f\xd0\x40\xba\xc2\xb1\xd0\x9d\x90\xb8\xf8\x1e\x00\xe8\x5a\x08\xd1\xb4\xdf\x5f\x5c\x34\x9f\x44\x5c\xed\x0c\x72\x98\xf8\xd5\x06\x20\x08\xe0\x6c\x33\x98\x24\xb4\x5b\x89\xeb\x15\xd6\xe7\xfe\x5c\x9f\x0c\x26\x0e\x37\x7d\x72\x98\xe4\x91\x22\x25\x82\xe0\xe9\xe3\xf4\x19\x91\x3f\xbe\x15\x8e\x9b\xb5\x46\x64\x4f\xd2\xbf\xda\x2d\x05\x9b\x25\x97\xd1\x73\x57\x69\x3d\x2f\xb3\xf5\x74\x61\x7c\x4c\xe7\x16\xf1\x2e\xaa\xdf\xe0\x4e\x99\x5a\xfa\x0d\xaa\x3d\xa9\x4d\x71\x34\x09\x96\xf3\x8f\xd8\xf7\x6e\xbf\x01\x00\x00\xff\xff\x4e\xa5\x97\xca\x68\x02\x00\x00")

func shaderFragBytes() ([]byte, error) {
	return bindataRead(
		_shaderFrag,
		"shader.frag",
	)
}

func shaderFrag() (*asset, error) {
	bytes, err := shaderFragBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shader.frag", size: 616, mode: os.FileMode(438), modTime: time.Unix(1603110892, 0)}
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
	"shader.vert": shaderVert,
	"shader.frag": shaderFrag,
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
	"shader.frag": &bintree{shaderFrag, map[string]*bintree{}},
	"shader.vert": &bintree{shaderVert, map[string]*bintree{}},
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
