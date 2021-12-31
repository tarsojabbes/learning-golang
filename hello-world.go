package main
import "fmt"
func main(){
    
    var name string
   
    # Reads a line and assigns the result to name
    fmt.Scanln(&name)
    
    var year_of_birth int

    # Reads another line and assigns the result to year_of_birth
    fmt.Scanln(&year_of_birth)

    # Calculates the age based on the read line
    age := 2021 - year_of_birth

    fmt.Println("Hello " + name)
    fmt.Println(age)
}
