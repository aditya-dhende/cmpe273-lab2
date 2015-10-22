package main 


import (
    "encoding/json" 
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"  
)

type Request struct {
    Name string  `json:"name"`
   
}

type Response struct {
    Greeting string `json:"greeting"`
}


func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))

}   

func hello2(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

    op := Request{}
    //json.NewDecoder(req.Body)

    decoder := json.NewDecoder(req.Body)

   err := decoder.Decode(&op)
   if err != nil {
       panic("Error decoding")
}

    op2 := Response{}
    op2.Greeting = "Hello," + op.Name + "!"

    mj , _ := json.Marshal(op2)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", mj)


}


func main() {
    
    router := httprouter.New()  
    router.GET("/hello/:name", hello)
    router.POST("/hello", hello2)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: router,
    }
    server.ListenAndServe()
}
