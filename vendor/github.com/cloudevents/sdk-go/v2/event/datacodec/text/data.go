// Text codec converts []byte or string to string and vice-versa.
package text

import (
	"context"
	"fmt"
)

func Decode(_ context.Context, in []byte, out interface{}) error {
	p, _ := out.(*string)
	if p == nil {
		return fmt.Errorf("text.Decode out: want *string, got %T", out)
	}
	*p = string(in)
	return nil
}

func Encode(_ context.Context, in interface{}) ([]byte, error) {
	s, ok := in.(string)
	if !ok {
		return nil, fmt.Errorf("text.Encode in: want string, got %T", in)
	}
	return []byte(s), nil
}
