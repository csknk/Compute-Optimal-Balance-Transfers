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
	inputs := []int64{2020, 4000, 8000}
	outputs, err := UnderlyingForSigning(target, inputs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", inputs)
	fmt.Printf("%v\n", outputs)

}

func UnderlyingForSigning(target int64, inputs []int64) (outputs []int64, err error) {
	fee := int64(10)
	outputs, sendingAmount, _ := buildInputsOutputs(target, inputs, fee)
	fmt.Printf("sendingAmount: %d\n", sendingAmount)
	fmt.Printf("outputs: %v\n", outputs)

	return
}

func buildInputsOutputs(target int64, inputs []int64, fee int64) (outputs []int64, totalSendingAmount int64, err error) {
forLoop:
	for i, inputAmount := range inputs {
		if target == 0 {
			break
		}

		remainder := target % existentialDeposit
		maxSpendRetainingSpendableBalance := inputAmount - existentialDeposit - fee
		canPayAllRetainingED := maxSpendRetainingSpendableBalance >= target
		canPayAllRetainingDust := inputAmount-fee > target && !canPayAllRetainingED
		canPayAllRetainingZero := inputAmount-fee == target
		canPaySomeRetainingED := maxSpendRetainingSpendableBalance < target && maxSpendRetainingSpendableBalance >= existentialDeposit // ACTUAL ED
		//		canPaySomeRetainingZero := inputAmount-fee < target && inputAmount-fee >= existentialDeposit && remainder == 0
		canPaySomeRetainingZero := inputAmount-fee < target && inputAmount-fee >= existentialDeposit

		fmt.Printf("input %d, inputAmount: %d target amount: %d\n", i, inputAmount, target)
		//		sendFromThisInput := int64(remainder)
		sendFromThisInput := int64(0)

		switch {
		case canPayAllRetainingED:
			fmt.Println("canPayAllRetainingED")
			sendFromThisInput = target

		case canPayAllRetainingDust:
			// unimplemented, dust will cause unacceptable reaping
			fallthrough

		case canPayAllRetainingZero:
			fmt.Println("canPayAllRetainingZero")
			sendFromThisInput = target

		case canPaySomeRetainingZero:
			fmt.Println("canPaySomeRetainingZero")
			sendFromThisInput = inputAmount - fee
			//			if remainder == 0 {
			//				// must send multiple of ED
			//				if inputAmount-fee < 2*existentialDeposit {
			//					continue forLoop
			//				}
			//			}
			// max spend
			//			sendFromThisInput = (((inputAmount - fee) / existentialDeposit) - 1) * existentialDeposit

		case canPaySomeRetainingED:
			fmt.Println("canPaySomeRetainingED")
			nEds := (inputAmount / existentialDeposit) - 1
			if remainder != 0 {
				if inputAmount-fee < 2*existentialDeposit {
					continue forLoop
				}
				nEds--
			}
			sendFromThisInput += nEds * existentialDeposit

		default:
			fmt.Println("this input cannot pay")
		}

		target -= sendFromThisInput
		totalSendingAmount += sendFromThisInput
		inputs[i] = inputAmount - sendFromThisInput - fee
		outputs = append(outputs, sendFromThisInput)
		fmt.Printf("sendFromThisInput: %d\n", sendFromThisInput)
		fmt.Printf("remainingThisInput: %d\n", inputAmount-sendFromThisInput-fee)
		fmt.Println("--------------------------------")
	}
	fmt.Printf("outputs: %v\n", outputs)
	fmt.Printf("totalSendingAmount: %v\n", totalSendingAmount)

	return
}
