package market 
	
import (

 )



type Algo struct {
  PacaStream
  Analysis
  // Feature Metric
  // Feature Map(string)[Decision]
  // Features map[string]Decision
  // Quotes []string
  // TradeChan chan<- Decision
  Trade
}

func NewAlgo(a Analysis, s PacaStream, tr Trade) Algo {
  al:=Algo{}
  
  // a := Algo{Analysis: m,
  //           Trade: tr,
  //           TradeChan: make(chan Decision)}
  return al
}

func (a Algo) InitPosition() {
}

func (a Algo) Sell() {
}




