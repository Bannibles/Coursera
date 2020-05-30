package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//Animal is an interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow struct
type Cow struct {
	food, locomotion, noise string
}

//Eat cow
func (a Cow) Eat() {
	fmt.Println(a.food)
}

//Move cow
func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

//Speak cow
func (a Cow) Speak() {
	fmt.Println(a.noise)
}

//Bird struct
type Bird struct {
	food, locomotion, noise string
}

//Eat bird
func (a Bird) Eat() {
	fmt.Println(a.food)
}

//Move bird
func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

//Speak bird
func (a Bird) Speak() {
	fmt.Println(a.noise)
}

//Snake struct
type Snake struct {
	food, locomotion, noise string
}

//Eat snake
func (a Snake) Eat() {
	fmt.Println(a.food)
}

//Move snake
func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

//Speak snake
func (a Snake) Speak() {
	fmt.Println(a.noise)
}

//UserInput sees what user types
func UserInput() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Seperate commands by a space \n")
	fmt.Printf("For a query: enter 'query', name of animal, and its action \n")
	fmt.Printf("For a new animal: enter 'newanimal', name of new animal, and its properties of a cow/snake/bird \n >")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	textSlice := strings.Split(text, " ")
	return textSlice
}

func animalAction(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("invalid action: " + action)
	}
}

func newAnimal(kind string) (Animal, error) {
	var animal Animal
	switch kind {
	case "cow":
		animal = Cow{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo"}
	case "bird":
		animal = Bird{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep"}
	case "snake":
		animal = Snake{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss"}
	default:
		return nil, errors.New("invalid animal kind: " + kind)
	}
	return animal, nil
}

func main() {
	animals := make(map[string]Animal)
	for {
		//userInput := []string{"newanimal", "foo", "bird"}
		userInput := UserInput()
		fmt.Printf("%v\n", userInput)
		if len(userInput) != 3 {
			fmt.Printf("Invalid number of inputs \n")
			continue
		}
		switch userInput[0] {
		case "query":
			name := userInput[1]
			action := userInput[2]
			animal := animals[name]
			if animal != nil {
				animalAction(animal, action)
			}
			break
		case "newanimal":
			name := userInput[1]
			kind := userInput[2]
			a, err := newAnimal(kind)
			if err != nil {
				fmt.Println(err.Error())
				continue
			} else {
				animals[name] = a
				fmt.Println("Created it!")
			}
			break
		default:
			fmt.Println("invalid input for first command")
			continue
		}
	}
}
