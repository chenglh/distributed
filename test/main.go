package main

import "fmt"

// 定义
type Registration struct {
	ServiceName string
	ServiceURL  string
}

func main() {
	registration := []Registration{
		{
			"service1",
			"url1",
		},
	}

	rst := []Registration{}

	url := "url1"
	for idx, _ := range registration {
		if registration[idx].ServiceURL == url {
			rst = append(registration[:idx], registration[idx+1:]...)
		}
	}

	fmt.Println(rst)
	//[]
}
