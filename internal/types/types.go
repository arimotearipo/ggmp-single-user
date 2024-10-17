package types

type URI struct {
	Id  int
	Uri string
}

type PasswordGeneratorConfig struct {
	UppercaseLength int
	LowercaseLength int
	SpecialLength   int
	NumericLength   int
	TotalLength     int
}
