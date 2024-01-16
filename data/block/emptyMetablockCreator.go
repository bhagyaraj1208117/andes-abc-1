package block

import "github.com/bhagyaraj1208117/andes-abc-1/data"

type emptyMetaBlockCreator struct{}

// NewEmptyMetaBlockCreator is able to create empty metablock instances
func NewEmptyMetaBlockCreator() *emptyMetaBlockCreator {
	return &emptyMetaBlockCreator{}
}

// CreateNewHeader creates a new empty metablock
func (creator *emptyMetaBlockCreator) CreateNewHeader() data.HeaderHandler {
	return &MetaBlock{}
}

// IsInterfaceNil returns true if there is no value under the interface
func (creator *emptyMetaBlockCreator) IsInterfaceNil() bool {
	return creator == nil
}
