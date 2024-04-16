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

package dataflow

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newJobsV1Beta3ClientHook clientHook

// JobsV1Beta3CallOptions contains the retry settings for each method of JobsV1Beta3Client.
type JobsV1Beta3CallOptions struct {
	CreateJob          []gax.CallOption
	GetJob             []gax.CallOption
	UpdateJob          []gax.CallOption
	ListJobs           []gax.CallOption
	AggregatedListJobs []gax.CallOption
	CheckActiveJobs    []gax.CallOption
	SnapshotJob        []gax.CallOption
}

func defaultJobsV1Beta3GRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("dataflow.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultJobsV1Beta3CallOptions() *JobsV1Beta3CallOptions {
	return &JobsV1Beta3CallOptions{
		CreateJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListJobs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		AggregatedListJobs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CheckActiveJobs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SnapshotJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultJobsV1Beta3RESTCallOptions() *JobsV1Beta3CallOptions {
	return &JobsV1Beta3CallOptions{
		CreateJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListJobs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		AggregatedListJobs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CheckActiveJobs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SnapshotJob: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalJobsV1Beta3Client is an interface that defines the methods available from Dataflow API.
type internalJobsV1Beta3Client interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateJob(context.Context, *dataflowpb.CreateJobRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	GetJob(context.Context, *dataflowpb.GetJobRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	UpdateJob(context.Context, *dataflowpb.UpdateJobRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	ListJobs(context.Context, *dataflowpb.ListJobsRequest, ...gax.CallOption) *JobIterator
	AggregatedListJobs(context.Context, *dataflowpb.ListJobsRequest, ...gax.CallOption) *JobIterator
	CheckActiveJobs(context.Context, *dataflowpb.CheckActiveJobsRequest, ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error)
	SnapshotJob(context.Context, *dataflowpb.SnapshotJobRequest, ...gax.CallOption) (*dataflowpb.Snapshot, error)
}

// JobsV1Beta3Client is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides a method to create and modify Google Cloud Dataflow jobs.
// A Job is a multi-stage computation graph run by the Cloud Dataflow service.
type JobsV1Beta3Client struct {
	// The internal transport-dependent client.
	internalClient internalJobsV1Beta3Client

	// The call options for this service.
	CallOptions *JobsV1Beta3CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *JobsV1Beta3Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *JobsV1Beta3Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *JobsV1Beta3Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateJob creates a Cloud Dataflow job.
//
// To create a job, we recommend using projects.locations.jobs.create with a
// [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.create is not recommended, as your job will always start
// in us-central1.
func (c *JobsV1Beta3Client) CreateJob(ctx context.Context, req *dataflowpb.CreateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.CreateJob(ctx, req, opts...)
}

// GetJob gets the state of the specified Cloud Dataflow job.
//
// To get the state of a job, we recommend using projects.locations.jobs.get
// with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.get is not recommended, as you can only get the state of
// jobs that are running in us-central1.
func (c *JobsV1Beta3Client) GetJob(ctx context.Context, req *dataflowpb.GetJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.GetJob(ctx, req, opts...)
}

// UpdateJob updates the state of an existing Cloud Dataflow job.
//
// To update the state of an existing job, we recommend using
// projects.locations.jobs.update with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.update is not recommended, as you can only update the state
// of jobs that are running in us-central1.
func (c *JobsV1Beta3Client) UpdateJob(ctx context.Context, req *dataflowpb.UpdateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.UpdateJob(ctx, req, opts...)
}

// ListJobs list the jobs of a project.
//
// To list the jobs of a project in a region, we recommend using
// projects.locations.jobs.list with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). To
// list the all jobs across all regions, use projects.jobs.aggregated. Using
// projects.jobs.list is not recommended, as you can only get the list of
// jobs that are running in us-central1.
func (c *JobsV1Beta3Client) ListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	return c.internalClient.ListJobs(ctx, req, opts...)
}

// AggregatedListJobs list the jobs of a project across all regions.
func (c *JobsV1Beta3Client) AggregatedListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	return c.internalClient.AggregatedListJobs(ctx, req, opts...)
}

// CheckActiveJobs check for existence of active jobs in the given project across all regions.
func (c *JobsV1Beta3Client) CheckActiveJobs(ctx context.Context, req *dataflowpb.CheckActiveJobsRequest, opts ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error) {
	return c.internalClient.CheckActiveJobs(ctx, req, opts...)
}

// SnapshotJob snapshot the state of a streaming job.
func (c *JobsV1Beta3Client) SnapshotJob(ctx context.Context, req *dataflowpb.SnapshotJobRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	return c.internalClient.SnapshotJob(ctx, req, opts...)
}

// jobsV1Beta3GRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type jobsV1Beta3GRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing JobsV1Beta3Client
	CallOptions **JobsV1Beta3CallOptions

	// The gRPC API client.
	jobsV1Beta3Client dataflowpb.JobsV1Beta3Client

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewJobsV1Beta3Client creates a new jobs v1 beta3 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides a method to create and modify Google Cloud Dataflow jobs.
// A Job is a multi-stage computation graph run by the Cloud Dataflow service.
func NewJobsV1Beta3Client(ctx context.Context, opts ...option.ClientOption) (*JobsV1Beta3Client, error) {
	clientOpts := defaultJobsV1Beta3GRPCClientOptions()
	if newJobsV1Beta3ClientHook != nil {
		hookOpts, err := newJobsV1Beta3ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := JobsV1Beta3Client{CallOptions: defaultJobsV1Beta3CallOptions()}

	c := &jobsV1Beta3GRPCClient{
		connPool:          connPool,
		jobsV1Beta3Client: dataflowpb.NewJobsV1Beta3Client(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *jobsV1Beta3GRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *jobsV1Beta3GRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *jobsV1Beta3GRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type jobsV1Beta3RESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing JobsV1Beta3Client
	CallOptions **JobsV1Beta3CallOptions
}

// NewJobsV1Beta3RESTClient creates a new jobs v1 beta3 rest client.
//
// Provides a method to create and modify Google Cloud Dataflow jobs.
// A Job is a multi-stage computation graph run by the Cloud Dataflow service.
func NewJobsV1Beta3RESTClient(ctx context.Context, opts ...option.ClientOption) (*JobsV1Beta3Client, error) {
	clientOpts := append(defaultJobsV1Beta3RESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultJobsV1Beta3RESTCallOptions()
	c := &jobsV1Beta3RESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &JobsV1Beta3Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultJobsV1Beta3RESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataflow.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://dataflow.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://dataflow.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *jobsV1Beta3RESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *jobsV1Beta3RESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *jobsV1Beta3RESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *jobsV1Beta3GRPCClient) CreateJob(ctx context.Context, req *dataflowpb.CreateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateJob[0:len((*c.CallOptions).CreateJob):len((*c.CallOptions).CreateJob)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.CreateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) GetJob(ctx context.Context, req *dataflowpb.GetJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetJob[0:len((*c.CallOptions).GetJob):len((*c.CallOptions).GetJob)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.GetJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) UpdateJob(ctx context.Context, req *dataflowpb.UpdateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateJob[0:len((*c.CallOptions).UpdateJob):len((*c.CallOptions).UpdateJob)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.UpdateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) ListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListJobs[0:len((*c.CallOptions).ListJobs):len((*c.CallOptions).ListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.Job, string, error) {
		resp := &dataflowpb.ListJobsResponse{}
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
			resp, err = c.jobsV1Beta3Client.ListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

func (c *jobsV1Beta3GRPCClient) AggregatedListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "project_id", url.QueryEscape(req.GetProjectId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).AggregatedListJobs[0:len((*c.CallOptions).AggregatedListJobs):len((*c.CallOptions).AggregatedListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.Job, string, error) {
		resp := &dataflowpb.ListJobsResponse{}
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
			resp, err = c.jobsV1Beta3Client.AggregatedListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

func (c *jobsV1Beta3GRPCClient) CheckActiveJobs(ctx context.Context, req *dataflowpb.CheckActiveJobsRequest, opts ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CheckActiveJobs[0:len((*c.CallOptions).CheckActiveJobs):len((*c.CallOptions).CheckActiveJobs)], opts...)
	var resp *dataflowpb.CheckActiveJobsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.CheckActiveJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *jobsV1Beta3GRPCClient) SnapshotJob(ctx context.Context, req *dataflowpb.SnapshotJobRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SnapshotJob[0:len((*c.CallOptions).SnapshotJob):len((*c.CallOptions).SnapshotJob)], opts...)
	var resp *dataflowpb.Snapshot
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobsV1Beta3Client.SnapshotJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateJob creates a Cloud Dataflow job.
//
// To create a job, we recommend using projects.locations.jobs.create with a
// [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.create is not recommended, as your job will always start
// in us-central1.
func (c *jobsV1Beta3RESTClient) CreateJob(ctx context.Context, req *dataflowpb.CreateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetJob()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs", req.GetProjectId(), req.GetLocation())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetReplaceJobId() != "" {
		params.Add("replaceJobId", fmt.Sprintf("%v", req.GetReplaceJobId()))
	}
	if req.GetView() != 0 {
		params.Add("view", fmt.Sprintf("%v", req.GetView()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateJob[0:len((*c.CallOptions).CreateJob):len((*c.CallOptions).CreateJob)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.Job{}
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

// GetJob gets the state of the specified Cloud Dataflow job.
//
// To get the state of a job, we recommend using projects.locations.jobs.get
// with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.get is not recommended, as you can only get the state of
// jobs that are running in us-central1.
func (c *jobsV1Beta3RESTClient) GetJob(ctx context.Context, req *dataflowpb.GetJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v", req.GetProjectId(), req.GetLocation(), req.GetJobId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetView() != 0 {
		params.Add("view", fmt.Sprintf("%v", req.GetView()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetJob[0:len((*c.CallOptions).GetJob):len((*c.CallOptions).GetJob)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.Job{}
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

// UpdateJob updates the state of an existing Cloud Dataflow job.
//
// To update the state of an existing job, we recommend using
// projects.locations.jobs.update with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). Using
// projects.jobs.update is not recommended, as you can only update the state
// of jobs that are running in us-central1.
func (c *jobsV1Beta3RESTClient) UpdateJob(ctx context.Context, req *dataflowpb.UpdateJobRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetJob()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v", req.GetProjectId(), req.GetLocation(), req.GetJobId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateJob[0:len((*c.CallOptions).UpdateJob):len((*c.CallOptions).UpdateJob)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.Job{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewReader(jsonReq))
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

// ListJobs list the jobs of a project.
//
// To list the jobs of a project in a region, we recommend using
// projects.locations.jobs.list with a [regional endpoint]
// (https://cloud.google.com/dataflow/docs/concepts/regional-endpoints (at https://cloud.google.com/dataflow/docs/concepts/regional-endpoints)). To
// list the all jobs across all regions, use projects.jobs.aggregated. Using
// projects.jobs.list is not recommended, as you can only get the list of
// jobs that are running in us-central1.
func (c *jobsV1Beta3RESTClient) ListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	it := &JobIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.Job, string, error) {
		resp := &dataflowpb.ListJobsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs", req.GetProjectId(), req.GetLocation())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != 0 {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetView() != 0 {
			params.Add("view", fmt.Sprintf("%v", req.GetView()))
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
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

// AggregatedListJobs list the jobs of a project across all regions.
func (c *jobsV1Beta3RESTClient) AggregatedListJobs(ctx context.Context, req *dataflowpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	it := &JobIterator{}
	req = proto.Clone(req).(*dataflowpb.ListJobsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataflowpb.Job, string, error) {
		resp := &dataflowpb.ListJobsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/jobs:aggregated", req.GetProjectId())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != 0 {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetLocation() != "" {
			params.Add("location", fmt.Sprintf("%v", req.GetLocation()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req.GetView() != 0 {
			params.Add("view", fmt.Sprintf("%v", req.GetView()))
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
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

// CheckActiveJobs check for existence of active jobs in the given project across all regions.
func (c *jobsV1Beta3RESTClient) CheckActiveJobs(ctx context.Context, req *dataflowpb.CheckActiveJobsRequest, opts ...gax.CallOption) (*dataflowpb.CheckActiveJobsResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetProjectId() != "" {
		params.Add("projectId", fmt.Sprintf("%v", req.GetProjectId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CheckActiveJobs[0:len((*c.CallOptions).CheckActiveJobs):len((*c.CallOptions).CheckActiveJobs)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.CheckActiveJobsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("", baseUrl.String(), nil)
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

// SnapshotJob snapshot the state of a streaming job.
func (c *jobsV1Beta3RESTClient) SnapshotJob(ctx context.Context, req *dataflowpb.SnapshotJobRequest, opts ...gax.CallOption) (*dataflowpb.Snapshot, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/jobs/%v:snapshot", req.GetProjectId(), req.GetLocation(), req.GetJobId())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()), "job_id", url.QueryEscape(req.GetJobId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).SnapshotJob[0:len((*c.CallOptions).SnapshotJob):len((*c.CallOptions).SnapshotJob)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.Snapshot{}
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
