package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func Run(rootPath string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(fmt.Sprint(rootPath, "/public"), r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.ServeFile(
				w,
				r,
				fmt.Sprint(rootPath, "/public/index.html"),
			)
			return
		}
		http.FileServer(http.Dir(fmt.Sprint(rootPath, "/public"))).
			ServeHTTP(w, r)
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
