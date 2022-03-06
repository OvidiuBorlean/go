package main

import (
	"fmt";
        "net/http";
 	"log"
)

func main(){
   

   resp, err := http.Get("http://google.ro")
   if err != nil {
	log.Fatalln(err)
   }
   fmt.Println(resp)
}
