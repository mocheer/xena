package amap

import (
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/provider"
)

type AMap struct {
	provider.Provider
}

func New(typeName string) {

}

func NewNormal() *AMap {
	return &AMap{
		Provider: provider.Provider{
			URL:        "http://webrd0{s}.is.autonavi.com/appmaptile?lang=zh_cn&style=8&x={x}&y={y}&z={z}",
			Loader:     ctp.New(),
			Subdomains: subdomains,
		},
	}
}

func NewSatellite() *AMap {
	return &AMap{
		Provider: provider.Provider{
			URL:        "http://wprd0{s}.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=2&style=6",
			Loader:     ctp.New(),
			Subdomains: subdomains,
		},
	}
}
func NewSatelliteA() *AMap {
	return &AMap{
		Provider: provider.Provider{
			URL:        "http://wprd0{s}.is.autonavi.com/appmaptile?x={x}&y={y}&z={z}&lang=zh_cn&size=1&scl=1&style=8",
			Loader:     ctp.New(),
			Subdomains: subdomains,
		},
	}
}
