// go get ./...

package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
	"io/ioutil"
	"sync"
)


var myData = make(map[string]string)
var mutex sync.Mutex

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/{.*}", getValueByKey).Methods("GET")
	router.HandleFunc("/{.*}", updateCreateKeyValue).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func updateCreateKeyValue(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/text")
	vars := mux.Vars(r)
	key := vars[".*"]

	 b, _ := ioutil.ReadAll(r.Body)

	fmt.Println(key)
	fmt.Println(string(b))

	mutex.Lock()
	if _, ok := myData[key]; ok{
		myData[key] = string(b)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(myData[key]))
	} else{
		w.WriteHeader(http.StatusCreated)
		myData[key] = string(b)
		w.Write([] byte("Created"))
	}
	mutex.Unlock()

}


func getValueByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars[".*"]
	fmt.Println("Nandeshwar")
	//jsonData,_ := json.Marshal(myData)


	if value, ok := myData[key]; ok{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(value))
	} else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([] byte("Not Found"))
	}
}















