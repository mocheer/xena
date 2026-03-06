package graph

import (
	"log"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/qmuntal/gltf"
)

// Save
func (m *Graph) Save(filename string) error {
	ds.CreateDirFromFilename(filename)
	if ds.IsExist(filename) {
		log.Println("File already exists:", filename)
	}
	return gltf.Save(m.Document, filename)
}

// SaveBinary
func (m *Graph) SaveBinary(filename string) error {
	ds.CreateDirFromFilename(filename)
	if ds.IsExist(filename) {
		log.Println("File already exists:", filename)
	}
	return gltf.SaveBinary(m.Document, filename)
}
