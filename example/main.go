package main

import "github.com/ThingiverseIO/console"

func main() {
	c := console.New()
	c.Println("Hello Console")
	s := c.AskString("Write something and hit enter: ")
	c.Println("Great, you wrote: ", s)

	cpts := map[string]interface{}{
		"Jean-Luc":  "Good Choice!",
		"Ben":       "Badass guy.",
		"Catherine": "Well...",
		"Jonathan":  "Really?",
		"Jim":       "Oldschool!",
	}

	cpt, opinion, aborted := c.AskOptionValue("Who is the best captain?", cpts)
	if aborted {
		c.Println("Don't like Star Trek? A Shame!")
	} else {
		c.Printf("You have chosen '%s'. %s\n", cpt, opinion)
	}
	c.AskEnter("Hit return to exit")
}
