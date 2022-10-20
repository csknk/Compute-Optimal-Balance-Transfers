package main

import (
	"fmt"
	"log"
)

const (
	existentialDeposit = int64(1000)
	edsToRetain        = 2
)

func main() {
	target := int64(5001)
	inputs := []int64{4000, 2000}
	outputs, err := UnderlyingForSigning(target, inputs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", inputs)
	fmt.Printf("%v\n", outputs)

}

func UnderlyingForSigning(target int64, inputs []int64) (outputs []int64, err error) {

	remainder := int64(target % existentialDeposit)
	fee := int64(1000)
	sendingAmount := int64(0)
	// split amount
	// EDs in input[0]
	edsInTarget := target / existentialDeposit
	firstAmount := (edsInTarget - edsToRetain) + remainder
	secondAmount := target - firstAmount
	outputsFirst, _ := buildInputsOutputs(firstAmount, inputs)
	outputsSecond, _ := buildInputsOutputs(secondAmount, inputs)

}

func buildInputsOutputs(target int64, inputs []int64) (outputs []int64, err error) {

	for i, input := range inputs {

		canPayAllRetainingED := inputAmount-fee-existentialDeposit >= target
		canPayAllRetainingDust := inputAmount-fee > target && !canPayAllRetainingED
		canPayAllRetainingZero := inputAmount-fee == target
		canPaySomeRetainingED := inputAmount-fee-existentialDeposit < target
		canPaySomeRetainingZero := inputAmount-fee < target

	}
}
