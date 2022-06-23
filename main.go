package main

import "salary-calculator/tax"

func main() {

	startSalary := 3000
	diff := 600
	for i := 1; i < 20; i++ {
		salary := startSalary + (i * diff)
		selectedI, realSalary := tax.Calc(float64(salary))
		tax.PrintResult(float64(salary), selectedI, realSalary)
	}
}
