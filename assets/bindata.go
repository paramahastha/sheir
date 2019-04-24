// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/sql/001_create_users_table.sql

package assets


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataAssetsSql001createuserstablesql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x4d\x4e\xeb\x30\x14\xc5\xf1\xb9\x57\x71\x94\x49\xdf\x13\x64\x05" +
	"\x1d\x05\x6a\xa4\x88\xf4\x83\xc4\x95\xda\x51\x65\xe2\x5b\x6a\xe1\xd8\x96\xed\x28\x2c\x1f\xa5\x88\xb6\x83\x80\x98" +
	"\xfa\x77\xac\x2b\xfd\xf3\x1c\x77\x9d\x7e\x0b\x32\x11\xb6\x9e\xe5\x39\x9a\x97\x0a\xda\x22\x52\x9b\xb4\xb3\x98\x6d" +
	"\xfd\x0c\x3a\x82\x3e\xa8\xed\x13\x29\x0c\x27\xb2\x48\x27\x1d\xf1\xf5\x6f\x1c\xe9\x08\xe9\xbd\xd1\xa4\xd8\x63\xcd" +
	"\x0b\xc1\x21\x8a\x87\x8a\xa3\x7c\xc2\x6a\x2d\xc0\x77\x65\x23\x1a\xf4\x91\x42\xc4\x3f\x06\x64\x5a\x65\x68\x78\x5d" +
	"\x16\xd5\x79\xb0\xda\x56\x15\x36\x75\xb9\x2c\xea\x3d\x9e\xf9\xfe\x7e\xdc\x1c\x75\x88\xe9\x60\x65\x47\x19\x04\xdf" +
	"\x89\xcb\xf2\xac\x46\xfe\x82\xd4\x49\x6d\xa6\xc0\xcb\x18\x07\x17\xd4\x94\xb5\xce\x1e\x75\xe8\xa6\x28\x38\x33\x79" +
	"\xa7\x0d\x24\x13\xa9\x83\x4c\x19\x44\xb9\xe4\x8d\x28\x96\x9b\x2b\xf7\x5e\xfd\xc0\xec\xff\x9c\xb1\xdb\xf6\x0b\x37" +
	"\xd8\xef\xfa\x97\xf4\xe3\xe3\x9f\xe2\x07\x67\x0c\x29\xbc\xca\xf6\x9d\x2d\xea\xf5\xe6\x9a\xff\x36\xfd\x9c\x7d\x06" +
	"\x00\x00\xff\xff\x9d\xfa\xc2\xb1\xee\x01\x00\x00")

func bindataAssetsSql001createuserstablesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsSql001createuserstablesql,
		"assets/sql/001_create_users_table.sql",
	)
}



func bindataAssetsSql001createuserstablesql() (*asset, error) {
	bytes, err := bindataAssetsSql001createuserstablesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "assets/sql/001_create_users_table.sql",
		size: 494,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1555917964, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"assets/sql/001_create_users_table.sql": bindataAssetsSql001createuserstablesql,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"assets": {Func: nil, Children: map[string]*bintree{
		"sql": {Func: nil, Children: map[string]*bintree{
			"001_create_users_table.sql": {Func: bindataAssetsSql001createuserstablesql, Children: map[string]*bintree{}},
		}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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