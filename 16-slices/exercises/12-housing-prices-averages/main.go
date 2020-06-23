// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locations                  []string
		sizes, beds, baths, prices []int
	)
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		tmp := strings.Split(row, separator)

		locations = append(locations, tmp[0])
		num, _ := strconv.Atoi(tmp[1])

		sizes = append(sizes, num)
		num, _ = strconv.Atoi(tmp[2])

		beds = append(beds, num)
		num, _ = strconv.Atoi(tmp[3])

		baths = append(baths, num)
		num, _ = strconv.Atoi(tmp[4])
		prices = append(prices, num)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	var sum int
	denom := float64(len(rows))
	fmt.Printf("%-15s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s", "")
	for _, v := range sizes {
		sum += v
	}

	sum = 0
	fmt.Printf("%-15.2f", float64(sum)/denom)
	for _, v := range beds {
		sum += v
	}

	sum = 0
	fmt.Printf("%-15.2f", float64(sum)/denom)
	for _, v := range baths {
		sum += v
	}

	sum = 0
	fmt.Printf("%-15.2f", float64(sum)/denom)
	for _, v := range prices {
		sum += v
	}
	fmt.Printf("%-15.2f", float64(sum)/denom)

	fmt.Println()
}
