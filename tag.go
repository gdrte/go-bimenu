package main

import (

	"encoding/json"

	"strconv"
)

// Tag represents a single tag.
type Tag struct {
	Name    string
	File    string
	Address int
	Type    TagType
	Fields  map[TagField]string
}

// TagField represents a single field in a tag line.
type TagField string

// Tag fields.
const (
	Access        TagField = "access"
	Signature     TagField = "signature"
	TypeField     TagField = "type"
	ReceiverType  TagField = "ctype"
	Line          TagField = "line"
	InterfaceType TagField = "ntype"
	Language      TagField = "language"
)

// TagType represents the type of a tag in a tag line.
type TagType string

// Tag types.
const (
	Package     TagType = "p"
	Import      TagType = "i"
	Constant    TagType = "c"
	Variable    TagType = "v"
	Type        TagType = "t"
	Interface   TagType = "n"
	Prototype   TagType = "o"
	Field       TagType = "w"
	Embedded    TagType = "e"
	Method      TagType = "m"
	Constructor TagType = "r"
	Function    TagType = "f"
)

// NewTag creates a new Tag.
func NewTag(name, file string, line int, tagType TagType) Tag {
	return Tag{
		Name:    name,
		File:    file,
		Address: line,
		Type:    tagType,
		Fields:  map[TagField]string{Line: strconv.Itoa(line)},
	}
}

func (t Tag) String() string {
	jbytes, err := json.Marshal(t)
	if err != nil {
		return "{}"
	}
	return string(jbytes)
}


func (t Tag) This() Tag {
	return t
}
