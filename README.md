Compute Optimal Balance Transfers
=================================

Project to compute sending amounts when transactions are batched to achieve a target balance on a receiving address when there are restrictions on:

* Minimum account balance
* Mimimum send balance

Terms
-----
* `ExistentialDeposit`: if an account has a balance that is less than the network-determined `ExistentialDeposit`, it will be _reaped_ - the balance and nonce of the account will be set to zero
* `inputAmount`: the layerone balance of a Settlement input
* `amountRequired`: the overall amount to be sent
* `fee`: Fee to be paid 

Constraints
-----------
* Each component amount must be than or equal to the `ExistentialDeposit`
* Each component amount must leave the sending account with either zero balance or at least the `ExistentialDeposit`. Otherwise the sending account will be reaped.

Algorithm
---------

Assume that:
* Inputs have been sorted in descending order (largest first)

```
BuildTransactions(A, Target)
INPUT: An Array, Inputs, of transaction inputs ordered by ascending magnitude
INPUT: Total balance to send, Target
OUTPUT: An Array of sending amounts, Outputs, to be used in component transactions; totalling the Target
IF Target < ExistentialDeposit
  return Target too low error
Outputs = []int64
FOR: i = 0; i < len(Inputs); i++
	IF: Target == 0
		break
	sendingAmount = 0

	IF: Input can pay all AND retain ExistentialDeposit
		sendingAmount = Target
		Target = 0
		Outputs = append(Outputs, sendingAmount)
		break

	IF: Input can pay all and be reduced to zero
		sendingAmount = Target
		Target = 0
		Outputs = append(Outputs, sendingAmount)
		break

	IF: Input can pay some and is consumed completely (retains exactly Zero balance)
		sendingAmount = Inputs[i] - fee
		Target -= sendingAmount

	IF: Input can pay some AND retain ExistentialDeposit
		sendingAmount = maximum n*ExistentialDeposit
		Outputs = append(Outputs, sendingAmount)
		Target -= sendingAmount
ENDFOR
RETURN Outputs
```
