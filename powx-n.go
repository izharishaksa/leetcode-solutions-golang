//https://leetcode.com/problems/powx-n/

package leetcode_solutions_golang

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}

//2^7
//2*(2*2*2)*(2*2*2)=2*(2*2)^3
