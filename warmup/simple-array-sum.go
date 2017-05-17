package main
import (
    "fmt"   
)

func main() {   
    var numElements int
    fmt.Scanf("%d", &numElements)
    
    var array = make([]int, numElements)
    var suma = 0
    
    for i := 0; i < len(array); i++ {
        fmt.Scanf("%d", &array[i])
        suma += array[i]
    }
    
    fmt.Println(suma)
}