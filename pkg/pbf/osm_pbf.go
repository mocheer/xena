package pbf

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/qedus/osmpbf"
)

// ReadOSM 专门用来读取osm的pbf文件
// tag  name:zh highway oneway lanes
func ReadOSM(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)

	// start decoding with several goroutines, it is faster
	err = d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}

	var nc, wc, rc uint64
	var nodeList []*osmpbf.Node
	var wayList []*osmpbf.Way
	var relationList []*osmpbf.Relation
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			// 点
			case *osmpbf.Node:
				nodeList = append(nodeList, v)
				nc++
			// 非闭合线，闭合线，区域
			case *osmpbf.Way:
				wayList = append(wayList, v)
				wc++
			case *osmpbf.Relation:
				relationList = append(relationList, v)
				// Process Relation v.
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)

	data, _ := json.Marshal(wayList[0:20])
	fmt.Println(string(data))

	data, _ = json.Marshal(nodeList[0:20])
	fmt.Println(string(data))

	data, _ = json.Marshal(relationList[0:20])
	fmt.Println(string(data))

}
