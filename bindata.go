// Code generated by go-bindata. DO NOT EDIT.
// sources:
// template/template.html (1.666kB)
// template/template1.html (1.598kB)
// template/template2.html (1.416kB)
// template/template3.html (1.446kB)

package pdf

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templateTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x49\x72\xeb\x36\x10\xdd\xe3\x14\xfd\xe9\x8d\x9d\x0a\x49\x91\x89\x33\x50\x14\x13\x95\x6d\x55\x76\xf1\x42\x59\x64\x09\x12\x10\x09\x9b\x00\x18\x00\xd4\x60\x97\xee\x91\x5d\xce\x91\x5b\xe5\x08\x29\x80\x83\x26\xab\xea\x2f\x6c\x5a\x25\xa1\x5f\xf7\xeb\x7e\x0f\x80\x94\x7e\x79\xfc\xfd\x61\xf9\xe7\xf3\x13\x54\x86\xd7\x19\x4a\x39\x35\x18\x2a\x63\x1a\x9f\xfe\xd5\xb2\xf5\xcc\x7b\x90\xc2\x50\x61\xfc\xe5\xae\xa1\x1e\x14\xdd\x6a\xe6\x19\xba\x35\xa1\xad\x99\x42\x51\x61\xa5\xa9\x99\xfd\xb1\x5c\xf8\x3f\x79\x19\x4a\xb5\xd9\xd5\x34\x43\xb9\x24\x3b\x78\x47\x00\x1c\xab\x92\x89\x04\x26\x53\x04\xd0\x60\x42\x98\x28\xfb\x55\x8e\x8b\xd7\x52\xc9\x56\x10\xbf\x90\xb5\x54\x09\xdc\x2c\xe6\xf6\xb1\xe0\x4a\x0a\x93\x40\x14\x37\x06\xbc\x25\xae\x24\xc7\xde\x14\xed\x11\xfa\xc6\xb1\xe6\x72\xeb\x6b\xf6\xe6\xb8\x72\xa9\x08\x55\x7e\x2e\xb7\xb6\xce\xe7\xf2\xcd\xbf\x06\xef\x11\x0a\x1a\x5c\x52\xc7\xb1\x61\xc4\x54\x09\xc4\x51\xc1\x6d\x21\x67\xc2\xaf\x28\x2b\x2b\x93\x40\xfc\x73\xf0\x63\x17\x1d\x27\x8e\xfb\xac\x5e\x4e\x54\x70\xc0\xad\x91\x4e\x87\xeb\x90\x40\xd4\x6c\xe1\xe6\xf1\x3b\xfb\x80\x96\x35\x23\x07\xd0\x57\x98\xb0\x56\x27\x70\xdf\x6c\x4f\xa5\x27\xb0\xa9\x98\xa1\xd3\x41\x54\x85\x89\xdc\x24\x30\x81\x89\xcd\x05\x55\xe6\xf8\x76\xf2\x2d\xf4\xff\x41\x74\xd7\xa9\xd0\x6d\x3e\x0a\x19\x67\xec\x95\x0c\xf3\xb8\x7a\x4a\x0e\xb3\x8c\xf2\xee\x7f\xe0\x2e\x53\xb6\xa6\x66\x82\x3a\x75\x70\xb3\x58\x3c\xcd\x9f\xe6\x43\xfa\x1e\xa1\x5f\xc7\x16\x9a\xbd\xd1\x04\xe6\xdf\x4f\x4f\x77\xd4\xe6\x70\x4a\x18\x86\x46\x31\x61\x5c\xea\xc1\xe0\xb3\xcd\x3f\x0c\xc6\x04\x33\x0c\xd7\xc7\xc1\xd1\xa0\x13\xac\xdf\xa2\x93\xd8\xf1\x3e\x9d\x11\x1d\xec\x3b\x05\x8e\xcc\x3e\x01\xec\xa4\x7e\xae\x28\x7e\xf5\xf1\xca\xd8\xc9\x70\xbd\xc1\x3b\x6d\xd1\x3d\xda\x23\xc6\xcb\x80\x72\xf9\xc2\x3a\x39\x43\xd7\x88\x3a\xf3\x86\xe9\x86\xe5\x28\x16\x82\xc9\x3d\xe5\xf6\x7d\x80\xd6\x54\x19\x56\xe0\xda\xc7\x35\x2b\x45\x02\xfe\xa4\x83\xf6\x28\x0d\xfb\x0b\x93\xea\x42\xb1\xc6\x80\x56\xc5\xcc\xb3\x77\x50\x27\x61\x68\x36\xae\x7d\xc0\xf1\xb6\x20\x22\x28\x24\x0f\xe3\x43\x90\x89\xe0\x45\xff\x12\x45\x41\xec\x65\x69\xd8\xd5\x8f\x44\xd9\x86\x09\x22\x37\x81\x14\xb5\xc4\x04\x66\xb0\x6a\x45\x61\x98\x14\x70\x7b\x07\xef\x30\x90\x34\xf6\xfa\xde\x12\x59\xb4\x9c\x0a\x13\xd8\x6b\x7b\x37\x85\xfd\x11\x5d\xf7\xdd\x80\x52\x0b\x65\xc8\xf9\x96\x12\xb6\x86\xa2\xc6\x5a\xcf\xbc\x5c\xca\x57\x2f\x73\xe1\x73\xc8\xda\x7b\x04\x9d\xc3\xfd\x11\x3e\xcb\xb8\xe4\x27\xbb\x0f\x52\xec\xdf\xb2\x62\x1a\x98\x06\xbe\x83\x15\x53\xda\x40\x23\xb5\xf9\x02\xff\xfd\xf3\xf7\xbf\x1f\xbd\x2e\xdb\x84\x84\xad\x2f\xa9\x2f\x02\xcf\xf6\x40\x47\x61\x8c\xae\x14\xa3\x2b\x84\x9f\xe3\x45\x15\x65\xbf\x51\x45\xad\x50\x2c\xa4\xa9\xa8\x72\x3a\xd3\xb0\x8a\x3e\xc3\x39\x4d\x0b\x29\x48\x67\xdd\xd7\x3a\xe4\x0c\x89\xaf\x1b\x72\xb6\xec\x3f\x22\x84\xd2\xb0\x3f\x45\xa9\xfb\xf9\xc8\xfe\x0f\x00\x00\xff\xff\xf8\x06\x29\x70\x82\x06\x00\x00")

func templateTemplateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateTemplateHtml,
		"template/template.html",
	)
}

func templateTemplateHtml() (*asset, error) {
	bytes, err := templateTemplateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/template.html", size: 1666, mode: os.FileMode(0644), modTime: time.Unix(1643161198, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xec, 0x6a, 0xd8, 0xf, 0xa, 0x2a, 0x98, 0xb6, 0x85, 0xd8, 0x31, 0x3, 0x9a, 0x26, 0xf2, 0x22, 0xe4, 0x7d, 0x82, 0x7f, 0xfd, 0x43, 0x6b, 0x53, 0x90, 0x2a, 0xb2, 0xb, 0xdb, 0x8b, 0x5e, 0x49}}
	return a, nil
}

var _templateTemplate1Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4b\x72\xdb\x38\x10\xdd\xe3\x14\x6d\x79\x63\x4f\x0d\x49\x91\x33\x9e\xa9\x50\xb4\x12\x55\x6c\x57\x76\xf1\xc2\x9b\x2c\x41\x02\x22\xdb\x26\x00\x16\x00\x59\x92\x5d\xba\x47\x76\x39\x47\x6e\x95\x23\xa4\x00\x7e\xf4\xb3\xaa\xb2\xb0\x28\x95\x88\x7e\xdd\xaf\xfb\x3d\x00\xd9\xd9\xcd\xd7\xcf\x0f\xdf\xee\x6f\xa1\xb2\xa2\x9e\x92\xcc\xd8\x75\xcd\xa7\x24\x57\x6c\x0d\xaf\x04\x40\x50\x5d\xa2\x4c\x61\x3c\x21\x00\x0d\x65\x0c\x65\xd9\xad\x72\x5a\x3c\x95\x5a\x2d\x24\x0b\x0a\x55\x2b\x9d\xc2\xf9\xdd\xcc\x3d\x0e\x9c\x2b\x69\x53\x88\x93\xc6\xc2\xe8\x81\x56\x4a\xd0\xd1\x84\x6c\x08\xf9\xcb\xb3\xe6\x6a\x15\x18\x7c\xf1\x5c\xb9\xd2\x8c\xeb\x20\x57\x2b\x57\x17\x08\xf5\x12\x9c\x82\x37\x84\x84\x0d\x2d\xb9\xe7\x58\x22\xb3\x55\x0a\x49\x5c\x08\x57\x28\x50\x06\x15\xc7\xb2\xb2\x29\x24\x1f\xc2\xff\xdb\xe8\x30\x71\xd2\x65\x75\x72\xe2\x42\x00\x5d\x58\xe5\x75\xf8\x0e\x29\xc4\xcd\x0a\xce\x6f\xfe\x71\x0f\x18\x55\x23\xdb\x82\x81\xa6\x0c\x17\x26\x85\xab\x66\xb5\x2f\x3d\x85\x65\x85\x96\x4f\x7a\x51\x15\x65\x6a\x99\xc2\x18\xc6\x2e\x17\x74\x99\xd3\x8b\xf1\xdf\xd0\x7d\xc3\xf8\xb2\x55\x61\x16\xf9\x20\x64\x98\xb1\x53\xd2\xcf\xe3\xeb\x39\xdb\xce\x32\xc8\xbb\xfa\x4f\xf8\x4c\xb5\xb0\x35\x4a\xee\xd5\xc1\xf9\xdd\xdd\xed\xec\x76\xd6\xa7\x6f\x08\xf9\x34\xb4\x30\xf8\xc2\x53\x98\xfd\x3b\xd9\xdf\x51\x97\x23\x38\x43\x0a\x8d\x46\x69\x7d\xea\xd6\xe0\x83\xcd\xdf\x0e\x86\x12\x2d\xd2\x7a\x37\x38\x18\xb4\x87\x75\x5b\xb4\x17\xdb\xdd\xa7\x03\xa2\xad\x7d\xfb\xc0\x8e\xd9\x7b\x80\x9b\x34\xc8\x35\xa7\x4f\x01\x9d\x5b\x37\x19\xad\x97\x74\x6d\x1c\xba\x21\x1b\x82\xa2\x0c\xb9\x50\x8f\xd8\xca\xe9\xbb\xc6\xdc\x9b\xd7\x4f\xd7\x2f\x07\xb1\x10\x8e\xaf\xb8\x70\xff\x3d\xf4\xcc\xb5\xc5\x82\xd6\x01\xad\xb1\x94\x29\x04\xe3\x16\xda\x90\x2c\xea\x2e\x4c\x66\x0a\x8d\x8d\x05\xa3\x8b\xeb\x51\x65\x6d\x63\xd2\x28\xb2\x4b\xdf\x3e\x14\x74\x55\x30\x19\x16\x4a\x44\xc9\x36\x88\x32\x7c\x34\x1f\xe3\x38\x4c\x46\xd3\x2c\x6a\xeb\x07\xa2\xe9\x12\x25\x53\xcb\x50\xc9\x5a\x51\x06\xd7\x30\x5f\xc8\xc2\xa2\x92\x70\x71\x09\xaf\xd0\x93\x34\x54\x1b\x7e\xc1\x54\xb1\x10\x5c\xda\xd0\x5d\xdb\xcb\x09\x6c\x76\xe8\xda\x7b\x4d\x32\x07\x4d\x89\xf7\x2d\x63\xf8\x0c\x45\x4d\x8d\xb9\x1e\xe5\x4a\x3d\x8d\xa6\x3e\x7c\x08\x39\x7b\x77\xa0\x43\xb8\x3b\xc2\x07\x19\xc7\xfc\x6c\xfd\x46\x8a\xfb\x3c\x54\x68\x00\x0d\x88\x35\xcc\x51\x1b\x0b\x8d\x32\xf6\x0c\x7e\xfd\xf8\xfe\xf3\xad\xdf\x71\x9b\x88\xe1\xf3\x31\xf5\x51\xe0\xde\x1d\xe8\x38\x4a\xc8\x89\x62\x72\x82\xf0\x7d\xbc\xa8\xe2\xe9\x17\xae\xb9\x13\x4a\xa5\xb2\x15\xd7\x5e\x67\x16\x55\xf1\x7b\x38\x67\x78\xa1\x24\x6b\xad\xfb\x53\x87\xbc\x21\xc9\x69\x43\x0e\x96\xdd\x2b\x21\x24\x8b\xba\x53\x94\x45\xfe\x58\xfd\x0e\x00\x00\xff\xff\x8b\xc8\x2a\x1d\x3e\x06\x00\x00")

func templateTemplate1HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateTemplate1Html,
		"template/template1.html",
	)
}

func templateTemplate1Html() (*asset, error) {
	bytes, err := templateTemplate1HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/template1.html", size: 1598, mode: os.FileMode(0644), modTime: time.Unix(1643161313, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4c, 0xd9, 0xa5, 0xbc, 0x1, 0x18, 0x49, 0xf0, 0x80, 0xbf, 0xdd, 0xd9, 0xd4, 0xb8, 0x11, 0x1f, 0x4c, 0xea, 0x3b, 0x40, 0x8b, 0xd8, 0xc0, 0xea, 0xb, 0xe, 0x45, 0xcd, 0xa8, 0xa4, 0x6b, 0x16}}
	return a, nil
}

var _templateTemplate2Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4d\x6e\xdb\x3c\x10\xdd\xf3\x14\x13\x65\xf3\x7d\x45\x19\x5b\x6a\xd3\x1f\x59\x16\x6a\x24\x31\xba\x6b\x16\xee\xa2\x4b\x4a\x64\x44\x22\x12\xa9\x92\x54\x6c\xa5\xf0\x3d\xba\xeb\x39\x7a\xab\x1e\xa1\x20\xad\xc8\x96\x1c\x03\x5d\xc4\xb2\x01\x73\xde\x9b\xe1\xbc\x37\x63\x27\x67\xd7\x5f\xae\x56\xdf\x6e\x6f\x80\xdb\xaa\x4c\x51\x52\x31\x4b\x80\x5b\x5b\x63\xf6\xbd\x11\x0f\xf3\xe0\x4a\x49\xcb\xa4\xc5\xab\xb6\x66\x01\xe4\xbb\xd3\x3c\xb0\x6c\x63\x27\x2e\x67\x06\x39\x27\xda\x30\x3b\xff\xba\x5a\xe2\x0f\x41\x8a\x12\x63\xdb\x92\xa5\x28\x53\xb4\x85\x1f\x08\xa0\x22\xba\x10\x32\x86\xe9\x0c\x01\xd4\x84\x52\x21\x8b\xee\x94\x91\xfc\xbe\xd0\xaa\x91\x14\xe7\xaa\x54\x3a\x86\xf3\xe5\xc2\x3d\x0e\xbc\x53\xd2\xc6\x10\x46\xb5\x85\x60\x45\xb8\xaa\x48\x30\x43\x5b\x84\x5e\xf9\xaa\x99\xda\x60\x23\x1e\x7d\xad\x4c\x69\xca\x34\xce\xd4\xc6\xe5\xe1\x4a\x3d\xe2\x53\xf0\x16\xa1\x8b\x9a\x14\xcc\xd7\x58\x0b\x6a\x79\x0c\x51\x98\x57\x2e\xb1\x12\x12\x73\x26\x0a\x6e\x63\x88\x3e\x5e\xbc\xdf\x45\xfb\x8e\xa3\x8e\xd5\xc9\x09\xf3\x0a\x48\x63\x95\xd7\xe1\x6f\x88\x21\xac\x37\x70\x7e\xfd\xc6\x3d\x60\x54\x29\xe8\x1e\xc4\x9a\x50\xd1\x98\x18\x2e\xeb\xcd\x50\x7a\x0c\x6b\x2e\x2c\x9b\x3d\x89\xe2\x84\xaa\x75\x0c\x53\x98\x3a\x2e\xe8\x22\x23\xff\x4d\x5f\x43\xf7\xbe\x08\xff\xdf\xa9\x30\x4d\xd6\x0b\xe9\x7b\xec\x94\x3c\xf5\xe3\xf3\x19\xdd\xf7\xd2\xcb\xbb\x7c\x57\x79\xa6\x6a\x6c\x29\x24\xf3\xea\xe0\x7c\xb9\xbc\x59\xdc\x2c\x9e\xe8\x5b\x84\x3e\xf5\x57\x18\xf1\xc8\x62\x58\xbc\x9d\x0d\x27\xea\x38\x15\xa3\x82\x40\xad\x85\xb4\x9e\xba\x37\x78\x34\xfc\x7d\x63\x42\x0a\x2b\x48\x79\x18\xec\x0d\x1a\x60\xdd\x88\x06\xb1\xc3\x39\x8d\x0a\xed\xed\x1b\x02\x07\x66\x0f\x00\xd7\x29\xce\x34\x23\xf7\x98\xdc\x59\xd7\x19\x29\xd7\xa4\x35\x0e\xdd\x3a\x71\xc9\xa4\x5b\xe7\x64\xf7\x0b\x41\x89\xdb\xeb\x14\xf9\xec\x84\x8a\x07\xc8\x4b\x62\xcc\x3c\xc8\x94\xba\x0f\x52\x1f\x1e\x43\xee\x92\x03\x68\x0c\x77\x83\x1c\x31\x8e\xeb\xd3\xf6\x19\x8a\x7b\xad\xb8\x30\x20\x0c\x54\x2d\xdc\x09\x6d\x2c\xd4\xca\xd8\x33\xf8\xf3\xeb\xe7\xef\xe7\x3e\xc7\xd7\x4c\xa8\x78\x38\x2e\x7d\x14\xb8\x75\x63\x0d\x27\x11\x3a\x91\x8c\x4e\x14\x7c\x19\x2f\x78\x98\x7e\x66\x9a\x39\xa1\x44\x2a\xcb\x99\xf6\x3a\x93\x09\x0f\x5f\xc2\x39\xc3\x72\x25\xe9\xce\xba\x7f\x75\xc8\x1b\x12\x9d\x36\x64\x74\xec\xbe\x22\xb7\x55\xdd\x16\x25\xfe\x4f\x34\xfd\x1b\x00\x00\xff\xff\x73\xb5\x67\x5f\x88\x05\x00\x00")

func templateTemplate2HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateTemplate2Html,
		"template/template2.html",
	)
}

func templateTemplate2Html() (*asset, error) {
	bytes, err := templateTemplate2HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/template2.html", size: 1416, mode: os.FileMode(0644), modTime: time.Unix(1643161353, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x42, 0xda, 0x1c, 0x3d, 0x38, 0x18, 0x58, 0x5f, 0xe6, 0x74, 0x71, 0x17, 0x43, 0x7e, 0x4, 0x9c, 0x1e, 0xc4, 0xbe, 0xc2, 0x6e, 0x15, 0x71, 0x83, 0xbe, 0x9c, 0xf9, 0x53, 0x94, 0x33, 0x77, 0x3d}}
	return a, nil
}

var _templateTemplate3Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\xcd\x72\xd3\x3c\x14\xdd\xeb\x29\x6e\xd3\xcd\xf7\x31\x38\xfe\x81\xc2\xe0\x98\x0c\x19\xda\x0e\x3b\xba\xe8\x86\xa5\x6c\xa9\xf6\xa5\x96\xe5\x91\x94\xe6\x87\xf1\x7b\xb0\xe3\x39\x78\x2b\x1e\x81\x91\xfc\x93\xd8\x69\x66\x58\x34\x4e\x26\xd6\x3d\x47\x57\xe7\x1c\x49\xc9\xc5\xf5\xd7\xcf\xf7\xdf\xee\x6e\xa0\x30\xa2\x5c\x92\x44\x9b\x5d\xc9\x97\x24\x95\x6c\x07\x3f\x08\x80\xa0\x2a\xc7\x2a\x86\x60\x41\x00\x6a\xca\x18\x56\x79\x37\x4a\x69\xf6\x98\x2b\xb9\xae\x98\x97\xc9\x52\xaa\x18\x2e\x6f\x57\xf6\xb1\xe0\x83\xac\x4c\x0c\x61\x54\x1b\x98\xdd\xd3\x42\x0a\x3a\x5b\x90\x86\x90\x57\xae\x6b\x2a\xb7\x9e\xc6\xbd\xeb\x95\x4a\xc5\xb8\xf2\x52\xb9\xb5\xf3\x3c\x21\xf7\xde\x39\xb8\x21\x64\x5e\xd3\x9c\xbb\x1e\x1b\x64\xa6\x88\x21\x0a\x33\x61\x27\x0a\xac\xbc\x82\x63\x5e\x98\x18\xa2\x0f\xf3\xf7\x6d\x75\x50\x1c\x75\xac\xce\x4e\x98\x09\xa0\x6b\x23\x9d\x0f\xb7\x42\x0c\x61\xbd\x85\xcb\xeb\x37\xf6\x01\x2d\x4b\x64\x07\xd0\x53\x94\xe1\x5a\xc7\x70\x55\x6f\xc7\xd6\x63\xd8\x14\x68\xf8\xa2\x37\x55\x50\x26\x37\x31\x04\x10\x58\x2e\xa8\x3c\xa5\xff\x05\xaf\xa1\xfb\xce\xc3\xff\x5b\x17\x7a\x9d\x0e\x46\x06\x8d\x9d\x93\x5e\x8f\x9b\xcf\xd9\x41\xcb\x60\xef\xea\x9d\x70\x4c\xb9\x36\x25\x56\xdc\xb9\x83\xcb\xdb\xdb\x9b\xd5\xcd\xaa\xa7\x37\x84\x7c\x1a\x96\xd0\xb8\xe7\x31\xac\xde\x2e\xc6\x3b\x6a\x39\x82\x33\xa4\x50\x2b\xac\x8c\xa3\x1e\x02\x9e\x6c\xfe\x41\x18\x56\x68\x90\x96\xc7\xc5\x21\xa0\x11\xd6\x6d\xd1\xa8\x76\xbc\x4f\x93\x46\x87\xf8\xc6\xc0\x51\xd8\x23\xc0\x2a\xf5\x52\xc5\xe9\xa3\x47\x1f\x8c\x55\x46\xcb\x0d\xdd\x69\x8b\x36\xa4\x21\x28\xf2\x39\x17\xf2\x3b\xb6\x76\xfa\x55\x43\xee\xc2\xeb\xd5\xf5\xc3\xc1\x2c\xcc\x83\x2b\x2e\xec\x7f\x0f\x3d\x71\x65\x30\xa3\xa5\x47\x4b\xcc\xab\x18\xbc\xa0\x85\x1a\x92\xf8\xdd\x85\x49\xda\xfb\x43\x12\x7b\x73\x96\xc4\xe9\x4b\x18\x3e\x41\x56\x52\xad\x3f\xce\x52\x29\x1f\x67\x4b\x57\x9e\x42\xd6\xc6\x11\x34\x85\xbb\xa3\x32\x61\x9c\xf6\x67\xbb\x67\x28\xf6\x73\x5f\xa0\x06\xd4\x20\x76\xf0\x80\x4a\x1b\xa8\xa5\x36\x17\xf0\xe7\xd7\xcf\xdf\xcf\xfd\x4e\x97\xf1\x19\x3e\x9d\xb6\x3e\x29\xdc\xd9\x83\x13\xfa\x11\x39\x33\x99\x9c\x69\xf8\x32\x59\x14\xe1\xf2\x0b\x57\xdc\x1a\xa5\x95\x34\x05\x57\xce\x67\xe2\x17\xe1\x4b\x24\xa7\x79\x26\x2b\xd6\x46\xf7\xaf\x09\xb9\x40\xa2\xf3\x81\x4c\x86\xdd\x2b\x21\x24\xf1\xbb\x53\x94\xf8\xee\x58\xfd\x0d\x00\x00\xff\xff\x73\xb1\x06\x61\xa6\x05\x00\x00")

func templateTemplate3HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templateTemplate3Html,
		"template/template3.html",
	)
}

func templateTemplate3Html() (*asset, error) {
	bytes, err := templateTemplate3HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/template3.html", size: 1446, mode: os.FileMode(0644), modTime: time.Unix(1643161752, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xdb, 0xa3, 0x6c, 0x78, 0xe7, 0xa6, 0xa9, 0xdb, 0xb3, 0xa, 0xd6, 0xff, 0xf0, 0xae, 0xfe, 0x4d, 0x7c, 0xa, 0x4e, 0xae, 0x3a, 0x2d, 0x4c, 0x70, 0xd2, 0x46, 0xae, 0xc, 0x70, 0x18, 0x80}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"template/template.html":  templateTemplateHtml,
	"template/template1.html": templateTemplate1Html,
	"template/template2.html": templateTemplate2Html,
	"template/template3.html": templateTemplate3Html,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"template": {nil, map[string]*bintree{
		"template.html": {templateTemplateHtml, map[string]*bintree{}},
		"template1.html": {templateTemplate1Html, map[string]*bintree{}},
		"template2.html": {templateTemplate2Html, map[string]*bintree{}},
		"template3.html": {templateTemplate3Html, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
