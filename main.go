package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Restaurant struct {
    Name    string `json:"name"`
    Address struct {
        FirstLine  string `json:"firstLine"`
        City       string `json:"city"`
        PostalCode string `json:"postalCode"`
    } `json:"address"`
    Rating struct {
        StarRating float64 `json:"starRating"`
    } `json:"rating"`
    Cuisines []Cuisine `json:cuisines"`
}

type Cuisine struct {
    Name string `json:"cuisine"`
}

func main() {
    //ask user to input a postal code
    var postalCode string
    fmt.Println("Enter a Postal Code")
    fmt.Scanln(&postalCode)

    //retrieve restaurant data from API
    restaurants, err := getRestaurantByPostalCode(postalCode)
    if err != nil {
        log.Fatalf("failed to retrieve restaurants: %v", err)
    }

    //display first 10 restaurant info
    for i, restaurant := range restaurants {
        if i >= 10 {
            break
        }
        fmt.Printf("Restaurant %d:\n", i+1)
        fmt.Printf("Name: %s\n", restaurant.Name)
        fmt.Printf("Address: %s, %s, %s\n", restaurant.Address.FirstLine, restaurant.Address.City, restaurant.Address.PostalCode)
        fmt.Printf("Cuisines: %s\n", restaurant.Cuisines)
        fmt.Printf("Rating: %.2f\n", restaurant.Rating.StarRating)
    }
}

func getRestaurantByPostalCode(postalCode string) ([]Restaurant, error) {
    //send GET request to the API
    resp, err := http.Get("https://uk.api.just-eat.io/discovery/uk/restaurants/enriched/bypostcode/" + postalCode)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve data from API: %v", err)
    }
    defer resp.Body.Close()

    //decode the JSON response
    var data struct {
        Restaurants []Restaurant `json:"restaurants"`
    }

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, fmt.Errorf("failed to decode JSON response: %v", err)
    }

    return data.Restaurants, nil

}