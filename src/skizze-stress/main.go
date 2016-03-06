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
	zipf := rand.NewZipf(r, 1.1, 1.1, uint64(len(words)-1))
	totalAdds := 0
	for i := 0; i < 100000; i++ {
		fill := make([]string, 1000, 1000)
		for j := 0; j < len(fill); j++ {
			k := zipf.Uint64()
			fill[j] = words[k]
		}
		totalAdds += len(fill)

		t := time.Now()
		if err := client.AddToDomain(domainName, fill...); err != nil {
			fmt.Println(err)
			return
		}
		end += time.Since(t)
		if end.Seconds() > 0 {
			fmt.Printf("Added %d values (%d unique) in %ds (avg. %d v/s)\n", totalAdds, len(words), int(end.Seconds()), totalAdds/int(end.Seconds()+1))
		}
	}

	client.Close()
	fmt.Printf("Added %d values (%d unique) in %ds (avg. %d v/s)\n", totalAdds, len(words), int(end.Seconds()), totalAdds/int(end.Seconds()+1))
}
