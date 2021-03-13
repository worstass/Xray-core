package encoding

import (
	"context"

	"google.golang.org/grpc"
)

func ServerDesc(name string) grpc.ServiceDesc {
	return grpc.ServiceDesc{
		ServiceName: name,
		HandlerType: (*GRPCServiceServer)(nil),
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "TunMulti",
				Handler:       _GRPCService_TunMulti_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "grpc.proto",
	}
}

func (c *gRPCServiceClient) TunCustomName(ctx context.Context, name string, opts ...grpc.CallOption) (GRPCService_TunMultiClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerDesc(name).Streams[0], "/"+name+"/TunMulti", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCServiceTunMultiClient{stream}
	return x, nil
}

type GRPCServiceClientX interface {
	TunCustomName(ctx context.Context, name string, opts ...grpc.CallOption) (GRPCService_TunMultiClient, error)
	TunMulti(ctx context.Context, opts ...grpc.CallOption) (GRPCService_TunMultiClient, error)
}

func RegisterGRPCServiceServerX(s *grpc.Server, srv GRPCServiceServer, name string) {
	desc := ServerDesc(name)
	s.RegisterService(&desc, srv)
}
