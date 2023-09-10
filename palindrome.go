package main
import "fmt"
func isPalindrome(num int) bool {
	rev := 0
	temp := num
	for temp != 0 {
		rev = rev*10 + temp%10
		temp /= 10
	}
	return rev == num
}
func largestPalindrome(lower, upper int) (int, int, int) {
	large := 0
	num1 := 0
	num2 := 0
	product := 0
	for i := upper; i >= lower; i-- {
		for j := upper; j >= lower; j-- {
			product = i * j
			if product >= large && isPalindrome(product) {
				large = product
				num1 = i
				num2 = j
			}
		}
	}
	return large, num1, num2
}
func main() {
	var lower, upper int
	fmt.Print("Enter the lower bound of the range (e.g., 100): ")
	fmt.Scan(&lower)
	fmt.Print("Enter the upper bound of the range (e.g., 999): ")
	fmt.Scan(&upper)
	if lower < 100 || upper > 999 || lower > upper {
		fmt.Println("Invalid range. Please enter valid 3-digit bounds.")
		return
	}
	large, num1, num2 := largestPalindrome(lower, upper)
	fmt.Println("The largest Palindrome product: ", large)
	fmt.Println("The Numbers: ", num1, num2)
}
