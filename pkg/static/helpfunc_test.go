package static

import (
	"net"
	"testing"
)

func TestHF_GetHttpIP(t *testing.T) {
	ip := "2408:824e:8904:4ea1:6af8:9655:7710:a7b6"
	// 统一去掉端口（兼容 IPv6 中括号格式）
	if host, _, err := net.SplitHostPort(ip); err == nil {
		t.Logf("host: %v", host)
		return
	}
	t.Logf("ip: %v", ip)
}
