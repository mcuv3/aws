package main

import (
	"log"
	"math"
	"os"
	"strconv"
)

func PolicyTest() {
	js, err := os.Open("./policy.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer js.Close()

	// bytes, _ := ioutil.ReadAll(js)

	// policy, err := auth.ParsePolicy(string(bytes))

	if err != nil {
		log.Fatal(err)
		return
	}

	// splited := policy.Statements[2].Resource[0].GetService()
	// fmt.Println(len(splited))
	// fmt.Println(splited)

	// for _, p := range splited {
	// 	fmt.Println(p)
	// }

}

func getTarget(n, p int) int {
	s := strconv.Itoa(n)
	var total float64
	factor := float64(p)

	for _, d := range s {
		num, _ := strconv.Atoi(string(d))
		total += math.Pow(float64(num), factor)
		factor++
	}

	return int(total)
}
