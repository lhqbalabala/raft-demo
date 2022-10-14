package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxlen := 0
	len := len(s)
	a := make([]int, 200)
	l := 0
	r := 0
	for {
		for {
			if 0 != a[s[r]] {
				break
			}
			a[s[r]]++
			if maxlen < r-l+1 {
				maxlen = r - l + 1
			}
			r++
			if r == len {
				break
			}
		}
		for {
			a[s[l]]--
			l++
			if l == r {
				break
			}
		}
		if r == len {
			break
		}
	}
	return maxlen
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abca"))
	//var channelA chan struct{}
	//var channelB chan struct{}
	//var channelC chan struct{}
	//channelA = make(chan struct{}, 1)
	//channelB = make(chan struct{})
	//channelC = make(chan struct{})
	//channelA <- struct{}{}
	//go func() {
	//	for {
	//		<-channelA
	//		fmt.Println("A")
	//		channelB <- struct{}{}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//go func() {
	//	for {
	//		<-channelB
	//		fmt.Println("B")
	//		channelC <- struct{}{}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//go func() {
	//	for {
	//		<-channelC
	//		fmt.Println("C")
	//		channelA <- struct{}{}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//
	//select {}
}
