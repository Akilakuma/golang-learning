package main

import (
	"net/http"
	"fmt"
	_ "reflect"
	"encoding/json"
	s "./format"
	"github.com/mozillazg/request"
	"github.com/bitly/go-simplejson"

)

func main() {
	testAPI()
	testAPI2()
}

func testAPI() *simplejson.Json {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get("https://api.nasa.gov/neo/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&api_key=DEMO_KEY")
	j, err := resp.Json()
	defer resp.Body.Close()  // Don't forget close the response body

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(reflect.TypeOf(j))

	// 用Get的方式取到值
	nextValue := j.Get("links").Get("next").MustString()
	fmt.Println(nextValue)

	// 用Get的方式取到值
	count := j.Get("element_count").MustInt()
	fmt.Println(count)

	// 可以連續用Get的方式取到值，但好像遇到了object array，所以卡住無法再往下取到東西
	neo := j.Get("near_earth_objects").Get("2015-09-08").Get("neo_reference_id").MustString()
	fmt.Println(neo)

	dateArray := j.Get("near_earth_objects").Get("2015-09-08").MustArray()
	for _, aValue := range dateArray {
		// 要繼續解出更裡面的東西實在非常麻煩
		fmt.Println(aValue)
	}

	return j
}



func testAPI2() s.Planetary {
	c := new(http.Client)
	req := request.NewRequest(c)
	resp, err := req.Get("https://api.nasa.gov/neo/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&api_key=DEMO_KEY")

	var response s.Planetary
	json.Unmarshal(resp, &response)
	defer resp.Body.Close()  // Don't forget close the response body

	if err != nil {
		fmt.Println(err.Error())
	}


	return j
}
