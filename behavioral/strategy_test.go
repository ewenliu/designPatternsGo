package behavioral

import (
	"fmt"
	"testing"
)

func TestCashStrategy(t *testing.T){
	money := 199.00
	strategy := MakeCashContext("八折")
	fmt.Println(strategy.GetMoney(money))
	strategy = MakeCashContext("满100返20")
	fmt.Println(strategy.GetMoney(money))
	strategy = MakeCashContext("")
	fmt.Println(strategy.GetMoney(money))
}