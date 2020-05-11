// Code generated for package stage by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../stage/codefresh/assets/assets.md
// ../stage/codefresh/certs/tls.md
// ../stage/codefresh/config.yaml
// ../stage/codefresh/images/images-list
// ../stage/codefresh/images/images.md
// ../stage/k8s-agent/config.yaml
package stage

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

var _codefreshAssetsAssetsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xca\x41\x0a\x02\x20\x10\x05\xd0\xbd\xa7\xf8\xe0\x49\x22\x87\x8a\x40\x21\xbc\x80\xe1\x17\x05\x53\x98\xb1\xfb\x47\xfb\xe7\xbd\xc7\x4d\xa2\xbc\x2e\x59\x02\x9e\xdf\x37\x75\xf1\xd0\x50\x56\xc5\x9d\xf3\x83\x62\xc6\x63\x68\x5b\x71\xdd\x95\x4d\x69\xdd\xb9\x90\x10\x53\x86\x84\x47\x46\x1b\x93\x86\xb1\x70\xfa\xf8\xc3\x59\xa9\xbf\x00\x00\x00\xff\xff\xd6\xa5\xce\x4f\x58\x00\x00\x00")

func codefreshAssetsAssetsMdBytes() ([]byte, error) {
	return bindataRead(
		_codefreshAssetsAssetsMd,
		"codefresh/assets/assets.md",
	)
}

func codefreshAssetsAssetsMd() (*asset, error) {
	bytes, err := codefreshAssetsAssetsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codefresh/assets/assets.md", size: 88, mode: os.FileMode(420), modTime: time.Unix(1588502364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codefreshCertsTlsMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x51\x3d\x8f\x1a\x31\x10\xed\xfd\x2b\x9e\xd8\x26\x69\x36\x3d\x0a\x69\x68\xa3\xa4\x40\xa7\x6b\x31\xde\xf1\xee\x88\xc1\x5e\x79\x66\xe1\xf6\xdf\x9f\x6c\x38\x28\xae\xb1\xc6\x7e\xd6\xfb\x9a\xae\xeb\xf0\x4e\x27\x1c\x0e\x7f\xb1\xa7\x62\x1c\x39\x78\x23\x45\xcc\x05\xfb\x3c\x50\x2c\xa4\x13\x38\xa9\x79\x11\x2a\xee\x39\x21\xe4\x14\x79\x5c\x0a\x29\x38\x8d\x85\x54\x61\xa2\x98\xbd\xf9\x0b\x19\x15\x85\x0f\x21\x97\xc8\x69\x84\x65\x60\x63\xa2\x1b\x9c\x69\x05\x27\x5c\xbd\x2c\xa4\xfd\xea\x2f\xe2\xdc\xf1\x78\x6c\x43\x87\x81\xa2\x5f\xc4\x1e\xb0\x33\xd1\xad\x03\x94\x24\x1e\x78\x4c\x34\x6c\x11\xbd\x28\x39\x20\x50\xb1\x6d\x3b\xf5\x97\xaa\xf4\xa1\x98\x43\x65\xff\x7a\x9c\x0b\x5f\xbd\x51\x7f\xa6\xb5\x0a\x38\xc7\x11\xf5\xe3\x8b\x6c\xd7\xb8\xf0\xe3\x21\xfa\xf3\x95\xb2\xea\xf3\xd0\x7a\xf0\x69\xc0\xa2\xa4\x0f\x4b\xc8\x77\x96\xaa\xd1\xb0\x7a\x39\xd3\xda\xbb\x7b\x7b\xbe\xd5\xd7\x90\x9a\x34\xb2\x90\x42\xa7\xbc\xc8\x00\xfa\x60\xb5\x1a\xde\x26\x82\xce\x14\x38\x32\x0d\x90\x1c\xbc\x71\x4e\xbd\xfb\x6f\x13\x95\x1b\x2b\xe1\xbb\x57\x2b\x0b\x81\x0d\x23\x25\x2a\xcd\x59\x05\xb5\x81\x2d\xf1\x73\x73\x37\xb6\x09\xfb\x7f\xbb\xdf\xa3\xe4\x93\x97\xde\xcf\xf3\x5b\x91\x3f\x9f\x01\x00\x00\xff\xff\xee\xcd\x17\x47\xea\x01\x00\x00")

func codefreshCertsTlsMdBytes() ([]byte, error) {
	return bindataRead(
		_codefreshCertsTlsMd,
		"codefresh/certs/tls.md",
	)
}

func codefreshCertsTlsMd() (*asset, error) {
	bytes, err := codefreshCertsTlsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codefresh/certs/tls.md", size: 490, mode: os.FileMode(420), modTime: time.Unix(1588502364, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codefreshConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x6f\x1b\xb9\xf1\x7f\xbf\x9f\x62\x10\xbd\xf8\xdf\xfd\xe1\x5d\x49\x49\xee\x92\x5b\xf4\xc9\xa7\xf8\xee\x82\xe0\x62\xc3\xb2\xd3\x0b\x8a\xc2\xa0\xb8\x23\x89\x15\x97\x5c\x93\x5c\xd9\x4a\xd1\xef\x5e\xcc\x90\xfb\x20\x27\x40\x93\xb8\xe8\x9b\x16\x36\x20\x72\x86\xf3\x3c\x24\x7f\xdc\x1a\x83\xa8\x44\x10\x65\x06\xb0\x53\xa6\x2a\x41\xda\x0a\xd7\x0e\xfd\x36\x03\x50\xc6\x07\xa1\x35\x3a\x62\x03\x4c\x20\x1c\x1a\x2c\x21\x4d\x00\x9e\xd8\x06\x9d\x08\xd6\x3d\x81\x1c\x44\xd3\xe8\xc3\x20\x0e\xd2\x55\x50\xe1\x5a\x19\x15\x94\x35\x83\xcc\x16\x75\x4d\xeb\x93\xf2\x69\xdb\x6c\x9c\xa8\x10\x88\x0e\x72\x2b\x5c\x80\xb5\xb3\x35\x48\xad\xd0\x04\x96\x8b\x66\x3b\x63\x4c\xea\x26\xd1\x33\x80\x89\xaa\xc5\x06\x4b\xd8\x48\x57\x28\x3b\xed\xdd\xc8\xd1\x04\x74\x8d\x53\x1e\x07\xe2\x54\xae\x73\x6b\x1a\x87\x75\xde\xeb\xd1\x22\xa0\x0f\x9d\x36\x8f\x6e\xaf\x24\x9e\x4a\x69\x5b\x13\xde\x8a\x1a\x7b\x43\x7e\xa7\x9a\xc5\xe5\xab\x38\x27\xa7\x3b\x0e\xfb\x3e\x4a\x60\x11\x36\x1f\xa0\x93\x72\xd8\xd8\x6e\xbc\x47\xe7\x95\x35\x65\x96\x01\x4c\x26\x13\xf8\x85\x22\xdf\x0b\xdd\xa2\xcf\x36\xda\xae\x84\x2e\x13\x67\xd1\x67\xf3\xb4\x69\xa0\xb2\xb5\x50\x06\x8c\xa8\x31\x03\xca\xf7\xb5\xd3\x5c\x0d\xd1\x34\x17\xce\x06\x2b\xad\x2e\x61\x1b\x42\xe3\xb3\x6c\x02\xcb\x60\x9d\xd8\x20\x48\x2d\xbc\x87\xb5\x75\x20\xb4\x86\x86\x8c\xfb\x80\x26\x40\x0a\xd2\xc3\x37\xd6\xe8\x03\x28\x03\x52\x78\x84\x83\x6d\x1d\xf8\x23\x61\xdf\x36\x8d\x75\xc1\x83\x68\x83\xad\x45\x50\x12\xf6\x56\xb7\x35\x42\xe3\xec\x5e\x51\x34\xca\x6c\xbe\x2d\xe0\xc7\x03\xd5\x5c\xb4\x3a\x8c\x7c\xbf\x53\x5a\x43\xeb\x11\xc2\x16\x7b\xf6\xb1\x01\x69\xcd\x5a\x6d\x5a\x87\x15\xb9\xc1\x1e\x48\xdd\xfa\x80\xee\x04\xd4\x9a\x05\x8f\x05\x0e\xb6\x85\x3b\x61\x02\x04\xcb\xaa\x95\x37\xff\x47\x11\x05\x10\xbe\xb7\xd1\xa9\xda\xbd\xf4\x9d\x3a\x68\x34\x52\x90\xd1\x1d\xe5\x61\x2f\x9c\x12\x2b\x8d\xa4\xa8\x42\x2f\x9d\x5a\x21\xa8\x90\x4d\xa0\xb3\xb8\x20\x83\x25\xd4\x87\x3c\x11\x72\x76\x81\x52\xfc\x2a\x19\x32\xb6\xc2\x25\x6a\x94\xc1\x3a\xce\x74\xe7\x6c\x63\x2b\x5f\xc0\xb5\xc7\x75\xab\xfb\x04\xdb\x35\x68\x2b\x85\x4e\x39\xf4\x23\x5b\x17\xb6\x7a\x3b\xd2\x55\x12\x0b\x60\xd7\xae\xd0\x19\x0c\xe8\xa9\xb7\xb7\xd6\x07\xea\x81\xb2\x13\xca\x89\x92\xcf\xe6\x59\x46\x3d\xf3\xeb\xd5\x35\xbc\xa3\x66\x62\x47\x2a\x2b\x77\xe8\x38\xab\xab\x56\xe9\x0a\x1d\x08\x53\x81\x6b\x8d\x41\x47\xda\xeb\xd0\x96\x30\x7f\x3e\x9b\x45\xe9\x33\xb3\x57\xce\x9a\x9a\xfa\xa3\x4b\x8d\xe7\xbd\xad\xb0\xa2\x1c\x71\x0f\xd9\x8a\x9d\x46\xb3\x4f\x1e\xfe\x72\x75\x75\x71\x73\x71\x79\xfe\xdb\xfb\x12\x9e\x50\x07\x96\xd3\x69\x7d\x68\x9c\xbd\x3f\x14\xb1\x6d\x0b\x69\xeb\xf2\xe5\xec\xe5\xec\x49\x94\xa0\x45\x37\xbc\xe0\x73\x25\xc8\xc6\xf2\xcb\x8d\xf8\x2f\xb3\xf2\xf6\xbc\x37\x31\x7f\xfa\xa2\x98\x15\xb3\x62\x7e\xc2\xe5\xa2\x2c\x9f\x8c\x4a\x91\xba\xac\xf0\x7b\x79\x52\x0c\xbb\x9e\xa6\xf3\xd9\xac\xf8\xfe\x39\x8b\xce\xbf\xff\xa1\x78\xfa\xdd\xf3\x22\xfd\x9e\xc8\x75\x9e\x0a\x41\x43\xb9\x16\x8d\xa2\x41\x23\x0e\x94\x73\x1f\x89\x2d\xd3\xf8\x40\xa9\x5b\x8f\x6d\xdd\x4f\x7d\x5e\x0b\x23\x36\x49\x3a\xf6\x74\xce\x9b\xb0\x42\x17\xa5\xad\xf1\xad\x1e\x46\x79\x52\x66\x4d\xc0\xfb\x70\x24\xee\xac\x69\x59\x86\x8e\xb1\x9c\x8e\xa8\x31\x7b\x8b\xae\x46\x66\x2b\xb3\x71\xe8\x3d\xeb\x70\x96\x2e\x84\x31\x95\xd2\x9a\xaf\x84\xdc\xa1\xa9\x88\x4e\x29\xca\x95\x09\xb8\x71\x82\x8e\x7e\xa2\xd5\xd6\x6c\x6c\xb5\xa2\xa1\x11\x31\x4a\x63\x6b\x65\x79\xd4\xa8\x06\xb5\x32\x38\x36\xde\x58\x1f\x48\xfd\x2d\x47\xe2\xc4\x6a\xa5\x42\x7d\xcb\x63\xac\x94\x8f\x83\x8d\xf2\xc1\x1d\x78\xcc\xfd\x9c\x46\x41\xd5\x98\xe3\xd0\xc8\x63\xbd\xb4\x65\x90\x06\x41\xf8\x1d\xba\x7c\xa8\x67\xaa\xbf\xb1\x7d\xbf\xfc\xaf\xfe\xff\x5d\xf5\xe7\x03\x10\xef\x1b\xeb\x71\x84\x63\x28\x12\x67\x84\x86\xce\x5e\x84\x27\x29\xfa\x8c\xef\xf6\xc8\xb8\xb0\x74\xff\x3f\x7f\xfe\x8c\x88\x51\xcf\x65\x62\x95\x10\x5c\x8b\xd1\x42\x77\xd7\x29\xb3\x01\xbc\x4f\xca\x39\x41\x40\x50\x6c\x45\x57\x44\xba\x9a\x49\x51\x4a\xdd\xa5\xb5\xe1\xda\xa3\xe3\x8b\x88\x69\xb9\xa8\x6a\x65\xf2\xd6\xa3\x7b\xb0\xee\x42\x78\x7f\x67\x5d\x55\xc2\x01\x6f\xaf\xf0\xdd\xdd\xed\x3b\xf1\xc3\xed\xab\x5b\x5c\xdd\xf6\x2b\xaf\x2f\x5f\x97\x90\x64\xca\xe9\xd4\xdb\x1a\x49\x53\x59\x5f\xa9\xdb\x3b\x3c\xf5\xd5\xdd\x9f\x7a\x43\x5d\x9f\xf9\xad\x70\x55\x3e\x9b\xd1\xff\xfe\x83\xba\x9d\x17\x49\x41\x61\x30\x94\x4f\x5f\xcc\xe6\x2f\xa6\x7f\xf4\x5e\xff\x9e\x83\xed\x2c\x2d\x77\xaa\x21\xcf\x17\x0e\xb9\x1f\x52\x2e\x3a\xf6\x2b\x6c\xb4\x3d\x94\xb0\x16\xda\x23\x83\xc4\x4a\xf9\x78\xed\x54\xcc\xe2\xbb\xc8\xae\x87\x42\xc4\x5c\x75\x29\xca\x86\x12\x12\x6c\xa2\x9b\xd1\x37\x42\xe2\x31\x94\x9d\xa4\xfe\x67\xbc\x34\x21\x91\x58\x86\x08\xde\xb2\xa0\x59\xd8\xa3\x5e\x2f\xd5\xc6\x60\xd5\x3b\x29\x91\x41\x1d\xba\xe0\xa7\xde\xeb\x42\x3a\x02\x89\x3b\x3c\x74\xc4\xc6\xa9\xbd\x08\x58\xec\xf0\x90\x65\x8c\x42\x59\x55\x6f\xbc\x6b\x81\xa5\x28\xc1\x8b\xe2\x6f\x9e\xc1\xf0\x84\x10\xc8\x45\x14\x1d\x9a\x84\x73\xc0\xdc\xe6\x01\x2b\x8b\xe8\x59\x54\x15\x75\x5d\x37\xa5\x8a\x99\x04\x4e\x69\xde\x74\x85\x4f\xf3\x18\xe3\x4f\x4a\xf3\x0a\xad\x7c\x60\xd1\x1c\xa2\x9f\xd3\xf8\x93\x13\x23\x03\xe0\xee\xfc\x49\x39\x4c\x2d\x28\x1d\x86\x6c\x9d\xe6\x4b\x9e\x96\x40\x3d\x0c\x8b\xd6\x07\x5b\x83\x30\xc6\x06\xae\x69\xc4\x97\x8b\xd1\xa6\xe1\xdd\x01\x0e\xbd\x6d\x9d\x24\xac\x25\x02\xd8\x3d\x3a\xa7\xaa\x1e\x04\xfa\x6c\x32\x52\xc1\x50\xe2\x18\xea\x24\x35\x85\x8c\x10\xcc\x6c\x94\xb9\xcf\x87\xb2\x92\x2b\x17\x9f\xc2\xb3\xdd\x71\x03\xe9\x30\x84\xee\x38\xf1\x27\x40\xe7\x47\x38\x81\x78\xa4\x7c\xdb\xef\x46\xf6\x21\x9b\x30\xb0\x94\xc2\x0c\x88\xf4\x08\x06\x46\x4c\x75\x30\xa2\x7e\x08\x81\x19\x01\x83\x75\xd0\x38\xca\x94\x08\x58\x01\xde\x2b\x1f\x94\xd9\x5c\xec\x65\x44\xed\x93\x23\x92\xdf\xda\x56\xa7\x55\xb0\xc2\xb5\x75\x08\x5a\xb4\x46\x6e\x49\x15\x61\x5e\x65\xe8\x89\xc4\xae\x31\x6c\x0b\x62\x87\x9e\x2d\x60\x85\x46\x22\x67\xf4\xc8\xc1\x6c\x42\xc8\xbf\x41\xa9\xd6\x07\x86\xa5\xd4\xd3\x11\x97\xaa\x35\x64\x13\x38\xbb\x17\x75\xa3\x11\xe6\x27\xdd\xee\x87\x3b\x15\xb6\x5f\x12\x25\x15\xaa\x3b\x39\xf8\xd2\x3c\x06\xca\x12\x9b\x6d\xde\x58\xab\xf3\xf9\x98\xbb\x54\x1f\xb0\x84\x97\x3f\x2b\x76\xb1\x73\xe3\xe9\x50\x9b\x5b\x0d\xd6\x8c\xd3\xd7\xec\x25\x7b\x32\x06\xcb\xb4\x44\x72\x78\xbd\x6b\xd9\x84\xe2\x1a\xb4\x44\x9f\x46\x89\x2e\x61\x74\xa1\xf8\x5c\xef\x79\x81\xf9\x18\x67\x7f\x3e\xd0\xee\xc3\xff\x44\x78\x0f\xf3\x11\x6f\x83\x63\x6f\x92\x3c\xb1\x8e\xfd\xe0\x7d\xfb\xd9\x6e\x8c\x83\xfe\xd7\x9e\x7c\xca\x91\x41\xc3\x63\x7d\x89\x5b\xed\x23\x3f\xe6\x9f\xe7\x47\xc2\x2a\xb3\xc7\x7a\xc1\xe8\xe0\x6b\x93\xc1\xc2\x8f\xf6\x20\x61\x95\xaf\x76\x22\xc9\x3f\xba\x1e\x8c\xee\xbe\xba\x1e\x2c\xfd\x68\x1f\x06\x04\xfb\xb5\xe9\x18\xa9\x78\xac\x37\x11\xd0\x92\x54\xea\x92\xc9\x04\x96\x18\x20\xd2\x79\x3d\xf6\x57\x68\x01\xaf\x03\x28\x0f\xb5\x30\x95\x08\xd6\x1d\x18\x28\xb0\xe0\x00\xaf\xf6\xe8\x0e\x1e\x65\xeb\xb0\x13\xe3\x55\x9f\x88\xf3\xe3\x48\x3f\x11\x2b\x7b\x10\x59\x0f\xe3\xfc\xd2\x5d\x30\xc0\x86\xe3\xea\xcf\x66\x9f\xbb\x15\xa2\x86\x47\xe6\x7c\x02\xaf\x8d\x47\x17\x40\x46\xc4\xd0\x63\xe8\xa3\x6b\x17\xbe\xe1\x87\x7a\x39\x9d\x56\x56\xfa\x22\x7e\xb3\xa0\x87\xf9\xb4\x5b\x3f\x3d\x5a\x3f\xfd\x96\xcf\xeb\x8e\xb9\x48\x48\x2e\x9e\xe1\xdd\xd7\x35\x98\x15\xf3\x44\xd2\xb6\xe7\x02\x68\xdc\xa3\x2e\xa1\xc2\x55\xbb\xe9\x89\x6b\x85\xba\xf2\xc3\x22\xe8\xc0\x44\xd9\x5b\x49\xbc\x14\xe1\xb0\x54\x0a\xb9\xc5\xb1\xe4\x4a\xdb\x55\xfc\x92\xd4\x50\xc2\x40\x99\x1a\x6b\xdb\x2b\x00\xf0\xcf\xc6\xcb\xd9\x00\xf9\xfb\xfe\xfc\xfa\xf2\xe6\xf2\xec\xe7\xd7\xe7\x6f\xc7\xec\x55\x2b\x77\x04\xbc\x98\xfd\xe3\xf5\xe2\xcd\xd9\xd5\xcd\xdb\xd3\x5f\xcf\xc6\x6b\x84\x94\xe8\x3d\x83\xd2\xd3\x3f\x2f\x6f\x4e\x17\x8b\xb3\xe5\xf2\xe6\xcd\xd9\xfb\xf1\xa2\x08\xe8\xfa\x45\xcb\xb3\xc5\xe5\xd9\xd5\x68\x11\x7f\x1e\xe9\x05\x08\x69\x96\x50\x7e\x37\x9b\xcd\x7a\xda\x16\x05\xbd\x32\xc7\xde\xff\x96\x2f\x08\x58\x9b\x90\x5f\x1d\x1a\xcc\xcf\x9b\x08\xe5\xe0\x2f\xc6\x7a\xa3\xd6\xeb\xbf\x76\xca\x51\xe8\xb0\x1d\x24\x53\x1a\x2b\xa7\xf6\xe8\xc6\x0a\xd1\x10\xf6\xef\xf0\xf7\x40\x67\xf4\xbf\x17\x9a\x7a\xd8\x8f\xe8\x61\x4b\x50\xd0\xea\xaa\x84\x67\x90\x65\xe9\x6d\xcd\x2f\x81\x71\xdb\xc2\xdf\xff\xd1\x6d\x77\x7a\x1e\x42\xb0\xe0\x5a\x93\x3e\x8f\x81\xd4\x28\x0c\x3a\xa0\xbb\x3a\x92\x16\x91\xb2\x70\xdc\x49\x30\x83\xff\xa7\x3f\x52\x71\xde\x81\xd8\xee\x7b\xda\xc5\x3b\xe0\x8f\xdc\x42\x83\x57\x1f\xf0\xab\xb7\x5c\xd2\x97\xcf\xb2\x2c\x3e\x67\xff\x13\x41\x44\x4b\xff\xb6\x18\xa2\x3a\x0a\x61\x02\x2b\x21\x77\x6d\x93\xba\x65\xf2\x71\x5d\xc5\x9d\x3f\xe5\xbe\x7d\x43\x2d\xd9\x91\xe2\x3b\xe3\x21\xc3\x3f\xe3\x6f\xdd\xfe\x59\x39\x9d\xfe\x8e\xde\xaa\x79\xdc\x16\x7f\x88\xcf\xb7\x7f\x06\x00\x00\xff\xff\x3d\x91\xb8\xc4\xc7\x18\x00\x00")

func codefreshConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_codefreshConfigYaml,
		"codefresh/config.yaml",
	)
}

func codefreshConfigYaml() (*asset, error) {
	bytes, err := codefreshConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codefresh/config.yaml", size: 6343, mode: os.FileMode(420), modTime: time.Unix(1589219828, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codefreshImagesImagesList = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x6d\x72\xe2\x30\x0c\xfd\x9f\x53\x74\xa6\xbf\x6d\xf2\x01\xb4\xf8\x18\x7b\x81\x1d\x27\x11\xc6\x83\x23\x79\x65\x39\xb3\xbd\xfd\x4e\xa6\x85\xa5\x90\x52\xa6\xf9\x47\xfc\x9e\x9e\xe5\x67\x59\xe2\xf9\xf9\xe9\x17\x04\xb0\x09\x9e\x2a\x5d\xea\xdd\xae\xe8\xa8\x87\x3d\x43\x3a\xac\xba\xcc\xc1\x04\x2b\x90\xe4\x72\x75\xaf\x38\xa3\xf8\x01\x54\x17\xc0\x22\xf0\x2d\xc7\x01\x02\x5b\x01\xd5\x01\x4b\x9a\xd1\x08\xfe\x76\xd1\x3a\x40\x31\x49\x6c\x1b\xa0\x70\x1d\x6b\x4f\xab\x33\xaa\x00\x05\x38\xb2\x4f\xb0\xfa\x94\xcb\xf1\x35\xa9\x81\xd0\x0b\xb1\x59\xeb\xad\x6e\x2e\x14\xf7\x49\x51\x3c\x6f\xdf\x53\x77\x04\x36\xd5\xab\x2e\x77\xaa\xf7\xd8\x17\x1d\x61\xca\xc1\x4c\x07\x2f\x0b\xb4\x92\x4c\xa9\x77\x7a\x5d\x30\x38\x9f\x84\xdf\x4c\xad\xb7\x45\xeb\x05\xed\xe0\x57\x03\xa1\xa3\xbe\x35\x8d\xde\xea\xaa\x51\x5c\x16\x91\x92\x38\x86\x64\x76\x7a\xab\xeb\x33\x91\x6d\xdb\x7a\x19\xfe\x98\x46\xbf\xe8\x5a\x71\xf5\x1f\x81\xde\x27\xd3\xe8\x5a\xef\x14\xd7\x57\xcb\xeb\xc9\xfe\x49\xf6\x7a\xc3\x09\x68\x8a\xee\x60\x59\x86\x9c\x20\x0f\xab\x8b\xdf\x66\x2c\xf5\x46\xd7\x27\xbf\x1c\x91\x0b\xf0\xbb\x23\x14\xeb\x11\x38\xad\x7a\xd8\xdb\x1c\xa4\xb5\xdd\x11\xb0\x37\x95\x6e\x1e\xf5\x96\x09\x73\x32\xdd\x5e\x11\x46\x86\x41\x8d\x1f\x05\xf2\x58\xf8\x01\x78\x80\x9f\x87\x23\x0d\x9e\x7e\x1e\xde\xed\x55\xcb\x64\xfb\xce\x26\x01\x5e\x22\x63\xa3\x5f\x12\x2e\x21\xa9\xe4\x1d\x2e\xd1\xc8\x0b\x32\x98\x2a\x25\xa9\xc1\xa2\x75\x4b\x7c\x08\x79\xf2\x51\x45\xa6\xd1\xf7\xc0\x73\x17\xf3\xa8\x14\xa1\xc0\x5f\x59\x9c\xd2\x01\xc2\xa0\x18\x22\x2d\x56\x3a\xe6\x16\x94\x47\x01\xc7\x56\x3c\xfd\xfc\xaa\xa2\x8f\x10\x3c\xc2\x9d\x8c\x1e\x54\x3a\x35\x58\xc0\xd1\x33\xe1\x00\xb8\xdc\x30\xb1\xe9\x08\xac\xa6\xd3\x32\x82\xcc\xbe\xcd\x8b\xde\x09\xe8\x3c\x82\xa9\x74\x55\x6e\x74\x79\x0b\x5c\x07\x37\xea\x20\x12\x3f\xcf\x89\xf7\x8e\xab\x62\x4e\x87\xd9\xc4\xe7\xc9\x21\x3c\x4c\x6e\xb3\x0f\xfd\xf7\xec\x73\x3b\x54\x81\xdc\xbc\x87\x9f\xf8\xce\x8b\xea\x02\xe1\xfc\x15\xde\x39\xe2\xb8\x51\x29\x56\xf7\x19\xb3\xa8\x58\x77\x66\xd4\xf7\x7c\x19\x5f\xee\xa3\x5f\xee\x7f\xb2\x6a\xac\xb6\xdf\xe1\x2a\xc5\x6f\xfc\x2b\xa7\x91\xb4\xfd\xd2\xb4\xf1\x6b\x3f\xc7\xdd\x95\x3a\x0d\x91\xd2\x54\x67\xf5\x46\x97\x37\x3b\xf7\x10\x03\xbd\x5d\x16\x6d\x8a\xea\xfa\x80\xd7\x9c\x0b\x78\x1a\xf3\xef\x13\x5f\x6f\xd4\x58\xaf\x8b\xcb\xd7\x7a\xa4\x40\xa7\x3f\x07\x85\x0d\x71\x2a\xeb\x8f\xcf\xdc\x66\x94\x7c\xfa\xfa\x17\x00\x00\xff\xff\x1d\x1a\x12\x25\x25\x09\x00\x00")

func codefreshImagesImagesListBytes() ([]byte, error) {
	return bindataRead(
		_codefreshImagesImagesList,
		"codefresh/images/images-list",
	)
}

func codefreshImagesImagesList() (*asset, error) {
	bytes, err := codefreshImagesImagesListBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codefresh/images/images-list", size: 2341, mode: os.FileMode(420), modTime: time.Unix(1589219708, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _codefreshImagesImagesMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcb\xb1\x0d\xc2\x40\x0c\x05\xd0\xfe\xa6\xf8\x52\x6a\x58\x22\x34\x8c\x61\xee\x7e\x72\x86\x60\x23\xdb\x12\x62\x7b\x2a\x24\x06\x78\xcb\xb2\xe0\xe2\xfd\xc1\xc0\xf5\x29\x3b\x13\x87\x66\x41\x6c\xa0\x4b\x9f\x6c\x6d\x75\x2b\xe9\x85\xd5\x07\xb7\x60\x4e\x94\xc3\x6f\x25\x6a\x48\x39\xdf\xd3\x0d\x6f\xad\x89\x1e\x1c\xb4\x52\x39\x12\x9b\xc7\x1f\xa0\x15\xe3\x15\x9a\x44\x70\xd7\xac\xf8\xb4\x76\xfa\xe9\x6f\x00\x00\x00\xff\xff\x60\xca\xf8\xb0\x83\x00\x00\x00")

func codefreshImagesImagesMdBytes() ([]byte, error) {
	return bindataRead(
		_codefreshImagesImagesMd,
		"codefresh/images/images.md",
	)
}

func codefreshImagesImagesMd() (*asset, error) {
	bytes, err := codefreshImagesImagesMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "codefresh/images/images.md", size: 131, mode: os.FileMode(420), modTime: time.Unix(1589219544, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _k8sAgentConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4d\x2d\x49\x4c\x49\x2c\x49\xb4\xe2\x52\x50\xc8\xce\xcc\x4b\xb1\x52\xc8\xb6\x28\xd6\x4d\x4c\x4f\xcd\x2b\x01\x04\x00\x00\xff\xff\x73\xdb\xf5\x15\x1b\x00\x00\x00")

func k8sAgentConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_k8sAgentConfigYaml,
		"k8s-agent/config.yaml",
	)
}

func k8sAgentConfigYaml() (*asset, error) {
	bytes, err := k8sAgentConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s-agent/config.yaml", size: 27, mode: os.FileMode(420), modTime: time.Unix(1589212305, 0)}
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
	"codefresh/assets/assets.md":   codefreshAssetsAssetsMd,
	"codefresh/certs/tls.md":       codefreshCertsTlsMd,
	"codefresh/config.yaml":        codefreshConfigYaml,
	"codefresh/images/images-list": codefreshImagesImagesList,
	"codefresh/images/images.md":   codefreshImagesImagesMd,
	"k8s-agent/config.yaml":        k8sAgentConfigYaml,
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
	"codefresh": &bintree{nil, map[string]*bintree{
		"assets": &bintree{nil, map[string]*bintree{
			"assets.md": &bintree{codefreshAssetsAssetsMd, map[string]*bintree{}},
		}},
		"certs": &bintree{nil, map[string]*bintree{
			"tls.md": &bintree{codefreshCertsTlsMd, map[string]*bintree{}},
		}},
		"config.yaml": &bintree{codefreshConfigYaml, map[string]*bintree{}},
		"images": &bintree{nil, map[string]*bintree{
			"images-list": &bintree{codefreshImagesImagesList, map[string]*bintree{}},
			"images.md":   &bintree{codefreshImagesImagesMd, map[string]*bintree{}},
		}},
	}},
	"k8s-agent": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{k8sAgentConfigYaml, map[string]*bintree{}},
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
