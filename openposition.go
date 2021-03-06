package invt

import "fmt"

type OpenPosition struct {
	InstrumentId string
	Units        int32
	Price        float64
	Side         int
}

const (
	SIDE_BUY  = iota
	SIDE_SELL = iota
)

func ParseSide(sideStr string) int {
	if sideStr == "buy" {
		return SIDE_BUY
	}
	return SIDE_SELL
}

func StringOfSide(side int) string {
	if side == SIDE_BUY {
		return "buy"
	}
	return "sell"
}

func (pos *OpenPosition) FloatUnits() float64 {
	return float64(pos.Units)
}

func (pos *OpenPosition) Value() float64 {
	return pos.FloatUnits() * pos.Price
}

func (pos *OpenPosition) SwitchSide() {
	if pos.Side == SIDE_BUY {
		pos.Side = SIDE_SELL
	} else {
		pos.Side = SIDE_BUY
	}
}

func (pos *OpenPosition) SplitPosition(units int32) *OpenPosition {
	if units > pos.Units {
		panic("Cannot split position with this many units")
	}
	pos.Units -= units

	return &OpenPosition{pos.InstrumentId, units, pos.Price, pos.Side}
}

func (pos *OpenPosition) String() string {
	return fmt.Sprintf("u: %d\np: %f\ns: %d", pos.Units, pos.Price, pos.Side)
}
