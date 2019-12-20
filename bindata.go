// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// .docker/templates/alpine.tmpl
// .docker/templates/debian-slim.tmpl
// .docker/templates/docker-compose.tmpl
// .docker/templates/docker-entrypoint.tmpl
// .docker/templates/dockerignore.tmpl
// .docker/templates/makefile.tmpl
// .docker/templates/readme-alpine.tmpl
// .docker/templates/readme.tmpl
// .docker/templates/travis.tmpl
// .docker/templates/ubuntu.tmpl
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

var _DockerTemplatesAlpineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xd1\x4e\x1b\x3b\x10\x7d\xbe\xfe\x8a\x51\xb8\x82\x07\xae\x13\x50\xde\xb8\x97\xab\x46\x21\x6d\xa3\x96\x24\x5a\x42\x2a\x54\xaa\xca\xb1\x27\xd9\x29\xbb\xb6\x65\xcf\x2e\x89\x10\xff\x5e\xed\x6e\x02\x01\x05\xf1\xd0\xa7\xb5\x67\x8e\xec\x73\x8e\xe7\xec\xc7\x64\x7c\x09\x2a\xf3\x64\xf1\xac\xdb\x3e\x3d\x81\xde\x15\xcc\x0b\xca\x8c\x10\xdf\xc6\xc9\x97\x8b\x61\x02\x1d\xe7\xb9\xa3\xbc\x17\xe2\x00\x86\x36\xb2\xca\x32\x98\xac\x39\x75\x16\x94\x35\x80\x2b\xc6\x60\x55\x06\x06\x3d\x5a\x83\x56\x13\xc6\x7f\x80\xac\xce\x0a\x43\x76\x09\x29\x2a\x83\x21\xd6\xe0\x4f\xfd\xbe\x48\xae\x47\xa0\xfc\x1d\x28\x63\x40\x4a\xeb\xa4\x56\x3a\x45\xf0\xf5\x91\xdd\xed\x57\x1a\x2c\xc1\xaf\xbb\xd2\x93\x87\x8c\xe6\x8b\x05\x6d\x3e\x75\x27\x2f\x62\x56\x2f\x96\x5a\xc3\x92\x18\xb4\x92\x1a\x03\xd3\x82\xb4\x62\x8c\xe0\x3c\xda\x79\xa6\xe2\x2b\xf4\xf1\xf1\x0b\x19\xe4\xd1\x96\x35\x23\x4f\xbe\x0b\xb4\xa9\xfb\xa6\x2e\x0e\xa0\x1f\x50\x31\x82\x82\x92\x02\x17\x2a\x03\xb4\x25\x05\x67\x73\xb4\x5c\x2b\x52\x9a\xa9\xac\x20\xc4\xcd\x39\x1b\x19\x32\x87\x12\x6d\xd9\xb8\x57\xad\xc4\x60\x34\x83\x49\x6f\xfa\xf9\xbc\xf5\x54\xeb\xcc\xc9\x9e\xfd\x5d\x15\x5b\x70\x2b\xfe\x9a\x0d\x93\xe9\x75\xef\xeb\xcf\xc1\x68\xb6\x03\x6a\xed\x32\xde\x35\x19\xc8\xb2\x03\x4e\x71\x2f\xb9\x7b\xe2\x74\x57\x60\x6d\x52\xe6\x2c\x82\x94\x06\x3d\xa7\xe7\xa7\x20\xe7\xf0\xf0\xd0\x9e\x61\x88\xe4\xec\xe3\x23\xa4\xcc\x3e\x9e\x75\x3a\x4b\xe2\xb4\x98\xb7\xb5\xcb\x3b\x7c\x4f\x96\x7d\x70\xbf\x50\x73\xb3\x79\x1a\x88\x8a\xf2\xe1\x21\x68\xf3\xba\xf2\xc2\x4b\x29\x0b\xbf\x0c\xca\x60\x55\xde\x07\xd0\xcd\x2c\xed\xe9\xd8\x22\xf7\xeb\x7d\x8d\xb6\x10\xaf\x07\x57\x5c\xf6\x86\xa3\x69\x6f\x38\x1a\x24\xb0\x3a\x09\xab\x3b\x57\xc2\x7f\x9b\xc5\x07\x1f\x1c\x3b\x9b\x2b\xca\x2a\x4d\xff\xef\x3c\x6c\x28\x2c\x53\x8e\x50\x44\x0c\xb5\x4d\xf9\x9d\xa1\x00\xd2\xd7\x9a\x36\x77\x2b\x63\xaa\x3e\xc8\x0b\x68\x1c\x90\xe9\xb3\x64\x19\xa1\x7a\xc6\x4e\x4c\xe1\x56\x40\x05\x8f\xc5\x16\xa6\xe1\x68\xc7\x9d\x7f\x9f\x0f\x37\x8a\xd5\xd1\x7b\x81\xda\x72\x7b\xf1\xe6\xce\x66\xeb\xf7\x22\xb4\x89\xcb\x36\x02\xd5\x3e\xb2\xd1\xcd\xec\x5f\xdd\x13\xeb\x14\xd8\xd5\x9a\x41\x3b\xcb\xb8\x62\x71\x7d\x35\x48\x1a\xda\x7b\x83\xdf\x77\x7e\xfd\xe6\xa8\x2d\x82\xcb\xeb\xa6\x0f\x58\x92\x2b\x22\x50\xae\x96\x28\xfa\xe3\xc9\x0d\x48\x59\xb5\xcf\xeb\x9f\xca\x73\x1a\x76\x72\x21\x0e\xa0\xb7\x0d\xd1\x1b\x37\xfc\x49\x76\x06\xa3\x69\x72\x33\x19\x0f\x47\x53\xf8\xde\xaa\x05\xb6\x7e\xfc\x0e\x00\x00\xff\xff\x16\xea\xa4\x28\xf6\x04\x00\x00")

func DockerTemplatesAlpineTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesAlpineTmpl,
		".docker/templates/alpine.tmpl",
	)
}

func DockerTemplatesAlpineTmpl() (*asset, error) {
	bytes, err := DockerTemplatesAlpineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/alpine.tmpl", size: 1270, mode: os.FileMode(420), modTime: time.Unix(1576830973, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDebianSlimTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x51\xeb\xda\x30\x14\xc5\xdf\xf3\x29\x2e\x0e\x7c\x98\xa4\x19\xf8\x30\x90\x39\x94\xad\x1b\x65\xb3\x95\xac\x2a\x32\xc7\x88\x6d\x96\x66\xb6\x49\x48\x6e\x9d\x22\x7e\xf7\x61\xf5\xcf\x5f\x7d\xca\x39\xe7\x47\xe0\x9c\xfb\x85\x67\x33\x70\x47\xac\xac\x19\x0d\xa3\xf7\x34\xd4\xba\xa1\x01\xbd\xc4\xa2\x22\x64\x36\x4d\xd2\x7c\x9a\xa4\x31\x87\xc3\x3b\x7f\xd8\xd9\x3d\x7c\xb8\x89\x89\xf3\x16\xad\x69\x84\xae\xa3\xc2\x36\x1f\x09\x99\xf2\xaf\x90\xaf\x92\x34\xff\xbd\x8c\xf9\x8f\x24\x4b\xc7\xa7\x53\xb4\x94\x3e\x68\x6b\xce\x67\x42\x3e\x65\xf3\x35\x94\xb6\xd8\x49\x4f\xa5\x41\x7f\x74\x56\x1b\x8c\x42\x05\xec\xc1\x12\xbe\x48\xa1\xa8\x1a\x5b\xc2\xe0\xf0\xcc\x3a\xb8\x21\xc2\x21\x55\x12\xa1\x75\xa5\x40\x09\xfd\xfe\x5d\xa6\x4d\x40\x51\xd7\x40\x8f\xb0\x21\x4a\xe3\xcb\x1f\xa7\xdd\xf0\x15\xd2\xd6\x29\x2f\x4a\x09\x54\x82\xd2\x38\xa8\x10\x5d\x18\x31\xa6\x34\x56\xed\xf6\x32\x89\xe1\x3f\x6d\xd0\x79\xfb\x57\x16\x78\x35\x91\xd2\x38\xb9\x5f\xf5\x46\x2a\x35\xee\xd0\x73\xb3\xa2\x96\xc2\x80\x68\xd1\x5e\x55\x57\xd1\x37\x40\xfd\x1f\x60\x7b\xe1\x59\xad\xb7\x4c\x38\x64\xb5\x0e\x18\xd8\x5b\x60\xd8\xb8\xcb\x73\x61\x9d\x24\x24\x4e\x73\xbe\x9e\x67\x49\x9a\xc3\xcf\xde\xe3\x21\x7a\xbf\xc8\x32\xfb\xbe\x98\xc5\x70\x6d\x46\x56\x19\xff\xf6\x39\xe1\xc0\x82\xdf\xdf\xa2\xff\x01\x00\x00\xff\xff\x15\xde\x8c\x62\xdc\x01\x00\x00")

func DockerTemplatesDebianSlimTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDebianSlimTmpl,
		".docker/templates/debian-slim.tmpl",
	)
}

func DockerTemplatesDebianSlimTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDebianSlimTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/debian-slim.tmpl", size: 476, mode: os.FileMode(420), modTime: time.Unix(1576745957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerComposeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x4d\x6b\xe3\x3c\x10\xc7\xef\xfe\x14\x43\x9e\x43\x4f\x52\xec\x3e\x64\xdb\x0a\x72\xe8\xb2\x7b\xd9\xcb\x16\x5a\x4a\x6f\x46\xb1\x27\x8d\x88\x5e\xc2\x48\x76\x5c\x8c\xbf\xfb\x22\xcb\x0e\x49\x0b\xe9\x65\x6c\xcd\xfc\xf4\x9f\x37\x31\xc6\xb2\x16\xc9\x2b\x67\x05\xdc\xfc\xcf\xef\x6e\x32\x8f\xd4\xaa\x0a\xbd\xc8\x32\x80\x70\x54\x36\x88\x0c\x00\x40\x19\xf9\x8e\x02\xba\x9c\xba\xbd\x6b\x97\x29\xd2\xf7\xc0\x5f\xd3\x7d\x18\x86\xbe\x57\x5b\xe0\x3f\xa5\xc7\x61\x60\x7d\x8f\xb6\x8e\xbe\xe4\x81\x61\x18\x65\x36\x8d\xd2\x75\x52\x04\xa8\x9c\x0d\xd8\x05\x01\x7c\x72\xd4\xae\xda\x23\x6d\x95\x46\x01\xbf\x4e\xff\xd9\xcc\x4a\x65\x91\x4a\x2b\x0d\x8a\x54\x5a\x59\x69\x35\xd6\xe9\x88\x1d\x48\xb5\xae\xfb\x10\x57\xe8\xe0\x68\xa2\xce\x5b\x8a\x49\xe5\x3b\x2e\xcf\x45\xb4\x0c\xe8\xc3\x48\x1d\x1c\x05\x9f\x54\x19\x2c\x1e\xf2\x55\x2e\xa2\x59\xc0\x7f\xf0\xe2\x08\x0e\x34\xcb\xa5\x68\x11\xa3\xc5\x1c\x8d\x75\x90\xd3\xa3\xc8\x0c\xdd\x17\xc5\xbd\x88\x26\x42\x4f\x29\x21\x40\x6c\x03\xb5\xf4\x41\x55\x1e\x25\x55\x3b\xf1\xb9\x46\x24\x3e\x01\xbc\x72\xcb\x0b\xf6\xf2\x24\xee\x78\xce\x8b\x2b\x73\x98\xe8\x91\x40\xdb\x2a\x72\xd6\xe0\xbc\x68\x06\xd6\xd5\xc8\x23\xbf\xbe\x90\x9d\xa2\x95\x6e\x7c\x40\xe2\xca\xaa\xa0\xa4\x2e\x8d\x8c\xc7\x32\x5e\xf2\x57\x2f\x8c\x8a\xa9\x11\x36\xf9\x26\x62\xe3\x5c\xf0\x81\xe4\x81\x1b\x34\x8e\x3e\x4a\xed\xaa\xfd\x3a\x50\x83\xf3\xcc\x7e\x3f\x97\x7f\x1e\x5f\x1f\xcb\xbf\x4f\x2f\xcf\x6b\xf6\x66\xfc\xaa\xb8\x35\xc0\xde\x4c\x17\x7f\x16\x23\xd6\x68\x65\xd4\xbc\x2a\x00\x83\x26\xca\xcc\x47\x00\xef\xb6\x41\x00\x2b\x4e\x8e\x9d\xa4\xfa\xe4\x68\x9d\x6e\x0c\x9e\x16\x8d\xbe\x96\x41\xe6\x85\x58\x36\x9e\x96\x7e\x27\x09\x3f\x8d\x3c\xc6\xbf\x3e\x90\x87\xdb\x3c\x17\xd1\xc4\x7d\xee\xd5\x46\x5a\xf9\xcd\x22\x13\x34\x7d\xbe\x5d\x5d\xc2\xbe\xe6\x5d\xfd\xc8\x0b\x11\x4d\x96\x9d\x75\x72\xea\x62\xa4\x6a\x52\x2d\x92\x00\xed\x2a\xa9\xb3\xcc\x62\x38\x3a\xda\x8f\x60\x8d\x5b\xd9\xe8\xe9\x05\x60\x17\x90\xac\xd4\xf3\xe8\x52\x7e\x7b\x2c\xc7\x12\xfe\x05\x00\x00\xff\xff\x6f\xbd\xb4\xe4\x34\x04\x00\x00")

func DockerTemplatesDockerComposeTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerComposeTmpl,
		".docker/templates/docker-compose.tmpl",
	)
}

func DockerTemplatesDockerComposeTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerComposeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/docker-compose.tmpl", size: 1076, mode: os.FileMode(420), modTime: time.Unix(1576846847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerEntrypointTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\x52\x71\xe0\x02\x04\x00\x00\xff\xff\xc2\x78\x36\x2c\x0f\x00\x00\x00")

func DockerTemplatesDockerEntrypointTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerEntrypointTmpl,
		".docker/templates/docker-entrypoint.tmpl",
	)
}

func DockerTemplatesDockerEntrypointTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerEntrypointTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/docker-entrypoint.tmpl", size: 15, mode: os.FileMode(420), modTime: time.Unix(1576745907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerignoreTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xce\xc9\xcc\xe5\x4a\xcc\x29\xc8\xcc\x4b\xe5\xf2\x4d\xcc\x4e\x4d\xcb\xcc\x49\xe5\x4a\xc9\x4f\xce\x4e\x2d\xd2\x4d\xce\xcf\x2d\xc8\x2f\x4e\xd5\xab\xcc\xcd\x41\x17\xd2\x02\x0b\xea\xa5\x67\x96\x80\x09\x7d\x08\xa9\x05\xa5\xb4\xb8\x00\x01\x00\x00\xff\xff\x78\xa9\x74\xfb\x57\x00\x00\x00")

func DockerTemplatesDockerignoreTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerignoreTmpl,
		".docker/templates/dockerignore.tmpl",
	)
}

func DockerTemplatesDockerignoreTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerignoreTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/dockerignore.tmpl", size: 87, mode: os.FileMode(420), modTime: time.Unix(1576745854, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\xdf\x0f\xd2\x30\x10\x7e\xee\xfd\x15\x97\x6c\x66\xf8\xb0\x15\x66\x82\xd0\x48\x82\x89\x44\x79\xe0\x47\x94\x68\x7c\xd1\xd4\xad\xb0\x65\xdd\xba\xac\x1d\xa8\xc0\xff\x6e\xda\x4d\xc7\x03\xbe\x5c\x7b\xdf\x7d\xfd\xee\xbb\xde\x7a\xf3\xf6\xfd\x0a\xd9\x02\x7f\x8e\x9b\xdf\x85\x3a\x53\x73\xc9\x2b\x13\xa6\x2a\x29\x44\x03\x9f\x57\x1f\x3f\xad\x77\x5b\xb6\x40\x7f\xa4\x33\x21\x25\x9e\x1a\x51\xe3\xe1\xcb\x7a\x7b\xf8\xde\x17\xf1\x9d\xe3\x1e\x73\x29\xf0\x86\xfc\x52\x60\x70\xad\x9b\xbc\x32\xe8\xc7\xf7\x00\x6f\x98\xb4\x06\xc3\x14\x83\x45\x80\xe1\x11\xe3\x97\x00\x9e\x87\x46\x68\x43\x08\x23\xf6\x8c\xc0\x46\x06\xc4\x34\xad\x70\xd5\xb3\x68\x74\xae\x2a\xc2\x48\x9a\xeb\x5a\xf2\x5f\x7f\x91\x08\xfa\x0b\x03\xb2\x14\x49\xa6\xd0\x1f\xf5\x3e\x3a\xdd\xbc\xe4\x27\x61\x85\x7f\xb4\xb9\x4c\xbb\x14\x79\x95\xa2\xe1\x27\x34\x99\x28\x23\x88\xf6\x1f\x76\xdb\xaf\xac\xab\x81\x8b\x56\xad\x1b\x19\xbb\x77\xa1\x41\xff\xea\xfe\xe6\xce\xfc\x6b\xdf\xe1\x8e\xd1\xc0\xb3\x7a\xcf\x28\xff\x30\xc9\xed\x54\xce\x54\xdd\xea\x2c\xec\x9c\x31\x62\x13\xec\x45\x1c\x36\x38\x1a\x78\x30\x5c\x1f\xbc\xb9\xa7\x4f\x9a\xfe\x8f\xf1\x60\x21\x13\xb2\xb6\xdf\xb2\x77\x9b\x49\x54\x59\xf2\x2a\xd5\x0e\x1e\xfa\xdb\x0c\x6c\x40\x86\x1b\x5e\x08\xbb\x53\x20\x4b\x2d\x52\x0c\x2b\x0c\x34\xfd\xe6\x79\x94\xd6\x01\xfa\x6f\x00\x3c\xcc\x8c\xa9\x35\xa3\x54\x1b\x9e\x14\xea\x2c\x9a\xa3\x54\x97\x28\x51\x25\xe5\x74\x1a\xbf\x7e\x35\x1b\xcf\xe9\x64\x16\x4f\x27\xe3\x39\xbc\xb0\x63\x30\xf8\x13\x00\x00\xff\xff\xa4\xdc\xfd\x9f\x6f\x02\x00\x00")

func DockerTemplatesMakefileTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesMakefileTmpl,
		".docker/templates/makefile.tmpl",
	)
}

func DockerTemplatesMakefileTmpl() (*asset, error) {
	bytes, err := DockerTemplatesMakefileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/makefile.tmpl", size: 623, mode: os.FileMode(420), modTime: time.Unix(1576745814, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesReadmeAlpineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\x1b\x37\x10\xbd\xeb\x57\x3c\xc8\x01\x12\x03\xa6\x2c\x7f\x01\x92\x81\x1c\xda\x38\x69\x83\x06\x68\x9a\xb8\xed\xd5\xa3\xdd\x59\x2d\x2b\x2e\xb9\x25\xb9\x92\x55\x41\xfd\xed\x05\xc9\xd5\x97\x13\xa7\x41\xbc\xf1\x41\x06\xb9\x8f\xf3\xde\x0c\x67\x38\xe4\x11\xfc\x42\x6a\x2f\x72\x93\xcd\xd8\xf6\x7a\x47\x47\xf8\xd9\x2c\xe0\x0d\x1a\xc7\x61\x78\x84\x37\xd2\x3a\x9f\x86\x6f\x8c\x45\xd1\x0e\x69\xca\x27\x58\x9a\x06\x9a\x39\x0f\x0b\x26\x8d\x54\x39\x64\x45\x53\x46\x32\x37\xe8\xf5\xee\xee\xee\x5c\xc9\x4a\xf5\xd2\x0c\xea\x46\x29\xdc\x0f\xed\x3f\x33\x33\x3f\x8d\xd4\xd7\xab\x15\x06\x7f\xb0\x75\xd2\x68\xac\xd7\x61\x45\xaf\x67\x2c\x3e\x59\x2b\x32\x53\xd5\xc6\x71\x22\x4a\xb8\xa3\x20\xf0\x95\x65\xf2\x0c\xd2\x20\x25\xc9\xed\xd6\xc5\x61\x72\xf0\x65\xbf\x15\x60\x1b\x0d\xe1\x25\x84\xb0\x15\xc4\x1c\xcf\x5e\xd4\x8b\xfc\xf8\x34\x27\x4f\xd7\xa7\xa6\xf6\xa7\x54\xd7\x71\xf4\x45\x95\xab\x95\x2c\x30\xf8\x91\x1c\xaf\xd7\x62\xb5\x62\x9d\x87\xb9\x34\x83\xf5\xba\x1f\x34\xe0\xf0\x2f\x05\xf3\xb5\x22\xe7\x65\xe6\x98\x6c\x56\x82\x74\x8e\x5f\xe4\x84\x34\xf5\x7a\x1f\x3d\x59\x1f\x03\x5f\x83\x15\x1d\x82\x66\x2d\xe8\xb1\x90\x34\x35\x44\x1e\x96\xed\x19\x6f\xd7\x6c\xe2\x84\x5b\x63\x51\x5b\x73\xbf\x3c\xe4\xf2\xc6\xfe\x9f\xdd\x00\xd9\x9a\x79\x7d\xcf\x59\xe3\x19\xb7\x21\x2c\xc8\x4c\x55\x91\xce\x1f\xb7\x10\xe3\x3d\xc7\xb3\xf7\x7f\xde\xb4\x91\x3c\x8c\x72\x9c\xdb\xb7\x4e\x55\xad\xd8\xc1\x14\x3b\xdb\x3f\xa0\xe0\x05\x9c\x0c\x5f\xc0\x1b\x84\x37\x28\x59\xd5\x31\x07\x1b\x9d\xb3\x75\x3e\x84\xca\x97\x8c\x09\x39\x99\xb9\xeb\x27\xaa\x6a\x7f\x45\x13\x92\xdf\x6a\xaa\x18\x02\x1f\x33\x4b\x35\x83\x94\x8a\x4c\xb7\x0b\x66\xef\x50\x58\x53\x45\xd4\x73\x07\x2f\x2b\x56\x52\xf3\xa0\x33\x52\x87\x5a\x6a\xa6\x3a\xb8\x7f\xa0\x60\x9f\x3d\xa8\x79\xa0\x00\x99\xd1\x9e\xa4\x96\x7a\xba\xb3\xf0\x34\x59\x87\x52\x5e\x19\xa5\x38\xf3\xe0\x39\xdb\x65\x52\xf3\x59\xce\xa4\x30\xa2\x8c\xe6\xe7\xae\x15\xde\x5d\x84\xc4\x92\xc9\xe2\x7c\x78\x76\xb9\xa7\xaa\x8d\x8e\x2f\xc9\x63\xc1\x96\xe1\xc3\x04\xe7\x98\x70\x61\x2c\x47\x78\x87\x12\x9c\xd4\x59\x34\x7a\x25\xce\xce\xc5\xf9\xf0\x6b\x94\x3c\x5c\xd3\x9d\x1c\x83\x42\x2a\x1e\xf8\x7b\xbf\x4b\x99\x56\x46\xa8\x12\x47\x73\x0e\x15\xb4\x01\x75\x4e\x9c\xb9\x39\x84\x88\xbf\x8f\xd1\x93\x03\x21\x20\xe2\x82\xee\x36\x82\x2b\x92\x0a\x42\xd4\xa5\xd1\xb1\x60\x4a\xb3\x38\xd8\x82\x4a\x4e\x4b\x8f\x32\x68\x48\x18\xdd\x54\x13\xb6\x0e\xc6\x22\x2d\xa6\x3c\xb7\xec\x1c\x3f\x31\x45\x1d\xfa\x37\x46\x93\xca\x71\x6b\x9b\xaa\xee\x43\x88\x39\x5b\x59\x48\xce\x21\x70\x23\x5d\xad\x68\xb9\x91\x36\x59\x62\xfb\x31\x78\xd3\x8a\xbd\x6d\x53\x85\x26\xa6\xf1\xd8\x37\xf7\x34\x6d\xd3\x97\xfd\xcb\xd1\x60\x34\x1a\x0e\x2f\x47\x27\xe7\x83\x8b\xd1\xd5\xf8\x62\x7c\x72\x36\xab\xfa\x5f\xb1\x87\xb1\x9e\x09\x96\x72\xd9\xc4\xb3\xfa\x6c\x56\x81\xac\x69\x74\x0e\x42\xad\x28\x63\x48\x8d\xf7\x64\x65\xda\x70\xbe\xaf\x4d\x68\x39\x25\x57\xf8\x36\xd5\xe6\xbb\x24\x0b\x3b\x28\x93\x91\x2a\x8d\xf3\xd7\xe3\xf3\x61\xa8\xda\x5f\x1b\x5f\x37\xbb\xa2\x35\x87\x4d\xbb\xeb\x42\xf9\xcb\x19\x0d\x21\xd2\xbf\x2f\x96\x4a\x84\x74\x5c\x2b\xe1\xf3\x24\x5c\x5a\x7c\x3a\x91\xf3\x49\x10\x11\x28\x77\xee\x13\x3e\xfe\xf6\x4e\x7a\xc6\x06\xdc\x21\x7f\x61\x94\x32\x8b\x90\xec\xbb\xc6\x16\xae\x16\xde\xb3\xdd\xb4\xb3\x2d\xa6\x73\xde\xd0\xa4\xb6\xbc\x8b\xd2\x3c\xe0\x6e\x99\x3b\xe5\xa5\xb9\xb1\xd2\xb3\xdb\xeb\x0e\x0f\xee\x12\x94\xb8\x4b\x72\xd8\xa0\xf3\xef\xe3\xb9\x08\xd3\xa2\x08\x57\xf2\x9d\x9a\x38\x8c\x0a\xa4\x2e\x8c\xad\xc8\x87\x2b\x2f\xa1\x66\x1b\xf3\x2f\x85\xa4\x3b\x3d\xb5\x35\x21\xa7\x37\x32\x7e\x77\x21\x03\x9c\x32\x8b\x13\x4c\x1a\x0f\x2e\x0a\xce\xbc\x9c\x33\x2a\xf6\xa5\x89\x2f\x8d\x29\xf9\x92\xed\x83\xb3\xa8\x4d\x96\xd6\x1c\x5e\xfc\x14\x41\x0e\xff\x5e\x84\xaa\x4e\xd8\x93\x6f\x3b\x7b\xde\xea\x4c\x35\x79\x08\xd9\x07\x4e\x75\x72\xdc\xe1\x86\xd8\xd6\xe6\xd6\xf9\xbf\x1b\x99\xcd\x3e\x75\x37\xa4\x48\x38\x88\x30\xde\xfa\x83\x17\xb1\x45\xc8\xa8\x8f\x1d\x36\xa6\x8e\x3f\x1f\x94\x4e\x45\xbb\xa6\x62\x9c\x0d\xc7\x97\xc3\x8b\xd1\xf8\x6a\x74\x31\xbc\x1a\x41\xe0\x43\x9a\x27\xb4\x8f\x10\x17\x5e\x1a\x21\x72\xdb\x3b\xaa\xab\x39\x4b\x4d\x2e\xdd\x17\xdf\xde\x0c\x76\xd7\xff\x1b\xf2\xe4\xd2\x3b\x33\x08\x70\x98\xb2\x66\x1b\x5e\x77\x93\x65\xd2\x91\xde\x9c\x19\x85\x4c\x0c\xdd\x26\x08\x51\x1e\x29\x33\xf3\x10\xa6\xf8\xa0\xe8\xbd\x63\x1f\x1c\x0f\xfd\x55\x9b\x05\xae\x8f\xff\x0b\x00\x00\xff\xff\x29\x77\x44\x5d\xea\x0e\x00\x00")

func DockerTemplatesReadmeAlpineTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesReadmeAlpineTmpl,
		".docker/templates/readme-alpine.tmpl",
	)
}

func DockerTemplatesReadmeAlpineTmpl() (*asset, error) {
	bytes, err := DockerTemplatesReadmeAlpineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/readme-alpine.tmpl", size: 3818, mode: os.FileMode(420), modTime: time.Unix(1576846946, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesReadmeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdd\x6e\x1b\x37\x13\xbd\xd7\x53\x1c\xc8\x01\x12\x03\xa6\x2c\xff\x01\x92\x81\x5c\x7c\x5f\x9c\xb4\x41\x03\x34\x8d\xdd\xf6\xd6\xa3\xdd\x59\x2d\x2b\x2e\xb9\x25\xb9\x92\x55\x41\x7d\xf6\x82\xe4\xea\xcf\x89\xd3\x20\xde\xfa\x42\x06\xb9\xc3\x39\x67\x0e\x67\xc8\xe1\x11\xfc\x42\x6a\x2f\x72\x93\xcd\xd8\xf6\x7a\x47\x47\xf8\xd1\x2c\xe0\x0d\x1a\xc7\x61\x78\x84\x77\xd2\x3a\x9f\x86\xef\x8c\x45\xd1\x0e\x69\xca\x27\x58\x9a\x06\x9a\x39\x0f\x0b\x26\x8d\x54\x39\x64\x45\x53\x46\x72\x37\xe8\xf5\xee\xef\xef\x5d\xc9\x4a\xf5\xd2\x0c\xea\x46\x29\x3c\x0c\xed\x5f\x33\x33\x3f\x8d\xd0\xd7\xab\x15\x06\xbf\xb1\x75\xd2\x68\xac\xd7\x61\x45\xaf\x67\x2c\x3e\x5b\x2b\x32\x53\xd5\xc6\x71\x02\x4a\x76\x47\x81\xe0\x1b\xcb\xe4\x19\xa4\x41\x4a\x92\xdb\xad\x8b\xc3\x14\xe0\xeb\x7e\x4b\xc0\x36\x1a\xc2\x4b\x08\x61\x2b\x88\x39\x5e\xbc\xaa\x17\xf9\xf1\x69\x4e\x9e\xae\x4f\x4d\xed\x4f\xa9\xae\xe3\xe8\xab\x2c\x57\x2b\x59\x60\xf0\x7f\x72\xbc\x5e\x8b\xd5\x8a\x75\x1e\xe6\xd2\x0c\xd6\xeb\x7e\xe0\x80\xc3\xbf\x24\xe6\x5b\x45\xce\xcb\xcc\x31\xd9\xac\x04\xe9\x1c\x3f\xc9\x09\x69\xea\xf5\x6e\x3d\x59\x1f\x85\xaf\xc1\x8a\x0e\x8d\x66\xad\xd1\x53\x92\x34\x35\x44\x1e\x96\xed\x39\x6f\xd7\x6c\x74\xc2\x9d\xb1\xa8\xad\x79\x58\x1e\x62\x79\x63\xff\xcd\x6f\x30\xd9\xba\x79\xfb\xc0\x59\xe3\x19\x77\x41\x16\x64\xa6\xaa\x48\xe7\x4f\x7b\x88\x7a\xcf\xf1\xe2\xe3\xef\x37\xad\x92\xa7\xce\xb6\xa2\xa6\xad\xd9\x77\x4d\x55\xad\xd8\xc1\x14\x3b\xc7\xff\x43\xc1\x0b\x38\x19\xbe\x80\x37\x16\xde\xa0\x64\x55\xc7\x04\x6c\x74\xce\xd6\xf9\xa0\x93\x2f\x19\x13\x72\x32\x73\xd7\xcf\xa1\xd4\xfe\x8a\x26\xa4\xbd\xd5\x54\x31\x04\x6e\x33\x4b\x35\x83\x94\x8a\x30\x77\x0b\x66\xef\x50\x58\x53\x45\xab\x97\x0e\x5e\x56\xac\xa4\xe6\x41\x37\x88\x0e\xb5\xd4\x4c\x75\x08\xfc\x00\x7e\x1f\x3a\x50\x79\x04\x8f\xcc\x68\x4f\x52\x4b\x3d\xdd\x79\x78\x06\xa7\x43\x1e\x6f\x8c\x52\x9c\x79\xf0\x9c\xed\x32\x51\xf9\x22\x60\xa2\x17\xad\x8c\xe6\x97\xae\x65\xdd\x91\x36\x62\xc9\x64\x71\x3e\x3c\xbb\xdc\xa3\xd4\xea\xe2\x4b\xf2\x58\xb0\x65\xf8\x30\xc1\x39\x26\x5c\x18\xcb\xd1\xbc\x2b\x7c\x27\x75\x16\x3d\x5e\x89\xb3\x73\x71\x3e\xfc\x16\x1a\x8f\xd7\x74\xc4\xc5\xa0\x90\x8a\x07\xfe\xc1\xef\xd2\xa4\xe5\x10\x6a\xc2\xd1\x9c\x43\xbd\x6c\x8c\xba\x45\xcd\xdc\x1c\x42\xc4\xdf\xa7\xb0\xc9\x81\x10\x2c\xe2\x82\x8e\xf4\xe7\x8a\xa4\x82\x10\x75\x69\x74\x2c\x8f\xd2\x2c\x0e\x94\xaf\xe4\xb4\xf4\x28\x03\x81\x64\xa3\x9b\x6a\xc2\xd6\xc1\x58\xa4\xc5\x94\xe7\x96\x9d\xe3\xe7\xe4\xa4\x43\xff\xc6\x68\x52\x39\xee\x6c\x53\xd5\x7d\x08\x31\x67\x2b\x0b\xc9\x39\x04\x6e\xa4\xab\x15\x2d\x37\xbc\x26\x4b\x6c\x3f\x86\x50\x5a\xa6\x77\x6d\x7a\xd0\xc4\x34\x1e\xfb\xee\x9e\x41\x6c\xfa\xba\x7f\x39\x1a\x8c\x46\xc3\xe1\xe5\xe8\xe4\x7c\x70\x31\xba\x1a\x5f\x8c\x4f\xce\x66\x55\xff\x1b\xb6\x2e\x96\x2e\xc1\x52\x2e\x9b\x78\x1a\x9f\xcd\x2a\x90\x35\x8d\xce\x41\xa8\x15\x65\x0c\xa9\xf1\x91\xac\x4c\xfb\xcc\x0f\xb5\x09\x37\x4a\xc9\x15\xbe\x83\xb2\xe9\x3e\x41\xd8\x41\x99\x8c\x54\x69\x9c\xbf\x1e\x9f\x0f\x43\x81\xfe\xdc\xf8\xba\xd9\xd5\xa7\x39\xbc\x8d\x3b\x2d\x8b\x3f\x9c\xd1\x10\x22\xfd\xfb\x6a\x61\x44\x93\x2e\x2b\x23\x34\x2f\x93\xd0\x87\xf8\x74\xe0\xe6\x93\xc0\x20\xe0\xed\x02\x27\xdc\xfe\xf2\x41\x7a\xc6\xc6\xb8\x2b\xf0\xc2\x28\x65\x16\x21\xb5\x77\x37\x56\x68\x15\xbc\x67\xbb\xb9\xa7\xb6\x36\xdd\x82\x86\xdb\x67\x0b\xba\x28\xcd\x23\xe0\x16\xb6\x3b\x50\x9a\x1b\x2b\x3d\xbb\xbd\x93\xff\x51\x6f\x40\x09\xb8\x24\x87\x8d\x75\xfe\x1f\xc4\x2c\xc2\xb4\x28\x42\x67\xbd\xa3\x12\x87\x11\x5e\xea\xc2\xd8\x8a\x7c\xe8\x5c\x09\x35\xdb\x98\x70\x49\x8c\x8e\xc8\xd4\xd6\x84\x0c\xde\x70\xf8\xd5\x85\x5d\x77\xca\x2c\x4e\x30\x69\x3c\xb8\x28\x38\xf3\x72\xce\xa8\xd8\x97\x26\xbe\x16\xa6\xe4\x4b\xb6\x8f\x0e\x9c\x36\x41\x5a\x77\x78\xf5\x43\x34\x72\xf8\xfb\x22\x14\x70\xb2\x3d\xf9\x8e\x03\xe6\xbd\xce\x54\x93\x07\xb1\x3e\x71\xaa\x8a\xe3\xae\xf6\xc1\xb6\x0e\xb7\x61\xff\xd9\xc8\x6c\xf6\x79\xa0\x21\x2d\xc2\x69\x83\xf1\x36\x12\xbc\x8a\xc7\xbf\x8c\xe4\xd8\x61\xe3\xea\xf8\xcb\x72\x74\xc7\xd8\x35\x15\xe3\x6c\x38\xbe\x1c\x5e\x8c\xc6\x57\xa3\x8b\xe1\xd5\x08\x02\x9f\xd2\x3c\xa1\x7d\x3f\xb8\xf0\x48\x08\x9a\x6d\xfb\x4c\x57\x73\x96\x6e\xaf\xd4\xf6\xbd\xbf\x19\xec\x9a\xf7\x1b\xf2\xe4\xd2\x13\x31\x1c\x2a\x0e\x53\xd6\x6c\xc3\xc3\x6c\xb2\x4c\x3c\xd2\x73\x31\xa3\x90\x7d\xe1\x26\x09\x44\x94\x47\xca\xc6\x3c\x68\x14\x9f\x03\xbd\x0f\xec\x43\xd4\xe1\xe2\xd4\x66\x81\xeb\xe3\x7f\x02\x00\x00\xff\xff\xbc\x2a\xc8\x6f\xa5\x0e\x00\x00")

func DockerTemplatesReadmeTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesReadmeTmpl,
		".docker/templates/readme.tmpl",
	)
}

func DockerTemplatesReadmeTmpl() (*asset, error) {
	bytes, err := DockerTemplatesReadmeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/readme.tmpl", size: 3749, mode: os.FileMode(420), modTime: time.Unix(1576846940, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesTravisTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x4f\x6b\xdc\x30\x10\xc5\xef\xfa\x14\x83\x59\xc8\xa1\xc8\x6a\x7a\xe8\xc1\x90\xc3\xd2\x5d\x4a\x68\x93\x2d\xbb\x34\xbd\x14\x16\xad\x35\x6b\xab\x91\x25\x47\x23\x39\x4b\x5c\x7f\xf7\x12\xdb\xc9\xba\xcd\x9f\xe6\x66\xf4\xde\xef\xc9\x33\x9a\x91\xfb\x80\x7e\x4b\xb9\xd7\x75\xc8\x18\x00\x07\xe5\xf2\x6b\xf4\xa0\x2b\x59\x20\x31\x46\x51\xb9\x0c\x3c\xde\x44\xed\x51\x31\xb6\xc3\xbd\xf3\xb8\xd5\x96\x82\x34\x66\x40\xee\x3d\xe0\x2b\xe0\x7b\x10\x91\xbc\x30\x2e\x97\x46\xec\xb4\x15\x43\x18\x27\xa3\xab\xb7\x3b\x39\xa1\x25\xe7\x7b\x20\x8f\xde\x00\xff\x0a\x65\x08\x35\x65\x42\x14\x3a\x94\x71\x97\xe6\xae\x9a\x12\x7f\x7d\x7b\x34\x28\x09\x49\x28\x77\x6b\x8d\x93\x4a\x9c\xa6\x1f\x3e\xa6\xa7\x42\x69\x0a\x5b\xa3\x6d\x3c\xa4\x41\xfa\xb4\xb8\x03\xce\x5d\x0c\x75\x0c\x30\xe1\x47\xad\xbf\x3d\x48\x0f\x87\x66\xff\x92\x9c\x97\x95\x53\xf0\xee\x00\xc7\xe4\x27\x15\xff\xc7\x33\xad\xb5\x6f\x4e\xd5\xbc\xe0\xfc\xa7\x61\x6f\x20\xc6\xec\xe7\x40\xcc\x4b\x07\x27\x6d\x82\x87\x1a\xbd\xae\xd0\x06\x69\x92\x2c\xf8\x88\xdd\x09\xfc\x1e\x72\x03\x22\x08\x0c\xf9\x18\x29\x94\xc4\xca\xd9\xf4\x17\xb9\xc9\xe5\x84\xbe\xd1\x39\x3e\x4c\x8d\x47\x0a\xd2\x87\xc7\x31\x99\x0e\x56\xae\x46\xd7\x5e\x1b\x24\x91\xcc\xae\x96\xeb\xcd\xf9\xea\x32\xe9\xd5\xf3\x8b\xf9\xe7\xe5\x59\x72\x78\xef\xef\xae\x5d\x23\xc2\xad\xb6\x21\x9b\xb5\xa3\x47\xfc\x14\x82\x77\x09\x63\x68\x9b\xac\x6d\xbd\xb4\x05\xc2\xac\x91\x06\xb2\x33\x48\xaf\xd0\x93\x76\x96\xba\xae\x8f\x1a\x99\xb3\xb6\xed\x2d\xe9\x42\x7b\x78\x55\x12\xd2\xd4\xda\xe2\x6b\x8e\xfb\x7e\xb6\x2d\x5a\xd5\x75\x8c\x19\x69\x8b\x28\x0b\xcc\x60\x27\xa9\x64\xec\xe9\xfa\x0c\x2f\xd6\x0c\xff\x35\x5d\xab\x5d\xd4\x46\x01\xe7\x74\x13\x25\x95\xc0\x03\x24\xb3\xbe\xf4\x04\xd2\x63\x57\xa7\x21\x03\xf1\xe0\x7a\x66\x45\x27\x27\xc6\x15\xda\x02\x8f\x30\x5b\xac\x3e\x7d\x59\xae\xb7\xdf\x37\xcb\xf5\xe5\xfc\x62\x09\xbc\x7e\x3c\xfb\x36\xdf\x6c\x7e\xac\xd6\x8b\x29\x58\x47\x2a\x8f\x77\xb0\xf1\x55\x29\x1b\xf5\x3f\x01\x00\x00\xff\xff\xe0\x56\xfa\xe3\x26\x04\x00\x00")

func DockerTemplatesTravisTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesTravisTmpl,
		".docker/templates/travis.tmpl",
	)
}

func DockerTemplatesTravisTmpl() (*asset, error) {
	bytes, err := DockerTemplatesTravisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/travis.tmpl", size: 1062, mode: os.FileMode(420), modTime: time.Unix(1576830941, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesUbuntuTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xdf\x6e\xd3\x30\x18\xc5\xef\xfd\x14\x9f\x86\x34\x01\x53\x62\xa6\x5d\x31\x31\xb4\x0a\x02\x44\xd0\x64\xf2\xd2\x56\x15\x45\xc8\x49\x8c\x63\x48\x6c\xcb\xfe\x5c\x48\xab\x3e\x10\xcf\xc1\x8b\xa1\xa6\x9d\xfa\xe7\xca\xe7\x9c\x9f\x65\x9d\xe3\x0f\x2c\x1f\x43\x28\x83\xc6\x70\x7b\xfd\x3a\xbe\x7e\x45\xc8\x78\x94\x66\xc5\x28\xcd\x12\x06\x8f\xff\xfe\x96\xdc\xa3\x12\x1a\x3e\x99\xb0\x12\x08\xcf\x7b\x63\x56\xca\xc4\x95\xe9\x5e\xc0\x1b\x2f\xf6\xf8\xfe\x10\xbf\x25\x64\xc4\x3e\x42\x31\x4b\xb3\xe2\xfb\x34\x61\x8f\x69\x9e\xdd\xad\xd7\xf1\x54\x38\xaf\x8c\xde\x6c\x08\x79\x97\x3f\xcc\xa1\x36\xd5\x2f\xe1\x22\xa1\xd1\xf5\xd6\x28\x8d\xb1\x6f\x80\x9e\x58\xc2\x26\x19\x54\x4d\x67\x6a\xb8\xfa\x73\xce\x06\xb8\x20\xdc\x62\x24\x05\x42\xb0\x35\x47\x01\x97\x97\x47\x99\xd2\x1e\x79\xdb\x42\xd4\xc3\x82\x48\x85\xb0\x20\xb6\xc7\xc6\xe8\x9b\xc8\x2a\xfb\xf4\x82\x55\xf6\xe6\x70\x35\x0a\x56\x3a\x5e\x0b\x88\x04\x48\x85\x57\x0d\xa2\xf5\xb7\x94\x4a\x85\x4d\x28\xb7\x03\x29\xfe\x56\x1a\xad\x33\x3f\x45\x85\x3b\x13\x4b\x85\xf7\xc7\x1b\x9f\x09\x29\xef\x06\x74\xde\xb3\x6a\x05\xd7\xc0\x03\x9a\x9d\x1a\x0a\xbb\x0e\x22\xf7\x03\xe8\x92\x3b\xda\xaa\x92\x72\x8b\xb4\x55\x1e\x3d\x7d\x09\x14\x3b\xbb\x3d\xb6\x6c\x90\x84\x24\x59\xc1\xe6\x0f\x79\x9a\x15\xf0\xf5\xe2\xf4\x5b\x2e\xbe\x91\x69\xfe\x65\x32\x4e\x60\xd7\x8c\xcc\x72\xf6\xf9\x7d\xca\x80\x7a\xb7\xdc\x47\xff\x03\x00\x00\xff\xff\x4f\xe2\x8d\x91\xf4\x01\x00\x00")

func DockerTemplatesUbuntuTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesUbuntuTmpl,
		".docker/templates/ubuntu.tmpl",
	)
}

func DockerTemplatesUbuntuTmpl() (*asset, error) {
	bytes, err := DockerTemplatesUbuntuTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/ubuntu.tmpl", size: 500, mode: os.FileMode(420), modTime: time.Unix(1576745927, 0)}
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
	".docker/templates/alpine.tmpl":            DockerTemplatesAlpineTmpl,
	".docker/templates/debian-slim.tmpl":       DockerTemplatesDebianSlimTmpl,
	".docker/templates/docker-compose.tmpl":    DockerTemplatesDockerComposeTmpl,
	".docker/templates/docker-entrypoint.tmpl": DockerTemplatesDockerEntrypointTmpl,
	".docker/templates/dockerignore.tmpl":      DockerTemplatesDockerignoreTmpl,
	".docker/templates/makefile.tmpl":          DockerTemplatesMakefileTmpl,
	".docker/templates/readme-alpine.tmpl":     DockerTemplatesReadmeAlpineTmpl,
	".docker/templates/readme.tmpl":            DockerTemplatesReadmeTmpl,
	".docker/templates/travis.tmpl":            DockerTemplatesTravisTmpl,
	".docker/templates/ubuntu.tmpl":            DockerTemplatesUbuntuTmpl,
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
	".docker": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"alpine.tmpl":            &bintree{DockerTemplatesAlpineTmpl, map[string]*bintree{}},
			"debian-slim.tmpl":       &bintree{DockerTemplatesDebianSlimTmpl, map[string]*bintree{}},
			"docker-compose.tmpl":    &bintree{DockerTemplatesDockerComposeTmpl, map[string]*bintree{}},
			"docker-entrypoint.tmpl": &bintree{DockerTemplatesDockerEntrypointTmpl, map[string]*bintree{}},
			"dockerignore.tmpl":      &bintree{DockerTemplatesDockerignoreTmpl, map[string]*bintree{}},
			"makefile.tmpl":          &bintree{DockerTemplatesMakefileTmpl, map[string]*bintree{}},
			"readme-alpine.tmpl":     &bintree{DockerTemplatesReadmeAlpineTmpl, map[string]*bintree{}},
			"readme.tmpl":            &bintree{DockerTemplatesReadmeTmpl, map[string]*bintree{}},
			"travis.tmpl":            &bintree{DockerTemplatesTravisTmpl, map[string]*bintree{}},
			"ubuntu.tmpl":            &bintree{DockerTemplatesUbuntuTmpl, map[string]*bintree{}},
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
