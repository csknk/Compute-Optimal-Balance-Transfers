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

	fee := int64(1000)
	//	sendingAmount := int64(0)
	// split amount
	// EDs in input[0]
	//	edsInTarget := target / existentialDeposit
	//	firstAmount := (existentialDeposit * (edsInTarget - edsToRetain)) + remainder
	//	secondAmount := target - firstAmount
	outputs, sendingAmount, _ := buildInputsOutputs(target, inputs, fee)
	//	outputsFirst, sendingAmountSubtotal1, _ := buildInputsOutputs(firstAmount, inputs, fee)
	//	outputsSecond, sendingAmountSubtotal2, _ := buildInputsOutputs(secondAmount, inputs, fee)
	//	outputs = append(outputs, outputsFirst...)
	//	outputs = append(outputs, outputsSecond...)

	fmt.Printf("sendingAmount: %d\n", sendingAmount)
	fmt.Printf("outputs: %v\n", outputs)

	return
}

func buildInputsOutputs(target int64, inputs []int64, fee int64) (outputs []int64, totalSendingAmount int64, err error) {
	fmt.Printf("target: %d\n", target)
	remainder := target % existentialDeposit

	for i, inputAmount := range inputs {
		canPayAllRetainingED := inputAmount-fee-existentialDeposit >= target
		canPayAllRetainingDust := inputAmount-fee > target && !canPayAllRetainingED
		canPayAllRetainingZero := inputAmount-fee == target
		canPaySomeRetainingED := inputAmount-fee-existentialDeposit < target
		canPaySomeRetainingZero := inputAmount-fee < target

		fmt.Printf("input %d\n", i)
		fmt.Printf("canPayAllRetainingED: %v\n", canPayAllRetainingED)
		fmt.Printf("canPayAllRetainingDust: %v\n", canPayAllRetainingDust)
		fmt.Printf("canPayAllRetainingZero: %v\n", canPayAllRetainingZero)
		fmt.Printf("canPaySomeRetainingED: %v\n", canPaySomeRetainingED)
		fmt.Printf("canPaySomeRetainingZero: %v\n", canPaySomeRetainingZero)
		fmt.Println("--------------------------------")

		sendFromThisInput := int64(0)
		if remainder != 0 {
			sendFromThisInput = remainder
		}

		if canPayAllRetainingED {
			totalSendingAmount = target
			inputs[i] = inputAmount - target - fee
			target = 0
			break
		}

		if canPayAllRetainingZero {
			sendFromThisInput := target - fee
			totalSendingAmount += sendFromThisInput
			inputs[i] = sendFromThisInput
			target -= sendFromThisInput
		}

		if canPaySomeRetainingED {
			sendFromThisInput += (inputAmount/existentialDeposit)*existentialDeposit - existentialDeposit - fee
			totalSendingAmount += sendFromThisInput
			inputs[i] = sendFromThisInput
			target -= sendFromThisInput
		}
	}
	fmt.Printf("outputs: %v\n", outputs)
	fmt.Printf("totalSendingAmount: %v\n", totalSendingAmount)

	return
}
