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

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// ThriftTest_TestStringMap_Args represents the arguments for the ThriftTest.testStringMap function.
//
// The arguments for testStringMap are sent and received over the wire as this struct.
type ThriftTest_TestStringMap_Args struct {
	Thing map[string]string `json:"thing,omitempty"`
}

// ToWire translates a ThriftTest_TestStringMap_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestStringMap_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Thing != nil {
		w, err = wire.NewValueMap(_Map_String_String_MapItemList(v.Thing)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestStringMap_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestStringMap_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestStringMap_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestStringMap_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TMap {
				v.Thing, err = _Map_String_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestStringMap_Args
// struct.
func (v *ThriftTest_TestStringMap_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", v.Thing)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestStringMap_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestStringMap_Args match the
// provided ThriftTest_TestStringMap_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestStringMap_Args) Equals(rhs *ThriftTest_TestStringMap_Args) bool {
	if !((v.Thing == nil && rhs.Thing == nil) || (v.Thing != nil && rhs.Thing != nil && _Map_String_String_Equals(v.Thing, rhs.Thing))) {
		return false
	}

	return true
}

// GetThing returns the value of Thing if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestStringMap_Args) GetThing() (o map[string]string) {
	if v.Thing != nil {
		return v.Thing
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testStringMap" for this struct.
func (v *ThriftTest_TestStringMap_Args) MethodName() string {
	return "testStringMap"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ThriftTest_TestStringMap_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ThriftTest_TestStringMap_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testStringMap
// function.
var ThriftTest_TestStringMap_Helper = struct {
	// Args accepts the parameters of testStringMap in-order and returns
	// the arguments struct for the function.
	Args func(
		thing map[string]string,
	) *ThriftTest_TestStringMap_Args

	// IsException returns true if the given error can be thrown
	// by testStringMap.
	//
	// An error can be thrown by testStringMap only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for testStringMap
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// testStringMap into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by testStringMap
	//
	//   value, err := testStringMap(args)
	//   result, err := ThriftTest_TestStringMap_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from testStringMap: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(map[string]string, error) (*ThriftTest_TestStringMap_Result, error)

	// UnwrapResponse takes the result struct for testStringMap
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if testStringMap threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := ThriftTest_TestStringMap_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ThriftTest_TestStringMap_Result) (map[string]string, error)
}{}

func init() {
	ThriftTest_TestStringMap_Helper.Args = func(
		thing map[string]string,
	) *ThriftTest_TestStringMap_Args {
		return &ThriftTest_TestStringMap_Args{
			Thing: thing,
		}
	}

	ThriftTest_TestStringMap_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ThriftTest_TestStringMap_Helper.WrapResponse = func(success map[string]string, err error) (*ThriftTest_TestStringMap_Result, error) {
		if err == nil {
			return &ThriftTest_TestStringMap_Result{Success: success}, nil
		}

		return nil, err
	}
	ThriftTest_TestStringMap_Helper.UnwrapResponse = func(result *ThriftTest_TestStringMap_Result) (success map[string]string, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// ThriftTest_TestStringMap_Result represents the result of a ThriftTest.testStringMap function call.
//
// The result of a testStringMap execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type ThriftTest_TestStringMap_Result struct {
	// Value returned by testStringMap after a successful execution.
	Success map[string]string `json:"success,omitempty"`
}

// ToWire translates a ThriftTest_TestStringMap_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestStringMap_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_String_String_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestStringMap_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestStringMap_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestStringMap_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestStringMap_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestStringMap_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_String_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestStringMap_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestStringMap_Result
// struct.
func (v *ThriftTest_TestStringMap_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestStringMap_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestStringMap_Result match the
// provided ThriftTest_TestStringMap_Result.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestStringMap_Result) Equals(rhs *ThriftTest_TestStringMap_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Map_String_String_Equals(v.Success, rhs.Success))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestStringMap_Result) GetSuccess() (o map[string]string) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "testStringMap" for this struct.
func (v *ThriftTest_TestStringMap_Result) MethodName() string {
	return "testStringMap"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ThriftTest_TestStringMap_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
