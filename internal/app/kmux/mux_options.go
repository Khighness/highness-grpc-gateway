package kmux

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Author Chen Zikang
// @Email  parakovo@gmail.com
// @Since  2022-09-11

func DefaultServeMuxOptions() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithMarshalerOption(runtime.MIMEWildcard, GRPCGatewayMarshaler),
		runtime.WithIncomingHeaderMatcher(CustomIncomingHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(CustomOutgoingHeaderMatcher),
		runtime.WithMetadata(RequestMetaHandler),
		runtime.WithErrorHandler(CustomErrorHandler),
		runtime.WithRoutingErrorHandler(ErrorRoutingHandler),
		runtime.WithForwardResponseOption(RedirectFilter),
	}
}

var GRPCGatewayMarshaler = &runtime.HTTPBodyMarshaler{
	Marshaler: &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	},
}
