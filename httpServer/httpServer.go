package httpServer

import (
	"fmt"
	"net/http"
	"os"
)

func InitHttpServer() {

	path := fmt.Sprintf("/%s/", os.Getenv("SERVICE_PATH"))

	fmt.Println(path)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		response := `{"test":"Test"}`

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})
}
