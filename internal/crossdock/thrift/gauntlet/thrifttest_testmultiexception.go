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

// ThriftTest_TestMultiException_Args represents the arguments for the ThriftTest.testMultiException function.
//
// The arguments for testMultiException are sent and received over the wire as this struct.
type ThriftTest_TestMultiException_Args struct {
	Arg0 *string `json:"arg0,omitempty"`
	Arg1 *string `json:"arg1,omitempty"`
}

// ToWire translates a ThriftTest_TestMultiException_Args struct into a Thrift-level intermediate
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
func (v *ThriftTest_TestMultiException_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Arg0 != nil {
		w, err = wire.NewValueString(*(v.Arg0)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Arg1 != nil {
		w, err = wire.NewValueString(*(v.Arg1)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestMultiException_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestMultiException_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestMultiException_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestMultiException_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg0 = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg1 = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestMultiException_Args
// struct.
func (v *ThriftTest_TestMultiException_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Arg0 != nil {
		fields[i] = fmt.Sprintf("Arg0: %v", *(v.Arg0))
		i++
	}
	if v.Arg1 != nil {
		fields[i] = fmt.Sprintf("Arg1: %v", *(v.Arg1))
		i++
	}

	return fmt.Sprintf("ThriftTest_TestMultiException_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestMultiException_Args match the
// provided ThriftTest_TestMultiException_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestMultiException_Args) Equals(rhs *ThriftTest_TestMultiException_Args) bool {
	if !_String_EqualsPtr(v.Arg0, rhs.Arg0) {
		return false
	}
	if !_String_EqualsPtr(v.Arg1, rhs.Arg1) {
		return false
	}

	return true
}

// GetArg0 returns the value of Arg0 if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestMultiException_Args) GetArg0() (o string) {
	if v.Arg0 != nil {
		return *v.Arg0
	}

	return
}

// GetArg1 returns the value of Arg1 if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestMultiException_Args) GetArg1() (o string) {
	if v.Arg1 != nil {
		return *v.Arg1
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testMultiException" for this struct.
func (v *ThriftTest_TestMultiException_Args) MethodName() string {
	return "testMultiException"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ThriftTest_TestMultiException_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ThriftTest_TestMultiException_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testMultiException
// function.
var ThriftTest_TestMultiException_Helper = struct {
	// Args accepts the parameters of testMultiException in-order and returns
	// the arguments struct for the function.
	Args func(
		arg0 *string,
		arg1 *string,
	) *ThriftTest_TestMultiException_Args

	// IsException returns true if the given error can be thrown
	// by testMultiException.
	//
	// An error can be thrown by testMultiException only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for testMultiException
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// testMultiException into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by testMultiException
	//
	//   value, err := testMultiException(args)
	//   result, err := ThriftTest_TestMultiException_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from testMultiException: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*Xtruct, error) (*ThriftTest_TestMultiException_Result, error)

	// UnwrapResponse takes the result struct for testMultiException
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if testMultiException threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := ThriftTest_TestMultiException_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ThriftTest_TestMultiException_Result) (*Xtruct, error)
}{}

func init() {
	ThriftTest_TestMultiException_Helper.Args = func(
		arg0 *string,
		arg1 *string,
	) *ThriftTest_TestMultiException_Args {
		return &ThriftTest_TestMultiException_Args{
			Arg0: arg0,
			Arg1: arg1,
		}
	}

	ThriftTest_TestMultiException_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *Xception:
			return true
		case *Xception2:
			return true
		default:
			return false
		}
	}

	ThriftTest_TestMultiException_Helper.WrapResponse = func(success *Xtruct, err error) (*ThriftTest_TestMultiException_Result, error) {
		if err == nil {
			return &ThriftTest_TestMultiException_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *Xception:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ThriftTest_TestMultiException_Result.Err1")
			}
			return &ThriftTest_TestMultiException_Result{Err1: e}, nil
		case *Xception2:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ThriftTest_TestMultiException_Result.Err2")
			}
			return &ThriftTest_TestMultiException_Result{Err2: e}, nil
		}

		return nil, err
	}
	ThriftTest_TestMultiException_Helper.UnwrapResponse = func(result *ThriftTest_TestMultiException_Result) (success *Xtruct, err error) {
		if result.Err1 != nil {
			err = result.Err1
			return
		}
		if result.Err2 != nil {
			err = result.Err2
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// ThriftTest_TestMultiException_Result represents the result of a ThriftTest.testMultiException function call.
//
// The result of a testMultiException execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type ThriftTest_TestMultiException_Result struct {
	// Value returned by testMultiException after a successful execution.
	Success *Xtruct    `json:"success,omitempty"`
	Err1    *Xception  `json:"err1,omitempty"`
	Err2    *Xception2 `json:"err2,omitempty"`
}

// ToWire translates a ThriftTest_TestMultiException_Result struct into a Thrift-level intermediate
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
func (v *ThriftTest_TestMultiException_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.Err1 != nil {
		w, err = v.Err1.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Err2 != nil {
		w, err = v.Err2.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestMultiException_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Xception2_Read(w wire.Value) (*Xception2, error) {
	var v Xception2
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a ThriftTest_TestMultiException_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestMultiException_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestMultiException_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestMultiException_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Xtruct_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Err1, err = _Xception_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Err2, err = _Xception2_Read(field.Value)
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
	if v.Err1 != nil {
		count++
	}
	if v.Err2 != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestMultiException_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestMultiException_Result
// struct.
func (v *ThriftTest_TestMultiException_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.Err1 != nil {
		fields[i] = fmt.Sprintf("Err1: %v", v.Err1)
		i++
	}
	if v.Err2 != nil {
		fields[i] = fmt.Sprintf("Err2: %v", v.Err2)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestMultiException_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestMultiException_Result match the
// provided ThriftTest_TestMultiException_Result.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestMultiException_Result) Equals(rhs *ThriftTest_TestMultiException_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.Err1 == nil && rhs.Err1 == nil) || (v.Err1 != nil && rhs.Err1 != nil && v.Err1.Equals(rhs.Err1))) {
		return false
	}
	if !((v.Err2 == nil && rhs.Err2 == nil) || (v.Err2 != nil && rhs.Err2 != nil && v.Err2.Equals(rhs.Err2))) {
		return false
	}

	return true
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestMultiException_Result) GetSuccess() (o *Xtruct) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// GetErr1 returns the value of Err1 if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestMultiException_Result) GetErr1() (o *Xception) {
	if v.Err1 != nil {
		return v.Err1
	}

	return
}

// GetErr2 returns the value of Err2 if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestMultiException_Result) GetErr2() (o *Xception2) {
	if v.Err2 != nil {
		return v.Err2
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "testMultiException" for this struct.
func (v *ThriftTest_TestMultiException_Result) MethodName() string {
	return "testMultiException"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ThriftTest_TestMultiException_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
