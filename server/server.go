package server

import (
	"fmt"
	"net/http"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn"
	"github.com/celer-network/im-executor/svc"
)

type RestServer struct {
	port        int
	sgn         *sgn.SgnClient
	executorSvc *svc.ExecutorService
}

func NewRestServer(port int, sgnClient *sgn.SgnClient, executorSvc *svc.ExecutorService) *RestServer {
	server := &RestServer{
		port:        port,
		sgn:         sgnClient,
		executorSvc: executorSvc,
	}
	server.registerRoutes()
	return server
}

func (s *RestServer) Start() {
	if s == nil {
		log.Fatalf("server uninitialized")
		return
	}
	log.Infof("Start http server and listen on localhost:%d", s.port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

func (s *RestServer) registerRoutes() {
	http.HandleFunc("/get-message", handler(s.handleGetMessage))
	http.HandleFunc("/unstuck", handler(s.handleUnstuckTx))
	http.HandleFunc("/revert-execution-status", handler(s.handleRevertExecutionStatus))
}
