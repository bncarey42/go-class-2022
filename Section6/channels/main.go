package main

import (
   "fmt"
   "net/http"
   "time"
)


func main() {
   links := []string {
      "http://google.com",
      "http://facebook.com",
      "http://golang.com",
      "http://amazon.com",
      "http://bncarey42.github.io",
   }

   c := make(chan string)

   for _, link := range links {
      //fmt.Println("checking", link)
      go checkLink(link, c)
      //fmt.Println("done with", link)
   }

   // for i := 0; i < len(links); i++ {
   for l := range c {
      // fmt.Println(<-c)
      // go checkLink(l, c)
      go func(link string) {
         time.Sleep(5*time.Second)
         go checkLink(link, c)
      }(l)
   }
}


func checkLink(link string, cannel chan string) {
   _, err := http.Get(link)
   if err != nil {
      fmt.Println(link, "She's dead Jim!")
      cannel <- link
      return
   }
   fmt.Println(link, "IT's ALIVE!")
      cannel <- link
}
