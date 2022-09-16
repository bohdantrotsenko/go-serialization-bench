package btbench

type Envelope3 struct {
	Msg []*SampleMsg3 `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
}

type SampleMsg3 struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func FromMsg3(m *SampleMsg) *SampleMsg3 {
	var res SampleMsg3
	res.Id = m.Id
	res.Payload = m.Payload
	return &res
}

func FromEnvelope3(e *Envelope) *Envelope3 {
	var res Envelope3
	res.Msg = make([]*SampleMsg3, 0, len(e.Msg))
	for _, msg := range e.Msg {
		res.Msg = append(res.Msg, FromMsg3(msg))
	}
	return &res
}
