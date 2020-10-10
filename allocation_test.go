package main

import (
	"reflect"
	"testing"
)

func TestAllocateFunds(t *testing.T) {
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
	actual := AllocateFunds(depositPlans, funds)
	expected := []Portfolio{
		{94}, {126},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Test failed, expected %v, got %v", expected, actual)
	} else {
		t.Logf("Test Success, expected %v and got %v", expected, actual)
	}
}
