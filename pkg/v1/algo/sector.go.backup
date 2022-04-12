package algo

// type SectorTrend struct {
//   // in <-chan Feature
// }

type SectorTrend map[string]Trend

type Trade struct{}

// Compute will take stock trend feature and return trade decision
func (s *SectorTrend) Compute(in <-chan Feature) (out chan<- Feature) {
  o := make(chan<- Feature)
  go func() { 
  for {
    m := <-in
    stock := m.GetSymbol()
    trends[stock] = m.GetTrend()
    count :=0
    ups:=0
    downs:=0
    sames:=0
    for _, trend := range trends {
      count++
      switch trend {
        case Up: ups++
        case Down: downs++
        case Same: sames++
      }
    }
    // var upPct float64
    upPct := float64(ups/count)
    switch {
      case (upPct > 0.75): o <- Hold
      case (upPct < 0.75): o <- Sell
    }
  }
  }()
  return o
}


