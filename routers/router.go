package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	_ "go-rest-mongodb/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// Routers godoc
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath /
func Routers() *mux.Router {
	//StrictSlash defines the trailing slash behavior for new routes. The initial value is false.
	//When true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	r := mux.NewRouter().StrictSlash(true)

	s := r.PathPrefix("/api").Subrouter()
	AddPlacesRouter(s)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
