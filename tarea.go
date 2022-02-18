package main

import "fmt"

func basedehechos(sujeto string) {
	hombres := make(map[string]string)
	hombres["luis"] = "bailar"
	hombres["pedro"] = "bailar"
	hombres["mateo"] = "bailar"
	hombres["juan"] = "comer"

	mujeres := make(map[string]string)
	mujeres["ana"] = "comer"
	mujeres["maria"] = "bailar"

	for sujeto, gusta := range hombres {
		if gusta == "bailar" {
			fmt.Println("maria estima a ", sujeto)
		} else {
			fmt.Println(false)
		}
	}

}

func main() {
	basedehechos("luis")

}
