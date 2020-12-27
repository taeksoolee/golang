package util

import (
	"fmt"
)

// 같은 패키지 내의 함수 사용 가능
func Print1() {
	fmt.Println("hello util1~")
	Print2()
}
