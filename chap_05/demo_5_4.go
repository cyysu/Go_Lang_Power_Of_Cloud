
package chap_05
//package main

import "fmt"

type myError struct {}

func (_ *myError) Error() string {
	return "To err is human"
}

func toErr(ok bool) error {
	var e *myError = nil
	if ok { e = &myError{} }
	return e
}

func noErr(ok bool) error{
	if !ok { return &myError{} }
	return nil
}

func Demo_5_4(){
	fmt.Println(toErr(true))	
	fmt.Println(toErr(false))
	fmt.Println(noErr(true))
	fmt.Println(noErr(false))
}

/*
func main(){
	Demo_5_4()
}
*/
