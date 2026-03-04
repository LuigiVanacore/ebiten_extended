package example_statemachine

import (
	"fmt"
)

// Instance singleton patterns are removed to avoid shared state across different Miners
// We can instantiate states directly or store them in the Miner struct.

type EnterMinerAndDigForNuggetState struct{}

func (e *EnterMinerAndDigForNuggetState) Enter(miner *Miner) {
	if miner.GetLocation() != GOLDMINE {
		fmt.Printf("\n%s: Walkin' to the goldmine", "Miner")
		miner.ChangeLocation(GOLDMINE)
	}
}

func (e *EnterMinerAndDigForNuggetState) Execute(miner *Miner) {
	miner.AddToGoldCarried(1)
	miner.IncreaseFatigue()
	fmt.Printf("\n%s: Pickin' up a nugget", "Miner")

	if miner.IsPocketsFull() {
		miner.GetStateMachine().ChangeState(&VisitBankAndDepositGoldState{})
	}

	if miner.IsThirsty() {
		miner.GetStateMachine().ChangeState(&QuenchThirstState{})
	}
}

func (e *EnterMinerAndDigForNuggetState) Exit(miner *Miner) {
	fmt.Printf("\n%s: Ah'm leavin' the goldmine with mah pockets full o' sweet gold", "Miner")
}

type VisitBankAndDepositGoldState struct{}

func (e *VisitBankAndDepositGoldState) Enter(miner *Miner) {
	if miner.GetLocation() != BANK {
		fmt.Printf("\n%s: Goin' to the bank. Yes siree", "Miner")
		miner.ChangeLocation(BANK)
	}
}

func (e *VisitBankAndDepositGoldState) Execute(miner *Miner) {
	miner.SetMoneyInTheBank(miner.GetMoneyInTheBank() + miner.GetGoldCarried())
	miner.SetGoldCarried(0)
	fmt.Printf("\n%s: Depositing gold. Total savings now: %d", "Miner", miner.GetMoneyInTheBank())

	if miner.GetMoneyInTheBank() >= ComfortLevel {
		fmt.Printf("\n%s: WooHoo! Rich enough for now. Back home to mah li'lle lady", "Miner")
		miner.GetStateMachine().ChangeState(&GoHomeAndSleepTilRestedState{})
	} else {
		miner.GetStateMachine().ChangeState(&EnterMinerAndDigForNuggetState{})
	}
}

func (e *VisitBankAndDepositGoldState) Exit(miner *Miner) {
	fmt.Printf("\n%s: Leavin' the bank", "Miner")
}

type GoHomeAndSleepTilRestedState struct{}

func (e *GoHomeAndSleepTilRestedState) Enter(miner *Miner) {
	if miner.GetLocation() != HOUSE {
		fmt.Printf("\n%s: Walkin' home", "Miner")
		miner.ChangeLocation(HOUSE)
	}
}

func (e *GoHomeAndSleepTilRestedState) Execute(miner *Miner) {
	if !miner.IsFatigued() {
		fmt.Printf("\n%s: What a God darn fantastic nap! Time to find more gold", "Miner")
		miner.GetStateMachine().ChangeState(&EnterMinerAndDigForNuggetState{})
	} else {
		miner.DecreaseFatigue()
		fmt.Printf("\n%s: ZZZZ... ", "Miner")
	}
}

func (e *GoHomeAndSleepTilRestedState) Exit(miner *Miner) {
	fmt.Printf("\n%s: Leaving the house", "Miner")
}

type QuenchThirstState struct{}

func (e *QuenchThirstState) Enter(miner *Miner) {
	if miner.GetLocation() != SALOON {
		miner.ChangeLocation(SALOON)
		fmt.Printf("\n%s: Boy, ah sure is thusty! Walking to the saloon", "Miner")
	}
}

func (e *QuenchThirstState) Execute(miner *Miner) {
	if miner.IsThirsty() {
		miner.BuyAndDrinkAWhiskey()
		fmt.Printf("\n%s: That's mighty fine sippin liquer", "Miner")
		miner.GetStateMachine().ChangeState(&EnterMinerAndDigForNuggetState{})
	} else {
		fmt.Printf("\nERROR!\nERROR!\nERROR!")
	}
}

func (e *QuenchThirstState) Exit(miner *Miner) {
	fmt.Printf("\n%s: Leaving the saloon, feelin' good", "Miner")
}
