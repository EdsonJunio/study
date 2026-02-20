package easy_exercises

import "fmt"

func main() {
	var value int
	fmt.Scan(&value)
	var num, numOdd, numPeers int
	numOdd = num
	numPeers = num

	for i := 0; i < value; i++ {
		fmt.Scan(&num)
		if num%2 == 0 {
			numOdd++
		} else {
			numPeers++
		}
	}

	fmt.Println("Pares:", numOdd)
	fmt.Println("Impares:", numPeers)

}
