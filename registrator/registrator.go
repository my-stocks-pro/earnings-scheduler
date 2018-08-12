package registrator

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func GetServices() (*map[string][]string, error) {

	var Services map[string][]string

	res, err := http.Get("127.0.0.1:8500/v1/catalog/services")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	blob, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	err = json.Unmarshal(blob, &Services)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return &Services, nil
}
