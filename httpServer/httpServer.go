package httpServer

import (
	"fmt"
	"net/http"
	"os"
)

func InitHttpServer() {

	path := fmt.Sprintf("/%s", os.Getenv("SERVICE_PATH"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

	})
}
