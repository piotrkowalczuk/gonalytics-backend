package lib

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
