package main

/** Seems like precision of 10 is ok, 
  * i.e. plots looks reasonable and meaningful */
func Round(n float64, precision int) float64 {
	buf := int(n * float64(precision))
	return float64(buf) / float64(precision)
}
