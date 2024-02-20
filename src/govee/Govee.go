package Govee

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Gov() {

	url := "https://developer-api.govee.com/v1/devices"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Govee-API-Key", "XXXXXXXXXXXXXXX")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
