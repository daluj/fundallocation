# GO Allocation
![Go](https://github.com/daluj/fundallocation/workflows/Go/badge.svg)

## Introduction
### Problem:
- Customers can decide to open multiple portfolios.
- Customers deposit funds into portfolios by bank transfer and have to include personal reference code when submitting transfer to bank.
- Reference code is unique per customer and the same for all of customer's portfolios.
- Customers can set up "deposit plan" on either a 1-Time and/or monthly basis to specify how funds are allocated amongs portfolios.

### Solution:
Build a method that receives a list of both 1-time and monthly deposit plans as well as a list of funds deposits for a particular customer and allocates the funds amongst the customer's portfolios.

## Installation
Project requirements
--------------------
* GO: ^1.13

How to run it
-------------
1. Clone the repo:
```
$ git clone https://github.com/daluj/fundallocation.git
```

2. Execute the tests:
```
$ go test
```
