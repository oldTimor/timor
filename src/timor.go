package timor

import "fmt"

type Income interface {
	calculate() int
	souce() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

func (f FixedBilling) calculate() int {
	return f.biddedAmount
}

func (f FixedBilling) souce() string {
	return f.projectName
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (t TimeAndMaterial) calculate() int {
	return t.noOfHours * t.hourlyRate
}

func (t TimeAndMaterial) souce() string {
	return t.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = %d\n", income.souce(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of Oranisation = %d", netincome)
}
