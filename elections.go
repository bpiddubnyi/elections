package main

import (
	"fmt"
)

func main() {
	regions, err := GetRegions()
	if err != nil {
		fmt.Errorf("Failed to get regions info: %v\n", err)
		return
	}

	for i, region := range regions {
        fmt.Printf("%d: %s [%d]\n", i + 1, region.Name, region.Id)
	}
}
