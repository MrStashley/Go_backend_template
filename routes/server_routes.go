package routes

import (
	"fmt"
	"net/http"
)

func RouteHello (w http.ResponseWriter, req *http.Request) {
	fmt.Println("/hello received");

	fmt.Fprintf(w, "hello\n")
}

func RouteHeaders (w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h);
		}
	}
}