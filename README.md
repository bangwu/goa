# goa is a go library for developer to easy to start up a server

```go

package main

import (
	"github.com/goa"
	"log"
	"net/http"
)


func main()  {
	app := goa.NewGoa()
	app.Route("/foo", fooHandler)
	app.Route("/bar", barHandler)
	_ = app.ListenAndServe(":3001")
}

func fooHandler(response http.ResponseWriter, _ *http.Request)  {
	lengthOfTheResponse,_ := response.Write([]byte("foo"))
	log.Print(lengthOfTheResponse)
}

func barHandler(response http.ResponseWriter, _ *http.Request)  {
	lengthOfTheResponse,_ := response.Write([]byte("bar"))
	log.Print(lengthOfTheResponse)
}


```
