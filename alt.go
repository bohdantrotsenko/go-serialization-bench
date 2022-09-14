package btbench

//go:generate ffjson $GOFILE

type Envelope2 struct {
	Msg []*SampleMsg2 `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
}

type SampleMsg2 struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func FromMsg(m *SampleMsg) *SampleMsg2 {
	var res SampleMsg2
	res.Id = m.Id
	res.Payload = m.Payload
	return &res
}

func FromEnvelope(e *Envelope) *Envelope2 {
	var res Envelope2
	res.Msg = make([]*SampleMsg2, 0, len(e.Msg))
	for _, msg := range e.Msg {
		res.Msg = append(res.Msg, FromMsg(msg))
	}
	return &res
}
