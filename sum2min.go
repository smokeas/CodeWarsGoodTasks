/*Create a function that returns the sum of the two lowest positive numbers given an array of minimum 4 positive integers. No floats or non-positive integers will be passed.

For example, when an array is passed like [19, 5, 42, 2, 77], the output should be 7.

[10, 343445353, 3453445, 3453545353453] should return 3453455.*/
func sumTwoSmallestNumbers(numbers []int) int {
    min1 := numbers[0]
    min2 := numbers[1]
    
    if min1 > min2 {
        min1, min2 = min2, min1
    }
    
    for i := 2; i < len(numbers); i++ {
        num := numbers[i]//тукущий элемент по индекс 
        
        if num < min1 {//Проверяем, меньше ли текущее число самого маленького найденного числа
            min2 = min1
            min1 =  //Устанавливаем новое минимальное значение
        } else if num < min2 {
            min2 = num
        }
    }
    
    sum := min1 + min2
    return sum
}
