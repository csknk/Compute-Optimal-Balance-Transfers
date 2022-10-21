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
	target := int64(10001)
	inputs := []int64{8000, 2000, 1010}
	outputs, err := UnderlyingForSigning(target, inputs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", inputs)
	fmt.Printf("%v\n", outputs)

}

func UnderlyingForSigning(target int64, inputs []int64) (outputs []int64, err error) {
	fee := int64(5)
	outputs, sendingAmount, _ := buildInputsOutputs(target, inputs, fee)
	fmt.Printf("sendingAmount: %d\n", sendingAmount)
	fmt.Printf("outputs: %v\n", outputs)

	return
}

func buildInputsOutputs(target int64, inputs []int64, fee int64) (outputs []int64, totalSendingAmount int64, err error) {
	for i, inputAmount := range inputs {
		if target == 0 {
			break
		}

		maxSpendRetainingSpendableBalance := inputAmount - existentialDeposit - fee

		canPayAllRetainingED := maxSpendRetainingSpendableBalance >= target
		canPayAllRetainingDust := inputAmount-fee > target && !canPayAllRetainingED
		canPayAllRetainingZero := inputAmount-fee == target
		canPaySomeRetainingED := maxSpendRetainingSpendableBalance < target && maxSpendRetainingSpendableBalance >= existentialDeposit // ACTUAL ED
		canPaySomeRetainingZero := inputAmount-fee < target

		remainder := target % existentialDeposit
		fmt.Printf("input %d, target amount: %d\n", i, target)
		sendFromThisInput := int64(remainder)

		switch {
		case canPayAllRetainingED:
			fmt.Println("canPayAllRetainingED")
			sendFromThisInput = target

		case canPayAllRetainingDust:
			// unimplemented
			fallthrough

		case canPayAllRetainingZero:
			fmt.Println("canPayAllRetainingZero")
			sendFromThisInput = target

		case canPaySomeRetainingED:
			fmt.Println("canPaySomeRetainED")
			nEds := (inputAmount / existentialDeposit) - 1
			if sendFromThisInput != 0 {
				nEds--
			}
			sendFromThisInput += nEds * existentialDeposit

		case canPaySomeRetainingZero:
			fmt.Println("canPaySomeRetainingZero")
			sendFromThisInput = inputAmount - fee

		default:
			fmt.Println("this input cannot pay")
		}

		target -= sendFromThisInput
		totalSendingAmount += sendFromThisInput
		inputs[i] = inputAmount - sendFromThisInput - fee
		outputs = append(outputs, sendFromThisInput)
		fmt.Printf("sendFromThisInput: %d\n", sendFromThisInput)
		fmt.Println("--------------------------------")
	}
	fmt.Printf("outputs: %v\n", outputs)
	fmt.Printf("totalSendingAmount: %v\n", totalSendingAmount)

	return
}
