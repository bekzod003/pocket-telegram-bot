package server

import "net/http"

type AuthorizationServer struct {
	server *http.Server
}

// constructor
func NewAuthorizationServer(addr string, handler http.Handler) *AuthorizationServer {
	return &AuthorizationServer{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}
