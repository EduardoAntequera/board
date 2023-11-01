package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Position struct {
	Xval int `json:"xval"`
	Yval int `json:"yval"`
}

func board(width, height int) [][]int {
	var result [][]int
	for h := 0; h < height; h++ {
		var row []int
		for w := 0; w < width; w++ {
			row = append(row, 0)
		}
		result = append(result, row)
	}
	return result
}

func position() Position {
	resp, err := http.Get("http://localhost:8080/position")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var position Position

	json.Unmarshal(body, &position)
	return position
}

func main() {
	board := board(10, 10)
	position := position()
	board[position.Xval][position.Yval] = 1

	for i := 0; i < len(board[0]); i++ {
		fmt.Println(board[i])
	}
}
