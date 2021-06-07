package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/jsonapi"
	"github.com/leg100/ots"
)

// ShutdownTimeout is the time given for outstanding requests to finish before shutdown.
const ShutdownTimeout = 1 * time.Second

type Server struct {
	server *http.Server
	router *mux.Router
	ln     net.Listener
	err    chan error

	// Listening Address in the form <ip>:<port>
	Addr string

	OrganizationService ots.OrganizationService
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		err:    make(chan error),
	}

	http.Handle("/", NewRouter(s))

	return s
}

func NewRouter(server *Server) *mux.Router {
	router := mux.NewRouter()

	// Filter json-api requests
	sub := router.Headers("Accept", jsonapi.MediaType).Subrouter()

	sub.HandleFunc("/organizations", server.ListOrganizations).Methods("GET")
	sub.HandleFunc("/organizations/{name}", server.GetOrganization).Methods("GET")
	sub.HandleFunc("/organizations/{name}", server.CreateOrganization).Methods("POST")
	sub.HandleFunc("/organizations/{name}", server.UpdateOrganization).Methods("PATCH")
	sub.HandleFunc("/organizations/{name}", server.DeleteOrganization).Methods("DELETE")
	sub.HandleFunc("/organizations/{name}/entitlement-set", server.GetEntitlements).Methods("GET")

	return router
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() (err error) {
	if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
		return err
	}

	// Begin serving requests on the listener. We use Serve() instead of
	// ListenAndServe() because it allows us to check for listen errors (such as
	// trying to use an already open port) synchronously.
	go func() {
		s.err <- s.server.Serve(s.ln)
	}()

	return nil
}

// Port returns the TCP port for the running server.
// This is useful in tests where we allocate a random port by using ":0".
func (s *Server) Port() int {
	if s.ln == nil {
		return 0
	}
	return s.ln.Addr().(*net.TCPAddr).Port
}

// Wait blocks until server stops listening or context is cancelled.
func (s *Server) Wait(ctx context.Context) error {
	select {
	case err := <-s.err:
		return err
	case <-ctx.Done():
		return s.server.Close()
	}
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}