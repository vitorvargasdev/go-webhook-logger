package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	port := strings.Split(os.Args[1:][0], "--port=")[1]

	http.HandleFunc("/listen", func(w http.ResponseWriter, r *http.Request) {
		text := strings.Builder{}

		text.WriteString("Headers:\n\n")
		text.WriteString("```text\n")

		for k, v := range r.Header {
			text.WriteString(k + ": " + v[0] + "\n")
		}

		text.WriteString("```\n")

		text.WriteString("\nBody:\n\n")
		text.WriteString("```json\n")

		body, err := io.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		prettyJSON, err := prettyJSON(body)

		text.WriteString(prettyJSON)

		text.WriteString("\n```\n")

		w.WriteHeader(http.StatusOK)

		date := strconv.FormatInt(time.Now().Unix(), 10)

		os.WriteFile("./logs/"+date+".md", []byte(text.String()), 0644)
	})

	http.ListenAndServe(":"+port, nil)
}

func prettyJSON(body []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "\t")

	if err != nil {
		log.Println("JSON parse error: ", err)
		return "", err
	}

	return string(prettyJSON.Bytes()), nil
}
