package main

import "testing"

func makeToken(isOp bool, op byte, val float64) Token {
	return Token{
		IsOp: isOp,
		Op: op,
		Val: val,
	}
}

func TestTokenizing(t *testing.T) {
	testStr := "(1+2+3)"

	// [ '(', '1', '+', '2', '+', '3', ')' ]
	expectedRes := []Token{ 
		makeToken(true, '(', 0),
		makeToken(false, ' ', 1),
		makeToken(true, '+', 0),
		makeToken(false, ' ', 2),
		makeToken(true, '+', 0),
		makeToken(false, ' ', 3),
		makeToken(true, ')', 0),
	}

	for idx, elem := range parseTokens(testStr) {
		if expectedRes[idx] != elem {
			t.Errorf(
				"Wrong tokenizing at index %d: wanted %v, got %v",
				idx,
				expectedRes[idx],
				elem,
			)
		}
	}
}

func TestShuntingYard(t *testing.T) {
	// [ '(', '1', '+', '2', ')', '*', '3' ]
	inTokens := []Token{
		makeToken(true, '(', 0),
		makeToken(false, ' ', 1),
		makeToken(true, '+', 0),
		makeToken(false, ' ', 2),
		makeToken(true, ')', 0),
		makeToken(true, '*', 0),
		makeToken(false, ' ', 3),
	}

	// [ '1', '2', '+', '3', '*' ]
	expectedRes := []Token{
		makeToken(false, ' ', 1),
		makeToken(false, ' ', 2),
		makeToken(true, '+', 0),
		makeToken(false, ' ', 3),
		makeToken(true, '*', 0),
	}

	for idx, elem := range toRpn(inTokens) {
		if expectedRes[idx] != elem {
			t.Errorf(
				"Wrong tokenizing at index %d: wanted %v, got %v",
				idx,
				expectedRes[idx],
				elem,
			)
		}
	}
}

func TestRpnEval(t *testing.T) {
	cases := []struct {
		expr			string
		in				[]Token
		expectedOut		float64
	}{
		{
			expr: "1 2 + 3 *",
			in: []Token{
				makeToken(false, ' ', 1),
				makeToken(false, ' ', 2),
				makeToken(true, '+', 0),
				makeToken(false, ' ', 3),
				makeToken(true, '*', 0),
			},
			expectedOut: 9,
		},
		{
			expr: "1 2 * 3 +",
			in: []Token{
				makeToken(false, ' ', 1),
				makeToken(false, ' ', 2),
				makeToken(true, '*', 0),
				makeToken(false, ' ', 3),
				makeToken(true, '+', 0),
			},
			expectedOut: 5,
		},
		{
			expr: "1 2 -",
			in: []Token{
				makeToken(false, ' ', 1),
				makeToken(false, ' ', 2),
				makeToken(true, '-', 0),
			},
			expectedOut: -1,
		},
		{
			expr: "4 2 /",
			in: []Token{
				makeToken(false, ' ', 4),
				makeToken(false, ' ', 2),
				makeToken(true, '/', 0),
			},
			expectedOut: 2,
		},
	}

	for idx, testCase := range cases {
		res := evalRpn(testCase.in)
		if res != testCase.expectedOut {
			t.Errorf(
				"Fail: test case %d, wrong evaluation of %s: wanted %f, got %f",
				idx,
				testCase.expr,
				testCase.expectedOut,
				res,
			)
		}
	}
}

func TestUnaryEval(t *testing.T) {
	expr := "-4\n"
	expectedRes := -4.0

	res := evalInf(expr)

	if res != expectedRes {
		t.Errorf(
			"Wrong evaluation of unary minus (%s), wanted %f, got %f",
			expr,
			expectedRes,
			res,
		)
	}
}

func TestInfEval(t *testing.T) {
	exprs := []string{
		"((1+5)*6)/4\n",
		"3 + 1 / (-4 * (-1/2)) - 4\n",
	}
	// expr := "((1+5)*6)/4\n"
	expectedResults := []float64{
		9.0,
		-0.5,
	}

	for idx, expr := range exprs {
		res := evalInf(expr)
		
		if res != expectedResults[idx] {
			t.Errorf(
				"Wrong evaluation of %s, wanted %f, got %f",
				expr,
				expectedResults[idx],
				res,
			)
		}
	}
}