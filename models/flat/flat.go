package flat

type Flat struct {
	id          int64
	street      string
	home        string
	structure   uint
	flatNumber  uint
	state       flatState
	floor       uint
	isCorner    bool
	flatType    string
	description string
}

type flatState string

const (
	Good   flatState = "хорошее"
	Euro   flatState = "евро"
	Soviet flatState = "советское"
)
