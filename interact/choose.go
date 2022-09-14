package interact

import (
	"fmt"
)

// ChoiceHanlder defines function signature of a vaild choice function.
type ChoiceHanlder func() error

// ToPrevious implements error interface that let app step out current choose loop.
type ToPrevious struct {
	Err error
}

func (e *ToPrevious) Error() string {
	return "exit the loop and go back to previous choose node"
}

// NewToPreviouswithErr warps error to ToPrevious.
func NewToPreviouswithErr(err error) error {
	return &ToPrevious{Err: err}
}

// DefaultChoice implements Choice interface with essential fields.
type DefaultChoice struct {
	Description string
	Function    ChoiceHanlder
}

// Desc returnn description
func (c *DefaultChoice) Desc() string {
	return c.Description
}

// Run the choose.
func (c *DefaultChoice) Run() error {
	return c.Function()
}

// NewChoice returns a Choice node.
func NewChoice(desc string, function func() error) Choice {
	return &DefaultChoice{
		Description: desc,
		Function:    function,
	}
}

// Choice defines the basic abilities of Choice node.
//
// This interface might be extended in the future.
type Choice interface {
	Desc() string
	Run() error
}

// Choose defines essential fields to let this interaction work.
type Choose struct {
	Title                string
	Choices              []Choice
	Loop                 bool
	ToPreviousChooseDesc string

	// TODO
	ToNextPageDesc *Choice
	// TODO
	ToPreviousPageDesc *Choice
}

// ToClosure warps the Choose interaction to a ChoiceHanlder.
// It can be used to create a nesting choose interaction.
func (c *Choose) ToClosure() ChoiceHanlder {
	return func() error {
		for {
			fmt.Printf("\n%v\n\n", c.Title)
			for i := range c.Choices {
				fmt.Printf("(%v): %v\n", i+1, c.Choices[i].Desc())
			}

			fmt.Printf("\n(0): %v\n\n", c.ToPreviousChooseDesc)

			op := ScanInputAndCheck()
			if op > len(c.Choices) || op < 0 {
				fmt.Println("Invaild input")
				continue
			}
			if op == 0 {
				break
			}
			err := c.Choices[op-1].Run()

			if c.Loop {
				if err != nil {
					if warp, ok := err.(*ToPrevious); ok {
						return warp.Err
					}
					fmt.Printf("Err: %v\n", err)
				}
				continue
			}
			break
		}
		return nil
	}
}

// Do the interaction.
func (c *Choose) Do() error {
	if c.ToPreviousChooseDesc == "" {
		c.ToPreviousChooseDesc = "Exit"
	}
	return c.ToClosure()()
}
