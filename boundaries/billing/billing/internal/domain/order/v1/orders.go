package v1

// Order list
type Orders struct {
	// order list
	list []*Order `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

// GetList
func (o *Orders) GetList() []*Order {
	return o.list
}
