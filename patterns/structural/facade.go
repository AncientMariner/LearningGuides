package structural

import "fmt"

// Facade: simplifies a complex subsystem behind a single entry point.

type CPU struct{}
type Memory struct{}
type HardDrive struct{}

func (c *CPU) Freeze()        { fmt.Println("CPU: freeze") }
func (c *CPU) Jump(pos int64) { fmt.Printf("CPU: jump to %d\n", pos) }
func (c *CPU) Execute()       { fmt.Println("CPU: execute") }

func (m *Memory) Load(pos int64, data []byte) {
	fmt.Printf("Memory: load %d bytes at %d\n", len(data), pos)
}

func (h *HardDrive) Read(lba int64, size int) []byte {
	fmt.Printf("HardDrive: read %d bytes at sector %d\n", size, lba)
	return make([]byte, size)
}

// Computer is the facade.
type Computer struct {
	cpu  *CPU
	mem  *Memory
	hd   *HardDrive
}

func NewComputer() *Computer {
	return &Computer{cpu: &CPU{}, mem: &Memory{}, hd: &HardDrive{}}
}

func (c *Computer) Start() {
	c.cpu.Freeze()
	data := c.hd.Read(0, 512)
	c.mem.Load(0, data)
	c.cpu.Jump(0)
	c.cpu.Execute()
}

func ExampleFacade() {
	computer := NewComputer()
	computer.Start()
	// Output:
	// CPU: freeze
	// HardDrive: read 512 bytes at sector 0
	// Memory: load 512 bytes at 0
	// CPU: jump to 0
	// CPU: execute
}
