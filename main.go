package main

import (
	"fmt"
	"strings"
)

// var remoteRouteMap map[string][]string

func main() {

	subsetAddrs := make([]string, 0)
	fmt.Println(subsetAddrs)
	a := "I am a"
	b := "slice-rbac-deployment-role"

	subsetAddrs = append(subsetAddrs, a)
	fmt.Println(subsetAddrs)

	rolebindingNameArr := strings.Split(b, "-")
	var rolebindingName []string
	fmt.Println(rolebindingNameArr)

	for i, name := range rolebindingNameArr {
		if i == len(rolebindingNameArr)-1 {
			continue
		}
		rolebindingName = append(rolebindingName, name)
	}
	newRolebindingName := strings.Join(rolebindingName, "-")
	fmt.Println(newRolebindingName + "-rolebinding")

	// a := "I am a"
	// b := "I am b"
	// c := "I am c"

	// var strarr = []string{"abcd", "efgh"}

	// remoteRouteMap = make(map[string][]string)

	// for i := 0; i < len(strarr)-1; i++ {
	// 	fmt.Println(strarr[0])
	// 	remoteRouteMap[a] = append(remoteRouteMap[a], strarr[0])
	// }

	// remoteRouteMap[a] = append(remoteRouteMap[a], b)
	// remoteRouteMap[a] = append(remoteRouteMap[a], c)

	// var count = 0

	// for i := 0; i < len(remoteRouteMap[a]); i++ {
	// 	count++
	// 	// fmt.Println(remoteRouteMap[a][0])

	// 	if reflect.DeepEqual(remoteRouteMap[a][i], c) {
	// 		fmt.Println("Matched")
	// 	}

	// }

	// fmt.Println(remoteRouteMap)

	// for _, val := range remoteRouteMap {
	// 	fmt.Println(len(val))
	// }

}
