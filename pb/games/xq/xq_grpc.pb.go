// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xq

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

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GuestLogin(ctx context.Context, in *GuestLoginRequest, opts ...grpc.CallOption) (*GuestLoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	MyStatus(ctx context.Context, in *MyStatusRequest, opts ...grpc.CallOption) (Game_MyStatusClient, error)
	GetTables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesReply, error)
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error)
	JoinTable(ctx context.Context, in *JoinTableRequest, opts ...grpc.CallOption) (*JoinTableResponse, error)
	LeaveTable(ctx context.Context, in *LeaveTableRequest, opts ...grpc.CallOption) (*LeaveTableResponse, error)
	StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error)
	Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error)
	GetMyProfile(ctx context.Context, in *GetMyProfileRequest, opts ...grpc.CallOption) (*GetMyProfileResponse, error)
	GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error)
	MakeFriend(ctx context.Context, in *MakeFriendRequest, opts ...grpc.CallOption) (*MakeFriendResponse, error)
	AcceptFriend(ctx context.Context, in *AcceptFriendRequest, opts ...grpc.CallOption) (*AcceptFriendResponse, error)
	StarPlayer(ctx context.Context, in *StarPlayerRequest, opts ...grpc.CallOption) (*StarPlayerResponse, error)
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GuestLogin(ctx context.Context, in *GuestLoginRequest, opts ...grpc.CallOption) (*GuestLoginResponse, error) {
	out := new(GuestLoginResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/GuestLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) MyStatus(ctx context.Context, in *MyStatusRequest, opts ...grpc.CallOption) (Game_MyStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Game_ServiceDesc.Streams[0], "/xq.Game/MyStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameMyStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Game_MyStatusClient interface {
	Recv() (*MyStatusResponse, error)
	grpc.ClientStream
}

type gameMyStatusClient struct {
	grpc.ClientStream
}

func (x *gameMyStatusClient) Recv() (*MyStatusResponse, error) {
	m := new(MyStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameClient) GetTables(ctx context.Context, in *TablesRequest, opts ...grpc.CallOption) (*TablesReply, error) {
	out := new(TablesReply)
	err := c.cc.Invoke(ctx, "/xq.Game/GetTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error) {
	out := new(CreateTableResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) JoinTable(ctx context.Context, in *JoinTableRequest, opts ...grpc.CallOption) (*JoinTableResponse, error) {
	out := new(JoinTableResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/JoinTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) LeaveTable(ctx context.Context, in *LeaveTableRequest, opts ...grpc.CallOption) (*LeaveTableResponse, error) {
	out := new(LeaveTableResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/LeaveTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error) {
	out := new(StartGameResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error) {
	out := new(MoveResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetMyProfile(ctx context.Context, in *GetMyProfileRequest, opts ...grpc.CallOption) (*GetMyProfileResponse, error) {
	out := new(GetMyProfileResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/GetMyProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*GetPlayerResponse, error) {
	out := new(GetPlayerResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/GetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) MakeFriend(ctx context.Context, in *MakeFriendRequest, opts ...grpc.CallOption) (*MakeFriendResponse, error) {
	out := new(MakeFriendResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/MakeFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) AcceptFriend(ctx context.Context, in *AcceptFriendRequest, opts ...grpc.CallOption) (*AcceptFriendResponse, error) {
	out := new(AcceptFriendResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/AcceptFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) StarPlayer(ctx context.Context, in *StarPlayerRequest, opts ...grpc.CallOption) (*StarPlayerResponse, error) {
	out := new(StarPlayerResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/StarPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*DeleteFriendResponse, error) {
	out := new(DeleteFriendResponse)
	err := c.cc.Invoke(ctx, "/xq.Game/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GuestLogin(context.Context, *GuestLoginRequest) (*GuestLoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	MyStatus(*MyStatusRequest, Game_MyStatusServer) error
	GetTables(context.Context, *TablesRequest) (*TablesReply, error)
	CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error)
	JoinTable(context.Context, *JoinTableRequest) (*JoinTableResponse, error)
	LeaveTable(context.Context, *LeaveTableRequest) (*LeaveTableResponse, error)
	StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error)
	Move(context.Context, *MoveRequest) (*MoveResponse, error)
	GetMyProfile(context.Context, *GetMyProfileRequest) (*GetMyProfileResponse, error)
	GetPlayer(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error)
	MakeFriend(context.Context, *MakeFriendRequest) (*MakeFriendResponse, error)
	AcceptFriend(context.Context, *AcceptFriendRequest) (*AcceptFriendResponse, error)
	StarPlayer(context.Context, *StarPlayerRequest) (*StarPlayerResponse, error)
	DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error)
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGameServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedGameServer) GuestLogin(context.Context, *GuestLoginRequest) (*GuestLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestLogin not implemented")
}
func (UnimplementedGameServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGameServer) MyStatus(*MyStatusRequest, Game_MyStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method MyStatus not implemented")
}
func (UnimplementedGameServer) GetTables(context.Context, *TablesRequest) (*TablesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTables not implemented")
}
func (UnimplementedGameServer) CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedGameServer) JoinTable(context.Context, *JoinTableRequest) (*JoinTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinTable not implemented")
}
func (UnimplementedGameServer) LeaveTable(context.Context, *LeaveTableRequest) (*LeaveTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveTable not implemented")
}
func (UnimplementedGameServer) StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedGameServer) Move(context.Context, *MoveRequest) (*MoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedGameServer) GetMyProfile(context.Context, *GetMyProfileRequest) (*GetMyProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
func (UnimplementedGameServer) GetPlayer(context.Context, *GetPlayerRequest) (*GetPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (UnimplementedGameServer) MakeFriend(context.Context, *MakeFriendRequest) (*MakeFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeFriend not implemented")
}
func (UnimplementedGameServer) AcceptFriend(context.Context, *AcceptFriendRequest) (*AcceptFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFriend not implemented")
}
func (UnimplementedGameServer) StarPlayer(context.Context, *StarPlayerRequest) (*StarPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StarPlayer not implemented")
}
func (UnimplementedGameServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GuestLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GuestLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/GuestLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GuestLogin(ctx, req.(*GuestLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_MyStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MyStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServer).MyStatus(m, &gameMyStatusServer{stream})
}

type Game_MyStatusServer interface {
	Send(*MyStatusResponse) error
	grpc.ServerStream
}

type gameMyStatusServer struct {
	grpc.ServerStream
}

func (x *gameMyStatusServer) Send(m *MyStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Game_GetTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/GetTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetTables(ctx, req.(*TablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_JoinTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).JoinTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/JoinTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).JoinTable(ctx, req.(*JoinTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_LeaveTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).LeaveTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/LeaveTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).LeaveTable(ctx, req.(*LeaveTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).StartGame(ctx, req.(*StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Move(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/GetMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetMyProfile(ctx, req.(*GetMyProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/GetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetPlayer(ctx, req.(*GetPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_MakeFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).MakeFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/MakeFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).MakeFriend(ctx, req.(*MakeFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_AcceptFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).AcceptFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/AcceptFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).AcceptFriend(ctx, req.(*AcceptFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_StarPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StarPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).StarPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/StarPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).StarPlayer(ctx, req.(*StarPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xq.Game/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xq.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Game_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Game_Logout_Handler,
		},
		{
			MethodName: "GuestLogin",
			Handler:    _Game_GuestLogin_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Game_Register_Handler,
		},
		{
			MethodName: "GetTables",
			Handler:    _Game_GetTables_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _Game_CreateTable_Handler,
		},
		{
			MethodName: "JoinTable",
			Handler:    _Game_JoinTable_Handler,
		},
		{
			MethodName: "LeaveTable",
			Handler:    _Game_LeaveTable_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _Game_StartGame_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _Game_Move_Handler,
		},
		{
			MethodName: "GetMyProfile",
			Handler:    _Game_GetMyProfile_Handler,
		},
		{
			MethodName: "GetPlayer",
			Handler:    _Game_GetPlayer_Handler,
		},
		{
			MethodName: "MakeFriend",
			Handler:    _Game_MakeFriend_Handler,
		},
		{
			MethodName: "AcceptFriend",
			Handler:    _Game_AcceptFriend_Handler,
		},
		{
			MethodName: "StarPlayer",
			Handler:    _Game_StarPlayer_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _Game_DeleteFriend_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MyStatus",
			Handler:       _Game_MyStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/games/xq/xq.proto",
}
