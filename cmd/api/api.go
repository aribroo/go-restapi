package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aribroo/go-ecommerce/repository"
	"github.com/aribroo/go-ecommerce/service/user"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{addr, db}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()

	userRepository := repository.NewUserRepository(s.db)

	userHandler := user.NewHandler(userRepository)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
