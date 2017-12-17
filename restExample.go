// go get ./...
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

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


/*
GET /{key}
PUT /{key}

PUT - value will be in body. if key doesn't exist create, if it does overwrite
GET - return value in body if key exists, if not return 404
/key1


GEt:
	 /1
	    404


	    POST - create



PUT /1

  	body - string
    map --key , value


   Mutex / channel   /rw mutex
     Map

     1 hr video tutorial
     study rest call - 2 hr
     2 hr


	 github.com/gorilla/mux - much better than standard library muxer

*/












