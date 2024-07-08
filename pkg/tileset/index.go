package tileset

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/mocheer/pluto/pkg/ts/ctp"
)

func New() *Tileset {
	return &Tileset{Asset: Asset{Generator: "charon"}}
}

// FromPath
// FromPath=FromURL+FromFile 这里根据url的前缀识别是网络路径还是本地路径
func FromPath(url string) (*Tileset, error) {
	if strings.HasPrefix(url, "http") {
		return FromURL(url)
	} else {
		return FromFile(url)
	}
}

// FromURL
// url路径有可能是"http://domain/tileset.json?a=b/c" 这个时候的相对路径会出现错误
func FromURL(url string) (*Tileset, error) {
	data, err := ctp.Get(url)
	if err == nil {
		t := FromBytes(data)
		return t, nil
	}
	return nil, err
}

// FromFile
func FromFile(filename string) (*Tileset, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := FromReader(f)
	return t, err
}

// FromBytes 从bytes数据中实例化 Tileset 对象
func FromBytes(bs []byte) *Tileset {
	return FromReader(bytes.NewBuffer(bs))
}

// FromReader
func FromReader(reader io.Reader) *Tileset {
	var ts *Tileset
	json.NewDecoder(reader).Decode(&ts)
	return ts
}
