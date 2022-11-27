package soliddesignprinciples

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Single Responsibility Principle
// define a journal structure
type Journal struct {
	entries []string
}

// printListOfJournal log the current state of the journal
func PrintListOfJournal(Journal Journal) {
	log.Println(Journal)
}

// declare a global var for entry count
var entryCount int

// func AddEntry adds entries to the journal
func (j *Journal) AddEntry(text string) int {
	// increment the count of every entry
	entryCount++
	// concanate the index and input text
	strJoin := strconv.Itoa(entryCount) + "." + text
	j.entries = append(j.entries, strJoin)

	//format string and return count of entries
	s := fmt.Sprintf("%s", strJoin)
	log.Println(s)
	return entryCount
}

// func removeEntry remove entry from journal
func (j *Journal) RemoveEntry(index int) {
	// range over the entries and skip entry to romove
	for ind, _ := range j.entries {
		if ind == index {
			j.entries = append(j.entries[:ind], j.entries[ind+1:]...)
			break
		}
	}
}

/*Break The Single Responsibility  Principle*/
func (j *Journal) SaveToFile(filename string) {
	err := ioutil.WriteFile(filename, []byte(strings.Join(j.entries, "\n")), 0644)
	if err != nil {
		log.Println(err)
	}
}

// CORRECT METHOD: Creating another structure to handle persistence.
type Persistance struct {
	lineSeparator string
}

func (p *Persistance) SaveToFileCorrect(j *Journal, filename string) {
	err := ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
	if err != nil {
		log.Println(err)
	}
}
