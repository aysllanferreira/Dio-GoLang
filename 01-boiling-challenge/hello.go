package main

import "fmt"

func main() {
	const boilingPointKelvin int = 373
	const kelvinToCelsiusDiff int = 273

	temperatureCelsius := boilingPointKelvin - kelvinToCelsiusDiff
	fmt.Printf("The boiling point of water is %d°C.\n", temperatureCelsius)
}
