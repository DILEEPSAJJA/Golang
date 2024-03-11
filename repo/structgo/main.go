package main
import ("fmt")

func main() {
	fmt.Println("Struct Introduction")

	dileep := User{"Dileep","dileep@gmail.com",true,20}
	fmt.Println(dileep)

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}
