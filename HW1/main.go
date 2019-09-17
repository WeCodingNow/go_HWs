package main

import ( 
	"fmt"
	// "bufio"
	// "sort"
	"flag"
	"os"
)

func main() {
	// regPtr := flag.Bool("f", false, "Ignore case")
	// firstPtr := flag.Bool("u", false, "Output only the first of uniques")
	// reversePtr := flag.Bool("r", false, "Reverse sort")
	// numericPtr := flag.Bool("n", false, "Sort by numbers")
	// rowPtr := flag.Int("k", -1, "Sort by column")
	// filenamePtr := flag.String("o", "", "Output to file")

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

}