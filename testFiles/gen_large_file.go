package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const mx int = 1e6

func main() {
	file, err := os.Create("large_file.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < mx; i++ {
		n := rand.Intn(mx)
		writer.WriteString(strconv.Itoa(n) + "\n")
	}
	fmt.Println("large_file.txt created with ", mx, " random numbers")
}
