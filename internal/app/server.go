package app

import (
	"github.com/gorilla/mux"
	"net/http"
	loggingCodes "user-service/internal/types/logging"
	logging "user-service/pkg/logger"
)

func CreateAppHttpServer(addr string, handler interface{}) {
	router := mux.NewRouter()

	router.Methods("OPTIONS") // Кросс доменные запросы

	router.HandleFunc("/health", handler).Methods("GET")
	router.HandleFunc("/health", handler).Methods("GET")
	router.HandleFunc("/health", handler).Methods("POST")
	router.HandleFunc("/health", handler).Methods("PUT")
	router.HandleFunc("/health", handler).Methods("DELETE")

	logger := logging.New(logging.GetDefaultConfig())

	logger.Info(loggingCodes.CodeInfoAppStarted, " user services app starting in addres "+addr, nil)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Error(loggingCodes.CodeErrorAppStartFailed, "user services failed to start", nil)
	}

}
