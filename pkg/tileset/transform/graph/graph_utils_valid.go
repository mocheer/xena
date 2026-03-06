package graph

import "log"

// Valid
func (m *Graph) Valid() error {
	m.EachUsedNodes(func(node *GraphNode) {
		node.EachUsedMesh(func(mesh *GraphMesh) {
			for _, p := range mesh.GetPrimitives() {
				if err := p.Valid(); err != nil {
					log.Println(err)
				}
			}
		})
	})
	return nil
}
