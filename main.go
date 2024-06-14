package main

import "fmt"

type Ebook struct {
	PandocAPIVersion []int `json:"pandoc-api-version"`
	Meta             struct {
		Author struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"author"`
		Date struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"date"`
		Identifier struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"identifier"`
		Language struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"language"`
		Rights struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"rights"`
		Source struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"source"`
		Title struct {
			T string `json:"t"`
			C []struct {
				T string `json:"t"`
				C string `json:"c"`
			} `json:"c"`
		} `json:"title"`
	} `json:"meta"`
	Blocks []struct {
		T string `json:"t"`
		C []struct {
			T string          `json:"t"`
			C [][]interface{} `json:"c"`
		} `json:"c"`
	} `json:"blocks"`
}

func main() {
	fmt.Println("Hello, World!")
	initAPI()

}
