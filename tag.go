package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

func (t Tag) ToJson() string {
	jbytes, err := json.Marshal(t)
	if err != nil {
		return "{}"
	}
	return string(jbytes)
}

type CompactTag struct {
	File  string
	Line  int
	Field string
	Type  TagType
}

func (t Tag) ToJsonCompact() *CompactTag {
	line := t.Address
	ifBlank := func(s string) string {
		if len(s) == 0 {
			return "%s"
		}
		return "(%s)"
	}

	switch t.Type {
	case Package:
	case Import, Constant, Interface:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s", t.Name), File: t.File}
	case Variable, Type:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s %s", t.Name, t.Fields["type"]), File: t.File}
	case Field:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s %s %s", t.Fields["ctype"], t.Name, t.Fields["type"]), File: t.File}
	case Embedded:
	case Method, Prototype:
		_type, ok := t.Fields["ctype"]
		if !ok {
			_type = t.Fields["ntype"]
		}
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s %s%s"+ifBlank(t.Fields["type"]), _type, t.Name, t.Fields["signature"], t.Fields["type"]), File: t.File}
	case Function:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s%s"+ifBlank(t.Fields["type"]), t.Name, t.Fields["signature"], t.Fields["type"]), File: t.File}

	}
	return nil
}
func (t Tag) This() Tag {
	return t
}

// The tags file format string representation of this tag.
func (t Tag) String() string {
	var b bytes.Buffer

	b.WriteString(t.Name)
	b.WriteByte('\t')
	b.WriteString(t.File)
	b.WriteByte('\t')
	b.WriteString(strconv.Itoa(t.Address))
	b.WriteString(";\"\t")
	b.WriteString(string(t.Type))
	b.WriteByte('\t')

	fields := make([]string, 0, len(t.Fields))
	i := 0
	for k, v := range t.Fields {
		if len(v) == 0 {
			continue
		}
		fields = append(fields, fmt.Sprintf("%s:%s", k, v))
		i++
	}

	sort.Sort(sort.StringSlice(fields))
	b.WriteString(strings.Join(fields, "\t"))

	return b.String()
}
