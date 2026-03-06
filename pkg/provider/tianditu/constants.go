package tianditu

import "strings"

const URL string = "http://t{s}.tianditu.gov.cn/DataServer?T={v}&x={x}&y={y}&l={z}&tk={t}"

var tokens = []string{
	"070a93160eddd5f891599e51a6b764ac",
	"60714a8ef3e4491df43827cd34c2aa22",
	"2c6c6c00cbb9f83e135243a4bf30162c",
	"34568012b0e7be57119fa5124bd7bdd6",
	"5d2a5431b4d468c54a025de39f039190",
	"ee1087d5e54c8e3933bd38ebdb9d8ad1",
	"171a747a409509e9ce89fb59845e09de",
	// "16554181e1d8f9f3b82ce84fe953c164",
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
