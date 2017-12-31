package strategy

type (
	Strategy struct {
		Signals []ISignaler
		IsRight bool
	}
)

type ISignaler interface {
	Calculating()
}

func NewStrategy() *Strategy {
	return &Strategy{}
}

func (s *Strategy) Attach(signal ISignaler) {
	s.Signals = append(s.Signals, signal)
}

func (s *Strategy) Start() {
	for _, singaler := range s.Signals {
		go singaler.Calculating()
	}
}
