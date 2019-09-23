package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Token struct {
	IsOp	bool
	Op		byte
	Val		float64
}

func isOneOf(in rune, out string) bool {
	for _, char := range out {
		if in == char {
			return true
		}
	}

	return false
}

func parseTokens(str string) []Token {
	retSl := make([]Token, 0)

	for i := 0; i < len(str); i++ {

		if isOneOf(rune(str[i]), " \n") {
			continue
		}

		if isOneOf(rune(str[i]), "+-*/()") {
			retSl = append(retSl, Token{
				IsOp: true,
				Op: str[i],
				Val: 0,
			})
		} else {
			numberEnd := 0

			for idx, char := range(str[i:]) {
				if !isOneOf(char, "1234567890") {
					numberEnd += idx
					break
				}
			}

			number, _ := strconv.Atoi(str[i:i + numberEnd])

			retSl = append(retSl, Token{
				IsOp: false,
				Op: ' ',
				Val: float64(number),
			})

			i += numberEnd - 1
		}
	}

	return retSl
}

func pop(stack []Token) (Token, []Token) {
	elem := stack[len(stack) - 1]
	return elem, stack[:len(stack) - 1]
}

func push(stack []Token, elem Token) []Token {
	return append(stack, elem)
}

// shunting yard algo
func toRpn(tokens []Token) []Token {
	var precedence map[byte]int = map[byte]int{
		'+':	0,
		'-':	0,
		'*':	1,
		'/':	1,
	}

	outStack := make([]Token, 0)
	opStack := make([]Token, 0)

	for _, t := range tokens {
		if !t.IsOp {
			outStack = push(outStack, t)
		} else {
			if t.Op == '(' {
				opStack = push(opStack, t)
			} else if t.Op == ')' {
				for {
					var top Token
					top, _ = pop(opStack)

					if top.Op != '(' {
						outStack = push(outStack, top)
					} else {
						_, opStack = pop(opStack)
						break
					}

					_, opStack = pop(opStack)
				}
			} else {
				for len(opStack) != 0 {
					top, _ := pop(opStack)

					if top.Op == '(' {
						break
					} else if isOneOf(rune(top.Op), "+-/*") {
						if precedence[top.Op] < precedence[t.Op] {
							break
						}
					}
					outStack = push(outStack, top)
					_, opStack = pop(opStack)
				}
				opStack = append(opStack, t)
			}
		}
	}

	for i := len(opStack) - 1; i != -1; i-- {
		outStack = push(outStack, opStack[i])
	}

	return outStack
}

func evalRpn(tokens []Token) float64 {
	stack := make([]float64, 0)

	for _, t := range tokens {
		if t.IsOp {
			lhs := stack[len(stack) - 2]
			rhs := stack[len(stack) - 1]

			stack = stack[:len(stack) - 2]

			switch t.Op {
				case '+':
					stack = append(stack, lhs + rhs)
				case '-':
					stack = append(stack, lhs - rhs)
				case '*':
					stack = append(stack, lhs * rhs)
				case '/':
					stack = append(stack, lhs / rhs)
			}

		} else {
			stack = append(stack, t.Val)
		}
	}

	return stack[0]
}

func evalInf(expr string) float64 {
	return evalRpn(toRpn(parseTokens(expr)))
}

func main() {
	stdin := bufio.NewReader(os.Stdin)

	expr, _ := stdin.ReadString('\n')

	fmt.Println(evalInf(expr))
}
