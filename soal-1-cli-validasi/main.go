package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func readLine(msg string) (string, error) {
	fmt.Print(msg)
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", nil
	}

	str = strings.TrimSpace(str)

	if len(str) == 0 {
		return "", fmt.Errorf("input: Jangan kosong dong")
	}

	return str, nil
}

func main() {
	for {
		fmt.Print("\n\n")

		name, err := readLine("Nama: ")
		if err != nil {
			fmt.Println(">> Error:", err)
			continue
		}

		age_str, err := readLine("Umur: ")
		if err != nil {
			fmt.Println(">> Error:", err)
			continue
		}

		age, err := strconv.Atoi(age_str)
		if err != nil {
			fmt.Println(">> Error: masukan bukan angka")
			continue
		}

		if age < 18 {
			fmt.Println(">> Error: umur tidak valid (minimal 18 tahun)")
			continue
		}

		fmt.Printf(">> Selamat datang, %s!", name)
		break
	}
}
