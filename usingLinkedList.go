package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd string
	description string
	nextnode *strorynode
	nextchoice *choices
}
type strorynode struct {
	text string
	choices *choices

}

func (node *strorynode) addchoice(cmd string, descriprion string, nextnode *strorynode) {
	choice:=&choices{cmd,descriprion,nextnode,nil}

	if node.choices==nil {
		node.choices=choice

	}else {
		currentchoice:=node.choices
		for currentchoice.nextchoice!=nil{
			currentchoice=currentchoice.nextchoice
		}
		currentchoice.nextchoice=choice
	}
}
func (node *strorynode) render()  {
	fmt.Println(node.text)
	currentchoice:=node.choices
	for currentchoice!=nil{
		fmt.Println(currentchoice.cmd,":",currentchoice.description)
		currentchoice=currentchoice.nextchoice
	}
}
func (node *strorynode) executecmd(cmd string)*strorynode  {
	currrentchoice:=node.choices
	for currrentchoice!=nil{
		if strings.ToLower(currrentchoice.cmd)==strings.ToLower(cmd) {
			return currrentchoice.nextnode
		}
		currrentchoice=currrentchoice.nextchoice
	}
	fmt.Println("Sorry, I did not understand that.")
	return node
}
var scanner *bufio.Scanner

func (node *strorynode) play() {
	node.render()
	if node.choices!=nil {
		scanner.Scan()
		node.executecmd(scanner.Text()).play()
	}

}
func main() {
	scanner=bufio.NewScanner(os.Stdin)

	start:=strorynode{
		text:    "you are in large chamber,deep underground." +
			"you see three passages leading out. A north passage leads into darkness"+
			"To the south, a passage appears to head upward. The eastern passages appears flat and well traveled",}
	darkroom:=strorynode{
		text:    "its is pitch black. you cannot see a thing",}

	darkroomlit:=strorynode{
		text:    "the dark passage is now lit by your lantern. you can continue north or head back south",}

	grue:=strorynode{
		text:    "while stumbling around in the darkness,you are eaten by grue",}
	trap:=strorynode{
		text:    "you head down the well traveled path when suddenly a trap door opens and you fall into a pit",}
	treasure:= strorynode{
		text:    "you arrive at a small chamber, filled with treasure",}

	start.addchoice("N","Go North",&darkroom)
	start.addchoice("S","Go South",&darkroom)
	start.addchoice("E","Go East",&trap)

	darkroom.addchoice("S","Try to go south",&grue)
	darkroom.addchoice("O","turn on lantern",&darkroomlit)

	darkroomlit.addchoice("N","Go North",&treasure)
	darkroomlit.addchoice("S","Go south",&start)

	start.play()

	fmt.Println()
	fmt.Println("The End")



}
