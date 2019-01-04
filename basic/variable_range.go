package main

import (
  "fmt"
)

/*
range 用于 array slice channel map
*/
func main() {	
	slic_range();

	arr_range();
	
	map_range();

	str_range();
}

func slic_range() {
	var sli1 = make([]string,2,5);
	sli1[0] = "1";
	sli1[1] = "2";
	for i, v:= range sli1 {
		fmt.Println(i,v);
	}
	fmt.Println("----------------");
}

func arr_range() {
	arr1 := [2]string{"a","b"};
	for index, v:= range arr1 {
		fmt.Println(index,v);
	}
	fmt.Println("----------------");
}

func map_range() {
	var map1 =map[string]string{"a":"a1","b":"b1"};
	for k,v:= range map1 {
		fmt.Println(k,v);
	}	
	fmt.Println("----------------");
}

//index => unicode
func str_range() {
	s1 := "ab";
	for index,unicodeIndex := range s1 {
		fmt.Println(index,unicodeIndex);
	}
	fmt.Println("----------------");
}