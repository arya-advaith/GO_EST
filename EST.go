package main

import ( "fmt";"os" )

func  main(){
	fmt.Println("+-----------------------------------------------------------------------------------------------------+")
	fmt.Println("")
	fmt.Println("		This is an experimental code to calculate the HF of a given molecule	")
	fmt.Println("")
	fmt.Println("+-----------------------------------------------------------------------------------------------------+")
	fmt.Println("")
	fmt.Println("")
	if len(os.Args) < 2 {
		fmt.Println("				ERROR_l01: Please provide a file name			")
		return
	}else{
		fmt.Printf("Opening the file %s",os.Args[1])
	}

}
