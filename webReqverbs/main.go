package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Requesting get in golang ")
	performGetRequest()
	performPostRequest()
	performPostFormRequest()
}

 func performGetRequest() {
 	const myURL = "http://localhost:5000/get"

 	response, err := http.Get(myURL)

 	if err != nil {
 		panic(err)
 	}
 	defer response.Body.Close()

 	fmt.Println("Status Code : ", response.StatusCode)
 	fmt.Println("Content length is: ", response.ContentLength)

 	var responseString strings.Builder
 	content, _ := ioutil.ReadAll(response.Body)
 	byteCount, _ := responseString.Write(content)

 	fmt.Println("ByteCount is: ", byteCount)
 	fmt.Println(responseString.String())
 }

func performPostRequest(){
	const myURL = "http://localhost:5000/post"

	requestBody := strings.NewReader(`
		{
			"coursename": "golang",
			"price": 0,
			"platform": "Youtube"
		}
	`)
	response, err:= http.Post(myURL, "application/json",requestBody)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performPostFormRequest(){
	const myURL = "http://localhost:5000/postform"

	data := url.Values{}
	data.Add("firstname", "Dileep")
	data.Add("lastname", "sajja")
	data.Add("Age", "20")

	response, err := http.PostForm(myURL,data)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
