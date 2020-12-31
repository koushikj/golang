package main

import (
	"errors"
	"fmt"
	"log"
)

type custError struct {
	id      int
	message string
}

func (c custError) Error() string {
	return fmt.Sprintf("Error occured : %v , %v ", c.id, c.message)
}
func main() {
	customError()
	createNewError()
	createNewFormatedError()
}

func createNewFormatedError() {
	str, err := foo4()
	log.Println(str, err)
}

func createNewError() {
	abc, err := foo4()
	log.Printf("received string %v and error %v \n", abc, err)
}

func foo4() (string, error) {
	return "hi", errors.New(" error from foo4")
}

func foo5() (string, error) {
	return "hi", fmt.Errorf("error from foo5 and the reason is %v \n", 10201)
}

func customError() {
	err := throwCustomError()
	if err != nil {
		log.Println("error id:", err.(custError).id) // asserting the error type is custError type
		log.Println("error message :", err.(custError).message)
		//log.Fatal("fatal error :",err) // this will do os.exit()
	}
}

func throwCustomError() error {
	c := custError{10, "invalid type"}
	return c
}
