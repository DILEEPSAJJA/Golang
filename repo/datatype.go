package repo

import "fmt"

func main() {
	var i int;
	var f float64;
	var b bool;
	var str string;

	fmt.Printf("%T %T %T %T",i,f,b,str);
	fmt.Printf("%v %v %v %q",i,f,b,str);

}
