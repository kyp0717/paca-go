package market 



type Model struct {
  StockList []string
  data PacaStream
  algo Algo
  Trade string
}

func NewModel() Algo {
  t := Algo{}
  return t
}

func (a Model) AddAlgo(m Metric) {
  // a.Trend = m
}


func (a Model) Run() {

}
