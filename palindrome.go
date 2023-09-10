package main
import (
	"fmt"
)
func isPalindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	return original == reversed
}
func main() {
	var lower, upper int
	fmt.Print("Enter the lower bound of the range : ")
	fmt.Scan(&lower)
	fmt.Print("Enter the upper bound of the range : ")
	fmt.Scan(&upper)
	if lower < 100 || upper > 999 || lower > upper {
		fmt.Println("Invalid range. Please enter valid 3-digit bounds.")
		return
	}
	largestPalindrome := 0
	for i := lower; i <= upper; i++ {
		for j := lower; j <= upper; j++ {
			product := i * j
			if isPalindrome(product) && product > largestPalindrome {
				largestPalindrome = product
			}
		}
	}
	fmt.Println("The largest palindrome product of two", lower, "-digit numbers is:", largestPalindrome)
}

