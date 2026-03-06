package tianditu

import (
	"github.com/mocheer/xena/pkg/provider"
)

func NewNormal() *TiandituMap {
	return New(provider.Provider{
		URL:        URL,
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
			Subdomains: subdomains,
			Vars: map[string]string{
				"v": EPSG4490_TYPES["Satellite_A"],
			},
		},
	)
}
