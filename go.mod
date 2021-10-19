module github.com/mocheer/xena

go 1.17

require (
	github.com/go-spatial/proj v0.2.0
	github.com/google/tiff v0.0.0-20161109161721-4b31f3041d9a
	github.com/liuvigongzuoshi/go-kriging v0.0.1-alpha.14
	github.com/mocheer/pluto v1.0.0
	github.com/paulmach/go.geojson v1.4.0
	github.com/rubenv/topojson v0.0.0-20180822134236-13be738db397
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/imroc/req v0.3.0 // indirect
	github.com/paulmach/go.geo v0.0.0-20180829195134-22b514266d33 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tidwall/gjson v1.8.1 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.1.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	gonum.org/v1/gonum v0.8.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace (
	github.com/mocheer/pluto => ../pluto
	github.com/mocheer/vesta => ../vesta
	github.com/mocheer/xena => ../xena
)
