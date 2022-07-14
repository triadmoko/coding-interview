# Coding Interview HackerRank
## How To Use
clone module 
```git
go get github.com/triadmoko/coding-interview
```

call the coding interview function inside main.go
### Algorithms

#### simple array sum
example test `simple array sum`

```go
func main() {
	val := []int{12, 124, 32, 43, 423, 423}
	algorithms.SimpleArraySum(val)
}
```
#### AVeryBigSum
example test `simple a very big sum`

```go
func main() {
	val := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	algorithms.AVeryBigSum(val)
}
```

## Data Structures
### Array DS
example test `array DS`
```go
func main() {
	val := []int{12, 124, 32, 43, 423, 423}
	datastructures.ReverseArray(val)
}
```