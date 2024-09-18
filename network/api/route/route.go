package route

import (
	"net/http"
	"network/api/handler"
	"network/internal/service"
	"network/pkg/keytap"
	"network/pkg/logger"

	"github.com/sirupsen/logrus"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var routes []Route

func RegisterRoute(path string, method string, handler http.HandlerFunc) {
	routes = append(routes, Route{path, method, handler})
	http.HandleFunc(path, handler)
}

func ListRoutes(log *logrus.Logger) {
	for _, route := range routes {
		log.Infof("Registered route method: %s path: %s", route.Method, route.Path)
	}
}

func InitRoute() {
	logger := logger.InitLogger()
	// init keybonding for linux
	keybonding, err := keytap.NewKeyBonding()
	if err != nil {
		logger.Fatal(err)
	}
	// init vlc service
	vlcService := service.NewVlcService(logger, &keybonding)
	// init handler
	handler := handler.NewHandler(logger, vlcService)

	// register routes
	RegisterRoute("/vlc", http.MethodPost, handler.VlcCommand)
	ListRoutes(logger)

	// init router
	logger.Info("Init route on port 8080")
	logger.Fatal(http.ListenAndServe(":8080", nil))
}
