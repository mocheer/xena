package tileset

type Asset struct {
	Version        string `json:"version"`                  //表示tileset使用的3D Tiles规范的版本。
	TilesetVersion string `json:"tilesetVersion,omitempty"` //表示当前数据集的版本，用于控制缓存和更新。这个是指
	GltfUpAxis     string `json:"gltfUpAxis,omitempty"`     //表示GLTF模型的上轴（up-axis），通常是Z。
	Generator      string `json:"generator,omitempty"`      //不在规范之中
}
