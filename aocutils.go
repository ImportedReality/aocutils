package aocutils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// File Utils

// OpenFile attempts to open a file with the given filename.
// It will panic if there are any issues opening the file.
// It returns a pointer to the File.
func OpenFile(filename string) *os.File {
	f, err := os.Open(filename)
	CheckErr(err)
	return f
}

// ReadSingleLineFile attempts to read a single line from a file.
// It will panic if there are any issues opening or reading the file.
// It returns a string.
func ReadSingleLine(filename string) (line string) {
	file := OpenFile(filename)
	defer file.Close()
	line, err := bufio.NewReader(file).ReadString('\n')
	CheckErr(err)
	return
}

// ReadLinesInFile attempts to read all lines in a file.
// It will panic if there are any issues opening or reading the file.
// It returns a slice of strings.
func ReadLines(filename string) (lines []string) {
	file := OpenFile(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

// ReadGrid attempts to read a grid from a file usign a given delimeter.
// It will panic if there are any issues opening or reading the file.
// It returns a slice of slices of strings ([][]string)
func ReadGrid(filename string, delim string) (grid Grid[string]) {
	file := OpenFile(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), delim)
		grid = append(grid, row)
	}
	return
}

// ReadNumberGrid attempts to read a grid of numbers from a file using a given delimeter
// It will panic if there are any issues opening or reading the file.
// It returns a slice of slices of ints ([][]int).
func ReadNumberGrid(filename string, delim string) (grid Grid[int]) {
	file := OpenFile(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)
		line := strings.Split(scanner.Text(), delim)
		for _, val := range line {
			row = append(row, StrToInt(val))
		}
		grid = append(grid, row)
	}
	return
}

// Error Utils

// CheckErr checks if the given err is nil, panicing if it isn't.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Conversions

// StrToInt attempts to convert a given string to an int.
// It will panic if the string cannot be converted.
// It returns an int.
func StrToInt(s string) (num int) {
	num, err := strconv.Atoi(s)
	CheckErr(err)
	return
}

// IntToStr converts a given int to a string
// It returns a string.
func IntToStr(num int) (s string) {
	s = strconv.Itoa(num)
	return
}

// Math

// Abs returns an int representing the aboslute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Pow returns an int representing n to the m power
func Pow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

// Array Utils
// Shamelessly copied from https://go.dev/wiki/SliceTricks

// Cut removes a range of elements from a slice of type T
// It returns a new slice of type T.
func Cut[T any](slice []T, start, end int) []T {
	copy(slice[start:], slice[end:])
	for k, n := len(slice)-end+start, len(slice); k < n; k++ {
		slice[k] = *new(T)
	}
	return slice[:len(slice)-end+start]
}

// Delete removes an element at a given index from a slice of type T,
// while preserving the original order of the slice.
// It returns a new slice of type T.
func Delete[T any](slice []T, index int) []T {
	copy(slice[index:], slice[index+1:])
	slice[len(slice)-1] = *new(T)
	return slice[:len(slice)-1]
}

// Insert adds an element at a given index to a slice of type T
// It returns a new slice of type T.
func Insert[T any](slice []T, element T, index int) []T {
	slice = append(slice, *new(T))
	copy(slice[index+1:], slice[index:])
	slice[index] = element
	return slice
}

// A type representing a slice of type T.
type Stack[T any] []T

// Push adds an element to the end of a stack of type T.
func (s Stack[T]) Push(element T) {
	s = append(s, element)
}

// Pop removes an element from the end of a stack of type T.
// It returns the removed element.
func (s Stack[T]) Pop() T {
	element, s := s[len(s)-1], s[:len(s)-1]
	return element
}

// Unshift adds an element to the beginning of a stack of type T.
func (s Stack[T]) Unshift(element T) {
	s = append([]T{element}, s...)
}

// Shift removes an element from the beginning of a stack of type T.
// It returns the removed element.
func (s Stack[T]) Shift() T {
	element, s := s[0], s[1:]
	return element
}

// Grid Utils

// A type representing a slice of slices of type T
type Grid[T any] [][]T

// A type representing an X and Y coordinate pair
type Coordinate struct{ x, y int }

// InBounds checks if the given coordinates are in the bounds of a given grid.
// The grid is assumed to be square
// It returns a bool.
func InBounds[T any](grid Grid[T], coord Coordinate) bool {
	return coord.y > 0 && coord.x > 0 && coord.y < len(grid) && coord.x < len(grid[0])
}

// Trees

type TreeNode[T any] struct {
	element     T
	firstChild  *TreeNode[T]
	nextSibling *TreeNode[T]
}

type BTreeNode[T any] struct {
	element T
	left    *BTreeNode[T]
	right   *BTreeNode[T]
}

func (t TreeNode[T]) GetNodes() []TreeNode[T] {
	nodes := make([]TreeNode[T], 0)
	nodes = append(nodes, t)
	if t.nextSibling != nil {
		nodes = append(nodes, t.nextSibling.GetNodes()...)
	}
	if t.firstChild != nil {
		nodes = append(nodes, t.firstChild.GetNodes()...)
	}

	return nodes
}
