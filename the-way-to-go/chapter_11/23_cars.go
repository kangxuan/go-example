package main

import (
	"fmt"
)

type Any interface{}

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car

func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	// query:
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewCars := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && car.BuildYear > 2010
	})

	fmt.Printf("AllCars: %v\n", allCars)
	fmt.Printf("New BMWs: %v\n", allNewCars)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

// Process all cars with the given function f:
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// 根据条件查询数据
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// 排序
func MakeSortedAppender(manufacturers []string) (func(c *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}

	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}

//AllCars: [0xc000072180 0xc0000721b0 0xc0000721e0 0xc000072210]
//New BMWs: [0xc0000721b0]
//Map sortedCars:  map[Aston Martin:[] BMW:[0xc0000721b0 0xc000072210] Default:[0xc0000721e0] Ford:[0xc000072180] Jaguar:[] Land Rover:[]]
//We have  2  BMWs
