package gtif_test

import (
	"fmt"
	"testing"

	"github.com/mocheer/xena/alg/proj4"
	"github.com/mocheer/xena/gfs/gtif"
)

func TestTiff(t *testing.T) {
	fileName := "./testdata/SRTM3_V4_90m.tif" // GCS_WGS_1984
	tfs := gtif.Read(fileName)

	fmt.Println(tfs.Tif.IFDs()[0])
	// 107.85853701103325 34.12743138973333 107.34437034436667 33.849098056400045
	fmt.Println(tfs.BBox())
	// 100,100
	fmt.Println(tfs.GetColRow(107.42770367769998, 34.04409805640001))
	// 107.42770367769998, 34.04409805640001
	fmt.Println(tfs.GetLonLat(100, 100))
	// 1757
	fmt.Println(tfs.GetAlt(100, 100))
}

func TestBigTiff(t *testing.T) {
	// UTM 投影：经度采用6度分带，纬度采用8度分带，从80S到84N共20个纬度带（X带多4度），分别用C到X的字母来表示
	fileName := "./testdata/utm/54R_20200101-20210101.tif" //这是一个bigtiff格式的文件
	tfs := gtif.Read(fileName)

	tfs.IFD_Index = 0
	fmt.Println(len(tfs.Tif.IFDs()))
	fmt.Println(tfs.Origin())
	fmt.Println(tfs.BBox())
	fmt.Println(tfs.Scale())
	//
	fmt.Println(tfs.GetLonLat(100, 100))
	xy := tfs.GetLonLat(100, 100)
	fmt.Println(proj4.UTM_WGS84_ZONE(54).ConvertToWGS84([]float64{xy[0], xy[1]}))
}

func TestTiff2(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N53E125/ASTGTMV003_N53E125_dem.tif" // lzw 压缩
	tfs := gtif.Read(fileName)

	fmt.Println(len(tfs.Tif.IFDs()))
	fmt.Println(tfs.Width())
	fmt.Println(tfs.Height())
	fmt.Println(tfs.Origin())
	fmt.Println(tfs.BBox())
	fmt.Println(tfs.Scale())
	//
	fmt.Println(tfs.GetLonLat(100, 3600))
	fmt.Println(tfs.GetAlt(2, 0))
	fmt.Println(tfs.GetAltByLonLat(125, 53))
}

func TestTiff3(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N34E107/ASTGTMV003_N34E107_dem.tif"
	tfs := gtif.Read(fileName)

	fmt.Println(len(tfs.Tif.IFDs()))
	fmt.Println(tfs.Width())
	fmt.Println(tfs.Height())
	fmt.Println(tfs.Origin())
	fmt.Println(tfs.BBox())
	fmt.Println(tfs.Scale())
	//
	fmt.Println(tfs.GetLonLat(100, 3600))
	fmt.Println(tfs.GetAlt(2, 0))
	fmt.Println(tfs.GetAltByLonLat(107.42770367769998, 34.04409805640001))
}

func TestTiff4(t *testing.T) {
	fileName := "./testdata/ASTGTMV003_N43E093/ASTGTMV003_N43E093_dem.tif"
	tfs := gtif.Read(fileName)

	fmt.Println(len(tfs.Tif.IFDs()))
	fmt.Println(tfs.Width())
	fmt.Println(tfs.Height())
	fmt.Println(tfs.Origin())
	fmt.Println(tfs.BBox())
	fmt.Println(tfs.Scale())
	fmt.Println(tfs.Scale()[0] * 111110) //30.86 说明这个分辨率只有30米
	fmt.Println(tfs.GetColRow(93.878768, 43.074783))
	fmt.Println(tfs.GetAltByLonLat(93.878768, 43.074783)) // 1458 => 1455.0753793690708  Cesium.sampleTerrainMostDetailed(terrainProvider,[Cartographic.fromDegrees()])
}
