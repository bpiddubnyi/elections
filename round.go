package main

/** Seems like precision of 10 is ok, 
  * i.e. plots looks reasonable and meaningful */
func Round(n float64) float64 {
	buf := int32(n * 10)
	return float64(buf) / 10.0
}
