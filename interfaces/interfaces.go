package interfaces

import (
	oerror "github.com/betit/orion-go-sdk/error"
	"github.com/betit/orion-go-sdk/logger"
)

// Codec interface
type Codec interface {
	Encode(...interface{}) ([]byte, error)
	Decode([]byte, ...interface{}) error
}

// Transport interface
type Transport interface {
	Listen(func())
	Publish(string, []byte) error
	Subscribe(string, string, func([]byte)) error
	Handle(string, string, func([]byte) []byte) error
	Request(string, []byte, int) ([]byte, error)
	Close()
	OnClose(interface{})
}

// Response interface
type Response interface {
	GetError() *oerror.Error
	SetError(*oerror.Error) Response
	ParsePayload(interface{}) error
	SetPayload(interface{}) error
}

// Request interface
type Request interface {
	GetPath() string
	SetPath(string) Request
	GetID() string
	SetID(string) Request
	GetTracerData() map[string][]string
	SetTracerData(map[string][]string) Request
	GetTimeout() *int
	SetTimeout(int) Request
	GetMeta() map[string]string
	SetMeta(map[string]string) Request
	GetMetaProp(key string) string
	SetMetaProp(key, value string) Request
	GetParams() []byte
	ParseParams(interface{}) error
	SetParams(interface{}) error
	SetError(error) Request
}

// Tracer interface
type Tracer interface {
	Trace(Request) func()
}

// Logger interface
type Logger = logger.Logger
