package main

import "testing"

func TestJustSort(t *testing.T) {
	toSort := "c\nb\na"
	sorted := "a\nb\nc"
	result := MySort(toSort, -1, false, false, false, false)

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestIgnorecasedSort(t *testing.T) {
	toSort := "c\nB\na"
	sorted := "a\nB\nc"
	result := MySort(toSort, -1, true, false, false, false)

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestReverseSort(t *testing.T) {
	toSort := "a\nb\nc"
	sorted := "c\nb\na"
	result := MySort(toSort, -1, false, false, true, false)

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestColumnSort(t *testing.T) {
	toSort := "1 c\n2 b\n3 a"
	sorted := "3 a\n2 b\n1 c"
	result := MySort(toSort, 1, false, false, false, false)

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestNumericSort(t *testing.T) {
	toSort := "120\n99\n65\n32\n1"
	sorted := "1\n32\n65\n99\n120"
	result := MySort(toSort, -1, false, false, false, true)

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestUniqueSort(t *testing.T) {
	toSort := "b\nb\nc\na\na"
	sorted := "a\nb\nc"
	result := MySort(toSort, -1, false, true, false, false)

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}
