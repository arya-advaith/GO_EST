package main

import("fmt";"bufio";"os";"math")

type atom struct{
	Znuc float64
	x float64
	y float64 
	z float64 
}
func read_geo(filename string,nAt int,nO int)([][]float64, float64){

	file, ferr := os.Open(filename)
	if ferr != nil{
		panic(ferr)
	}
	line := bufio.NewScanner(file)

	for dumb:=0; dumb<3; dumb++{
		line.Scan()
	}
	nO++
	atom := make([]atom, nAt)
	for i:=0; i<nAt; i++{
		line.Scan()
		fmt.Sscanf(line.Text(), "  %f    %f          %f         %f",&atom[i].Znuc, &atom[i].x, &atom[i].y, &atom[i].z)
	}
	fmt.Printf("Atom Coordinates: %f\n", atom)

	ENuc := 0.0
	rAt := make([][]float64, nAt, nAt)
	for i:=0; i<nAt;i++{
		for j:=0; j<nAt;j++{
			xpart := math.Pow(atom[i].x - atom[j].x,2)
			ypart := math.Pow(atom[i].y - atom[j].y,2)
			zpart := math.Pow(atom[i].z - atom[j].z,2)
			rs := math.Sqrt(xpart + ypart + zpart)
			rAt = append(rAt, []float64{rs})
			if i!=j{
			ENuc = ENuc + atom[i].Znuc*atom[j].Znuc/math.Sqrt(rAt[i][j])
			}
		}
	}
	fmt.Printf("Distance Matrix:  %f\n", rAt)
	fmt.Printf("Nuclear Repulsion Energy: %f\n", ENuc)
	return rAt,ENuc
}
