package example_statemachine

import (
	"github.com/LuigiVanacore/ebiten_extended/stateMachine"
)

type locationType uint

const (
	GOLDMINE locationType = iota
	BANK
	HOUSE
	SALOON
)

const (
	//the amount of gold a miner must have before he feels comfortable
	ComfortLevel = 5
	//the amount of nuggets a miner can carry
	MaxNuggets         = 3;
//above this value a miner is thirsty
ThirstLevel        = 5;
//above this value a miner is sleepy
TirednessThreshold = 5;
)

type Miner struct {
	location locationType
	goldCarried uint
	moneyInTheBank uint
	thirst uint
	fatigue uint
	stateMachine statemachine.StateMachine
}

func NewMiner() *Miner {
	miner := &Miner{ goldCarried: 0, moneyInTheBank: 0, thirst: 0, location: HOUSE}
	return miner
}

func (miner *Miner) GetLocation() locationType {
	return miner.location
}

func (miner *Miner) ChangeLocation(location locationType) {
	miner.location = location
}

func (miner *Miner) GetGoldCarried() uint {
	return miner.goldCarried
}

func (miner *Miner) SetGoldCarried(value uint) {
	miner.goldCarried = value
}

func (miner *Miner) AddToGoldCarried(value uint) {
	miner.goldCarried += value
	if miner.goldCarried <= 0 {
		miner.goldCarried = 0
	}
}

func (miner *Miner) IsPocketsFull() bool {
	return miner.goldCarried >= MaxNuggets

}

func (miner *Miner) IsFatigued() bool {
	if miner.fatigue >= TirednessThreshold {
		return true
	}
	return false
}

func (miner *Miner) DecreaseFatigue() {
	miner.fatigue -= 1
}

func (miner *Miner) IncreaseFatigue() {
	miner.fatigue += 1
}

func (miner *Miner) GetMoneyInTheBank() uint {
	return miner.moneyInTheBank
}

func (miner *Miner) SetMoneyInTheBank(wealth uint)  {
	miner.moneyInTheBank = wealth
}

func (miner *Miner) IsThirsty() bool {
	if miner.thirst >= ThirstLevel {
		return true
	}
	return false
}

func (miner *Miner) BuyAndDrinkAWhiskey() {
	miner.thirst = 0
	miner.moneyInTheBank -= 2
}


func (miner *Miner) Update() {
	miner.thirst += 1

	miner.stateMachine.Update()
}