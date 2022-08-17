package tile

import (
	"fmt"
	"math"
	"math/rand"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/console"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/ts/awt"
	"github.com/mocheer/pluto/pkg/web_request"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/tile_arcgis"
)

type LoadConfig struct {
	URL        string
	DirName    string
	MinZoom    int
	MaxZoom    int
	Origin     string
	Subdomains []string
	SavePath   string
}

// LoadArcgis 下载arcgis本地图包中的所有瓦片数据
func LoadArcgis(localPath, dirName string) error {
	service, err := tile_arcgis.NewTileArcgis(filepath.Join(localPath, "conf.xml"))
	if err != nil {
		return err
	}

	for z := 0; z < 18; z++ {
		size := int(math.Pow(2, float64(z)))
		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				zoomRowColumn := fmt.Sprintf("%d/%d/%d", z, x, y)
				dirPath := filepath.Join(dirName, zoomRowColumn)
				data, _ := service.ReadTile(gm.Tile{
					Z: z, X: x, Y: y,
				})
				if len(data) != 0 {
					if !ds.IsExist(dirPath) {
						ds.Save(dirPath, data)
						msg := fmt.Sprintf("下载瓦片成功：%s", zoomRowColumn)
						fmt.Println(msg)
					}
				} else {
					msg := fmt.Sprintf("瓦片为空：%s", zoomRowColumn)
					fmt.Println(msg)
				}
			}
		}
	}

	return nil
}

// Load 下载瓦片服务中的所有瓦片数据
func Load(lc *LoadConfig) error {

	minZoom := lc.MinZoom
	maxZoom := lc.MaxZoom
	for i := minZoom; i < maxZoom; i++ {
		loadingByZoom(lc, i)
	}
	return nil
}

// loadingByZoom 下载tile
func loadingByZoom(lc *LoadConfig, z int) error {
	size := int(math.Pow(2, float64(z)))
	at := awt.New(10, 1000)
	//
	startX := 0
	startY := 0
	endX := size
	endY := size
	//
	if z >= 10 {
		p1 := gm.LonLat{73, 54}
		p2 := gm.LonLat{136, 3}
		t1, _ := p1.GetTileAndOffset(float64(z))
		t2, _ := p2.GetTileAndOffset(float64(z))
		startX = t1.X
		startY = t1.Y
		endX = t2.X
		endY = t2.Y
	}

	for x := startX; x < endX; x++ {
		for y := startY; y < endY; y++ {
			tile := &gm.Tile{X: x, Y: y, Z: z}
			at.Add(func(args ...any) {
				tile := args[0].(*gm.Tile)
				loadingByXYZ(lc, tile)
			}, tile)
		}
	}
	at.Wait()
	return nil
}

func loadingByXYZ(lc *LoadConfig, tile *gm.Tile) error {
	rootUrl := lc.URL
	dirName := lc.DirName
	origin := lc.Origin

	x, y, z := tile.X, tile.Y, tile.Z
	savePath := lc.SavePath

	if savePath == "" {
		savePath = "{z}/{x}/{y}.png"
	}
	fpath := fn.FmtString(savePath, map[string]any{"z": z, "x": x, "y": y})
	s := ""
	if lc.Subdomains != nil {
		s = lc.Subdomains[rand.Intn(len(lc.Subdomains))]
	}
	url := fn.FmtString(rootUrl, map[string]any{"z": z, "x": x, "y": y, "s": s})
	dirPath := filepath.Join(dirName, fpath)
	if !ds.IsExist(dirPath) {

		err := web_request.Save(url, dirPath, origin)
		if err == nil {
			msg := fmt.Sprintf("下载瓦片成功：%s", url)
			console.Log(msg)
			// 休眠一小会，防止被认为是爬虫抓取数据
			// time.Sleep(time.Millisecond * 100)
		} else {
			msg := fmt.Sprintf("下载瓦片失败：%s , %s ", url, err)
			console.Warn(msg)
		}
	}
	return nil
}
