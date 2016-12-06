package lints

type ResultEnum int

const (
	NA    ResultEnum = iota // 0
	NE                      // 1
	Pass                    // 2
	Info                    // 3
	Warn                    // 4
	Error                   // 5
	Fatal                   // 6
)

type ResultStruct struct {
	Result  ResultEnum `json:"Result"` //this is the ResultEnum enumeration and uses the values found there
	Details string     `json:"Details,omitempty"`
}