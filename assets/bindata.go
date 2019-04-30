// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/sql/001_create_users_table.sql
// assets/sql/002_create_roles_table.sql
// assets/sql/003_create_user_role_table.sql

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
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x41\x4f\x83\x30\x18\x86\xef\xfd\x15\xef\x6d\x1a\xe5\x17\xec\x84" +
	"\xae\x26\x8d\xc0\x18\x94\x64\x3b\x2d\x95\x7e\xba\x46\x28\x4d\x5b\x82\x3f\xdf\xe0\x22\x4e\xa3\x89\xd7\xf7\x79\x9f" +
	"\xcb\x93\x24\xb8\xe9\xcd\x8b\x57\x91\xd0\x38\x96\x24\xa8\x77\x19\x8c\x45\xa0\x36\x9a\xc1\x62\xd5\xb8\x15\x4c\x00" +
	"\xbd\x51\x3b\x46\xd2\x98\x4e\x64\x11\x4f\x26\xe0\xec\xcd\x27\x13\xa0\x9c\xeb\x0c\x69\x76\x5f\xf1\x54\x72\xc8\xf4" +
	"\x2e\xe3\x10\x0f\x28\xb6\x12\x7c\x2f\x6a\x59\x63\x0c\xe4\x03\xae\x18\x60\x34\x6a\x5e\x89\x34\xfb\xc0\x45\x93\x65" +
	"\x28\x2b\x91\xa7\xd5\x01\x8f\xfc\x70\xcb\x80\x67\xe3\x43\x3c\x5a\xd5\x13\x24\xdf\xcb\xe5\x37\xb3\x4e\xfd\x89\xa8" +
	"\x57\xa6\x3b\xcf\x4d\x21\x76\x0d\xff\x46\x9d\x0a\x61\x1a\xbc\xfe\xe1\x01\x00\x03\x5a\x4f\x2a\x92\x3e\xaa\x08\x29" +
	"\x72\x5e\xcb\x34\x2f\x17\x75\x74\xfa\x77\xc8\xae\xd7\x8c\x5d\x56\xdc\x0c\x93\xfd\xec\xb8\x44\x9c\xc7\x7f\x65\xf4" +
	"\x43\xd7\x91\xc6\x93\x6a\x5f\xd9\xa6\xda\x96\x5f\x21\x2f\x23\xae\xdf\x03\x00\x00\xff\xff\x35\x0a\xe2\x4f\xb7\x01" +
	"\x00\x00")

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
		size: 439,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1556598807, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataAssetsSql002createrolestablesql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\xcd\x4e\x02\x31\x14\xc5\xf1\xfd\x7d\x8a\xb3\x43\x23\xf3\x04\xac" +
	"\x46\xa9\x49\xe3\x30\x0c\xfd\x48\x60\x65\xea\xf4\x46\x1a\x87\x4e\x33\x2d\xc1\xc7\x37\x48\x44\x16\x2e\xdc\xde\xff" +
	"\xb9\x8b\x5f\x55\xe1\xe1\x10\xde\x27\x57\x18\x36\x51\x55\x41\x6f\x1a\x84\x88\xcc\x7d\x09\x63\xc4\xcc\xa6\x19\x42" +
	"\x06\x7f\x72\x7f\x2c\xec\x71\xda\x73\x44\xd9\x87\x8c\xcb\xdf\x79\x14\x32\x5c\x4a\x43\x60\x4f\x4f\x4a\xd4\x46\xc0" +
	"\xd4\x8f\x8d\x80\x7c\x46\xbb\x36\x10\x5b\xa9\x8d\xc6\x34\x0e\x9c\x71\x47\x40\xf0\xd0\x42\xc9\xba\xf9\xce\xad\x6d" +
	"\x1a\x74\x4a\xae\x6a\xb5\xc3\x8b\xd8\xcd\x09\x88\xee\xc0\x30\x62\x6b\x60\x5b\xb9\xb1\xe2\x3a\x9c\x03\x04\xf4\x13" +
	"\xbb\xc2\xfe\xd5\x15\x18\xb9\x12\xda\xd4\xab\xee\x92\x09\x38\x26\xff\x77\xa4\xfb\x05\x40\x74\x6b\x5e\x8e\xa7\xf8" +
	"\xa3\xbe\x92\xcf\xc7\x7f\xa1\xa7\x71\x18\xd8\xe3\xcd\xf5\x1f\xb4\x54\xeb\xee\x97\x7d\x4b\x5e\xd0\x57\x00\x00\x00" +
	"\xff\xff\xfe\x7f\x34\x9c\x66\x01\x00\x00")

func bindataAssetsSql002createrolestablesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsSql002createrolestablesql,
		"assets/sql/002_create_roles_table.sql",
	)
}



func bindataAssetsSql002createrolestablesql() (*asset, error) {
	bytes, err := bindataAssetsSql002createrolestablesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "assets/sql/002_create_roles_table.sql",
		size: 358,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1556081073, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataAssetsSql003createuserroletablesql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x8e\x9b\x30\x14\x45\xf7\xfe\x8a\xbb\x9b\x44\x1d\xbe\x60\x56" +
	"\x14\x5e\x66\xac\x82\x4d\x8d\x51\x3b\xdd\x20\x8a\xdd\x89\x15\x0a\x08\x88\xd2\xfe\x7d\x65\x02\x4d\x54\x25\x52\xa5" +
	"\x61\xc7\xbb\xe7\x3d\x1d\xd0\x0d\x02\x7c\xf8\xe9\xde\x86\x6a\xb2\x28\x7a\x16\x04\xc8\x3f\x27\x70\x2d\x46\x5b\x4f" +
	"\xae\x6b\xf1\x50\xf4\x0f\x70\x23\xec\x2f\x5b\x1f\x27\x6b\x70\xda\xdb\x16\xd3\xde\x8d\x38\xef\x79\xc8\x8d\xa8\xfa" +
	"\xbe\x71\xd6\xb0\x48\x51\xa8\x09\x3a\xfc\x98\x10\xf8\x0e\x42\x6a\xd0\x57\x9e\xeb\x1c\xc7\xd1\x0e\xe5\xd0\x35\x16" +
	"\x1b\x80\xe1\xfc\xee\x0c\xb8\xd0\xf4\x4c\x6a\x46\x45\x91\x24\x8f\x0c\xf0\xd8\xbd\xec\x6d\xa8\xda\xa9\x34\x5e\x59" +
	"\xf3\x94\x72\x1d\xa6\x19\xbe\x70\xfd\x22\x0b\x3d\x4f\xf0\x4d\x0a\xf2\x64\x3d\xd8\x6a\xb2\xa6\xac\xa6\x2b\x72\x3d" +
	"\x73\xec\xcd\xfd\x30\x53\x3c\x0d\xd5\x2b\x3e\xd1\x2b\x36\x8b\xe8\xe3\x6a\xb5\xf5\x44\x24\x45\xae\x55\xc8\x85\xbe" +
	"\x7c\x58\xb9\x00\xe5\x8f\x83\xfd\x8d\x9d\x54\xc4\x9f\xc5\xf9\xc6\xba\xca\x30\x3f\x8a\x76\xa4\x48\x44\x94\xcf\x47" +
	"\x47\x6c\x9c\xd9\x22\x0d\x75\xf4\x82\x9c\xa7\x59\x42\x0b\x28\x05\x8a\x2c\xf6\xbf\x54\x48\x84\x91\xe6\x52\xf8\x59" +
	"\x4c\x09\x5d\xcf\xee\x1a\x2d\xee\x37\x8c\x96\xe4\x86\x91\x4f\xde\x6d\xc4\xb6\x4f\x8c\x5d\xd7\x2b\xee\x4e\xed\x5a" +
	"\xb0\xbf\xed\xf2\xc3\xff\xea\xd7\xd0\x35\x8d\x35\xf8\x5e\xd5\x07\x16\x2b\x99\x5d\x1a\xf6\x6f\xbb\x9e\xd8\x9f\x00" +
	"\x00\x00\xff\xff\x83\x21\xd5\x73\xd5\x02\x00\x00")

func bindataAssetsSql003createuserroletablesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsSql003createuserroletablesql,
		"assets/sql/003_create_user_role_table.sql",
	)
}



func bindataAssetsSql003createuserroletablesql() (*asset, error) {
	bytes, err := bindataAssetsSql003createuserroletablesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "assets/sql/003_create_user_role_table.sql",
		size: 725,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1556099746, 0),
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
	"assets/sql/001_create_users_table.sql":     bindataAssetsSql001createuserstablesql,
	"assets/sql/002_create_roles_table.sql":     bindataAssetsSql002createrolestablesql,
	"assets/sql/003_create_user_role_table.sql": bindataAssetsSql003createuserroletablesql,
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
			"002_create_roles_table.sql": {Func: bindataAssetsSql002createrolestablesql, Children: map[string]*bintree{}},
			"003_create_user_role_table.sql": {Func: bindataAssetsSql003createuserroletablesql, Children: map[string]*bintree{}},
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
