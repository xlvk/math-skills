package main

import (
	
	"bufio"
	"fmt"
	"log"
	"os"
	// "strings"
	"math"
	"strconv"
)

func main() {

	var arr []int
	sum := 0.0
	mean := 0
	sd := 0.0
	readFile, err := os.Open("data.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ha, e := strconv.Atoi(fileScanner.Text())
    	if e != nil {
        	fmt.Printf("%T \n %v", ha, ha)
		}
		arr = append(arr, ha)
		sum = sum + float64(ha)
	}

	fmt.Print("The Average is: ")
	fmt.Println(math.Round(sum/float64(len(arr))))
	fmt.Print("The Median is: ")
	if len(arr)%2 == 0 {
		mean = int(math.Round(float64((arr[(len(arr)/2)-1]) + (arr[(len(arr)/2)])))/2)
	} else {
		mean = int(math.Round(float64(arr[(len(arr)/2)])))
	}
	fmt.Println(mean)
	fmt.Print("The Variance is: ")
	for j := 0; j < 10; j++ {
		sd += math.Pow(float64(arr[j]-mean), 2)
	}
	varr := math.Round(sd / float64(len(arr)))
	fmt.Println(varr)
	fmt.Print("The Standard Deviation is: ")
	
	sd = math.Round(math.Sqrt(sd / float64(len(arr))))
	fmt.Println(sd)

}

