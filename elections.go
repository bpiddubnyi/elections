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
        fmt.Printf("%s [%d-%d]:\n", region.Name, region.FirstDist, region.FirstDist + region.DistCount -1)
        for d, dist := range region.Districts {
            fmt.Printf("  ОВК-%d:\n", d)
            for _, prec := range dist.Precincts {
                fmt.Printf("    ВД-%d [%d/%d/%f%%]\n", prec.Number, prec.VotersTotal, prec.VotersVoted, prec.VotedPerc)
            }
        }
    }
}
