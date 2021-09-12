package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config   *Config
	logLevel *logrus.Logger
	router   *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config:   config,
		logLevel: logrus.New(),
		router:   mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logLevel.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logLevel.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.getTodoList())
}

func (s *APIServer) getTodoList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world")
	}
}
