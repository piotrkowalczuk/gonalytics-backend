package lib

import (
	"bytes"
	"sort"
)

const (
	// ConditionTypeEqual ...
	ConditionTypeEqual = "equal"
)

// Metric ...
type Metric struct {
	Dimensions []Dimension
}

// Dimension ...
type Dimension struct {
	Name      string    `xml:"name"`
	Condition Condition `xml:"condition"`
}

// Condition ...
type Condition struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// Meets ...
func (c *Condition) Meets(value string) bool {
	switch c.Type {
	case ConditionTypeEqual:
		return c.Value == value
	}

	return false
}

// SortDimensions ...
func (m *Metric) SortDimensions() {
	sort.Sort(dimensionsByName(m.Dimensions))
}

// DimensionsNamesAndDimensionsValues ...
func (m *Metric) DimensionsNamesAndDimensionsValues() (string, string) {
	m.SortDimensions()

	var dimensionNameBuffer bytes.Buffer
	var dimensionValueBuffer bytes.Buffer

	nbOfDimensions := len(m.Dimensions)

	for index, dimension := range m.Dimensions {
		dimensionNameBuffer.WriteString(dimension.Name)

		dimensionValueBuffer.WriteString("[")
		dimensionValueBuffer.WriteString(dimension.Condition.Type)
		dimensionValueBuffer.WriteString("=\"")
		dimensionValueBuffer.WriteString(dimension.Condition.Value)
		dimensionValueBuffer.WriteString("\"]")

		if index != nbOfDimensions-1 {
			dimensionNameBuffer.WriteString("|")
			dimensionValueBuffer.WriteString("|")
		}
	}

	return dimensionNameBuffer.String(), dimensionValueBuffer.String()
}

// DimensionsByName ...
type dimensionsByName []Dimension

// Len is part of sort.Interface.
func (dbn dimensionsByName) Len() int {
	return len(dbn)
}

// Swap is part of sort.Interface.
func (dbn dimensionsByName) Swap(i, j int) {
	dbn[i], dbn[j] = dbn[j], dbn[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (dbn dimensionsByName) Less(i, j int) bool {
	return dbn[i].Name < dbn[j].Name
}
