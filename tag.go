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
	Address string
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
	l := strconv.Itoa(line)
	return Tag{
		Name:    name,
		File:    file,
		Address: l,
		Type:    tagType,
		Fields:  map[TagField]string{Line: l},
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
	Line  int
	Field string
	Type  TagType
}

func (t Tag) ToJsonCompact() *CompactTag {
	line, _ := strconv.Atoi(t.Address)
	switch t.Type {
	case Package:
	case Import, Constant, Interface:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s", t.Name)}
	case Variable, Type:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s %s", t.Name, t.Fields["type"])}
	case Field:
		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s %s %s", t.Fields["ctype"], t.Name, t.Fields["type"])}
	case Embedded:
	case Method, Prototype:
		_type, ok := t.Fields["ctype"]
		if !ok {
			_type = t.Fields["ntype"]
		}
		var fs string
		_, ok = t.Fields["type"]
		if !ok {
			fs = "%s"
		} else {
			fs = "(%s)"
		}

		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s %s%s"+fs, _type, t.Name, t.Fields["signature"], t.Fields["type"])}
	case Function:
		var fs string
		_, ok := t.Fields["type"]
		if !ok {
			fs = "%s"
		} else {
			fs = "(%s)"
		}

		return &CompactTag{Type: t.Type, Line: line, Field: fmt.Sprintf("%s%s"+fs, t.Name, t.Fields["signature"], t.Fields["type"])}

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
	b.WriteString(t.Address)
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
