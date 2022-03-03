// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by cloud.google.com/go/internal/gapicgen/gensnippets. DO NOT EDIT.

// [START assuredworkloads_v1beta1_generated_AssuredWorkloadsService_UpdateWorkload_sync]

package main

import (
	"context"

	assuredworkloads "cloud.google.com/go/assuredworkloads/apiv1beta1"
	assuredworkloadspb "google.golang.org/genproto/googleapis/cloud/assuredworkloads/v1beta1"
)

func main() {
	ctx := context.Background()
	c, err := assuredworkloads.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &assuredworkloadspb.UpdateWorkloadRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/assuredworkloads/v1beta1#UpdateWorkloadRequest.
	}
	resp, err := c.UpdateWorkload(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END assuredworkloads_v1beta1_generated_AssuredWorkloadsService_UpdateWorkload_sync]
