package main

import (
   "fmt"
   "net/http"
   "os"
   "io"
)

type logWriter struct{}


// This is document for the main func
// It prints out the reponse from the openbrewerydb api
func main() {
   resp,err := http.Get("https://api.openbrewerydb.org/breweries")
   if err != nil {
      fmt.Println("ERROR:", err)
      os.Exit(1)
   }

   /*bs:= make([]byte, 99999)
   resp.Body.Read(bs)
   */

   lw:= logWriter{}

   io.Copy(lw, resp.Body)
//   fmt.Println(string(bs))
}


// Write takes a slice of bytes and returns an int64 and an error
func (logWriter) Write(bs []byte) (int, error) {
   fmt.Println(string(bs))
   fmt.Println("Just wrote this many bytes:", len(bs))
   return len(bs), nil
}
