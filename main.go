package main
import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r*http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful")
	name:= r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func hellohandler(w http.ResponseWriter, r*http.Request){
if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
}
if r.Method != "GET"{
		http.Error(w, " MEthod is not supported", http.StatusNotFound)
		return
}
fmt.Fprintf(w, "hello!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my Go web server!")
    })
	http.HandleFunc("/form.html", func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form.html")
})
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
