package main

import "fmt"

func appendTest(data []string) {
	lastCap := cap(data);
	fmt.Println("First cap: ", lastCap)
	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)


	if lastCap != cap(data) {
		capChg := float64(cap(data) - lastCap) / float64(lastCap) * 100

		lastCap = cap(data)
		fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
			&data[0],
			record,
			cap(data),
			capChg)

	}
}
}