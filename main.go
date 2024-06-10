package main

import (
	"fmt"
)

func interpretBrainfuck(code string) {
	memory := make([]byte, 30000)
	pointer := 0

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			pointer++
		case '<':
			pointer--
		case '+':
			memory[pointer]++
		case '-':
			memory[pointer]--
		case '.':
			fmt.Print(string(memory[pointer]))
		case ',':
			var input byte
			fmt.Scanf("%c", &input)
			memory[pointer] = input
		case '[':
			if memory[pointer] == 0 {
				loopCount := 1
				for loopCount > 0 {
					i++
					if code[i] == '[' {
						loopCount++
					} else if code[i] == ']' {
						loopCount--
					}
				}
			}
		case ']':
			if memory[pointer] != 0 {
				loopCount := 1
				for loopCount > 0 {
					i--
					if code[i] == ']' {
						loopCount++
					} else if code[i] == '[' {
						loopCount--
					}
				}
			}
		}
	}
}

func main() {
	code := "+[----->+++<]>.+++++++++++++.-----------."
	interpretBrainfuck(code)
}