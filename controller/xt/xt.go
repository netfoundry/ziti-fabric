package xt

import (
	"time"
)

type Registry interface {
	RegisterFactory(factory Factory)
	GetStrategy(name string) (Strategy, error)
}

type Factory interface {
	GetStrategyName() string
	GetStrategyAliases() []string
	NewStrategy() Strategy
}

type Terminator interface {
	GetId() string
	GetCost() uint16
	GetServiceId() string
	GetRouterId() string
	GetBinding() string
	GetAddress() string
	GetPeerData() map[uint32][]byte
	GetCreatedAt() time.Time
}

type WeightedTerminator interface {
	Terminator
	GetRouteWeight() uint32
}

type StrategyChangeEvent interface {
	GetServiceId() string
	GetCurrent() []Terminator
	GetAdded() []Terminator
	GetChanged() []Terminator
	GetRemoved() []Terminator
}

type Strategy interface {
	Select(terminators []WeightedTerminator, totalWeight uint32) (Terminator, error)
	HandleTerminatorChange(event StrategyChangeEvent) error
	NotifyEvent(event TerminatorEvent)
}

type Precedence interface {
	getCostIntervalStart() uint16
}

type TerminatorEvent interface {
	GetTerminator() Terminator
	Accept(visitor EventVisitor)
}

type EventVisitor interface {
	VisitDialFailed(event TerminatorEvent)
	VisitDialSucceeded(event TerminatorEvent)
	VisitSessionEnded(event TerminatorEvent)
}

type Costs interface {
	ClearCost(terminatorId string)
	GetCost(terminatorId string) uint16
	GetPrecedence(terminatorId string) Precedence
	SetPrecedence(terminatorId string, precedence Precedence)
	SetPrecedenceCost(terminatorId string, weight uint8)
	GetPrecedenceCost(terminatorId string) uint8
}
