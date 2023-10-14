// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"math"
	"math/rand"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Randombool_seeded_Name = "randombool_seeded"

var Randombool_seeded_ArgumentTypes = []value.Type{value.IntegerType, value.IntegerType, value.IntegerType}

func Randombool_seeded_Validate(args []value.Value) error {
	if len(args) != 3 {
		return errors.ArgumentNotEnough(Randombool_seeded_Name, 3, args)
	}
	for i := range args {
		if args[i].Type() != Randombool_seeded_ArgumentTypes[i] {
			return errors.TypeMismatch(Randombool_seeded_Name, i+1, Randombool_seeded_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of randombool_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool-seeded/
func Randombool_seeded(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Randombool_seeded_Validate(args); err != nil {
		return value.Null, err
	}

	numerator := value.Unwrap[*value.Integer](args[0])
	denominator := value.Unwrap[*value.Integer](args[1])
	seed := value.Unwrap[*value.Integer](args[2])

	rand.Seed(seed.Value)
	r := rand.Int63n(math.MaxInt64)

	return &value.Boolean{
		Value: r/math.MaxInt64 < numerator.Value/denominator.Value,
	}, nil
}