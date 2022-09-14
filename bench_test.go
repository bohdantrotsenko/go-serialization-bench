package btbench

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/pquerna/ffjson/ffjson"
	"google.golang.org/protobuf/proto"
)

var envelope *Envelope
var envelope2 *Envelope2

var bytesProto []byte
var bytesJSON []byte

func TestMain(m *testing.M) {
	log.Println("in TestMain")
	envelope = GenerateEnvelope()
	envelope2 = FromEnvelope(envelope)

	bytesProto, _ = proto.Marshal(envelope)
	_, _ = ffjson.Marshal(envelope2)
	_, _ = json.Marshal(envelope2)
	bytesJSON, _ = json.Marshal(envelope)
	var env Envelope
	var env2 Envelope2
	proto.Unmarshal(bytesProto, &env)
	json.Unmarshal(bytesJSON, &env)
	ffjson.Unmarshal(bytesJSON, &env2)
	json.Unmarshal(bytesJSON, &env2)
	m.Run()
}

func BenchmarkProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(envelope)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkProtoUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var env Envelope
		err := proto.Unmarshal(bytesProto, &env)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(envelope)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var env Envelope
		err := json.Unmarshal(bytesJSON, &env)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFFJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ffjson.Marshal(envelope2)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFFJSONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var env2 Envelope2
		err := ffjson.Unmarshal(bytesJSON, &env2)
		if err != nil {
			b.Error(err)
		}
	}
}
