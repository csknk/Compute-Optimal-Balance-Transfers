Compute Optimal Balance Transfers
=================================

Project to compute sending amounts when transactions are batched to achieve a target balance on a receiving address when there are restrictions on:

* Minimum account balance
* Mimimum send balance

Terms
-----
* `Existential Deposit`: if an account has a balance that is less than the network-determined `Existential Deposit`, it will be _reaped_ - the balance and nonce of the account will be set to zero
* `inputAmount`: the layerone balance of a Settlement input
* `amountRequired`: the overall amount to be sent
* `fee`: Fee to be paid 

Constraints
-----------
* Each component amount must be than or equal to the `Existential Deposit`
* Each component amount must leave the sending account with either zero balance or at least the `Existential Deposit`. Otherwise the sending account will be reaped.

