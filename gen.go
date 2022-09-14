package btbench

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const PayloadEntropyBytes = 500
const MessagesInEnvelope = 100

var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMicro())) // don't use in goroutines

func GenerateEnvelope() *Envelope {
	var res Envelope
	for i := 0; i < MessagesInEnvelope; i++ {
		res.Msg = append(res.Msg, GenerateMessage())
	}
	return &res
}

func GenerateMessage() *SampleMsg {
	var buf [PayloadEntropyBytes]byte
	rnd.Read(buf[:])

	return &SampleMsg{
		Id:      uuid.New().String(),
		Payload: base64.RawStdEncoding.EncodeToString(buf[:]),
	}
}
