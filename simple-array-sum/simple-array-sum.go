package main
import "fmt"

func getSum (numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
} 

func main() {
    var size int
    fmt.Scanf("%v\n", &size)
    numbersToSum := make([]int, size)
    for i := 0; i<size; i++ {
        fmt.Scanf("%v", &numbersToSum[i])
    }
    sum := getSum(numbersToSum)
    fmt.Println(sum)
}