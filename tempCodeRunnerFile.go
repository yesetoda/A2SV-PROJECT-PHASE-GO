tedProblem.contestId)+"problem/"+ acceptedProblem.index)
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
	}
}

func getSubmissions(handle string) Submissions {
	resp, err := http.Get("http://codeforces.com/api/user.status?handle=" + handle + "&from=1&count=10000")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("HTTP Request Error! Check Your Connection and Check whether codeforces is running or not. : ", err)
	}