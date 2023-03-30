package reader

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var verses map[string]interface{}

// InitVerses reads verses.json and stores it in verse
// It is called once in main.go when the server starts
func InitVerses() {
	file, err := os.Open("verses.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&verses)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetTodayVerse() interface{} {
	today := time.Now()
	todayVerse := verses[today.Format("2006-01-02")]

	return todayVerse
}
