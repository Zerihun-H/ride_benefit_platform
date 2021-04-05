package httprouter

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// Router data will be registered to http listener
type Router struct {
	Method  string
	Path    string
	Handler httprouter.Handle
}

type routing struct {
	host           string
	domain         string
	allowedOrigins string
	routers        []Router
}

// Routers contains the functions of http handler to clean payloads and pass it the service
type Routers interface {
	Serve()
}

// Initialize is for initialize the handler
func Initialize(host, allowedOrigins string, routers []Router, domain string) Routers {
	return &routing{
		host,
		domain,
		allowedOrigins,
		routers,
	}
}

// Serve is to start serving the HTTP Listener for every domain
func (r *routing) Serve() {
	server := httprouter.New()

	for _, router := range r.routers {
		// group.Add(router.Method, router.Path, router.Handler)
		server.Handle(router.Method, router.Path, router.Handler)
	}

	logrus.WithFields(logrus.Fields{
		"host":   r.host,
		"domain": r.domain,
	}).Info("Starts Serving on HTTP")

	log.Fatal(http.ListenAndServe(r.host, server))

}
