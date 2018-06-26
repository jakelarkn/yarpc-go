// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package echo

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "echo",
	Package:  "go.uber.org/yarpc/internal/crossdock/thrift/echo",
	FilePath: "echo.thrift",
	SHA1:     "c3e4e93d3bee132394d26e5ec61011e3f76b7f33",
	Raw:      rawIDL,
}

const rawIDL = "// Note that type definitions are being declared before the service\n// because Apache Thrift doesn't support forward references. ThriftRW\n// works just fine with the service defined up top, but we're generating\n// shapes for both libraries from this file.\n\nstruct Ping {\n    1: required string beep\n}\n\nstruct Pong {\n    1: required string boop\n}\n\nservice Echo {\n    Pong echo(1: Ping ping) (\n        ttlms = '100'\n    )\n}\n"
