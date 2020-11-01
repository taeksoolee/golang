package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력해주세요")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다. 연산자를 입력해주세요.\n", n1, n2)

	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	switch op {
	case "+":
		fmt.Println("n1 + n2 = ", n1+n2)
	case "-":
		fmt.Println("n1 - n2 = ", n1-n2)
	case "*":
		fmt.Println("n1 * n2 = ", n1*n2)
	case "/":
		fmt.Println("n1 / n2 = ", n1/n2)
	}

}
