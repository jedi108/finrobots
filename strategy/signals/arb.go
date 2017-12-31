package signals

type (
	Pair struct {
		ticker string
		Rate float64
		Amount float64
	}
	SignalArb struct {
		Pairs []Pair
		isOn bool
	}
)

func (signal *SignalArb) Calculating() {

}