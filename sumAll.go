package sumAll
//Package sumAll calculates the sum in array


var sum int = 0 //sum is a priavte variable to store sum


func CalculateSum(array []int) int {
//calculateSum is a global function
    for k := 0; k < 6 ; k++{
            sum+=array[k];
    }
return sum
}
