# POC line market

## Introduction

The intent of this repo is to illustrate the behaviour of a line market 

## Bet
```go
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
```

## Ladder
The ladder is the representation of a market when bewts are placed on.

Placing a bet will add a bet to the market

Outcomes will return a list of intervals of runs with the result of all placed bets

Exposure will return the exposure of that market (eg. Exposure of 400 means that maximum loss is 400) 

MaxWin will return the maximum a player can winn with placed bets.