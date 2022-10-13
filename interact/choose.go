package interact

import (
	"fmt"
)

// ChoiceHanlder defines function signature of a vaild choice function.
type ChoiceHanlder func() (func(), error)

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
func (c *DefaultChoice) Run() (func(), error) {
	return c.Function()
}

// NewChoice returns a Choice node.
func NewChoice(desc string, function func() (func(), error)) Choice {
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
	Run() (output func(), err error)
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
	return func() (func(), error) {
		var msg func()
		for i := 0; ; i++ {
			if i != 0 {
				fmt.Printf("\033[2J")
				fmt.Printf("\033[H")
			}
			if msg != nil {
				fmt.Printf("Output:\n\n")
				msg()
				fmt.Println()
			}
			fmt.Printf("%v\n\n", c.Title)
			for i := range c.Choices {
				fmt.Printf("(%v): %v\n", i+1, c.Choices[i].Desc())
			}

			fmt.Printf("\n(0): %v\n\n", c.ToPreviousChooseDesc)

			op := ScanInputAndCheck()
			if op > len(c.Choices) || op < 0 {
				msg = func() { fmt.Printf("Invaild input\n") }
				continue
			}
			if op == 0 {
				break
			}
			var err error
			msg, err = c.Choices[op-1].Run()

			if c.Loop {
				if err != nil {
					if warp, ok := err.(*ToPrevious); ok {
						return msg, warp.Err
					}
					msg = func() { fmt.Printf("Err: %v\n", err) }
				}
				continue
			}
			break
		}
		return nil, nil
	}
}

// Do the interaction.
func (c *Choose) Do() (func(), error) {
	if c.ToPreviousChooseDesc == "" {
		c.ToPreviousChooseDesc = "Exit"
	}
	return c.ToClosure()()
}
