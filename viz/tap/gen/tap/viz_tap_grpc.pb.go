// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tap

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TapClient is the client API for Tap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TapClient interface {
	// Deprecated: Do not use.
	Tap(ctx context.Context, in *TapRequest, opts ...grpc.CallOption) (Tap_TapClient, error)
	// Deprecated: Do not use.
	TapByResource(ctx context.Context, in *TapByResourceRequest, opts ...grpc.CallOption) (Tap_TapByResourceClient, error)
}

type tapClient struct {
	cc grpc.ClientConnInterface
}

func NewTapClient(cc grpc.ClientConnInterface) TapClient {
	return &tapClient{cc}
}

// Deprecated: Do not use.
func (c *tapClient) Tap(ctx context.Context, in *TapRequest, opts ...grpc.CallOption) (Tap_TapClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tap_ServiceDesc.Streams[0], "/linkerd2.tap.Tap/Tap", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapTapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tap_TapClient interface {
	Recv() (*TapEvent, error)
	grpc.ClientStream
}

type tapTapClient struct {
	grpc.ClientStream
}

func (x *tapTapClient) Recv() (*TapEvent, error) {
	m := new(TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Deprecated: Do not use.
func (c *tapClient) TapByResource(ctx context.Context, in *TapByResourceRequest, opts ...grpc.CallOption) (Tap_TapByResourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tap_ServiceDesc.Streams[1], "/linkerd2.tap.Tap/TapByResource", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapTapByResourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tap_TapByResourceClient interface {
	Recv() (*TapEvent, error)
	grpc.ClientStream
}

type tapTapByResourceClient struct {
	grpc.ClientStream
}

func (x *tapTapByResourceClient) Recv() (*TapEvent, error) {
	m := new(TapEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TapServer is the server API for Tap service.
// All implementations must embed UnimplementedTapServer
// for forward compatibility
type TapServer interface {
	// Deprecated: Do not use.
	Tap(*TapRequest, Tap_TapServer) error
	// Deprecated: Do not use.
	TapByResource(*TapByResourceRequest, Tap_TapByResourceServer) error
	mustEmbedUnimplementedTapServer()
}

// UnimplementedTapServer must be embedded to have forward compatible implementations.
type UnimplementedTapServer struct {
}

func (UnimplementedTapServer) Tap(*TapRequest, Tap_TapServer) error {
	return status.Errorf(codes.Unimplemented, "method Tap not implemented")
}
func (UnimplementedTapServer) TapByResource(*TapByResourceRequest, Tap_TapByResourceServer) error {
	return status.Errorf(codes.Unimplemented, "method TapByResource not implemented")
}
func (UnimplementedTapServer) mustEmbedUnimplementedTapServer() {}

// UnsafeTapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TapServer will
// result in compilation errors.
type UnsafeTapServer interface {
	mustEmbedUnimplementedTapServer()
}

func RegisterTapServer(s grpc.ServiceRegistrar, srv TapServer) {
	s.RegisterService(&Tap_ServiceDesc, srv)
}

func _Tap_Tap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TapServer).Tap(m, &tapTapServer{stream})
}

type Tap_TapServer interface {
	Send(*TapEvent) error
	grpc.ServerStream
}

type tapTapServer struct {
	grpc.ServerStream
}

func (x *tapTapServer) Send(m *TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Tap_TapByResource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TapByResourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TapServer).TapByResource(m, &tapTapByResourceServer{stream})
}

type Tap_TapByResourceServer interface {
	Send(*TapEvent) error
	grpc.ServerStream
}

type tapTapByResourceServer struct {
	grpc.ServerStream
}

func (x *tapTapByResourceServer) Send(m *TapEvent) error {
	return x.ServerStream.SendMsg(m)
}

// Tap_ServiceDesc is the grpc.ServiceDesc for Tap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkerd2.tap.Tap",
	HandlerType: (*TapServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tap",
			Handler:       _Tap_Tap_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TapByResource",
			Handler:       _Tap_TapByResource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "viz_tap.proto",
}