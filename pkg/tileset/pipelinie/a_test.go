package pipelinie_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/tileset"
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

func Test5(t *testing.T) {
	// 0.json
	b1 := &tileset.BoundingVolume{
		Box: &[12]float64{
			5.686390032293275,
			0.026497788494452834,
			-0.4193151416257024,
			4774.894325314407,
			0,
			0,
			0,
			4953.960673033112,
			0,
			0,
			0,
			106.17404315582553,
		},
	}
	// 0-1202.json
	b2 := &tileset.BoundingVolume{
		Box: &[12]float64{
			-2356.47313700011,
			-3689.311824395787,
			-2.4047730891034007,
			911.635355406991,
			0,
			0,
			0,
			1264.6223508488301,
			0,
			0,
			0,
			46.63389648645261,
		},
	}

	// 0-1313-2020.json
	b3 := &tileset.BoundingVolume{
		Box: &[12]float64{
			920.3539557913316,
			3148.1310253855772,
			-0.16326295724138618,
			563.6424102197316,
			0,
			0,
			0,
			605.2658713092942,
			0,
			0,
			0,
			53.62882183780039,
		},
	}
	r1 := pipelinie.IsBoundingBoxesOverlap(b1, b2, 0)
	r2 := pipelinie.IsBoundingBoxesOverlap(b1, b3, 0)
	r3 := pipelinie.IsBoundingBoxesOverlap(b2, b3, 0)
	t.Log(r1, r2, r3)
}
