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

package debugger_test

import (
	"context"

	debugger "cloud.google.com/go/debugger/apiv2"
	clouddebuggerpb "google.golang.org/genproto/googleapis/devtools/clouddebugger/v2"
)

func ExampleNewDebugger2Client() {
	ctx := context.Background()
	c, err := debugger.NewDebugger2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleDebugger2Client_SetBreakpoint() {
	ctx := context.Background()
	c, err := debugger.NewDebugger2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &clouddebuggerpb.SetBreakpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/clouddebugger/v2#SetBreakpointRequest.
	}
	resp, err := c.SetBreakpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDebugger2Client_GetBreakpoint() {
	ctx := context.Background()
	c, err := debugger.NewDebugger2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &clouddebuggerpb.GetBreakpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/clouddebugger/v2#GetBreakpointRequest.
	}
	resp, err := c.GetBreakpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDebugger2Client_DeleteBreakpoint() {
	ctx := context.Background()
	c, err := debugger.NewDebugger2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &clouddebuggerpb.DeleteBreakpointRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/clouddebugger/v2#DeleteBreakpointRequest.
	}
	err = c.DeleteBreakpoint(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleDebugger2Client_ListBreakpoints() {
	ctx := context.Background()
	c, err := debugger.NewDebugger2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &clouddebuggerpb.ListBreakpointsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/clouddebugger/v2#ListBreakpointsRequest.
	}
	resp, err := c.ListBreakpoints(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleDebugger2Client_ListDebuggees() {
	ctx := context.Background()
	c, err := debugger.NewDebugger2Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &clouddebuggerpb.ListDebuggeesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/devtools/clouddebugger/v2#ListDebuggeesRequest.
	}
	resp, err := c.ListDebuggees(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
