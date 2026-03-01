package main

import ("fmt";"bufio";"os")

func read_mol(inputs string)(int,int){

	file,ferr := os.Open(inputs)
	if ferr != nil{
		panic(ferr)
		}
	
	line := bufio.NewScanner(file)
		fmt.Println(line.Text())

	var nAt,nEl int
	line.Scan()
	if line.Text() == "# nAt nEl"{
		fmt.Println("Number of atoms and electrons found")
		line.Scan()
		fmt.Sscanf(line.Text(), "%d  %d", &nAt, &nEl)
	}

	if nEl%2!=0{
		fmt.Println("This is not a closed shell system!")
		return 0,0
	}
	var nO int
	nO = nEl/2
	fmt.Println("Number of atoms: ", nAt)
	fmt.Println("Number of occupied orbitals: ", nO)
	

	defer file.Close()
	
return nAt, nO
}
