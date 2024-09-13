package main
import (
	"fmt"
	"time"
)

func main() {
	// Time 

	presentTime:= time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate:= time.Date(2020, time.January, 1, 12, 30, 0, 0, time.Local)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}