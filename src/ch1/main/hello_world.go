package main

import ("fmt"
"os")

func main(){
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	fmt.Println("Hello World")
	// os.Exit(-1)
}