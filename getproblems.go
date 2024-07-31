package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strconv"
	// "strings"
)

func main() {
	fmt.Print("Type Your min rating:")
	var minRating string
	fmt.Scan(&minRating)
	fmt.Print("Type Your min rating:")
	var maxRating string
	fmt.Scan(&maxRating)
	fmt.Print("Type Your min rating:")
	the_problems := getProblems()
	// problems := make([]Problem, 0, len(submissions))
	for _, v := range the_problems {
		fmt.Println(v)
	}
}

// folderName := "./Codeforces Solutions"
// os.Mkdir(folderName, 0700)
// for _, acceptedProblem := range acceptedProblems {
// 	fmt.Println("https://codeforces.com/contest/" + strconv.Itoa(acceptedProblem.contestId) + "/problem/" + acceptedProblem.index)
// here use the acceptedproblems.ContestId to get the
// f, err := os.OpenFile(folderName+"/"+acceptedProblem.getFileName(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
// if err != nil {
// 	panic(err)
// }
// _, err = f.WriteString(acceptedProblem.getLink())
// if err != nil {
// 	panic(err)
// }
// fmt.Println("Created File:", acceptedProblem.getFileName())
// f.Close()
// }
// }

func getProblems() Submissions {
	resp, err := http.Get("https://codeforces.com/api/problemset.problems" + "&from=1&count=10000")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("HTTP Request Error! Check Your Connection and Check whether codeforces is running or not. : ", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	arr := status{}
	json.Unmarshal(body, &arr)
	if arr.Status == "FAILED" {
		fmt.Println("Codeforces handle incorrect! Please try again.")
	}
	return arr.Result
}

type Problems []struct {
	Problem             struct {
		ContestID int      `json:"contestId"`
		Index     string   `json:"index"`
		Name      string   `json:"name"`
		Type      string   `json:"type"`
		Rating    int      `json:"rating"`
		Points    float64  `json:"points"`
		Tags      []string `json:"tags"`
	} `json:"problem"`
}

type status struct {
	Status string      `json:"status"`
	Result Problems `json:"result"`
}

type Response struct {
	Status string `json:"status"`
	Result struct {
		Problems []Problem `json:"problems"`
	} `json:"result"`
}
