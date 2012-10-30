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

    for _, region := range regions {
        fmt.Println(region.Name)
    }
}
