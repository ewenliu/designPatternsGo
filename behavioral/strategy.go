package behavioral

type CashStrategy interface {
	AcceptMoney(money float64) float64
}

type CashNormal struct {
}

func (*CashNormal) AcceptMoney(money float64) float64 {
	return money
}

func NewCashNormal() *CashNormal {
	return &CashNormal{}
}

type CashDiscount struct {
	CashDiscount float64
}

func (c *CashDiscount) AcceptMoney(money float64) float64 {
	return c.CashDiscount * money
}

func NewCashDiscount(discount float64) *CashDiscount {
	return &CashDiscount{CashDiscount: discount}
}

type CashRebate struct {
	MoneyCondition float64
	Rebate         float64
}

func (c *CashRebate) AcceptMoney(money float64) float64 {
	if money >= c.MoneyCondition {
		rebateTimes := int(money / c.MoneyCondition)
		return money - float64(rebateTimes)*c.Rebate
	}
	return money
}

func NewCashRebate(condition, rebate float64) *CashRebate {
	return &CashRebate{MoneyCondition: condition, Rebate: rebate}
}

type CashContext struct {
	Strategy CashStrategy
}

func (c *CashContext) GetMoney(money float64) float64 {
	return c.Strategy.AcceptMoney(money)
}

func MakeCashContext(cashType string) *CashContext {
	c := new(CashContext)
	switch cashType {
	case "八折":
		c.Strategy = NewCashDiscount(0.8)
	case "满100返20":
		c.Strategy = NewCashRebate(100.00, 20.00)
	default:
		c.Strategy = NewCashNormal()
	}
	return c
}
