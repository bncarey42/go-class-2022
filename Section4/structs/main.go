package main

import "fmt"

type contactinfo struct {
   email string
   zipCode int
}

type person struct {
   firstname string
   lastname string
   contact contactinfo
}

func main() {
/*   geoff := person{"Geoff","Wicklowe"}
   gary := person{firstname: "Gary", lastname:"Garyson"}

   fmt.Println(geoff)
   fmt.Println(gary)

   var jim person
   fmt.Println(jim)
   fmt.Printf("%+v\n",jim)
   jim.firstname = "Jim"
   jim.lastname = "Jimson"
   fmt.Printf("%+v\n",jim)*/

   geoff := person{
      firstname: "Geoff",
      lastname: "Wicklowe",
      contactinfo:contactinfo{

         email: "geoff@geoff.com",
         zipCode: 56842,
      },
   }
   geoff.print()
   geoff.updateName("Geoffrey")
   geoff.print()
}

func (p person) print() {
   fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
   (*pointerToPerson).firstname = newFirstName
}
