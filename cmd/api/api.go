package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thangsuperman/bee-happy/services/like"
	"github.com/thangsuperman/bee-happy/services/post"
	"github.com/thangsuperman/bee-happy/services/upload"
	"github.com/thangsuperman/bee-happy/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	postStore := post.NewStore(s.db)
	postHandler := post.NewHandler(postStore, userStore)
	postHandler.RegisterRoutes(subrouter)

	likeStore := like.NewStore(s.db)
	likeHandler := like.NewHandler(likeStore, userStore)
	likeHandler.RegisterRoutes(subrouter)

	uploadHandler := upload.NewHandler(userStore)
	uploadHandler.RegisterRoutes(subrouter)

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	)).Methods(http.MethodGet)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
