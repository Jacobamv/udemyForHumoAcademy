package main

import (
	"fmt"
	"udemy/pkg/mock"
	"udemy/pkg/services"
)

func main() {
	mock.FillAll()

	avg := services.GetAverageUserAgeByCourseIndex(0)

	fmt.Println(avg)

}
