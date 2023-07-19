package tile_arcgis_test

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	o := int64(binary.LittleEndian.Uint64([]byte{60, 0, 0, 0, 0}))
	fmt.Println(o)
}
