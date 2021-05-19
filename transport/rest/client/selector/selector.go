package selector

import (
	"errors"

	"github.com/chaokw/serviceMesh/registry"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrNoneAvailable = errors.New("none available")
)

type Selector interface {
	Options() Options
	Next(service string) (*registry.Node, error)
	Close() error
	String() string
}
