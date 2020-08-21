package main

import "fmt"
import "net/http"
import "io/ioutil"
func main(){
	fmt.Println("server creation")

	mux := http.NewServeMux()

	mux.HandleFunc("/hello",hello)

	err := http.ListenAndServe(":8080", mux)
	if err!= nil{
		fmt.Println("Failed")
	}

}

func hello(w http.ResponseWriter, req *http.Request){
	fmt.Println("Request Received!")

	body,_ := ioutil.ReadAll(req.Body)
	responseString := "Hello, "+ string(body)

	w.Write([]byte(responseString))
}