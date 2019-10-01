package main

import (
	"testing"
	"strings"
)

func makeOpts(col int, reg, first, reverse, numeric bool) sortOpts {
	return sortOpts {
		reg:		reg,
		first:		first,
		reverse:	reverse,
		numeric:	numeric,
		col:		col,
	}
}

func TestJustSort(t *testing.T) {
	toSort := "c\nb\na"
	sorted := "a\nb\nc"

	opts := makeOpts(-1, false, false, false, false)
	result := strings.Join(MySort(strings.Split(toSort, "\n"), opts), "\n")

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestIgnorecasedSort(t *testing.T) {
	toSort := "c\nB\na"
	sorted := "a\nB\nc"

	opts := makeOpts(-1, true, false, false, false)
	result := strings.Join(MySort(strings.Split(toSort, "\n"), opts), "\n")

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestReverseSort(t *testing.T) {
	toSort := "a\nb\nc"
	sorted := "c\nb\na"

	opts := makeOpts(-1, false, false, true, false)
	result := strings.Join(MySort(strings.Split(toSort, "\n"), opts), "\n")

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestColumnSort(t *testing.T) {
	toSort := "1 c\n2 b\n3 a"
	sorted := "3 a\n2 b\n1 c"

	opts := makeOpts(1, false, false, false, false)
	result := strings.Join(MySort(strings.Split(toSort, "\n"), opts), "\n")

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestNumericSort(t *testing.T) {
	toSort := "120\n99\n65\n32\n1"
	sorted := "1\n32\n65\n99\n120"

	opts := makeOpts(-1, false, false, false, true)
	result := strings.Join(MySort(strings.Split(toSort, "\n"), opts), "\n")

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}

func TestUniqueSort(t *testing.T) {
	toSort := "b\nb\nc\na\na"
	sorted := "a\nb\nc"

	opts := makeOpts(-1, false, true, false, false)
	result := strings.Join(MySort(strings.Split(toSort, "\n"), opts), "\n")

	if result != sorted {
		t.Errorf("wrong sort!!! got: \n%s \nwanted: \n%s", result, sorted)
	}
}
