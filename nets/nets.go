package nets

import (
	"net/http"
	"strconv"
	"strings"
)

type T struct {
	v string
}

func (t *T) String() string {
	return t.v
}

func (t *T) Int64() int64 {
	bits := strings.Split(t.v, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// IP 获取IP地址
func IP(req *http.Request) *T {
	if len(req.Header.Get("x-forwarded-for")) > 0 {
		// 代理转发ip 格式：180.158.93.171,128.18.31.52
		ipList := strings.Split(req.Header.Get("x-forwarded-for"), ",")
		if len(ipList) > 0 {
			return &T{strings.TrimSpace(ipList[0])}
		}
		// 防止被突破
	} else if len(req.Header.Get("X-App-Real-IP")) > 0 {
		// 转发过程强制设置的一个变量
		return &T{req.Header.Get("X-App-Real-IP")}
	} else if len(req.Header.Get("X-Real-Ip")) > 0 {
		// 转发过程强制设置的一个变量
		return &T{req.Header.Get("X-Real-Ip")}
	}
	// "IP:port" "192.168.1.150:8889"
	return &T{strings.Split(req.RemoteAddr, ":")[0]}
}

func IpInt(ip string) *T {
	return &T{ip}
}
