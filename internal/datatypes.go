package internal

type Data struct {
	Board        string `json:"board" binding:"required"`
	SoilHumidity int    `json:"soil" binding:"required"`
}

type Action string

const (
	TurnPumpForSec Action = "TurnPumpForSec"
)

type Decision struct {
	Action Action `json:"action"`
	Value  int    `json:"value"`
}

type DataID struct {
	BoardID        int64
	ReceivedDataID int64
}
