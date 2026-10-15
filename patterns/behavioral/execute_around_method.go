package behavioral

import "fmt"

// Execute Around Method: ensures setup/teardown always runs around an operation.
// defer is the idiomatic Go form; this shows the explicit higher-order function variant.

type Transaction struct {
	id string
}

func (t *Transaction) Begin()    { fmt.Printf("tx %s: begin\n", t.id) }
func (t *Transaction) Commit()   { fmt.Printf("tx %s: commit\n", t.id) }
func (t *Transaction) Rollback() { fmt.Printf("tx %s: rollback\n", t.id) }

func WithTransaction(tx *Transaction, fn func() error) error {
	tx.Begin()
	if err := fn(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func ExampleExecuteAroundMethod() {
	tx := &Transaction{id: "42"}
	_ = WithTransaction(tx, func() error {
		fmt.Println("doing work")
		return nil
	})
	// Output:
	// tx 42: begin
	// doing work
	// tx 42: commit
}
