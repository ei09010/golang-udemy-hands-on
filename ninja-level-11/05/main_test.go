package main

import "testing"

func TestMyProfileFinder(t *testing.T){
	expectedReturn := "a given person profile has the name tony and stark with age 28"

	retFromExecution := myProfileFinder(person{"tony","stark", 28 })

	if expectedReturn != retFromExecution{
		t.Errorf("Expected %v, got %v", expectedReturn, retFromExecution )
	}
}
