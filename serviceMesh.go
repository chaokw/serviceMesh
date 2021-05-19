package serviceMesh

import (
	"net/http"
	"github.com/chaokw/serviceMesh/registry"
	"github.com/chaokw/serviceMesh/transport/rest"
	restSelector "github.com/chaokw/serviceMesh/transport/rest/client/selector"
)

func NewRestServer(rg registry.Registry, handler http.Handler, opts ...rest.ServerOption) *rest.Server {
	return rest.NewSever(rg, handler, opts...)
}

func NewRestClient(name string, s restSelector.Selector, opt ...rest.ClientOption) (*rest.Client, error) {
	return rest.NewClient(name, s, opt...)
}
