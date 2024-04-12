package tianditu

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/provider"
)

type TiandituMap struct {
	provider.Provider
}

func New(config provider.Provider) *TiandituMap {
	config.Tokens = tokens
	config.Loader = createLoader()
	return &TiandituMap{Provider: config}
}

func (m *TiandituMap) LoadTile(t *gm.Tile) ([]byte, error) {
	return m.retry(0, t)
}

func (m *TiandituMap) retry(index int, t *gm.Tile) ([]byte, error) {
	data, err := m.Provider.LoadTile(t)
	if err != nil {
		// fmt.Println(index, err)
		// time.Sleep(time.Millisecond * 500)
		if index < 7 {
			return m.retry(index+1, t)
		}
		fmt.Println(index, err)
	}
	return data, err
}

func createLoader() *ctp.Ctp {
	loader := ctp.New()
	loader.WithTransport(&http.Transport{
		// 从环境变量中获取代理设置。简化代理配置的读取过程，使得Go程序能够根据环境变量中的代理设置来发送HTTP请求。如果没有配置 HTTP_PROXY、HTTPS_PROXY 等相关变量时，相当于没有代理
		// Proxy: http.ProxyFromEnvironment,
		//
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second, // 超时时间
			KeepAlive: 3 * time.Second, // KeepAlive 超时时间
			DualStack: true,            // 是否开启双协议栈(IPv4/IPv6)
		}).DialContext,
		MaxIdleConns:          10,               // 最大空闲连接数
		IdleConnTimeout:       30 * time.Second, // 空闲连接超时时间
		TLSHandshakeTimeout:   10 * time.Second, // TLS握手超时时间
		ExpectContinueTimeout: 3 * time.Second,  // 如果非零，如果指定请求包含“Expect: 100-continue”报头，则在写完请求报头后等待服务器第一个响应报头的时间。0表示没有超时，并立即发送正文，而无需等待服务器批准。这个时间不包括发送请求头的时间。
	})
	return loader
}
