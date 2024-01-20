package tile

import (
	"math/rand"

	"github.com/mocheer/xena/pkg/gm"
)

type LoadConfig struct {
	URL        string
	DirName    string
	MinZoom    int
	MaxZoom    int
	Bbox       gm.BBox
	Origin     string
	Subdomains []string
	SavePath   string
	FiberCount int
}

func (m LoadConfig) GetRandSubdomains() string {
	s := ""
	if m.Subdomains != nil {
		s = m.Subdomains[rand.Intn(len(m.Subdomains))]
	}
	return s
}
