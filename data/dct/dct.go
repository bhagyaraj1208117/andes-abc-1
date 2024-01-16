//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/bhagyaraj1208117/protobuf/protobuf  --gogoslick_out=. dct.proto
package dct

import (
	"math/big"

	dct "github.com/bhagyaraj1208117/andes-abc-1/data/dct/proto"
)

// New returns a new batch from given buffers
func New() *dct.DCToken {
	return &dct.DCToken{
		Value: big.NewInt(0),
	}
}
