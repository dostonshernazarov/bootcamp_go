package main

import (
)


func main(){
	groupCity := map[int][]string{
		10:   {"Деревня", "Село"},
		100:  {"Город", "Большой город"}, 
		1000: {"Миллионик"},
	 }
	 cityPopulation := map[string]int{
		"Село" : 100,
		"Миллионик" : 500,
	 }


	for _, city := range groupCity[1000] {
		delete(cityPopulation, city)
	}

	for _, city := range groupCity[10] {
		delete(cityPopulation, city)
	}
}