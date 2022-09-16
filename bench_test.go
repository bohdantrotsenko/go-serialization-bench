package btbench

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/amzn/ion-go/ion"
	"github.com/pquerna/ffjson/ffjson"
	"google.golang.org/protobuf/proto"
)

var envelope *Envelope
var envelope2 *Envelope2
var envelope3 *Envelope3

var bytesProto []byte
var bytesJSON []byte
var bytesION []byte

func TestMain(m *testing.M) {
	log.Println("in TestMain")
	envelope = GenerateEnvelope()
	envelope2 = FromEnvelope(envelope)
	envelope3 = FromEnvelope3(envelope)
	var err error

	bytesProto, _ = proto.Marshal(envelope)
	_, _ = ffjson.Marshal(envelope2)
	_, _ = json.Marshal(envelope2)
	_, _ = json.Marshal(envelope3)
	bytesJSON, _ = json.Marshal(envelope)
	bytesION, err = ion.MarshalBinary(envelope2)
	if err != nil {
		panic(err)
	}
	var env, envION Envelope
	var env2 Envelope2
	var env3 Envelope3
	proto.Unmarshal(bytesProto, &env)
	json.Unmarshal(bytesJSON, &env)
	json.Unmarshal(bytesJSON, &env3)
	ffjson.Unmarshal(bytesJSON, &env2)
	json.Unmarshal(bytesJSON, &env2)
	err = ion.Unmarshal(bytesION, &envION)
	if err != nil {
		panic(err)
	}
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

func BenchmarkION(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ion.MarshalBinary(envelope2)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkIONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var env Envelope2
		err := ion.Unmarshal(bytesION, &env)
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

func BenchmarkEasyJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(envelope3)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEasyJSONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var env Envelope3
		err := json.Unmarshal(bytesJSON, &env)
		if err != nil {
			b.Error(err)
		}
	}
}
