package handler

import (
	"encoding/json"
	"net/http"
	"network/internal/domain"
	"network/internal/service"

	"github.com/sirupsen/logrus"
)

// Handler is a interface for command handler
type Handler interface {
	VlcCommand(w http.ResponseWriter, r *http.Request)
}

// baseHandler is a base implementation of command handler
type baseHandler struct {
	logger     *logrus.Logger
	vlcService service.VlcService
}

// NewHandler creates a new command handler
func NewHandler(logger *logrus.Logger, service service.VlcService) Handler {
	return &baseHandler{
		logger:     logger,
		vlcService: service,
	}
}

// enableCors enables cors for the handler
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, or specify your allowed origins
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// NOTE: command handler for vlc
func (b *baseHandler) VlcCommand(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cmd domain.Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: validate command
	// process command
	b.logger.Infof("Received command: %s", cmd.Action)

	switch cmd.Action {
	case "start":
		err = b.vlcService.StartOrStop()
	case "volumeUp":
		err = b.vlcService.VolumeUp()
	case "volumeDown":
		err = b.vlcService.VolumeDown()
	case "next":
		err = b.vlcService.Next()
	case "previous":
		err = b.vlcService.Previous()
	case "previous5s":
		err = b.vlcService.Previous5s()
	case "next5s":
		err = b.vlcService.Next5s()
	default:
		http.Error(w, "Command not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		b.logger.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
