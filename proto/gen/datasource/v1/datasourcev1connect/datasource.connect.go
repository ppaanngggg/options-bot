// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: datasource/v1/datasource.proto

package datasourcev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ppaanngggg/option-bot/proto/gen/datasource/v1"
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
	// DataSourceServiceName is the fully-qualified name of the DataSourceService service.
	DataSourceServiceName = "datasource.v1.DataSourceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DataSourceServiceSetGlobalProcedure is the fully-qualified name of the DataSourceService's
	// SetGlobal RPC.
	DataSourceServiceSetGlobalProcedure = "/datasource.v1.DataSourceService/SetGlobal"
	// DataSourceServiceSearchSymbolsProcedure is the fully-qualified name of the DataSourceService's
	// SearchSymbols RPC.
	DataSourceServiceSearchSymbolsProcedure = "/datasource.v1.DataSourceService/SearchSymbols"
	// DataSourceServiceGetOptionExpirationsProcedure is the fully-qualified name of the
	// DataSourceService's GetOptionExpirations RPC.
	DataSourceServiceGetOptionExpirationsProcedure = "/datasource.v1.DataSourceService/GetOptionExpirations"
	// DataSourceServiceGetOptionChainsProcedure is the fully-qualified name of the DataSourceService's
	// GetOptionChains RPC.
	DataSourceServiceGetOptionChainsProcedure = "/datasource.v1.DataSourceService/GetOptionChains"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	dataSourceServiceServiceDescriptor                    = v1.File_datasource_v1_datasource_proto.Services().ByName("DataSourceService")
	dataSourceServiceSetGlobalMethodDescriptor            = dataSourceServiceServiceDescriptor.Methods().ByName("SetGlobal")
	dataSourceServiceSearchSymbolsMethodDescriptor        = dataSourceServiceServiceDescriptor.Methods().ByName("SearchSymbols")
	dataSourceServiceGetOptionExpirationsMethodDescriptor = dataSourceServiceServiceDescriptor.Methods().ByName("GetOptionExpirations")
	dataSourceServiceGetOptionChainsMethodDescriptor      = dataSourceServiceServiceDescriptor.Methods().ByName("GetOptionChains")
)

// DataSourceServiceClient is a client for the datasource.v1.DataSourceService service.
type DataSourceServiceClient interface {
	SetGlobal(context.Context, *connect.Request[v1.SetGlobalRequest]) (*connect.Response[v1.SetGlobalResponse], error)
	SearchSymbols(context.Context, *connect.Request[v1.SearchSymbolsRequest]) (*connect.Response[v1.SearchSymbolsResponse], error)
	GetOptionExpirations(context.Context, *connect.Request[v1.GetOptionExpirationsRequest]) (*connect.Response[v1.GetOptionExpirationsResponse], error)
	GetOptionChains(context.Context, *connect.Request[v1.GetOptionChainsRequest]) (*connect.Response[v1.GetOptionChainsResponse], error)
}

// NewDataSourceServiceClient constructs a client for the datasource.v1.DataSourceService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDataSourceServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) DataSourceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &dataSourceServiceClient{
		setGlobal: connect.NewClient[v1.SetGlobalRequest, v1.SetGlobalResponse](
			httpClient,
			baseURL+DataSourceServiceSetGlobalProcedure,
			connect.WithSchema(dataSourceServiceSetGlobalMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		searchSymbols: connect.NewClient[v1.SearchSymbolsRequest, v1.SearchSymbolsResponse](
			httpClient,
			baseURL+DataSourceServiceSearchSymbolsProcedure,
			connect.WithSchema(dataSourceServiceSearchSymbolsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOptionExpirations: connect.NewClient[v1.GetOptionExpirationsRequest, v1.GetOptionExpirationsResponse](
			httpClient,
			baseURL+DataSourceServiceGetOptionExpirationsProcedure,
			connect.WithSchema(dataSourceServiceGetOptionExpirationsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOptionChains: connect.NewClient[v1.GetOptionChainsRequest, v1.GetOptionChainsResponse](
			httpClient,
			baseURL+DataSourceServiceGetOptionChainsProcedure,
			connect.WithSchema(dataSourceServiceGetOptionChainsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// dataSourceServiceClient implements DataSourceServiceClient.
type dataSourceServiceClient struct {
	setGlobal            *connect.Client[v1.SetGlobalRequest, v1.SetGlobalResponse]
	searchSymbols        *connect.Client[v1.SearchSymbolsRequest, v1.SearchSymbolsResponse]
	getOptionExpirations *connect.Client[v1.GetOptionExpirationsRequest, v1.GetOptionExpirationsResponse]
	getOptionChains      *connect.Client[v1.GetOptionChainsRequest, v1.GetOptionChainsResponse]
}

// SetGlobal calls datasource.v1.DataSourceService.SetGlobal.
func (c *dataSourceServiceClient) SetGlobal(ctx context.Context, req *connect.Request[v1.SetGlobalRequest]) (*connect.Response[v1.SetGlobalResponse], error) {
	return c.setGlobal.CallUnary(ctx, req)
}

// SearchSymbols calls datasource.v1.DataSourceService.SearchSymbols.
func (c *dataSourceServiceClient) SearchSymbols(ctx context.Context, req *connect.Request[v1.SearchSymbolsRequest]) (*connect.Response[v1.SearchSymbolsResponse], error) {
	return c.searchSymbols.CallUnary(ctx, req)
}

// GetOptionExpirations calls datasource.v1.DataSourceService.GetOptionExpirations.
func (c *dataSourceServiceClient) GetOptionExpirations(ctx context.Context, req *connect.Request[v1.GetOptionExpirationsRequest]) (*connect.Response[v1.GetOptionExpirationsResponse], error) {
	return c.getOptionExpirations.CallUnary(ctx, req)
}

// GetOptionChains calls datasource.v1.DataSourceService.GetOptionChains.
func (c *dataSourceServiceClient) GetOptionChains(ctx context.Context, req *connect.Request[v1.GetOptionChainsRequest]) (*connect.Response[v1.GetOptionChainsResponse], error) {
	return c.getOptionChains.CallUnary(ctx, req)
}

// DataSourceServiceHandler is an implementation of the datasource.v1.DataSourceService service.
type DataSourceServiceHandler interface {
	SetGlobal(context.Context, *connect.Request[v1.SetGlobalRequest]) (*connect.Response[v1.SetGlobalResponse], error)
	SearchSymbols(context.Context, *connect.Request[v1.SearchSymbolsRequest]) (*connect.Response[v1.SearchSymbolsResponse], error)
	GetOptionExpirations(context.Context, *connect.Request[v1.GetOptionExpirationsRequest]) (*connect.Response[v1.GetOptionExpirationsResponse], error)
	GetOptionChains(context.Context, *connect.Request[v1.GetOptionChainsRequest]) (*connect.Response[v1.GetOptionChainsResponse], error)
}

// NewDataSourceServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDataSourceServiceHandler(svc DataSourceServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	dataSourceServiceSetGlobalHandler := connect.NewUnaryHandler(
		DataSourceServiceSetGlobalProcedure,
		svc.SetGlobal,
		connect.WithSchema(dataSourceServiceSetGlobalMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	dataSourceServiceSearchSymbolsHandler := connect.NewUnaryHandler(
		DataSourceServiceSearchSymbolsProcedure,
		svc.SearchSymbols,
		connect.WithSchema(dataSourceServiceSearchSymbolsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	dataSourceServiceGetOptionExpirationsHandler := connect.NewUnaryHandler(
		DataSourceServiceGetOptionExpirationsProcedure,
		svc.GetOptionExpirations,
		connect.WithSchema(dataSourceServiceGetOptionExpirationsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	dataSourceServiceGetOptionChainsHandler := connect.NewUnaryHandler(
		DataSourceServiceGetOptionChainsProcedure,
		svc.GetOptionChains,
		connect.WithSchema(dataSourceServiceGetOptionChainsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/datasource.v1.DataSourceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DataSourceServiceSetGlobalProcedure:
			dataSourceServiceSetGlobalHandler.ServeHTTP(w, r)
		case DataSourceServiceSearchSymbolsProcedure:
			dataSourceServiceSearchSymbolsHandler.ServeHTTP(w, r)
		case DataSourceServiceGetOptionExpirationsProcedure:
			dataSourceServiceGetOptionExpirationsHandler.ServeHTTP(w, r)
		case DataSourceServiceGetOptionChainsProcedure:
			dataSourceServiceGetOptionChainsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDataSourceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDataSourceServiceHandler struct{}

func (UnimplementedDataSourceServiceHandler) SetGlobal(context.Context, *connect.Request[v1.SetGlobalRequest]) (*connect.Response[v1.SetGlobalResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("datasource.v1.DataSourceService.SetGlobal is not implemented"))
}

func (UnimplementedDataSourceServiceHandler) SearchSymbols(context.Context, *connect.Request[v1.SearchSymbolsRequest]) (*connect.Response[v1.SearchSymbolsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("datasource.v1.DataSourceService.SearchSymbols is not implemented"))
}

func (UnimplementedDataSourceServiceHandler) GetOptionExpirations(context.Context, *connect.Request[v1.GetOptionExpirationsRequest]) (*connect.Response[v1.GetOptionExpirationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("datasource.v1.DataSourceService.GetOptionExpirations is not implemented"))
}

func (UnimplementedDataSourceServiceHandler) GetOptionChains(context.Context, *connect.Request[v1.GetOptionChainsRequest]) (*connect.Response[v1.GetOptionChainsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("datasource.v1.DataSourceService.GetOptionChains is not implemented"))
}
