package config

import "flag"

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-07

var (
	GRPC_PORT    = flag.Int("grpc_port", 10010, "port of hello service")
	GATEWAY_PORT = flag.Int("gateway_port", 10020, "port of gateway service")
)

func init() {
	flag.Parse()
}
