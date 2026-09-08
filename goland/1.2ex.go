package main

//1.1
import (
	"fmt"
	"os"
)
func main() {
	for _, arg := range os.Args[1:]{
		fmt.Println(arg)
	}
}

//1.2
import (
	"fmt"
	"os"
)
func main() {
	for i, arg := range os.Args[1:]{
		fmt.Println(i+1, arg)
	}
}

//1.3 в консоль time go run goland/1.2ex.go hello world
import (
	"fmt"
	"os"
	"strings"
)
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	}
