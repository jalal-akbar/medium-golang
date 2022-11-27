package main

import soliddesignprinciples "github.com/jalal-akbar/medium-golang/solid-design-principle"

func main() {
	// initialize journal
	j := &soliddesignprinciples.Journal{}
	// call AddEntru Func and inserting values
	j.AddEntry("This is my 1st journal")
	j.AddEntry("This is my 2st journal")
	j.AddEntry("This is my 3st journal")

	// call the RemoveEntry func to Remove index 2 from entry slice
	j.RemoveEntry(1)
	soliddesignprinciples.PrintListOfJournal(*j) // print what's left in the journal
}
