package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {

	//router := gin.Default()

	//Scheduler := scheduler.New()

	//go Scheduler.Run()

	//router.Run(":8002")

	r, e := http.Get("http://127.0.0.1:8500/v1/catalog/service/approved-scheduler")
	if e != nil {
		fmt.Println(e)
	}


	b, e1 := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if e1 != nil {
		fmt.Println(e1)
	}
	fmt.Println(string(b))

}
