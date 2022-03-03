// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package servicedirectory

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicedirectorypb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newLookupClientHook clientHook

// LookupCallOptions contains the retry settings for each method of LookupClient.
type LookupCallOptions struct {
	ResolveService []gax.CallOption
}

func defaultLookupGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("servicedirectory.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("servicedirectory.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://servicedirectory.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultLookupCallOptions() *LookupCallOptions {
	return &LookupCallOptions{
		ResolveService: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalLookupClient is an interface that defines the methods availaible from Service Directory API.
type internalLookupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ResolveService(context.Context, *servicedirectorypb.ResolveServiceRequest, ...gax.CallOption) (*servicedirectorypb.ResolveServiceResponse, error)
}

// LookupClient is a client for interacting with Service Directory API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service Directory API for looking up service data at runtime.
type LookupClient struct {
	// The internal transport-dependent client.
	internalClient internalLookupClient

	// The call options for this service.
	CallOptions *LookupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *LookupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *LookupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *LookupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ResolveService returns a service and its
// associated endpoints.
// Resolving a service is not considered an active developer method.
func (c *LookupClient) ResolveService(ctx context.Context, req *servicedirectorypb.ResolveServiceRequest, opts ...gax.CallOption) (*servicedirectorypb.ResolveServiceResponse, error) {
	return c.internalClient.ResolveService(ctx, req, opts...)
}

// lookupGRPCClient is a client for interacting with Service Directory API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type lookupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing LookupClient
	CallOptions **LookupCallOptions

	// The gRPC API client.
	lookupClient servicedirectorypb.LookupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewLookupClient creates a new lookup service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service Directory API for looking up service data at runtime.
func NewLookupClient(ctx context.Context, opts ...option.ClientOption) (*LookupClient, error) {
	clientOpts := defaultLookupGRPCClientOptions()
	if newLookupClientHook != nil {
		hookOpts, err := newLookupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := LookupClient{CallOptions: defaultLookupCallOptions()}

	c := &lookupGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		lookupClient:     servicedirectorypb.NewLookupServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *lookupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *lookupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *lookupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *lookupGRPCClient) ResolveService(ctx context.Context, req *servicedirectorypb.ResolveServiceRequest, opts ...gax.CallOption) (*servicedirectorypb.ResolveServiceResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 15000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ResolveService[0:len((*c.CallOptions).ResolveService):len((*c.CallOptions).ResolveService)], opts...)
	var resp *servicedirectorypb.ResolveServiceResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.lookupClient.ResolveService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
