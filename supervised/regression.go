package supervised

import "fmt"

// SimpleLinearRegression calculates the dependent var given independent var
func SimpleLinearRegression(x, thetaOne, thetaZero float32) (yHat float32) {
	yHat = thetaZero + (thetaOne*x)
	fmt.Println(yHat)
	return yHat
}

// BestFitParameters mathematically determines the best fit slope and intercept given x and y coordinates
func BestFitParameters(x, y *[]float32) (thetaOne, thetaZero float32) {
	xBar := Mean(x)
	yBar := Mean(y)
	numerator := float32(0)
	denominator := float32(0)
	for i := 0; i < len(*x); i++ {
		xNorm := (*x)[i]-xBar
		yNorm := (*y)[i]-yBar
		numerator += xNorm * yNorm
		denominator += xNorm * xNorm
	}
	thetaOne = numerator / denominator
	thetaZero = yBar - (thetaOne*xBar)
	return thetaOne, thetaZero
}

// Mean calculates the mean of a slice of numbers
func Mean(x *[]float32) float32 {
	runningSum := float32(0)
	n := len(*x)
	for _, v := range *x {
		runningSum += v
	}
	return runningSum/float32(n)
}