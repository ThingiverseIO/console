package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Console struct {
	in *bufio.Reader
}

func New() (c *Console) {
	c = &Console{
		in: bufio.NewReader(os.Stdin),
	}
	return
}

func (c *Console) Print(message ...interface{})   { c.print(fmt.Sprint(message...)) }

func (c *Console) Println(message ...interface{}) { c.print(fmt.Sprintln(message...)) }

func (c *Console) Printf(format string, values ...interface{}) {
	c.print(fmt.Sprintf(format, values...))
}

func (c *Console) AskEnter(prompt string) {
	c.Println(prompt)
	c.getInput()
}

func (c *Console) AskString(promt string) (s string) {
	c.Print(promt)
	return c.getInput()
}

func (c *Console) AskStringf(format string, values ...interface{}) (s string) {
	return c.AskString(fmt.Sprintf(format, values...))
}

func (c *Console) AskOption(prompt string, options ...string) (selected string, aborted bool) {
	for i, option := range options {
		fmt.Printf(" [%d] %s\n", i+1, option)
	}
	fmt.Println(" [q] abort")
	c.Println(prompt)
	for !aborted {
		answer := c.AskStringf("Select [1-%d], 'q' to abort: ", len(options))
		if answer == "q" {
			aborted = true
			return
		}
		i, err := strconv.ParseInt(answer, 10, 64)
		if err != nil || int(i) > len(options) || i < 1 {
			c.Printf("Invalid input '%s'. ", answer)
		} else {
			selected = options[i-1]
			return
		}
	}
	return
}

func (c *Console) AskOptionValue(prompt string, options map[string]interface{}) (selected string, value interface{}, aborted bool) {
	var selections []string
	for selection := range options {
		selections = append(selections, selection)
	}
	selected, aborted = c.AskOption(prompt,selections...)
	if !aborted {
		value = options[selected]
	}

	return
}

func (c *Console) print(msg string) {
	fmt.Print(msg)
}

func (c *Console) getInput() (input string) {
	input, _ = c.in.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return
}
