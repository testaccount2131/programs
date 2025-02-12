package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func greatestsecond(arr []int) int {
	max1 := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				max1 = arr[i]
			}
		}
	}
	return max1
}
func greatestsecondop(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}
func greatestsecondop2(arr []int) int {
	second := -1
	large := -1
	for _, i := range arr {
		if i > large {
			large = i
			second = large
		} else {
			if i > second && i < large {
				second = i
			}
		}
	}
	return second
}
func maxprofit(arr []int) int {
	min := arr[0]
	max := 0
	for _, val := range arr {
		if val < min {
			min = val
		} else {
			profit := val - min
			if profit > max {
				max = profit
			}
		}

	}
	return max
}
func findlengthofstring(s string) int {
	len1 := 0
	reset := false
	for i := 0; i < len(s); i++ {
		letter := s[i : i+1]
		if letter == " " {
			reset = true
		}
		if reset {
			len1 = 0
			reset = false
		}
		len1++
	}
	return len1

}
func isPalindrome(x int) bool {
	strnum := strconv.Itoa(x)
	length := len(strnum)
	if length == 1 {
		return true
	}

	for i := 0; i < length/2; i++ {
		if strnum[i] == strnum[length-i-1] {
			return true
		}
	}
	return false
}
func maxprofit2(arr []int) int {
	profit := 0

	for i := 0; i < len(arr)-1; i++ {

		if arr[i+1] > arr[i] {
			profit += arr[i+1] - arr[i]
		}
	}
	return profit
}
func sortedarray(num1 []int, num2 []int) []int {
	m := len(num1)
	n := len(num2)
	z := make([]int, m+n)
	k, i, j := 0, 0, 0
	for i < m && j < n {
		if num1[i] < num2[j] {
			z[k] = num1[i]
			i++
		} else {
			z[k] = num2[j]
			j++

		}
		k++

	}
	if i < m {
		z[k] = num1[i]
		i++
		k++
	}
	if j < n {
		z[k] = num2[j]
		j++
		k++
	}
	return z

}
func searchinsert(arr []int, target int) int {
	mid := len(arr) / 2

	if len(arr) == 0 {
		arr[0] = target
	}
	if target < mid {
		for i := mid - 1; i > 0; i-- {
			if arr[i] == target {
				return i
			} else {
				if arr[i] > target {
					return i + 1
				}

			}
		}

	}
	if target > mid {
		for i := mid; i < len(arr); i++ {
			if arr[i] == target {
				return i
			} else {
				if arr[i] < target {
					return i - 1
				}

			}
		}

	}
	return 0
}
func singleNumber(ar []int) int {
	l := len(ar)
	sort.Ints(ar)
	for i := 0; i < l; i++ {

		if i == l-1 {
			return ar[i]
		}

		if ar[i] == ar[i+1] {
			i++
		} else {
			return ar[i]
		}
	}
	return 0
}
func plusone(arr []int) []int {
	m := len(arr)
	for i := m - 1; i > 0; i-- {
		if arr[i] == 9 {
			arr[i] = 0
		} else {
			arr[i]++
			return arr
		}

	}
	return append([]int{1}, arr...)
}

func medianofarr(arr []int) int {
	n := len(arr)
	mid := n / 2
	if n/2 == 0 {
		med := (arr[mid] + arr[mid+1]) / 2
		return med
	}
	return arr[mid]
}
func rotatearrar(arr []int, k int) []int {
	num := len(arr)
	k = k % num // make sure the value doesnt exceed the length
	res := make([]int, num)
	for i := 0; i < num; i++ {
		res[(i+k)%num] = arr[i]
	}
	copy(arr, res)

	return arr
}
func rotatearrarrecur(arr []int, k int) {
	if k == 0 || len(arr) <= 1 {
		return
	}
	last := arr[len(arr)-1]

	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = last
	rotatearrarrecur(arr, k-1)
}

func tripletpro(arr []int) int {
	sort.Ints(arr)
	return (arr[len(arr)-1] * arr[len(arr)-2] * arr[len(arr)-3])
}
func removenon_letters(s string) string {
	var final string

	for _, val := range s {
		if unicode.IsLetter(val) || unicode.IsDigit((val)) {
			final += string(val)
		}
	}
	return final
}
func twoSum(numbers []int, target int) []int {
	var l, r = 0, len(numbers) - 1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}

		if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}

	}
	return nil
}
func haystack(hay string, needle string) bool {
	for i := 0; i < len(hay); i++ {
		if len(needle)+i > len(hay) {
			return false
		} else if hay[i:len(needle)+i] == needle {
			return true
		}
	}

	return false
}
func isSubsequence(large string, small string) bool {
	l1 := []rune(large)
	sm := []rune(small)
	length := len(small) - 1
	count := 0
	for count <= length {
		found := false
		for _, val := range l1 {
			if val == sm[count] {
				count++
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
func isSubsequence_2(s string, t string) bool {
	queue := []rune(s)
	for _, char := range t {
		if len(queue) == 0 {
			continue
		}

		if queue[0] == char {
			queue = queue[1:]
		}
	}

	return len(queue) == 0
}

func validpalindrone(s string) bool {
	s = removenon_letters(s)
	s = strings.ToLower(s)
	str := []rune(s)
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	c := "helloworld"
	d := "worlg"
	//q := 123123
	a := []int{2, 5, 5, 9}
	b := []int{1, 3, 5, 6}
	arr12 := []int{3, 5, 7, 1, 36, 5, 3, 2}
	stat := haystack(c, d)
	//w := isPalindrome(q)
	f := sortedarray(a, b)
	h := medianofarr(f)
	rotatearrarrecur(arr12, 4)
	r := singleNumber(a)
	s := "Heloo ffgfgfg     fbfbfbfbefb"
	value := findlengthofstring(s)
	e := plusone(a)
	fmt.Println(value)
	fmt.Println(f)
	fmt.Println(e)
	fmt.Println(r)
	fmt.Println(h)
	fmt.Println(stat)
	fmt.Println(arr12)
}
