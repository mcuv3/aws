package main

import (
	"log"
	"os"
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
