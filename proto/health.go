package proto

// HealthCheckRequest yêu cầu kiểm tra sức khỏe server
type HealthCheckRequest struct{}

func (x *HealthCheckRequest) Reset()         {}
func (x *HealthCheckRequest) String() string { return "" }
func (*HealthCheckRequest) ProtoMessage()    {}

// HealthCheckResponse trả về trạng thái server
type HealthCheckResponse struct {
	Healthy   bool   `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Version   string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *HealthCheckResponse) Reset()         {}
func (x *HealthCheckResponse) String() string { return "" }
func (*HealthCheckResponse) ProtoMessage()    {}

func (x *HealthCheckResponse) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *HealthCheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HealthCheckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HealthCheckResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HealthCheckResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}
