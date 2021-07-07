module github.com/mocheer/xena

go 1.16

require (
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/go-spatial/proj v0.2.0
	github.com/google/tiff v0.0.0-20161109161721-4b31f3041d9a
	github.com/liuvigongzuoshi/go-kriging v0.0.1-alpha.14
	github.com/mocheer/pluto v1.0.0
	github.com/paulmach/go.geo v0.0.0-20180829195134-22b514266d33 // indirect
	github.com/paulmach/go.geojson v1.4.0
	github.com/rubenv/topojson v0.0.0-20180822134236-13be738db397
	github.com/stretchr/testify v1.4.0
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect

)

replace (
	github.com/mocheer/pluto => ../pluto
	github.com/mocheer/vesta => ../vesta
	github.com/mocheer/xena => ../xena
)
