package coffee

import (
	"fmt"

	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CoffeeList struct {
	List []CoffeeDetails `json:"coffeelist"`
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return nil, err
	}

	fmt.Printf("Raw config: %+v\n", viper.AllSettings())

	coffeeListRaw := viper.Get("coffeelist")
	if coffeeListRaw == nil {
		fmt.Println("Warning: 'coffeelist' key not found in config")
		return &CoffeeList{}, nil
	}

	var coffeeList []CoffeeDetails
	err = viper.UnmarshalKey("coffeelist", &coffeeList)
	if err != nil {
		fmt.Printf("Error unmarshaling coffeelist: %v\n", err)
		return nil, err
	}

	fmt.Printf("Unmarshaled coffee list: %+v\n", coffeeList)

	Coffees.List = coffeeList

	return &Coffees, nil
}

func IsCofeeAvailable(coffeetype string) bool {
	for _, element := range Coffees.List {
		if element.Name == coffeetype {
			result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
			fmt.Println(result)
			return true
		}
	}
	return false
}
