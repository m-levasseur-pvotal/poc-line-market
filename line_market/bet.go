package line_market

import (
	"github.com/shopspring/decimal"
)

type Bet struct {
	ID string
	// Size the size of the bet
	Size int64
	// Odds expressed in percentage
	Odds int64
	// Runs the number of the runs you bet on
	Runs int64
	// RunsAndOver true if win when the result is equal or greater than runs value
	// false if win when the result is equal or lower than runs value
	RunsAndOver bool
}

func (b *Bet) PL() int64 {
	return decimal.NewFromInt(b.Odds).Div(decimal.NewFromInt(100)).Mul(decimal.NewFromInt(b.Size)).RoundBank(0).IntPart()
}

func (b *Bet) Outcomes() []*Outcome {
	if b.RunsAndOver {
		return []*Outcome{
			{
				selection: NewInterval(b.Runs, MaxInterval),
				value:     b.PL(),
			},
			{
				selection: NewInterval(0, b.Runs),
				value:     -b.Size,
			},
		}
	}
	return []*Outcome{
		{
			selection: NewInterval(b.Runs+1, MaxInterval),
			value:     -b.Size,
		},
		{
			selection: NewInterval(0, b.Runs+1),
			value:     b.PL(),
		},
	}
}
