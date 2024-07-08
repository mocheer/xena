package tileset

import (
	"path"
	"testing"
)

func Test(t *testing.T) {
	t.Log(path.Dir("http://domain/path?a=b/c"))
	t.Log(path.Ext("http://domain/path?a=b/c"))
}
