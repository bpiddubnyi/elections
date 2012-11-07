package main

func Round(n float64) float64 {
	buf := int32(n * 10)
	return float64(buf)/10.0
}
