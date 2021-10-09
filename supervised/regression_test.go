package supervised

import (
	"fmt"
	"testing"
)

var x1 = []float32{2.0, 2.4, 1.5, 3.5, 3.5, 3.5, 3.5, 3.7, 3.7}
var x2 = []float32{4, 4, 4, 6, 6, 6, 6, 6, 6}
var y = []float32{196, 221, 136, 255, 244, 230, 232, 255, 267}

var thetaOne, thetaZero = BestFitParameters(&x1, &y)

func TestMean(t *testing.T) {
	tables := []struct {
		input  *[]float32
		result string
	}{
		{&x1, "3.03"},
		{&x2, "5.33"},
		{&y, "226.22"},
	}
	for _, v := range tables {
		actual := Mean(v.input)
		result := fmt.Sprintf("%.2f", actual)
		if result != v.result {
			t.Error("Fail")
		}
	}
}

func TestBestFitParameters(t *testing.T) {
	th1, th0 := BestFitParameters(&x1, &y)
	res1 := fmt.Sprintf("%.2f", th1)
	res2 := fmt.Sprintf("%.2f", th0)
	if res1 != "43.98" {
		t.Error(res1)
	}
	if res2 != "92.80" {
		t.Error(res2)
	}
}

func TestSimpleLinearRegression(t *testing.T) {
	yHat := SimpleLinearRegression(2.4, thetaOne, thetaZero)
	if yHat != 198.36539 {
		t.Error(yHat)
	}
}