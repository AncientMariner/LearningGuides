package behavioral

import "fmt"

// Template Method: defines the skeleton of an algorithm; subclasses fill in steps.
// In Go, embed a struct and override methods, or use higher-order functions.

type DataMiner interface {
	ExtractData() []string
	ParseData(raw []string) []int
	AnalyzeData(data []int) string
}

// Mine runs the template algorithm.
func Mine(dm DataMiner) {
	raw := dm.ExtractData()
	data := dm.ParseData(raw)
	result := dm.AnalyzeData(data)
	fmt.Println("result:", result)
}

type CSVMiner struct{}

func (c *CSVMiner) ExtractData() []string { return []string{"1", "2", "3"} }
func (c *CSVMiner) ParseData(raw []string) []int {
	out := make([]int, len(raw))
	for i, s := range raw {
		fmt.Sscan(s, &out[i])
	}
	return out
}
func (c *CSVMiner) AnalyzeData(data []int) string {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return fmt.Sprintf("sum=%d", sum)
}

func ExampleTemplate() {
	Mine(&CSVMiner{})
	// Output:
	// result: sum=6
}
