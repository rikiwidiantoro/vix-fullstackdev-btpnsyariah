package main
import ("fmt")

func main() {
	fmt.Println("")
	fmt.Println("Hello World!")
	fmt.Println("")
}


// Goroutines
// package main
// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	fmt.Println("")
// 	go say("Alterra")
// 	say("Academy")
// 	fmt.Println("")
// }


// Pointers
// package main
// import "fmt"

// func main() {
// 	i, j := 42, 2701

// 	p := &i
// 	fmt.Println(*p)
// 	*p = 21
// 	fmt.Println(i)

// 	p = &j
// 	*p = *p / 37
// 	fmt.Println(j)
// }

// package main
// import "fmt"

// func main() {
// 	var fruits1 = []string{"apple", "banana", "manggo", "dunian", "pinaple"}

// 	fruits1 = append(fruits1[:3], "rambutan")
// 	fmt.Printf("%#v\n", fruits1)
// }

// func main() {
// 	var test interface{}

// 	test = 30

// 	test = test * 3

// 	fmt.Println(test)
// }