package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gin-gonic/gin"
	// "reflect"
	//"os"
	"strings"
	//"time"
	//"encoding/json"
)

// type PuzzleData struct {
// 	hints []int
// 	puzzle []int
// 	solution []int
// }

// type Level struct {
// 	day_of_week string
// 	difficulty string
// 	print_date string
// 	published string
// 	puzzle_id int
// 	version int
// 	puzzle_data PuzzleData
// }

// type Data struct {
// 	displayDate string
// 	easy Level
// 	hard Level
// 	medium Level
// }

type ReturnData struct {
	Easy string `json:"easy"`
	Medium string `json:"medium"`
	Hard string `json:"hard"`
}

func scrape() ReturnData {

	c := colly.NewCollector(
		colly.AllowedDomains("www.nytimes.com"),
	)


	var data ReturnData

	c.OnHTML("script", func (e *colly.HTMLElement) {

		if (strings.Contains(e.Text, "window.gameData")) {
			s := string(e.Text[18:])
			// fmt.Println("all: ", s)

			// Data := []byte(string(e.Text[18:]))
		
			// err := json.Unmarshal(Data, &data)

			// if err != nil {
			// 	fmt.Println(err)
			// }

			index := 0

			currIndex := 0

			var boardsArr [3]string

			for strings.Index(s[index:], "\"puzzle\"") != -1 {
				newIndex := strings.Index(s[index:], "\"puzzle\"")

				

				// fmt.Println(reflect.TypeOf(s[index + newIndex + 10 : index + newIndex + 171]))

				boardsArr[currIndex] = s[index + newIndex + 10 : index + newIndex + 171]				
				
				// fmt.Println(s[index + newIndex + 10 : index + newIndex + 171])

				index += newIndex + 1

				currIndex++
			}


			for i := 0; i < len(boardsArr); i++ {

				newString := ""
				currNum := 1

				for j := 0; j < len(boardsArr[i]); j += 2 {

					newString += string(boardsArr[i][j])

					currNum++

					if ((j != 0 && j != 160) && currNum % 10 == 0) {
						newString += "_"
						currNum = 1
					}

				}

				boardsArr[i] = newString
			}

			data.Easy = boardsArr[0]
			data.Hard = boardsArr[1]
			data.Medium = boardsArr[2]

		}
	

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.nytimes.com/puzzles/sudoku/easy")

	return data
}

func main() {
	r := gin.Default()

	r.GET("/getall", func(c *gin.Context){
		fmt.Println("Open")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		scrapedDate := scrape()
		c.JSON(200, scrapedDate)

	})


	r.Run()
	
}