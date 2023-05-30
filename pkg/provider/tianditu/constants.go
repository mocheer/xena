package tianditu

import "strings"

const URL string = "http://t{s}.tianditu.gov.cn/DataServer?T={v}&x={x}&y={y}&l={z}&tk=070a93160eddd5f891599e51a6b764ac"

var tokens = []string{
	"070a93160eddd5f891599e51a6b764ac",
	"60714a8ef3e4491df43827cd34c2aa22",
}

var subdomains = strings.Split("01234567", "")

var EPSG3857_TYPES = map[string]string{
	"Normal":      "vec_w", //矢量地图
	"Normal_A":    "cva_w", //矢量-标注图
	"Terrain":     "ter_w", //地形图
	"Terrain_A":   "cta_w", //地形-标注图
	"Satellite":   "img_w", //卫星影像图
	"Satellite_A": "cia_w", //卫星-标注图
}

var EPSG4490_TYPES = map[string]string{
	"Normal":      "vec_c", //矢量地图
	"Normal_A":    "cva_c", //矢量-标注图
	"Terrain":     "ter_c", //地形图
	"Terrain_A":   "cta_c", //地形-标注图
	"Satellite":   "img_c", //卫星影像图
	"Satellite_A": "cia_c", //卫星-标注图
}
