package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var con *Console = New()

type Console struct {
	in *bufio.Reader
}

func New() (c *Console) {
	c = &Console{
		in: bufio.NewReader(os.Stdin),
	}
	return
}

func Get() (c *Console) { return con }

func Print(message ...interface{})              { con.Print(message...) }
func (c *Console) Print(message ...interface{}) { c.print(fmt.Sprint(message...)) }

func Println(message ...interface{})              { con.Println(message...) }
func (c *Console) Println(message ...interface{}) { c.print(fmt.Sprintln(message...)) }

func Printf(format string, values ...interface{}) { con.Printf(format, values...) }
func (c *Console) Printf(format string, values ...interface{}) {
	c.print(fmt.Sprintf(format, values...))
}

func AskEnter(prompt string) { con.AskEnter(prompt) }
func (c *Console) AskEnter(prompt string) {
	c.Println(prompt)
	c.getInput()
}

func AskEnterOrAbort(prompt, abortWith string) (abort bool) {
	return con.AskEnterOrAbort(prompt, abortWith)
}
func (c *Console) AskEnterOrAbort(prompt, abortWith string) (abort bool) {
	c.Println(prompt)
	abort = c.getInput() == "q"
	return
}

func AskYesOrNo(prompt string, def bool) (yes bool) { return con.AskYesOrNo(prompt, def) }
func (c *Console) AskYesOrNo(promt string, def bool) (yes bool) {
	c.Print(promt)
	if def {
		c.Print(" [Y/n] ")
	} else {
		c.Print(" [y/N] ")
	}
	in := c.getInput()
	switch strings.ToLower(in) {
	case "y":
		yes = true
	case "n":
		yes = false
	default:
		yes = def
	}
	return
}

func AskString(prompt string) (s string) { return con.AskString(prompt) }
func (c *Console) AskString(promt string) (s string) {
	c.Print(promt)
	return c.getInput()
}

func AskStringf(format string, values ...interface{}) (s string) {
	return con.AskStringf(format, values...)
}
func (c *Console) AskStringf(format string, values ...interface{}) (s string) {
	return c.AskString(fmt.Sprintf(format, values...))
}

func AskOption(prompt string, options ...string) (selected string, aborted bool) {
	return con.AskOption(prompt, options...)
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

func AskOptionValue(prompt string, options map[string]interface{}) (selected string, value interface{}, aborted bool) {
	return con.AskOptionValue(prompt, options)
}
func (c *Console) AskOptionValue(prompt string, options map[string]interface{}) (selected string, value interface{}, aborted bool) {
	var selections []string
	for selection := range options {
		selections = append(selections, selection)
	}
	selected, aborted = c.AskOption(prompt, selections...)
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
