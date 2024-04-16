// Copyright 2024 Google LLC
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

package migration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	migrationpb "cloud.google.com/go/bigquery/migration/apiv2alpha/migrationpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateMigrationWorkflow []gax.CallOption
	GetMigrationWorkflow    []gax.CallOption
	ListMigrationWorkflows  []gax.CallOption
	DeleteMigrationWorkflow []gax.CallOption
	StartMigrationWorkflow  []gax.CallOption
	GetMigrationSubtask     []gax.CallOption
	ListMigrationSubtasks   []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquerymigration.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("bigquerymigration.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquerymigration.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://bigquerymigration.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListMigrationWorkflows: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		StartMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetMigrationSubtask: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListMigrationSubtasks: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		CreateMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListMigrationWorkflows: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		StartMigrationWorkflow: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetMigrationSubtask: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListMigrationSubtasks: []gax.CallOption{
			gax.WithTimeout(120000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from BigQuery Migration API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateMigrationWorkflow(context.Context, *migrationpb.CreateMigrationWorkflowRequest, ...gax.CallOption) (*migrationpb.MigrationWorkflow, error)
	GetMigrationWorkflow(context.Context, *migrationpb.GetMigrationWorkflowRequest, ...gax.CallOption) (*migrationpb.MigrationWorkflow, error)
	ListMigrationWorkflows(context.Context, *migrationpb.ListMigrationWorkflowsRequest, ...gax.CallOption) *MigrationWorkflowIterator
	DeleteMigrationWorkflow(context.Context, *migrationpb.DeleteMigrationWorkflowRequest, ...gax.CallOption) error
	StartMigrationWorkflow(context.Context, *migrationpb.StartMigrationWorkflowRequest, ...gax.CallOption) error
	GetMigrationSubtask(context.Context, *migrationpb.GetMigrationSubtaskRequest, ...gax.CallOption) (*migrationpb.MigrationSubtask, error)
	ListMigrationSubtasks(context.Context, *migrationpb.ListMigrationSubtasksRequest, ...gax.CallOption) *MigrationSubtaskIterator
}

// Client is a client for interacting with BigQuery Migration API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to handle EDW migrations.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateMigrationWorkflow creates a migration workflow.
func (c *Client) CreateMigrationWorkflow(ctx context.Context, req *migrationpb.CreateMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	return c.internalClient.CreateMigrationWorkflow(ctx, req, opts...)
}

// GetMigrationWorkflow gets a previously created migration workflow.
func (c *Client) GetMigrationWorkflow(ctx context.Context, req *migrationpb.GetMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	return c.internalClient.GetMigrationWorkflow(ctx, req, opts...)
}

// ListMigrationWorkflows lists previously created migration workflow.
func (c *Client) ListMigrationWorkflows(ctx context.Context, req *migrationpb.ListMigrationWorkflowsRequest, opts ...gax.CallOption) *MigrationWorkflowIterator {
	return c.internalClient.ListMigrationWorkflows(ctx, req, opts...)
}

// DeleteMigrationWorkflow deletes a migration workflow by name.
func (c *Client) DeleteMigrationWorkflow(ctx context.Context, req *migrationpb.DeleteMigrationWorkflowRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteMigrationWorkflow(ctx, req, opts...)
}

// StartMigrationWorkflow starts a previously created migration workflow. I.e., the state transitions
// from DRAFT to RUNNING. This is a no-op if the state is already RUNNING.
// An error will be signaled if the state is anything other than DRAFT or
// RUNNING.
func (c *Client) StartMigrationWorkflow(ctx context.Context, req *migrationpb.StartMigrationWorkflowRequest, opts ...gax.CallOption) error {
	return c.internalClient.StartMigrationWorkflow(ctx, req, opts...)
}

// GetMigrationSubtask gets a previously created migration subtask.
func (c *Client) GetMigrationSubtask(ctx context.Context, req *migrationpb.GetMigrationSubtaskRequest, opts ...gax.CallOption) (*migrationpb.MigrationSubtask, error) {
	return c.internalClient.GetMigrationSubtask(ctx, req, opts...)
}

// ListMigrationSubtasks lists previously created migration subtasks.
func (c *Client) ListMigrationSubtasks(ctx context.Context, req *migrationpb.ListMigrationSubtasksRequest, opts ...gax.CallOption) *MigrationSubtaskIterator {
	return c.internalClient.ListMigrationSubtasks(ctx, req, opts...)
}

// gRPCClient is a client for interacting with BigQuery Migration API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client migrationpb.MigrationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new migration service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to handle EDW migrations.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:    connPool,
		client:      migrationpb.NewMigrationServiceClient(connPool),
		CallOptions: &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions
}

// NewRESTClient creates a new migration service rest client.
//
// Service to handle EDW migrations.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://bigquerymigration.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://bigquerymigration.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://bigquerymigration.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://bigquerymigration.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) CreateMigrationWorkflow(ctx context.Context, req *migrationpb.CreateMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateMigrationWorkflow[0:len((*c.CallOptions).CreateMigrationWorkflow):len((*c.CallOptions).CreateMigrationWorkflow)], opts...)
	var resp *migrationpb.MigrationWorkflow
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetMigrationWorkflow(ctx context.Context, req *migrationpb.GetMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetMigrationWorkflow[0:len((*c.CallOptions).GetMigrationWorkflow):len((*c.CallOptions).GetMigrationWorkflow)], opts...)
	var resp *migrationpb.MigrationWorkflow
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListMigrationWorkflows(ctx context.Context, req *migrationpb.ListMigrationWorkflowsRequest, opts ...gax.CallOption) *MigrationWorkflowIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListMigrationWorkflows[0:len((*c.CallOptions).ListMigrationWorkflows):len((*c.CallOptions).ListMigrationWorkflows)], opts...)
	it := &MigrationWorkflowIterator{}
	req = proto.Clone(req).(*migrationpb.ListMigrationWorkflowsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*migrationpb.MigrationWorkflow, string, error) {
		resp := &migrationpb.ListMigrationWorkflowsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListMigrationWorkflows(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMigrationWorkflows(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) DeleteMigrationWorkflow(ctx context.Context, req *migrationpb.DeleteMigrationWorkflowRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteMigrationWorkflow[0:len((*c.CallOptions).DeleteMigrationWorkflow):len((*c.CallOptions).DeleteMigrationWorkflow)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) StartMigrationWorkflow(ctx context.Context, req *migrationpb.StartMigrationWorkflowRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StartMigrationWorkflow[0:len((*c.CallOptions).StartMigrationWorkflow):len((*c.CallOptions).StartMigrationWorkflow)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.StartMigrationWorkflow(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetMigrationSubtask(ctx context.Context, req *migrationpb.GetMigrationSubtaskRequest, opts ...gax.CallOption) (*migrationpb.MigrationSubtask, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetMigrationSubtask[0:len((*c.CallOptions).GetMigrationSubtask):len((*c.CallOptions).GetMigrationSubtask)], opts...)
	var resp *migrationpb.MigrationSubtask
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetMigrationSubtask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListMigrationSubtasks(ctx context.Context, req *migrationpb.ListMigrationSubtasksRequest, opts ...gax.CallOption) *MigrationSubtaskIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListMigrationSubtasks[0:len((*c.CallOptions).ListMigrationSubtasks):len((*c.CallOptions).ListMigrationSubtasks)], opts...)
	it := &MigrationSubtaskIterator{}
	req = proto.Clone(req).(*migrationpb.ListMigrationSubtasksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*migrationpb.MigrationSubtask, string, error) {
		resp := &migrationpb.ListMigrationSubtasksResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListMigrationSubtasks(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMigrationSubtasks(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateMigrationWorkflow creates a migration workflow.
func (c *restClient) CreateMigrationWorkflow(ctx context.Context, req *migrationpb.CreateMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetMigrationWorkflow()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v/workflows", req.GetParent())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateMigrationWorkflow[0:len((*c.CallOptions).CreateMigrationWorkflow):len((*c.CallOptions).CreateMigrationWorkflow)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &migrationpb.MigrationWorkflow{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetMigrationWorkflow gets a previously created migration workflow.
func (c *restClient) GetMigrationWorkflow(ctx context.Context, req *migrationpb.GetMigrationWorkflowRequest, opts ...gax.CallOption) (*migrationpb.MigrationWorkflow, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v", req.GetName())

	params := url.Values{}
	if req.GetReadMask() != nil {
		readMask, err := protojson.Marshal(req.GetReadMask())
		if err != nil {
			return nil, err
		}
		params.Add("readMask", string(readMask[1:len(readMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetMigrationWorkflow[0:len((*c.CallOptions).GetMigrationWorkflow):len((*c.CallOptions).GetMigrationWorkflow)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &migrationpb.MigrationWorkflow{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListMigrationWorkflows lists previously created migration workflow.
func (c *restClient) ListMigrationWorkflows(ctx context.Context, req *migrationpb.ListMigrationWorkflowsRequest, opts ...gax.CallOption) *MigrationWorkflowIterator {
	it := &MigrationWorkflowIterator{}
	req = proto.Clone(req).(*migrationpb.ListMigrationWorkflowsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*migrationpb.MigrationWorkflow, string, error) {
		resp := &migrationpb.ListMigrationWorkflowsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v2alpha/%v/workflows", req.GetParent())

		params := url.Values{}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetReadMask() != nil {
			readMask, err := protojson.Marshal(req.GetReadMask())
			if err != nil {
				return nil, "", err
			}
			params.Add("readMask", string(readMask[1:len(readMask)-1]))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetMigrationWorkflows(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// DeleteMigrationWorkflow deletes a migration workflow by name.
func (c *restClient) DeleteMigrationWorkflow(ctx context.Context, req *migrationpb.DeleteMigrationWorkflowRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// StartMigrationWorkflow starts a previously created migration workflow. I.e., the state transitions
// from DRAFT to RUNNING. This is a no-op if the state is already RUNNING.
// An error will be signaled if the state is anything other than DRAFT or
// RUNNING.
func (c *restClient) StartMigrationWorkflow(ctx context.Context, req *migrationpb.StartMigrationWorkflowRequest, opts ...gax.CallOption) error {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v:start", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// GetMigrationSubtask gets a previously created migration subtask.
func (c *restClient) GetMigrationSubtask(ctx context.Context, req *migrationpb.GetMigrationSubtaskRequest, opts ...gax.CallOption) (*migrationpb.MigrationSubtask, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v2alpha/%v", req.GetName())

	params := url.Values{}
	if req.GetReadMask() != nil {
		readMask, err := protojson.Marshal(req.GetReadMask())
		if err != nil {
			return nil, err
		}
		params.Add("readMask", string(readMask[1:len(readMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetMigrationSubtask[0:len((*c.CallOptions).GetMigrationSubtask):len((*c.CallOptions).GetMigrationSubtask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &migrationpb.MigrationSubtask{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListMigrationSubtasks lists previously created migration subtasks.
func (c *restClient) ListMigrationSubtasks(ctx context.Context, req *migrationpb.ListMigrationSubtasksRequest, opts ...gax.CallOption) *MigrationSubtaskIterator {
	it := &MigrationSubtaskIterator{}
	req = proto.Clone(req).(*migrationpb.ListMigrationSubtasksRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*migrationpb.MigrationSubtask, string, error) {
		resp := &migrationpb.ListMigrationSubtasksResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v2alpha/%v/subtasks", req.GetParent())

		params := url.Values{}
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetReadMask() != nil {
			readMask, err := protojson.Marshal(req.GetReadMask())
			if err != nil {
				return nil, "", err
			}
			params.Add("readMask", string(readMask[1:len(readMask)-1]))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := io.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetMigrationSubtasks(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
