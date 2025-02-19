// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: example/loan/tx.proto

package loan

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

const (
	Msg_UpdateParams_FullMethodName  = "/example.loan.Msg/UpdateParams"
	Msg_RequestLoan_FullMethodName   = "/example.loan.Msg/RequestLoan"
	Msg_ApproveLoan_FullMethodName   = "/example.loan.Msg/ApproveLoan"
	Msg_CancelLoan_FullMethodName    = "/example.loan.Msg/CancelLoan"
	Msg_RepayLoan_FullMethodName     = "/example.loan.Msg/RepayLoan"
	Msg_LiquidateLoan_FullMethodName = "/example.loan.Msg/LiquidateLoan"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	RequestLoan(ctx context.Context, in *MsgRequestLoan, opts ...grpc.CallOption) (*MsgRequestLoanResponse, error)
	ApproveLoan(ctx context.Context, in *MsgApproveLoan, opts ...grpc.CallOption) (*MsgApproveLoanResponse, error)
	CancelLoan(ctx context.Context, in *MsgCancelLoan, opts ...grpc.CallOption) (*MsgCancelLoanResponse, error)
	RepayLoan(ctx context.Context, in *MsgRepayLoan, opts ...grpc.CallOption) (*MsgRepayLoanResponse, error)
	LiquidateLoan(ctx context.Context, in *MsgLiquidateLoan, opts ...grpc.CallOption) (*MsgLiquidateLoanResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestLoan(ctx context.Context, in *MsgRequestLoan, opts ...grpc.CallOption) (*MsgRequestLoanResponse, error) {
	out := new(MsgRequestLoanResponse)
	err := c.cc.Invoke(ctx, Msg_RequestLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveLoan(ctx context.Context, in *MsgApproveLoan, opts ...grpc.CallOption) (*MsgApproveLoanResponse, error) {
	out := new(MsgApproveLoanResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelLoan(ctx context.Context, in *MsgCancelLoan, opts ...grpc.CallOption) (*MsgCancelLoanResponse, error) {
	out := new(MsgCancelLoanResponse)
	err := c.cc.Invoke(ctx, Msg_CancelLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RepayLoan(ctx context.Context, in *MsgRepayLoan, opts ...grpc.CallOption) (*MsgRepayLoanResponse, error) {
	out := new(MsgRepayLoanResponse)
	err := c.cc.Invoke(ctx, Msg_RepayLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LiquidateLoan(ctx context.Context, in *MsgLiquidateLoan, opts ...grpc.CallOption) (*MsgLiquidateLoanResponse, error) {
	out := new(MsgLiquidateLoanResponse)
	err := c.cc.Invoke(ctx, Msg_LiquidateLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	RequestLoan(context.Context, *MsgRequestLoan) (*MsgRequestLoanResponse, error)
	ApproveLoan(context.Context, *MsgApproveLoan) (*MsgApproveLoanResponse, error)
	CancelLoan(context.Context, *MsgCancelLoan) (*MsgCancelLoanResponse, error)
	RepayLoan(context.Context, *MsgRepayLoan) (*MsgRepayLoanResponse, error)
	LiquidateLoan(context.Context, *MsgLiquidateLoan) (*MsgLiquidateLoanResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) RequestLoan(context.Context, *MsgRequestLoan) (*MsgRequestLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLoan not implemented")
}
func (UnimplementedMsgServer) ApproveLoan(context.Context, *MsgApproveLoan) (*MsgApproveLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveLoan not implemented")
}
func (UnimplementedMsgServer) CancelLoan(context.Context, *MsgCancelLoan) (*MsgCancelLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLoan not implemented")
}
func (UnimplementedMsgServer) RepayLoan(context.Context, *MsgRepayLoan) (*MsgRepayLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepayLoan not implemented")
}
func (UnimplementedMsgServer) LiquidateLoan(context.Context, *MsgLiquidateLoan) (*MsgLiquidateLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidateLoan not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RequestLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestLoan(ctx, req.(*MsgRequestLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveLoan(ctx, req.(*MsgApproveLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelLoan(ctx, req.(*MsgCancelLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RepayLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRepayLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RepayLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RepayLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RepayLoan(ctx, req.(*MsgRepayLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LiquidateLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidateLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidateLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_LiquidateLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidateLoan(ctx, req.(*MsgLiquidateLoan))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.loan.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "RequestLoan",
			Handler:    _Msg_RequestLoan_Handler,
		},
		{
			MethodName: "ApproveLoan",
			Handler:    _Msg_ApproveLoan_Handler,
		},
		{
			MethodName: "CancelLoan",
			Handler:    _Msg_CancelLoan_Handler,
		},
		{
			MethodName: "RepayLoan",
			Handler:    _Msg_RepayLoan_Handler,
		},
		{
			MethodName: "LiquidateLoan",
			Handler:    _Msg_LiquidateLoan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/loan/tx.proto",
}
