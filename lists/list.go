package list

import "fmt"

type Product struct {
	title string
	id    string
	price string
}

func main() {
	// prices := []float64{10.99, 8.99}
	// fmt.Println(prices[0:1])
	// prices[1] = 9.99
	// prices = append(prices, 5.99)

	// fmt.Println(prices, prices)

	var productList []Product

	addedProduct := Product{
		title: "My Book",
		id:    "12345",
		price: "12.99",
	}

	productList = append(productList, addedProduct)

	addedProduct = Product{
		title: "My GOD",
		id:    "12345",
		price: "12.99",
	}

	productList = append(productList, addedProduct)

	fmt.Println("productList:", productList)

	var hobbies = [3]string{"Gaming", "Cricket", "Football"}

	fmt.Println("hobbies:", hobbies)
	fmt.Println(hobbies[0])
	newHobbies := hobbies[1:3]
	fmt.Println("newHobbies:", newHobbies)

	var slicesHobbies = hobbies[0:2]
	fmt.Println(slicesHobbies)
	var newSlicesHobbies = slicesHobbies[1:3]
	fmt.Println(newSlicesHobbies)

	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)
}

// var productNames [4]string = [4]string{"A Book"}

// prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// fmt.Println(prices)
// productNames[2] = "A Carpet"
// fmt.Println(productNames)
// fmt.Println(prices[2])

// featuredPrices := prices[1:]
// featuredPrices[0] = 199.99
// highlightedPrices := featuredPrices[:1]
// fmt.Println(highlightedPrices)
// fmt.Println(prices)
// fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// highlightedPrices = highlightedPrices[:3]
// fmt.Println(highlightedPrices)
// fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
