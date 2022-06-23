package tax

import (
	"fmt"
	"math"
)

var taxBases = [][]float64{
	{0, 1200},
	{1200, 4600},
	{4600, 8800},
	{8800, 15000},
	{15000, 3000},
	{30000, 5000},
	{50000, math.Inf(1)},
}

var taxRates = []float64{
	0.06, 0.15, 0.24, 0.35, 0.38, 0.4, 0.42,
}

var progressiveDeductions = []float64{
	0, 108, 522, 1490, 1940, 2540, 3540,
}

func Calc(salary float64) (selectdI int, realSalary float64) {
	for i, taxBase := range taxBases {
		if len(taxBase) > 2 {
			panic("taxBase is worng")
		}

		lowerLimit := taxBase[0]
		upperLimit := taxBase[1]
		if salary >= lowerLimit && salary < upperLimit {
			taxRate := taxRates[i]
			deduction := progressiveDeductions[i]

			return i, salary - (salary * taxRate) + deduction
		}
	}

	return 0, 0
}

func PrintResult(salary float64, selectdI int, realSalary float64) {
	lowerLimit := taxBases[selectdI][0]
	upperLimit := taxBases[selectdI][1]
	taxRate := taxRates[selectdI]
	deduction := progressiveDeductions[selectdI]
	realMonthSalary := realSalary / 12

	fmt.Printf(`
		-----
		연 임금: %d만원
		과세표준 : %d ~ %d만원,
		세율: %f
		누진공제: %d만원
			년 실수령액: %d만원
			월 실수령액: %d만원
		-----			
	`, int(salary), int(lowerLimit), int(upperLimit), taxRate, int(deduction), int(realSalary), int(realMonthSalary))

}
