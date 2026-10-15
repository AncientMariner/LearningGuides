package structural

import "fmt"

// Adapter: wraps an incompatible type to satisfy an interface.

type LegacyPrinter struct{}

func (l *LegacyPrinter) PrintBanner(text string) {
	fmt.Printf("*** %s ***\n", text)
}

// Printer is the interface new code expects.
type Printer interface {
	Print(text string)
}

type PrinterAdapter struct {
	legacy *LegacyPrinter
}

func NewPrinterAdapter(l *LegacyPrinter) Printer {
	return &PrinterAdapter{legacy: l}
}

func (a *PrinterAdapter) Print(text string) {
	a.legacy.PrintBanner(text)
}

func ExampleAdapter() {
	legacy := &LegacyPrinter{}
	printer := NewPrinterAdapter(legacy)
	printer.Print("hello")
	// Output:
	// *** hello ***
}
