package main

import (
	"fmt"
	"sort"
)

// Fund the amount of money deposited
type Fund struct {
	amount int
}

// DepositPlan the deposit plan
type DepositPlan struct {
	t int         // Type of deposit (1 = One Time / 2 = Monthly)
	p []Portfolio // Portfolios associated with the deposit plan
}

// Portfolio the portfolio
type Portfolio struct {
	amount int // Amount of money destined to that portfolio
}

// allocate method that allocates the funds to the portfolios
func (fund Fund) allocate(ratios ...int) []Portfolio {
	// Sum of ratios
	var sum int
	for _, ratio := range ratios {
		sum += ratio
	}

	var total int
	portfolios := make([]Portfolio, 0, len(ratios))

	for _, ratio := range ratios {
		portfolio := Portfolio{
			amount: fund.amount * ratio / sum,
		}

		portfolios = append(portfolios, portfolio)
		total += portfolio.amount
	}

	// Calculate remaining value
	lo := fund.amount - total
	sub := 1
	if lo < 0 {
		sub = -sub
	}

	// Remaining value to first portfolio.
	for i := 0; lo != 0; i++ {
		portfolios[i].amount = portfolios[i].amount + sub
		lo -= sub
	}

	return portfolios
}

// AllocateFunds allocate the funds to the different portfolios of the customer
func AllocateFunds(depositPlans []DepositPlan, funds []Fund) []Portfolio {

	// Sort the deposit plans (1 - Time plans first)
	sort.Slice(depositPlans, func(i, j int) bool {
		return depositPlans[i].t < depositPlans[j].t
	})

	var portfolios []Portfolio

	for _, dp := range depositPlans {
		portfolios = make([]Portfolio, len(dp.p))
		break
	}

	for _, fund := range funds {
		done := true                 // Variable to only allocate the funds once
		allocateMoney := fund.amount // Money to be allocate

		for _, dp := range depositPlans {
			if allocateMoney == 0 { // Check if we already allocated all the funds
				continue
			}

			/* Get the ratios to allocate the money */
			var ratios []int // Ratios for that deposit plan
			// Sum of ratios
			var sum int
			for _, ratio := range dp.p {
				ratios = append(ratios, ratio.amount)
				sum += ratio.amount
			}

			// Check if the money funded is more than the deposit plan money
			if fund.amount > sum && done && len(depositPlans) > 1 {
				// Temporal funds to be allocated to the 1 - Time Deposit plan 1st
				fundTemp := Fund{amount: sum}
				portfolioTemp1 := fundTemp.allocate(ratios...)
				// Add the allocated money to the portfolios
				for i, portfolio := range portfolioTemp1 {
					portfolios[i].amount += portfolio.amount
				}
				allocateMoney -= fundTemp.amount
				fund.amount -= sum
				done = false

				continue
			}

			/* Allocate the money  */
			portfolioTemp2 := fund.allocate(ratios...)
			for i, portfolio := range portfolioTemp2 {
				portfolios[i].amount += portfolio.amount
			}
			allocateMoney -= fund.amount
		}
	}

	return portfolios
}

func main() {
	/* Input */
	// 1 - Time deposit plan
	dp1 := DepositPlan{
		1, []Portfolio{
			//{33}, {33}, {33},
			{10}, {10},
		},
	}
	// Monthly deposit plan
	dp2 := DepositPlan{
		2, []Portfolio{
			{20}, {30},
			//{33}, {33}, {33},
		},
	}
	depositPlans := []DepositPlan{dp1, dp2}

	funds := []Fund{
		{100}, {50}, {70},
	}

	/* Output */
	// Allocate the funds on the different portfolios
	portfolios := AllocateFunds(depositPlans, funds)
	fmt.Println(portfolios)
}
