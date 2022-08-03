package gm_test

import (
	"strconv"
	"testing"
)

func TestXYZ1(t *testing.T) {
	var x uint64 = 54081
	var y uint64 = 26909
	var z uint64 = 16
	id := (x << 32) + (y << 5) + z
	t.Log(id) //232276127196080
	t.Log(strconv.FormatUint(id, 16))
	t.Log(strconv.FormatUint(id, 32))
}

func TestXYZ2(t *testing.T) {
	var xyz uint64 = 232276127196080
	x := xyz >> 32
	t.Log(x)
	y := (xyz << 32 >> 32) >> 5
	t.Log(y)
	z := xyz << 59 >> 59
	t.Log(z)
}
