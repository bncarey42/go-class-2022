package main

import "fmt"

func main() {
   colors:= map[string]string{
      "red": "#ff000",
      "green": "#4bf745",
      "white": "#fefefe",
   }

   colors["yellow"] = "#89a12e"
   printMap(colors)
   delete(colors, "yellow")

   if _, ok := colors["yellow"]; !ok {
      fmt.Println("NOPE")
   }
   printMap(colors)
}

func printMap(m map[string]string) {
   for color, hex := range m {
      fmt.Println("Hex for color", color, "is", hex)
   }
}
