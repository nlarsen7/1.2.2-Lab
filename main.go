package main

import"fmt"

func main() {
 var answer float64=37
 fmt.Print("choose a number between 50 and 99") 
 var user float64
 fmt.Scanf("%f",&user)

 var factor float64=99-answer
 var guess float64= factor + user
 
 var magic float64= user - ((guess-100)+1) 

 fmt.Println(magic)

 
 

}