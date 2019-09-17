package main

import ( 
	"fmt";
	"flag";
	"os";
	"bufio";
	"strings";
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func comparator(lhs, rhs string) bool {
		cutoff := min(len(lhs), len(rhs))

		if cutoff == 0 {
			if len(lhs) < len(rhs) {
				return true
			}
			return false
		}

		if strings.Compare(lhs[:cutoff - 1], rhs[:cutoff - 1]) < 0 {
			return true
		}
		return false
}

func wordSelector(words []string, k int) string {
	if k == -1 {
		return strings.Join(words, " ")
	}

	if k <= len(words) - 1 {
		return words[k]
	}

	return ""
}

func sortLines(stringArray [][]string, k int, reverse bool) [][]string {
	//если задано k, то сортируем по k-ому столбцу
	//иначе всю строку (поэтому нужно сделать полную строку из слайса строк)
	for idx := range stringArray[:len(stringArray) - 1] {
		min := idx

		for compIdx, compWords := range stringArray[idx + 1:] {
			comparing := wordSelector(compWords, k)
			minWord := wordSelector(stringArray[min], k)

			if comparator(comparing, minWord) {
				min = idx + 1 + compIdx
			}
		}

		if min != idx {
			temp := stringArray[idx]
			stringArray[idx] = stringArray[min]
			stringArray[min] = temp
		}
	}

	if reverse {
		for idx := range(stringArray[:len(stringArray) / 2]) {
			opp := len(stringArray) - idx - 1
			stringArray[idx], stringArray[opp] = stringArray[opp], stringArray[idx]

		}
	}
	return stringArray
}


func main() {
	regPtr := flag.Bool("f", false, "Ignore case")
	firstPtr := flag.Bool("u", false, "Output only the first of uniques")
	reversePtr := flag.Bool("r", false, "Reverse sort")
	numericPtr := flag.Bool("n", false, "Sort by numbers")
	rowPtr := flag.Int("k", -1, "Sort by column")
	filenamePtr := flag.String("o", "", "Output to file")

	flag.Parse()
	
	var inputFile = flag.Arg(0)
	if inputFile == "" {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	file, err := os.Open(inputFile)
	
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	
	defer file.Close()


	var lineBuf [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokensSlice := strings.Split(scanner.Text(), " ")
		lineBuf = append(lineBuf, tokensSlice)
	}

	lineBuf = sortLines(lineBuf, *rowPtr, *reversePtr)

	fmt.Println(*regPtr, *firstPtr, *reversePtr, *numericPtr, *rowPtr, *filenamePtr)
	fmt.Println(inputFile)

	fmt.Println(lineBuf)

}
