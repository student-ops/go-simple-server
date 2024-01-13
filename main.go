package main

import (
	"fmt"
	"log"

	"github.com/FarStep131/go-simple-server/server"
)

const (
	host = "0.0.0.0"
	port = "8080"
)

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// レスポンスを設定
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello"}`))
	})

	// サーバーを起動
	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Server stopped with error: %s\n", err)
	}
}
