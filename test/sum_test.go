package test

import (
    "testing"
    "github.com/Wilovy09/Go_desde_0/utils"
)
func TestSum(t *testing.T) {
	result := utils.Sum(5, 5)
	expected := 10
	if result != expected {
		t.Errorf("Sum(5,5) expected %d but got %d", expected, result)
	}
}
