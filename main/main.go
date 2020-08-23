package main

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "github.com/osamingo/checkdigit"

type UidDetails struct{
	Uid string           `json:"uid"`
	
}
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
	uidDetails := UidDetails{}
	_= json.Unmarshal(body,&uidDetails)
	v := checkdigit.NewVerhoeff()
	responseString:="No"
	if v.Verify(uidDetails.Uid)
	{
	responseString = "Yes"
	}

	w.Write([]byte(responseString))
}