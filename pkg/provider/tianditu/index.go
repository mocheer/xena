package tianditu

import (
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/provider"
)

type TiandituMap struct {
	provider.Provider
}

func New(config provider.Provider) *TiandituMap {
	config.Tokens = tokens
	return &TiandituMap{Provider: config}
}

func (m *TiandituMap) LoadTile(t *gm.Tile) ([]byte, error) {
	var retry func(index int) ([]byte, error)
	retry = func(index int) ([]byte, error) {
		data, err := m.Provider.LoadTile(t)
		if err != nil {
			index++
			if index < 7 {
				return retry(index)
			}
		}
		return data, err
	}
	return retry(0)
}

func retry(index int) {
	panic("unimplemented")
}

func NewNormal() *TiandituMap {
	return New(provider.Provider{
		URL:        URL,
		Loader:     ctp.New(),
		Subdomains: subdomains,
		Vars: map[string]string{
			"v": EPSG3857_TYPES["Normal"],
		},
	},
	)
}

func NewNormalA() *TiandituMap {
	return New(
		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG3857_TYPES["Normal_A"],
			},
		},
	)
}

func NewSatellite() *TiandituMap {
	return New(

		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG3857_TYPES["Satellite"],
			},
		},
	)
}

func NewSatelliteA() *TiandituMap {
	return New(
		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG3857_TYPES["Satellite_A"],
			},
		},
	)
}

func NewNormal4326() *TiandituMap {
	return New(
		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG4490_TYPES["Normal"],
			},
		},
	)
}

func NewNormalA4326() *TiandituMap {
	return New(
		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG4490_TYPES["Normal_A"],
			},
		},
	)
}

func NewSatellite4326() *TiandituMap {
	return New(
		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG4490_TYPES["Satellite"],
			},
		},
	)

}
func NewSatelliteA4326() *TiandituMap {
	return New(
		provider.Provider{
			URL:        URL,
			Loader:     ctp.New(),
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG4490_TYPES["Satellite_A"],
			},
		},
	)
}
