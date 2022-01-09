package main

import (
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("http start"))
	w.Header().Set("Content-Type","application/json")
	for k,v := range r.Header{
		w.Header().Add(k,v[0])
	}

}
func HealthzHandler(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}

func main()  {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/healthz", HealthzHandler)
	err:= http.ListenAndServe(":8000",nil)

	if err != nil{
		log.Fatal("ListenAndServe:", err)
	}



}