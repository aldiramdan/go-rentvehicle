package libs

import "math"

func CalculateNewRating(totalRent int, currentRating, newRating float64) float64 {

	newRating = (currentRating*float64(totalRent) + newRating) / float64(totalRent+1)

	roundRating := math.Round(newRating*10) / 10

	return roundRating

}
