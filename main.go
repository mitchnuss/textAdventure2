package main

// To Do
// Non player characters - talk to them, fight them
// NPC move around the graph
// Items that can be picked up or placed down
// Accept natural language as input
//

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, choice.description)
		}
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry, I didn't understand that.")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func main() {

	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	You are in a large chamber, deep underground.
	You see three passages leading out. A north passage leads into darkness.
	To the south,  a passage appears upward. The eastern passages appears
	flat and well traveled`}

	darkRoom := storyNode{text: "It is pitch black. You cannot see a thing."}

	darkRoomLit := storyNode{text: "The dark passage is now lit by your latern. You can continue north or head back south"}

	grue := storyNode{text: "While stumbling around in the darkness, you come into contact with a grue and you die."}

	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit"}

	treasure := storyNode{text: "You arrive at a small chamber, filled with treasure!"}

	fight := storyNode{text: "You wake up a grue and he starts Charging at you, you find a club next to you to pick up to fight or you can run away"}

	fightToDeath := storyNode{text: "You fight the grue with all your might but he ends up grabbing hold of you and rips you apart limb from limb"}

	runAway := storyNode{text: "You turn around and run way from the grue and you out run the grue and dies from exhaustion"}

	gasMatches := storyNode{text: "You find gas and matches so you use them to light a stick on fire"}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &fight)

	fight.addChoice("P", "Pick up club and fight", &fightToDeath)
	fight.addChoice("R", "Turn around and Run Away", &runAway)

	runAway.addChoice("L", "Turn Left", &trap)
	runAway.addChoice("R", "Turn Right", &gasMatches)

	gasMatches.addChoice("F", "Continue Forward", &darkRoomLit)
	gasMatches.addChoice("T", "Turn around and go the way you came", &trap)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	start.play()

	fmt.Println()
	fmt.Println("The End.")
}
