package filter

import (
	"net"
	"net/http"

	"github.com/Walker-PI/iot-gateway/gateway/agw_context"
)

type IPBlackFilter struct {
	baseFilter
}

func newIPBlackFilter() Filter {
	return &IPBlackFilter{}
}

func (f *IPBlackFilter) Name() FilterName {
	return PreIPBlackFilter
}

func (f *IPBlackFilter) Type() FilterType {
	return PreFilter
}

func (f *IPBlackFilter) Priority() int {
	return priority[PreIPBlackFilter]
}

func (f *IPBlackFilter) Run(ctx *agw_context.AGWContext) (Code int, err error) {

	if len(ctx.RouteInfo.IPBlackList) == 0 {
		return f.baseFilter.Run(ctx)
	}

	realIP := ctx.GetString("Real-IP")
	if realIP == "" {
		return http.StatusForbidden, nil
	}

	netIP := net.ParseIP(realIP)
	if netIP == nil {
		return http.StatusForbidden, nil
	}

	for _, blackIP := range ctx.RouteInfo.IPBlackList {
		if net.IP.Equal(blackIP, netIP) {
			return http.StatusForbidden, nil
		}
	}
	return f.baseFilter.Run(ctx)
}
