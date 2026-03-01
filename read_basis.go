package main

import ("fmt";"bufio";"os")

func main(){

	nAt,nO := read_mol("input/molecule")

	rAt,ENuc := read_geo("input/molecule",nAt,nO)
	fmt.Println("This is ENuc: ", ENuc)

	file,ferr := os.Open("input/basis")
	if ferr != nil{
		panic(ferr)
		}
	
	line := bufio.NewScanner(file)
	fmt.Println(line.Text())

	nShell := 0
	fmt.Printf("This is nAt and nO: %d %d\n", nAt, nO)
	
	var iAt,nShAt int
	var TotAngMomShell [50]int
	CenterShell :=  make([][]float64,0)
	var KShell [50]int
	var shelltype string
	var ExpShell [50][10]float64
	var DShell [50][10]float64
		
	for i:=0;i<nAt;i++{
		line.Scan()
		fmt.Sscanf(line.Text(), "%d %d", &iAt, &nShAt)
		fmt.Println("This is the Atom number and number of shells: ", iAt, nShAt)
		for j:=0;j<nShAt;j++{
			nShell++
		 	CenterShell= append(CenterShell,rAt[iAt])
			line.Scan()
			fmt.Sscanf(line.Text(), "%s %d", &shelltype, &KShell[nShell])
        
			if(shelltype == "S"){
          			TotAngMomShell[nShell] = 0
          			fmt.Printf("s-type shell with K = %d\n",KShell[nShell])
  			}else if(shelltype == "P"){
          			TotAngMomShell[nShell] = 1
          			fmt.Printf("p-type shell with K = %d\n",KShell[nShell])
  			}else if(shelltype == "D"){
        	  		TotAngMomShell[nShell] = 2
          			fmt.Printf("d-type shell with K = %d\n",KShell[nShell])
  			}else if(shelltype == "F"){
          			TotAngMomShell[nShell] = 3
          			fmt.Printf("f-type shell with K = %d\n",KShell[nShell])
  			}else if(shelltype == "G"){
          			TotAngMomShell[nShell] = 4
          			fmt.Printf("g-type shell with K = %d\n",KShell[nShell])
  			}else if(shelltype == "H"){
          			TotAngMomShell[nShell] = 5
          			fmt.Printf("h-type shell with K = %d\n",KShell[nShell])
  			}else if(shelltype == "I"){
          			TotAngMomShell[nShell] = 6
          			fmt.Printf("i-type shell with K = %d\n",KShell[nShell])
				}
        		fmt.Println("Exponents\tContraction")
			for k:=0;k<KShell[nShell];k++{
        		  	line.Scan()
				fmt.Sscanf(line.Text(),"%f %f",&ExpShell[nShell][k],&DShell[nShell][k])
          			fmt.Printf("%f\t%f\n",ExpShell[nShell][k],DShell[nShell][k])
  			}				
		}
		fmt.Println("+--------------------------------------------------------------------------------------------------+")
	}

 	fmt.Printf("Number of shells %d\n",nShell)
	fmt.Println("+--------------------------------------------------------------------------------------------------+")
  	fmt.Println("")
	file.Close()
	nBas := 0
  	for iShell:=0;iShell<nShell;iShell++{
    		nBas = nBas + (TotAngMomShell[iShell]*TotAngMomShell[iShell] + 3*TotAngMomShell[iShell] + 2)/2
  	}
	fmt.Println("+--------------------------------------------------------------------------------------------------+")
  	fmt.Printf("Number of basis functions %d\n",nBas)
  	fmt.Println("+--------------------------------------------------------------------------------------------------+")
	fmt.Println("")


	nV := nBas - nO 
	fmt.Printf("Number of virtual orbitals %d\n",nV)
}
