package utils

import (
	"fmt"
	"strings"
)

// UniqueName generates a new unique string from a give string
func uniqueName() {
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
}
