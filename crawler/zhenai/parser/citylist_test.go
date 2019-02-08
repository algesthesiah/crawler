package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//bytes, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	bytes, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", bytes)
	result := ParseCityList(bytes)
	fmt.Printf("%v\n", result)
}

