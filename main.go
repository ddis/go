package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	resp, err := client.Get("https://rozetka.com.ua/lenovo-82xf004mra/p379983096/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	f, err := os.Create("test.txt")

	if err != nil {
		return
	}

	_, err = f.WriteString(string(body))

	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	f.Close()
}
