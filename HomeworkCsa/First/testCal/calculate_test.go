package testCal

import (
	"testing"
)

func TestCalculate(t *testing.T)  {
	res := Calculate(9,2,"pow")
	if res != 81{
		t.Fatalf("Wrong result ! Expected value = %.2f,True value = %.2f",81.0,res)
	}
	t.Logf("Run successfully, the result is correct")
}
