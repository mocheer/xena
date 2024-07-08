package pipelinie_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/tileset/pipelinie"
	"github.com/mocheer/xena/pkg/tileset/pipelinie/option"
)

func TestXxx(t *testing.T) {
	pipelinie.From("./testdata/3dtiles-20230427112136/tileset.json").Transform(
		&option.UpgradeOption{Version: "1.1"},
		&option.SaveOption{Path: "./testdata/3dtiles_gc_wz"},
	)
}

func Test2(t *testing.T) {
	pipelinie.From("./testdata/1.0/tileset.json").Transform(
		&option.UpgradeOption{Version: "1.1"},
		&option.SaveOption{Path: "./testdata/1.1"},
	)
}

func Test3(t *testing.T) {
	matrix := pipelinie.CalcEnuToEcefMatrix(120, 30, 10)
	t.Log(matrix)
}

func Test4(t *testing.T) {
	pipelinie.From("./testdata/3dtiles_gc_wz/tileset.json").Rebuild("testdata/top")
}
