package router_test

import (
	"testing"

	"v2ray.com/core/app"
	"v2ray.com/core/app/dispatcher"
	dispatchers "v2ray.com/core/app/dispatcher/impl"
	"v2ray.com/core/app/dns"
	"v2ray.com/core/app/proxyman"
	. "v2ray.com/core/app/router"
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/testing/assert"
)

func TestSimpleRouter(t *testing.T) {
	assert := assert.On(t)

	config := &Config{
		Rule: []*RoutingRule{
			{
				Tag: "test",
				NetworkList: &v2net.NetworkList{
					Network: []v2net.Network{v2net.Network_TCP},
				},
			},
		},
	}

	space := app.NewSpace()
	space.BindApp(dns.APP_ID, dns.NewCacheServer(space, &dns.Config{}))
	space.BindApp(dispatcher.APP_ID, dispatchers.NewDefaultDispatcher(space))
	space.BindApp(proxyman.APP_ID_OUTBOUND_MANAGER, proxyman.NewDefaultOutboundHandlerManager())
	r := NewRouter(config, space)
	space.BindApp(APP_ID, r)
	assert.Error(space.Initialize()).IsNil()

	tag, err := r.TakeDetour(v2net.TCPDestination(v2net.DomainAddress("v2ray.com"), 80))
	assert.Error(err).IsNil()
	assert.String(tag).Equals("test")
}
