package line_market

import (
	"fmt"
	"testing"
)

func TestLadder_Outcomes(t *testing.T) {
	l := NewLadder()
	l.PlaceBet(&Bet{
		ID:          "a",
		Size:        200,
		Odds:        100,
		Runs:        9,
		RunsAndOver: false,
	})
	l.PlaceBet(&Bet{
		ID:          "b",
		Size:        100,
		Odds:        100,
		Runs:        11,
		RunsAndOver: true,
	})
	l.PlaceBet(&Bet{
		ID:          "c",
		Size:        100,
		Odds:        100,
		Runs:        15,
		RunsAndOver: true,
	})
	fmt.Printf("Outcomes : %v\n", l.Outcomes())
	fmt.Printf("Exposure: %d\n", l.Exposure())
	fmt.Printf("Max Wins: %d\n", l.MaxWin())

	l1 := NewLadder()
	l1.PlaceBet(&Bet{
		ID:          "a",
		Size:        200,
		Odds:        100,
		Runs:        25,
		RunsAndOver: false,
	})
	l1.PlaceBet(&Bet{
		ID:          "b",
		Size:        100,
		Odds:        100,
		Runs:        11,
		RunsAndOver: true,
	})
	l1.PlaceBet(&Bet{
		ID:          "c",
		Size:        100,
		Odds:        100,
		Runs:        15,
		RunsAndOver: true,
	})
	fmt.Printf("Outcomes : %v\n", l1.Outcomes())
	fmt.Printf("Exposure: %d\n", l1.Exposure())
	fmt.Printf("Max Wins: %d\n", l1.MaxWin())

}
