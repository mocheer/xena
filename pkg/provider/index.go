package provider

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/gm"
)

type Provider struct {
	URL        string
	Subdomains []string
	Vars       map[string]string
	Loader     *ctp.Ctp
	Tokens     []string
}

func (m Provider) GetTileURL(t *gm.Tile) string {
	data := map[string]any{"z": t.Z, "x": t.X, "y": t.Y, "s": m.GetRandSubdomains(), "t": m.GetRandToken()}
	for k, v := range m.Vars {
		data[k] = v
	}
	return fn.FormatByMap(m.URL, data)
}

func (m *Provider) LoadTile(t *gm.Tile) ([]byte, error) {
	url := m.GetTileURL(t)
	if m.Loader == nil {
		m.Loader = ctp.New()
		m.Loader.WithTransport(&http.Transport{
			Proxy: http.ProxyFromEnvironment, // 代理
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second, // 超时时间
				KeepAlive: 3 * time.Second, // KeepAlive 超时时间
				DualStack: true,            // 是否开启双协议栈
			}).DialContext,
			MaxIdleConns:          100,              // 最大空闲连接数
			IdleConnTimeout:       90 * time.Second, // 空闲连接超时时间
			TLSHandshakeTimeout:   10 * time.Second, // TLS握手超时时间
			ExpectContinueTimeout: 1 * time.Second,  // 如果非零，如果指定请求包含“Expect: 100-continue”报头，则在写完请求报头后等待服务器第一个响应报头的时间。0表示没有超时，并立即发送正文，而无需等待服务器批准。这个时间不包括发送请求头的时间。
		})
	}
	data, err := m.Loader.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("reptile '%s' error: %s", url, err))
	}
	return data, nil
}

func (m Provider) GetRandSubdomains() string {
	s := ""
	if m.Subdomains != nil {
		s = m.Subdomains[rand.Intn(len(m.Subdomains))]
	}
	return s
}

func (m Provider) GetRandToken() string {
	s := ""
	if m.Tokens != nil {
		s = m.Tokens[rand.Intn(len(m.Tokens))]
	}
	return s
}
