package ir

import (
	"fmt"
	"strconv"
)

// GlobalIdent is a global identifier.
type GlobalIdent struct {
	GlobalName string
	GlobalID   int64
}

// Ident returns the identifier associated with the global identifier.
func (i GlobalIdent) Ident() string {
	if i.IsUnnamed() {
		return GlobalID(i.GlobalID)
	}
	return GlobalName(i.GlobalName)
}

// Name returns the name of the global identifier.
//
// If unnamed, the global ID is returned. To distinguish numeric names from
// unnamed IDs, numeric names are quoted.
func (i GlobalIdent) Name() string {
	if i.IsUnnamed() {
		return strconv.FormatInt(i.GlobalID, 10)
	}
	if x, err := strconv.ParseInt(i.GlobalName, 10, 64); err == nil {
		// Print GlobalName with quotes if it is a number; e.g. "42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return i.GlobalName
}

// SetName sets the name of the global identifier.
func (i *GlobalIdent) SetName(name string) {
	i.GlobalName = name
	i.GlobalID = 0
}

// ID returns the ID of the global identifier.
func (i GlobalIdent) ID() int64 {
	return i.GlobalID
}

// SetID sets the ID of the global identifier.
func (i *GlobalIdent) SetID(id int64) {
	i.GlobalID = id
}

// IsUnnamed reports whether the global identifier is unnamed.
func (i GlobalIdent) IsUnnamed() bool {
	return len(i.GlobalName) == 0
}

// LocalIdent is a local identifier.
type LocalIdent struct {
	LocalName string
	LocalID   int64
}

// NewLocalIdent returns a new local identifier based on the given string. An
// unnamed local ID is used if ident is an integer, and a local name otherwise.
func NewLocalIdent(ident string) LocalIdent {
	if id, err := strconv.ParseInt(ident, 10, 64); err == nil {
		return LocalIdent{LocalID: id}
	}
	return LocalIdent{LocalName: ident}
}

// Ident returns the identifier associated with the local identifier.
func (i LocalIdent) Ident() string {
	if i.IsUnnamed() {
		return LocalID(i.LocalID)
	}
	return LocalName(i.LocalName)
}

// Name returns the name of the local identifier.
//
// If unnamed, the local ID is returned. To distinguish numeric names from
// unnamed IDs, numeric names are quoted.
func (i LocalIdent) Name() string {
	if i.IsUnnamed() {
		return strconv.FormatInt(i.LocalID, 10)
	}
	if x, err := strconv.ParseInt(i.LocalName, 10, 64); err == nil {
		// Print LocalName with quotes if it is a number; e.g. "42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return i.LocalName
}

// SetName sets the name of the local identifier.
func (i *LocalIdent) SetName(name string) {
	i.LocalName = name
	i.LocalID = 0
}

// ID returns the ID of the local identifier.
func (i LocalIdent) ID() int64 {
	return i.LocalID
}

// SetID sets the ID of the local identifier.
func (i *LocalIdent) SetID(id int64) {
	i.LocalID = id
}

// IsUnnamed reports whether the local identifier is unnamed.
func (i LocalIdent) IsUnnamed() bool {
	return len(i.LocalName) == 0
}

// Ident is a named variable.
type Ident interface {
	Named
	// ID returns the ID of the local identifier.
	ID() int64
	// SetID sets the ID of the local identifier.
	SetID(id int64)
	// IsUnnamed reports whether the local identifier is unnamed.
	IsUnnamed() bool
}
