package repo
import "fmt"

func main() {
	var a int = 1
	for a < 10 {
		fmt.Print("Value ", a, "\n")
		a++
		if a > 5 {
			break
		}
	}
}
