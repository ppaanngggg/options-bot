// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: bot/v1/bot.proto

package botv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ppaanngggg/option-bot/proto/gen/bot/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BotServiceName is the fully-qualified name of the BotService service.
	BotServiceName = "bot.v1.BotService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BotServiceCreateProcedure is the fully-qualified name of the BotService's Create RPC.
	BotServiceCreateProcedure = "/bot.v1.BotService/Create"
	// BotServiceGetProcedure is the fully-qualified name of the BotService's Get RPC.
	BotServiceGetProcedure = "/bot.v1.BotService/Get"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	botServiceServiceDescriptor      = v1.File_bot_v1_bot_proto.Services().ByName("BotService")
	botServiceCreateMethodDescriptor = botServiceServiceDescriptor.Methods().ByName("Create")
	botServiceGetMethodDescriptor    = botServiceServiceDescriptor.Methods().ByName("Get")
)

// BotServiceClient is a client for the bot.v1.BotService service.
type BotServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
}

// NewBotServiceClient constructs a client for the bot.v1.BotService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBotServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BotServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &botServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+BotServiceCreateProcedure,
			connect.WithSchema(botServiceCreateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		get: connect.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+BotServiceGetProcedure,
			connect.WithSchema(botServiceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// botServiceClient implements BotServiceClient.
type botServiceClient struct {
	create *connect.Client[v1.CreateRequest, v1.CreateResponse]
	get    *connect.Client[v1.GetRequest, v1.GetResponse]
}

// Create calls bot.v1.BotService.Create.
func (c *botServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Get calls bot.v1.BotService.Get.
func (c *botServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// BotServiceHandler is an implementation of the bot.v1.BotService service.
type BotServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error)
}

// NewBotServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBotServiceHandler(svc BotServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	botServiceCreateHandler := connect.NewUnaryHandler(
		BotServiceCreateProcedure,
		svc.Create,
		connect.WithSchema(botServiceCreateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	botServiceGetHandler := connect.NewUnaryHandler(
		BotServiceGetProcedure,
		svc.Get,
		connect.WithSchema(botServiceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/bot.v1.BotService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BotServiceCreateProcedure:
			botServiceCreateHandler.ServeHTTP(w, r)
		case BotServiceGetProcedure:
			botServiceGetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBotServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBotServiceHandler struct{}

func (UnimplementedBotServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bot.v1.BotService.Create is not implemented"))
}

func (UnimplementedBotServiceHandler) Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("bot.v1.BotService.Get is not implemented"))
}
