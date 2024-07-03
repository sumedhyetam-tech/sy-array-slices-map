package main

import "fmt"

func main() {
	//userNames := []string{}
	userNames := make([]string, 2, 5)
	fmt.Printf("Address of slices 0 %p\n", &userNames)
	userNames[0] = "Sumedh"
	userNames[1] = "Sumedh"
	fmt.Printf("Address of slices 1 %p\n", &userNames)
	userNames = append(userNames, "Max")
	fmt.Printf("Address of slices 2 %p\n", &userNames)
	userNames = append(userNames, "Manuel")
	fmt.Printf("Address of slices 3 %p\n", &userNames)
	userNames = append(userNames, "Max")
	fmt.Printf("Address of slices 4 %p\n", &userNames)
	userNames = append(userNames, "Manuel")
	fmt.Printf("Address of slices 5 %p\n", &userNames)
	userNames = append(userNames, "Manuel")
	fmt.Printf("Address of slices 6 %p\n", &userNames)
	userNames = append(userNames, "Manuel")
	fmt.Printf("Address of slices 7 %p\n", &userNames)
	fmt.Println(userNames)

	//courseRatings := map[string]float64{}
	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.6
	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

}
