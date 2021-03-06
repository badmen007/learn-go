package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "xz",
		"course":  "golang",
		"site":    "xxx",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 === enpty map

	var m3 map[string]int // m3 === nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	fmt.Println(m)

	// map的key
	//map使用哈希表, 必须可以比较相等
	//除了slice,map,function的内建类型都可以作为key
	//Struct类型不包含上述字段，也可以作为key
}
