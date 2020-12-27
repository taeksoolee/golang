package main

import (
	"fmt"
	"io"
	"os"
)

// 파일을 선택 복사하는 코드
func main() {
	fi, err := os.Open("C:\\temp\\1.txt")

	if err != nil {
		panic(err)
	}

	defer fi.Close()

	fo, err := os.Create("C:\\temp\\2.txt")

	if err != nil {
		panic(err)
	}

	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		fmt.Print(buff)

		if cnt == 0 {
			break
		}

		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}

}
