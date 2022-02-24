package pbf

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/qedus/osmpbf"
)

type P_Result struct {
	Nodes     []*osmpbf.Node
	Ways      []*osmpbf.Way
	Relations []*osmpbf.Relation
}

// ReadOSM 专门用来读取osm的pbf文件
// tag  name:zh highway oneway lanes
func ReadOSM(fileName string) *P_Result {
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
			// 非闭合线，闭合线，区域
			case *osmpbf.Way:
				wayList = append(wayList, v)
			// 关系
			case *osmpbf.Relation:
				relationList = append(relationList, v)
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	return &P_Result{
		Nodes:     nodeList,
		Ways:      wayList,
		Relations: relationList,
	}
}

// ReadOSMNode 专门用来读取osm的pbf文件
// tag  name:zh highway oneway lanes
func ReadOSMNode(fileName string, callback func(interface{})) {
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

	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			callback(v)
		}
	}
}
