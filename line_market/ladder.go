package line_market

import "fmt"

type Outcome struct {
	selection *Interval
	value     int64
}

func (o *Outcome) String() string {
	return fmt.Sprintf("%s:%d", o.selection.String(), o.value)
}

type Ladder struct {
	bets []*Bet
}

func NewLadder() *Ladder {
	return &Ladder{bets: make([]*Bet, 0)}
}

func (l *Ladder) PlaceBet(b *Bet) {
	l.bets = append(l.bets, b)
}

// Outcomes returns the list of outcomes giveen placed bets
func (l *Ladder) Outcomes() []*Outcome {
	var result []*Outcome
	for i, b := range l.bets {
		if i == 0 {
			result = b.Outcomes()
		} else {
			newResult := make([]*Outcome, 0)
			for _, o := range result {
				for _, bo := range b.Outcomes() {
					newResult = append(newResult, &Outcome{
						selection: o.selection.Intersection(bo.selection),
						value:     o.value + bo.value,
					})
				}
			}
			result = newResult
		}
	}
	// cleanup
	oldResult := result
	result = make([]*Outcome, 0)
	for _, o := range oldResult {
		if !o.selection.Empty() {
			result = append(result, o)
		}
	}
	return result
}

func (l *Ladder) Exposure() int64 {
	outcomes := l.Outcomes()
	if len(outcomes) == 0 {
		return 0
	}
	min := outcomes[0].value
	for _, o := range outcomes[1:] {
		if min > o.value {
			min = o.value
		}
	}
	return -min
}

func (l *Ladder) MaxWin() int64 {
	outcomes := l.Outcomes()
	if len(outcomes) == 0 {
		return 0
	}
	max := outcomes[0].value
	for _, o := range outcomes[1:] {
		if max < o.value {
			max = o.value
		}
	}
	return max
}

func (l *Ladder) Settle(runs int64) int64 {
	for _, o := range l.Outcomes() {
		if o.selection.Contains(runs) {
			return o.value
		}
	}
	return 0
}
