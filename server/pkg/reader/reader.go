package reader

import (
	"fmt"
	"os"
)

func ReadVerses() {
	jsonfile, err := os.Open("../data/verses.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonfile.Close()
}
