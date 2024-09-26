// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/function/shared"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Urlencode_Name = "urlencode"

var Urlencode_ArgumentTypes = []value.Type{value.StringType}

func Urlencode_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Urlencode_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Urlencode_ArgumentTypes[i] {
			return errors.TypeMismatch(Urlencode_Name, i+1, Urlencode_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of urlencode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/urlencode/
func Urlencode(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Urlencode_Validate(args); err != nil {
		return value.Null, err
	}

	input := value.Unwrap[*value.String](args[0]).Value
	enc, err := shared.UrlEncode(input)
	if err != nil {
		return &value.String{IsNotSet: true}, errors.New(Urlencode_Name, err.Error())
	}

	return &value.String{Value: enc}, nil
}
