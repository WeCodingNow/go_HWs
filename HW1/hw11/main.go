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

func wordSelector(words string, k int) string {
	retVal := ""

	if k == -1 {
		retVal = words
	} else {
		cols := strings.Split(words, " ")

		if k <= len(cols) - 1 {
			retVal = cols[k]
		}
	}

	return retVal
}

//MySort it's great m8
func MySort(content string, col int, ignorecase, unique, reverse, numeric bool) string {
	strContent := strings.Split(string(content), "\n")

	sort.SliceStable(strContent, func (i, j int) bool {
		lhs, rhs := strContent[i], strContent[j]

		if col != -1 {
			lhs, rhs = wordSelector(lhs, col), wordSelector(rhs, col)
		}

		if numeric {
			lhsInt, err := strconv.ParseInt(lhs, 10, 32)

			if err == nil {
				rhsInt, err := strconv.ParseInt(rhs, 10, 32)
				if err == nil {
					if lhsInt < rhsInt {
						return !reverse
					}
					return reverse
				}
			}
		}

		if ignorecase {
			lhs, rhs = strings.ToLower(lhs), strings.ToLower(rhs)
		}

		if lhs < rhs {
			return !reverse
		}

		return reverse
	})


	var output strings.Builder

	for idx, elem := range strContent {
		if (unique) {
			termCounter := 0
			for _, elem2 := range strContent[:idx] {
				lhs, rhs := elem, elem2
				
				if (ignorecase) {
					lhs, rhs = strings.ToLower(elem), strings.ToLower(elem2)
				}
				
				if lhs == rhs {
					termCounter++
				}
			}

			if termCounter == 0 {
				output.WriteString(elem + "\n")
			}

		} else {
			output.WriteString(elem + "\n")
		}
	}

	outStr := output.String()
	return outStr[:len(outStr) - 1]
}

func main() {
	regPtr := flag.Bool("f", false, "Ignore case")
	firstPtr := flag.Bool("u", false, "Output only the first of uniques")
	reversePtr := flag.Bool("r", false, "Reverse sort")
	numericPtr := flag.Bool("n", false, "Sort by numbers")
	colPtr := flag.Int("k", -1, "Sort by column")
	filenamePtr := flag.String("o", "", "Output to file")

	flag.Parse()
	
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

	strContent := MySort(string(content), *colPtr, *regPtr, *firstPtr, *reversePtr, *numericPtr)

	if *filenamePtr != "" {
		err = ioutil.WriteFile(*filenamePtr, []byte(strContent), 0677)
		if err != nil {
			fmt.Println("File writing error", err)
		}
	} else {
		fmt.Println(strContent)
	}
}
