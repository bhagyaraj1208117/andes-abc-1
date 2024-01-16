package block

import (
	"github.com/bhagyaraj1208117/andes-abc-1/core/check"
	"github.com/bhagyaraj1208117/andes-abc-1/data"
	"github.com/bhagyaraj1208117/andes-abc-1/marshal"
)

// GetHeaderFromBytes will unmarshal the header bytes based on the header type
func GetHeaderFromBytes(marshaller marshal.Marshalizer, creator EmptyBlockCreator, headerBytes []byte) (data.HeaderHandler, error) {
	if check.IfNil(marshaller) {
		return nil, data.ErrNilMarshalizer
	}
	if check.IfNil(creator) {
		return nil, data.ErrNilEmptyBlockCreator
	}

	header := creator.CreateNewHeader()
	err := marshaller.Unmarshal(header, headerBytes)
	if err != nil {
		return nil, err
	}

	return header, nil
}
