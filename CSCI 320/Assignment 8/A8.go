package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Actor struct {
	LastName  string
	FirstName string
	Series    string
}

func (act Actor) String() string {
	return fmt.Sprintf("%s, %s (%s)", act.LastName, act.FirstName, act.Series)
}


type Actors []Actor
// Define methods for the Actors type to implement the sort interface
func (actors Actors) Len() int {
	return len(actors)
}

func (actors Actors) Less(i, j int) bool {
	return actors[i].LastName < actors[j].LastName
}

func (actors Actors) Swap(i, j int) {
	actors[i], actors[j] = actors[j], actors[i]
}
// sorts the Actors slice by a specific field based on its index
func (actors Actors) SortByField(fieldIndex int) {
	switch fieldIndex {
	case 0:
		sort.SliceStable(actors, func(i, j int) bool {
			return actors[i].LastName < actors[j].LastName
		})
	case 1:
		sort.SliceStable(actors, func(i, j int) bool {
			return actors[i].FirstName < actors[j].FirstName
		})
	case 2:
		sort.SliceStable(actors, func(i, j int) bool {
			return actors[i].Series < actors[j].Series
		})
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("<filename> <field index>")
		os.Exit(1)
	}

	filename := os.Args[1]
	fieldIndexArg, err := strconv.Atoi(os.Args[2])
	if err != nil || fieldIndexArg < 0 || fieldIndexArg > 2 {
		fmt.Println("can only be index 0, 1, or 2.")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var actorsList Actors

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			fmt.Println("Invalid line:", line)
			continue
		}
		actor := Actor{LastName: fields[0], FirstName: fields[1], Series: fields[2]}
		actorsList = append(actorsList, actor)
	}

	actorsList.SortByField(fieldIndexArg)

	for _, actor := range actorsList {
		fmt.Println(actor)
	}
}
