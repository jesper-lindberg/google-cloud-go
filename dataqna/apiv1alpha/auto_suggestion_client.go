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

package dataqna

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	dataqnapb "cloud.google.com/go/dataqna/apiv1alpha/dataqnapb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newAutoSuggestionClientHook clientHook

// AutoSuggestionCallOptions contains the retry settings for each method of AutoSuggestionClient.
type AutoSuggestionCallOptions struct {
	SuggestQueries []gax.CallOption
}

func defaultAutoSuggestionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataqna.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("dataqna.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("dataqna.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataqna.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAutoSuggestionCallOptions() *AutoSuggestionCallOptions {
	return &AutoSuggestionCallOptions{
		SuggestQueries: []gax.CallOption{
			gax.WithTimeout(2000 * time.Millisecond),
		},
	}
}

func defaultAutoSuggestionRESTCallOptions() *AutoSuggestionCallOptions {
	return &AutoSuggestionCallOptions{
		SuggestQueries: []gax.CallOption{
			gax.WithTimeout(2000 * time.Millisecond),
		},
	}
}

// internalAutoSuggestionClient is an interface that defines the methods available from Data QnA API.
type internalAutoSuggestionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	SuggestQueries(context.Context, *dataqnapb.SuggestQueriesRequest, ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error)
}

// AutoSuggestionClient is a client for interacting with Data QnA API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This stateless API provides automatic suggestions for natural language
// queries for the data sources in the provided project and location.
//
// The service provides a resourceless operation suggestQueries that can be
// called to get a list of suggestions for a given incomplete query and scope
// (or list of scopes) under which the query is to be interpreted.
//
// There are two types of suggestions, ENTITY for single entity suggestions
// and TEMPLATE for full sentences. By default, both types are returned.
//
// Example Request:
//
// The service will retrieve information based on the given scope(s) and give
// suggestions based on that (e.g. “top item” for “top it” if “item” is a known
// dimension for the provided scope).
type AutoSuggestionClient struct {
	// The internal transport-dependent client.
	internalClient internalAutoSuggestionClient

	// The call options for this service.
	CallOptions *AutoSuggestionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AutoSuggestionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AutoSuggestionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AutoSuggestionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// SuggestQueries gets a list of suggestions based on a prefix string.
// AutoSuggestion tolerance should be less than 1 second.
func (c *AutoSuggestionClient) SuggestQueries(ctx context.Context, req *dataqnapb.SuggestQueriesRequest, opts ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error) {
	return c.internalClient.SuggestQueries(ctx, req, opts...)
}

// autoSuggestionGRPCClient is a client for interacting with Data QnA API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type autoSuggestionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AutoSuggestionClient
	CallOptions **AutoSuggestionCallOptions

	// The gRPC API client.
	autoSuggestionClient dataqnapb.AutoSuggestionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAutoSuggestionClient creates a new auto suggestion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This stateless API provides automatic suggestions for natural language
// queries for the data sources in the provided project and location.
//
// The service provides a resourceless operation suggestQueries that can be
// called to get a list of suggestions for a given incomplete query and scope
// (or list of scopes) under which the query is to be interpreted.
//
// There are two types of suggestions, ENTITY for single entity suggestions
// and TEMPLATE for full sentences. By default, both types are returned.
//
// Example Request:
//
// The service will retrieve information based on the given scope(s) and give
// suggestions based on that (e.g. “top item” for “top it” if “item” is a known
// dimension for the provided scope).
func NewAutoSuggestionClient(ctx context.Context, opts ...option.ClientOption) (*AutoSuggestionClient, error) {
	clientOpts := defaultAutoSuggestionGRPCClientOptions()
	if newAutoSuggestionClientHook != nil {
		hookOpts, err := newAutoSuggestionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AutoSuggestionClient{CallOptions: defaultAutoSuggestionCallOptions()}

	c := &autoSuggestionGRPCClient{
		connPool:             connPool,
		autoSuggestionClient: dataqnapb.NewAutoSuggestionServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *autoSuggestionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *autoSuggestionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *autoSuggestionGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type autoSuggestionRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing AutoSuggestionClient
	CallOptions **AutoSuggestionCallOptions
}

// NewAutoSuggestionRESTClient creates a new auto suggestion service rest client.
//
// This stateless API provides automatic suggestions for natural language
// queries for the data sources in the provided project and location.
//
// The service provides a resourceless operation suggestQueries that can be
// called to get a list of suggestions for a given incomplete query and scope
// (or list of scopes) under which the query is to be interpreted.
//
// There are two types of suggestions, ENTITY for single entity suggestions
// and TEMPLATE for full sentences. By default, both types are returned.
//
// Example Request:
//
// The service will retrieve information based on the given scope(s) and give
// suggestions based on that (e.g. “top item” for “top it” if “item” is a known
// dimension for the provided scope).
func NewAutoSuggestionRESTClient(ctx context.Context, opts ...option.ClientOption) (*AutoSuggestionClient, error) {
	clientOpts := append(defaultAutoSuggestionRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultAutoSuggestionRESTCallOptions()
	c := &autoSuggestionRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &AutoSuggestionClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultAutoSuggestionRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataqna.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://dataqna.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://dataqna.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataqna.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *autoSuggestionRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *autoSuggestionRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *autoSuggestionRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *autoSuggestionGRPCClient) SuggestQueries(ctx context.Context, req *dataqnapb.SuggestQueriesRequest, opts ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SuggestQueries[0:len((*c.CallOptions).SuggestQueries):len((*c.CallOptions).SuggestQueries)], opts...)
	var resp *dataqnapb.SuggestQueriesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.autoSuggestionClient.SuggestQueries(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SuggestQueries gets a list of suggestions based on a prefix string.
// AutoSuggestion tolerance should be less than 1 second.
func (c *autoSuggestionRESTClient) SuggestQueries(ctx context.Context, req *dataqnapb.SuggestQueriesRequest, opts ...gax.CallOption) (*dataqnapb.SuggestQueriesResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:suggestQueries", req.GetParent())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).SuggestQueries[0:len((*c.CallOptions).SuggestQueries):len((*c.CallOptions).SuggestQueries)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataqnapb.SuggestQueriesResponse{}
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
