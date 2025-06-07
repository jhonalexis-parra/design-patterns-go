package singleresponsibility

import (
	"fmt"
	"os"
	"strings"
)

// The single responsibility princciple state that a type should have one primary responsibility, and as a result,
// it should have one reason to chnage that reason beign somehow related to its primary responsabiilty

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// Some logic, journal responsibility
}

// Separatins of concerns mean different concerns or different problems that the system solves have to reside
// in different construct, wheteher attached to different structures or residing in different package (split up)

// God object take everything in the kitechne sink that you are doing and put it into a single package

func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}
