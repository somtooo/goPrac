package main

import (
	"fmt"
)


func main()  {
    var fruits [5]string;
    fruits[0] = "papaya"
    fruits[1] = "goat balls"
    fruits[2] = "goat lo"
    fruits[3] = "idiota"
    fruits[4] = "plum"

    // value semantic form of for range.
    for index, fruit := range fruits {
        //fruits[2] = "goat meat"
        fmt.Println(index,fruit)
    }

    fmt.Println(fruits[2])
    fmt.Println("\n")

    //pointer semantic of the for range
    for index := range fruits{
        fruits[1] = "goat fish"
        fmt.Println(fruits[index])
    }



    //Slice creation


    fmt.Println("\n")

    var nillSLice []string
    //var es struct{}
    //nillSLicee := []string{}

    fmt.Println("Nill slice creaiton", nillSLice)
    fmt.Println("\n")

    fmt.Println("Slice creation with more capacity")

    moreCap := make([]string,5,8)
    moreCap[0] = "Apple"
	moreCap[1] = "Orange"
	moreCap[2] = "Banana"
    inspect(moreCap)
    fmt.Println("\n")


    fmt.Println("Lets talk append")
    var data []string
    appendTest(data)



    fmt.Println("\nSlice Creation.....")
    slice := createSlice();
    fmt.Println("\nSlicing a Slice")
    sliceOfSlice := slice[2:4:4]
    sliceOfSlice = append(sliceOfSlice,"changed")
    // So the question becomes how do you modify a slice without changing the shared array.

    inspect(slice)
    fmt.Println("\n")
    inspect(sliceOfSlice)



}

//Create slice of fruits.
 func  createSlice() []string {
    fruit := make([]string, 5)
    fruit[0] = "Apple"
	fruit[1] = "Orange"
	fruit[2] = "Banana"
	fruit[3] = "Grape"
	fruit[4] = "Plum"

    return fruit;
}

func inspect(slice []string) {
    fmt.Printf("Lenght[%d], Capacity[%d]\n", len(slice), cap(slice))
    for index, v := range slice {
        fmt.Printf("[%d] %p %s \n", index, &slice[1], v)
    }
}

