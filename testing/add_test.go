package math //unit tests are written with same package

import "testing"  //in-build package for testing

//Naming:
//For function F -> TestF
//For struct R -> TestR
//For method M() -> TestT_M

func TestAdd(t *testing.T) { //testing function name should be func Testfuncname(t *testing.T){//..}
    operation := Add(2, 3)
    result := 5
    t.Logf("The addition is %d",operation) //just like printf()

    if operation != result {
        t.Errorf("Add() returned %d, expected %d",operation, result) //t.Error(log message with exit status)
    }

    operation = Sub(9, 3)
    result = 6
    t.Logf("The subtraction is %d",operation)

    if operation != result {
        t.Errorf("Sub() returned %d, expected %d",operation, result)
    }

}

//similarly we have t.Fail() -> wills how the testcase has failed; 
//T.FailNow() -> t.Fail() + safely exit; 
//t.Fatal() -> t.Log() + t.failNow()
