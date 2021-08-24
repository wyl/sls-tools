package test

import (
	"fmt"
	"testing"
)

func TestPPP(t *testing.T) {

	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	for country := range countryCapitalMap {
		vv, vvv := countryCapitalMap[country]
		fmt.Println(vv, vvv)
		//fmt.Println(country, "首都是", countryCapitalMap [country])
	}
}
