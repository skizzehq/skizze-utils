package main

import (
	"data"
	"fmt"
	"math/rand"
	"time"

	"github.com/skizzehq/goskizze/skizze"
)

func main() {
	words := data.GetData()
	// shuffle
	rand.Seed(time.Now().UnixNano())
	for i := range words {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}
	client, err := skizze.Dial("127.0.0.1:3596", skizze.Options{Insecure: true})

	if err != nil {
		fmt.Printf("Error connecting to Skizze: %s\n", err)
		return
	}
	domainName := "skizze_stress"
	if _, err := client.CreateDomain(domainName); err != nil {
		fmt.Println(err)
	}

	end := time.Duration(0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	zipf := rand.NewZipf(r, 1.1, 1.1, 500000)
	totalAdds := uint64(0)
	for i := 0; i < 100; i++ {
		word := words[i]
		n := zipf.Uint64() + 1
		fmt.Printf("%d Push: %s (%d times)\n", i, word, n)
		totalAdds += n
		fill := make([]string, n, n)
		for j := 0; j < len(fill); j++ {
			fill[j] = word
		}
		t := time.Now()
		if err := client.AddToDomain(domainName, fill...); err != nil {
			fmt.Println(err)
			return
		}
		end += time.Since(t)
	}

	client.Close()
	fmt.Printf("Added %d values (%d unique) in %ds\n", totalAdds, len(words), int(end.Seconds()))
}
