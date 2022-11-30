package main

import "fmt"

func main() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

//Map item: Capital of France is Paris
//Map item: Capital of Italy is Rome
//Map item: Capital of Japan is Tokyo
