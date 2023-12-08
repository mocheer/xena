package tile_arcgis

import "github.com/mocheer/xena/pkg/gm"

// TileArcgis implements TileCache for ESRI local files
// @see https://github.com/wthorp/AGES/tree/master/pkg/sources/tilecache  ==> 这个源码一堆错误，容错率低
// @see https://github.com/fuzhenn/tiler-arcgis-bundle/blob/master/index.js
type TileArcgis struct {
	BaseDirectory string // 根目录
	CacheFormat   string // 切片缓存模式
	TileFormat    string // 切片数据类型
	TileColSize   int    // 切片 column 大小
	TileRowSize   int    // 切片 row 大小
	ColsPerFile   int    //
	RowsPerFile   int    //
	WKID          int    // 切片服务wkid
	Bbox          gm.BBox
	BboxC         gm.BBox // 不一定准确，最大不一定是最大，最小不一定是最小
}
