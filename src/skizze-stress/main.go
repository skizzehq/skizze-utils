package main

import (
	"data"
	"fmt"
	"math/rand"

	"github.com/seiflotfy/goskizze/skizze"
)

func main() {
	words := data.GetData()
	// shuffle
	for i := range words {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}
	client, err := skizze.Dial("127.0.0.1:3596", skizze.Options{Insecure: true})

	if err != nil {
		fmt.Printf("Error connecting to Skizze: %s\n", err)
		return
	}
	domainName := "stress"
	if _, err := client.CreateDomain(domainName); err != nil {
		fmt.Println(err)
		return
	}

	/*
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		zipf := rand.NewZipf(r, 3.14, 2.72, 10000)
		for i := 0; i < len(words); i++ {
			word := words[i]
			n := zipf.Uint64() + 1
			fmt.Printf("%d Push: %s (%d times)\n", i, word, n)
			fill := make([]string, n, n)
			for j := 0; j < len(fill); j++ {
				fill[j] = word
			}
			if err := client.AddToDomain(domainName, fill...); err != nil {
				fmt.Println(err)
				return
			}
		}
		if err := client.DeleteDomain(domainName); err != nil {
			fmt.Println(err)
			return
		}
	*/
}
