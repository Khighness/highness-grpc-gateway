package kmux

import (
	"net/textproto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-11

func CustomIncomingHeaderMatcher(key string) (string, bool) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	return runtime.MetadataPrefix + key, true
}

func CustomOutgoingHeaderMatcher(key string) (string, bool) {
	return key, true
}
