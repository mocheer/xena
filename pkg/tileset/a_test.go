package tileset

import (
	"path"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	t.Log(path.Dir("http://domain/path?a=b/c"))
	t.Log(path.Ext("http://domain/path?a=b/c"))
}

func Test2(t *testing.T) {
	err := Load("http://112.81.89.197:9900/data/3dtiles/tileset.json", "data/3dtiles-test/tileset.json")
	panic(err)
}

func Test3(t *testing.T) {
	children := []int{4, 6, 7, 8, 9, 3}
	slices.SortFunc(children, func(a, b int) int {
		return a - b
	})
	t.Log(children)
}
