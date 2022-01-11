package main

import (
	"fmt"
	"net/http"
	"simple-go-server/routes"
)

func main() {
	http.HandleFunc("/hello", routes.RouteHello);
	http.HandleFunc("/headers", routes.RouteHeaders);

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html");
	} );

	fmt.Println("starting up server on port 8090")

	http.ListenAndServe(":8090", nil);

}
