package networking

import (
	"fmt"
	"net/http"
)

func StartServer(port int,messageHandler func(http.ResponseWriter,*http.Request)){
	http.HandleFunc("/message",messageHandler)
	fmt.Printf("Listing on port %d...\n",port)
	if error := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); error != nil {
        fmt.Printf("Failed to start server: %s\n", error)
    }
}