package main
import (
	"fmt"
	"net/http"
)
import "exercises/app/controller"
import 	"github.com/gorilla/mux"

func main() {

	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("").Subrouter()
	sub.Methods("GET").Path("/").HandlerFunc(controller.GetDate)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", router)

}