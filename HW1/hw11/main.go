package main

import ( 
	"strings"
	"sort"
	"os"
	"strconv"
	"flag"
	"io/ioutil"
	"fmt"
)

var regPtr = flag.Bool("f", false, "Ignore case")
var firstPtr = flag.Bool("u", false, "Output only the first of uniques")
var reversePtr = flag.Bool("r", false, "Reverse sort")
var numericPtr = flag.Bool("n", false, "Sort by numbers")
var colPtr = flag.Int("k", -1, "Sort by column")
var filenamePtr = flag.String("o", "", "Output to file")

type sortOpts struct {
	reg		bool
	first	bool
	reverse	bool
	numeric	bool
	col		int
}

func grabOpts() sortOpts {
	return sortOpts {
		reg:		*regPtr,
		first:		*firstPtr,
		reverse:	*reversePtr,
		numeric:	*numericPtr,
		col:		*colPtr,
	}
}

//MySort it's great m8
func MySort(content []string, options sortOpts) []string {
	wordSelector := getWordSelector(options.col)
	reverser := getReverser(options.reverse)

	sort.SliceStable(content, func (i, j int) bool {
		lhs, rhs := wordSelector(content[i]), wordSelector(content[j])

		if options.numeric {
			lhsInt, err := strconv.ParseInt(lhs, 10, 32)

			if err == nil {
				rhsInt, err := strconv.ParseInt(rhs, 10, 32)
				if err == nil {
					return reverser(lhsInt < rhsInt)
				}
			}
		}

		if options.reg {
			lhs, rhs = strings.ToLower(lhs), strings.ToLower(rhs)
		}

		return reverser(lhs < rhs)
	})

	if options.first {
		content = filterUnique(content, options.reg)
	}

	return content
}

func getWordSelector(k int) func(string) string {
	if k == -1 {
		return func(in string) string {
			return in
		}
	}

	return func(in string) string {
		retVal := ""

		cols := strings.Split(in, " ")

		if k <= len(cols) - 1 {
			retVal = cols[k]
		}

		return retVal
	}
}

func getReverser(reverse bool) func(bool) bool {
	return func(in bool) bool {
		if in {
			return !reverse
		}

		return reverse
	}
}

func filterUnique(in []string, ignorecase bool) []string {
	var output []string

	for idx, elem := range in {
		termCounter := 0
		for _, elem2 := range in[:idx] {
			lhs, rhs := elem, elem2
			
			if (ignorecase) {
				lhs, rhs = strings.ToLower(elem), strings.ToLower(elem2)
			}
			
			if lhs == rhs {
				termCounter++
			}
		}

		if termCounter == 0 {
			output = append(output, elem)
		}
	}

	return output
}

func main() {
	flag.Parse()
	options := grabOpts()
	
	var inputFile = flag.Arg(0)
	if inputFile == "" {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	
	content, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}

	inputSlices := strings.Split(string(content), "\n")
	sortedSlices := MySort(inputSlices, options)
	output := strings.Join(sortedSlices, "\n")

	if *filenamePtr != "" {
		err = ioutil.WriteFile(*filenamePtr, []byte(output), 0677)
		if err != nil {
			fmt.Println("File writing error", err)
		}
	} else {
		fmt.Println(output)
	}
}
