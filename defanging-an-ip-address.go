package leetcode_solutions_golang

import "strings"

//https://leetcode.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
