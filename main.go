package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"sync"
)

func SortFunc[T any](s []T, swap func(a T, b T) bool) {
	// Loop through whole slice starting at second element
	for i := 1; i < len(s); i++ {
		//
		for j := i; j > 0; j-- {
			if !swap(s[j], s[j-1]) {
				break
			}

			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func Sort(s []int) {
	// start at the second element
	for i := 1; i < len(s); i++ {
		// don't move forward in the step if the
		// item is already in the right place
		if s[i] > s[i-1] {
			continue
		}

		for j := i; j > 0; j-- {
			if s[j] > s[j-1] {
				break
			}
			// swap
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

// 1, 2, 3, 4, 5, 6, 7, 8, 9
func Bisect(s []int, item int, start int, end int) int {
	// grab the halfway point. if we're lucky, it's a match.
	half := int(math.Round(float64(end+start) / 2))

	if start > end {
		return -1
	}

	if item == s[half] {
		return half
	}

	if item < s[half] {
		end = half - 1
	}

	if item > s[half] {
		start = half + 1
	}

	return Bisect(s, item, start, end)
}

func FindIndex(s []int, item int) int {
	if len(s) == 0 {
		return -1
	}

	return Bisect(s, item, 0, len(s)-1)
}

func Change(s []int) {
	copy(s, []int{123, 456})
}

type Foo struct {
	Quantity int
}

func (f *Foo) Increase(n int) {
	f.Quantity += n
}

func (f *Foo) Decrease(n int) {
	f.Quantity -= n
}

func (f *Foo) String() string {
	return fmt.Sprintf("The quantity is %v", f.Quantity)
}

type Y struct {
	Name string
}

type X struct {
	Age int
	*Y
}

func ModifyY(x X) {
	x.Age = 5
	x.Y.Name = "CHANGED"
}

func LastIndexOf(nums []int, value int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != value {
			continue
		}

		return i
	}

	return -1
}

func singleNumber(nums []int) int {
	tracked := make(map[int]struct{}, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := tracked[nums[i]]; ok {
			continue
		}

		if LastIndexOf(nums, nums[i]) != i {
			tracked[nums[i]] = struct{}{}
			continue
		}

		return nums[i]
	}

	return -1
}

func isPalindrome(s string) bool {
	str := strings.ToLower(regexp.MustCompile(`[^a-zA-Z]`).ReplaceAllString(s, ""))

	for i := 0; i < len(str)/2; i++ {
		letter := str[i]
		if letter != str[len(str)-(i+1)] {
			return false
		}
	}

	return true
}

type YieldFunc[T any] func(data T)

type GeneratorHandler[T any] func(yield YieldFunc[T], done func())

type GeneratorChannel[T any] chan T

func (channel GeneratorChannel[T]) Next() T {
	return <-channel
}

func Generator[T any](handler GeneratorHandler[T]) GeneratorChannel[T] {
	wg := sync.WaitGroup{}
	wg.Add(1)
	channel := GeneratorChannel[T](make(chan T))

	// A function that will wait for the done signal, then close
	// the channel
	go func() {
		wg.Wait()
		close(channel)
	}()

	yield := func(data T) {
		channel <- data
	}
	done := func() {
		wg.Done()
	}

	// Run the handler in a goroutine
	go handler(yield, done)

	return channel
}

func main() {
	generator := Generator(func(yield YieldFunc[int], done func()) {
		for i := 0; i < 10; i++ {
			yield(i)
		}
		done()
	})

	fmt.Println(generator.Next())
	fmt.Println(generator.Next())
}
