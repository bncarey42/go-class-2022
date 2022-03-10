package main

import ("fmt")

type contactInfo struct {
   email string
   zip int
}

type person struct {
   first string
   last string
   contact contactInfo
}

func main() {
   // 
   geoff := &person{first: "Geoff", last:"Whicklowe", contact: contactInfo{email: "geoff@techonerds.io", zip: 55310}}
   fmt.Printf("%+v\n", *geoff)

   geoff.updatePersonName("Geoffrey")
   fmt.Printf("%+v\n", *geoff)
}


func (p *person) updatePersonName(name string) {
   p.first = name
}
