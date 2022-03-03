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

package servicecontrol

import (
	"context"
	"fmt"
	"math"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicecontrolpb "google.golang.org/genproto/googleapis/api/servicecontrol/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newQuotaControllerClientHook clientHook

// QuotaControllerCallOptions contains the retry settings for each method of QuotaControllerClient.
type QuotaControllerCallOptions struct {
	AllocateQuota []gax.CallOption
}

func defaultQuotaControllerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("servicecontrol.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("servicecontrol.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://servicecontrol.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultQuotaControllerCallOptions() *QuotaControllerCallOptions {
	return &QuotaControllerCallOptions{
		AllocateQuota: []gax.CallOption{},
	}
}

// internalQuotaControllerClient is an interface that defines the methods availaible from Service Control API.
type internalQuotaControllerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AllocateQuota(context.Context, *servicecontrolpb.AllocateQuotaRequest, ...gax.CallOption) (*servicecontrolpb.AllocateQuotaResponse, error)
}

// QuotaControllerClient is a client for interacting with Service Control API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Google Quota Control API (at /service-control/overview)
//
// Allows clients to allocate and release quota against a managed
// service (at https://cloud.google.com/service-management/reference/rpc/google.api/servicemanagement.v1#google.api.servicemanagement.v1.ManagedService).
type QuotaControllerClient struct {
	// The internal transport-dependent client.
	internalClient internalQuotaControllerClient

	// The call options for this service.
	CallOptions *QuotaControllerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *QuotaControllerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *QuotaControllerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *QuotaControllerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AllocateQuota attempts to allocate quota for the specified consumer. It should be called
// before the operation is executed.
//
// This method requires the servicemanagement.services.quota
// permission on the specified service. For more information, see
// Cloud IAM (at https://cloud.google.com/iam).
//
// NOTE: The client must fail-open on server errors INTERNAL,
// UNKNOWN, DEADLINE_EXCEEDED, and UNAVAILABLE. To ensure system
// reliability, the server may inject these errors to prohibit any hard
// dependency on the quota functionality.
func (c *QuotaControllerClient) AllocateQuota(ctx context.Context, req *servicecontrolpb.AllocateQuotaRequest, opts ...gax.CallOption) (*servicecontrolpb.AllocateQuotaResponse, error) {
	return c.internalClient.AllocateQuota(ctx, req, opts...)
}

// quotaControllerGRPCClient is a client for interacting with Service Control API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type quotaControllerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing QuotaControllerClient
	CallOptions **QuotaControllerCallOptions

	// The gRPC API client.
	quotaControllerClient servicecontrolpb.QuotaControllerClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewQuotaControllerClient creates a new quota controller client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Google Quota Control API (at /service-control/overview)
//
// Allows clients to allocate and release quota against a managed
// service (at https://cloud.google.com/service-management/reference/rpc/google.api/servicemanagement.v1#google.api.servicemanagement.v1.ManagedService).
func NewQuotaControllerClient(ctx context.Context, opts ...option.ClientOption) (*QuotaControllerClient, error) {
	clientOpts := defaultQuotaControllerGRPCClientOptions()
	if newQuotaControllerClientHook != nil {
		hookOpts, err := newQuotaControllerClientHook(ctx, clientHookParams{})
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
	client := QuotaControllerClient{CallOptions: defaultQuotaControllerCallOptions()}

	c := &quotaControllerGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		quotaControllerClient: servicecontrolpb.NewQuotaControllerClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *quotaControllerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *quotaControllerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *quotaControllerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *quotaControllerGRPCClient) AllocateQuota(ctx context.Context, req *servicecontrolpb.AllocateQuotaRequest, opts ...gax.CallOption) (*servicecontrolpb.AllocateQuotaResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AllocateQuota[0:len((*c.CallOptions).AllocateQuota):len((*c.CallOptions).AllocateQuota)], opts...)
	var resp *servicecontrolpb.AllocateQuotaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.quotaControllerClient.AllocateQuota(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
