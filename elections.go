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

    fmt.Printf("Regions total: %d\n", len(regions))
/*	for i, region := range regions {
		fmt.Printf("%d: %s [%d-%d]\n", i+1, region.Name, region.FirstDist, region.FirstDist+region.DistCount-1)

	}*/
}
