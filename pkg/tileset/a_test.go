package tileset

import (
	"testing"
)

func Test2(t *testing.T) {
	err := Load("http://112.81.89.197:9900/data/3dtiles/tileset.json", "data/3dtiles-test/tileset.json")
	panic(err)
}
