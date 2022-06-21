package main

import "fmt"
func main()  {
	route := route2.Route {
		ID: "1",
		CLientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}