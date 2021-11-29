package main

import (
	"errors"
	"fmt"
)

type myCustomError struct {
	status int
	msg    string
}

type myCustomError2 struct {
	status int
	msg    string
}

func devolverError() error {
	return errors.New("Esto es un error")
}

func (e myCustomError2) Error() string {
	if e.status == 200 {
		return "Error creado" + e.msg
	} else {

		return "Error creado 2" + e.msg
	}
}

func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func myCustomErrorTest(status int) (int, error) {

	if status >= 300 {
		return status, &myCustomError{
			status: status,
			msg:    "algo salió mal"}
	} else if status == 200 {
		return status, &myCustomError{
			status: status,
			msg:    "algo no salió tan mal"}
	}

	return 100, nil

}

// func (e myCustomError2) Saludar() string {
// 	return "Error creado"
// }

// func devolverError2() error {
// 	var errorcito myCustomError
// 	errorcito.msg = "Esto es un error creado"
// 	return errorcito
// }

// func devolverError3() error {
// 	var errorcito myCustomError2
// 	errorcito.msg = "Esto es un error creado"
// 	return errorcito
// }

func main() {
	_, err := myCustomErrorTest(350)
	fmt.Println(err)

	err2 := fmt.Errorf("Tiene un error interno %w", err)

	fmt.Println(err2)

	fmt.Println(errors.Unwrap(err2))
	fmt.Println(errors.Unwrap(err))

	// err = devolverError2()
	// fmt.Println(err)

	// err = devolverError3()
	// fmt.Println(err)
}
