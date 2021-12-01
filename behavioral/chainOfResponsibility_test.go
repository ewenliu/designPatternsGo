package behavioral

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	day := 8
	lm := &LineManager{}
	sm := &SeniorManger{}
	chief := &Chief{}
	lm.setSuperiorManager(sm)
	sm.setSuperiorManager(chief)
	lm.handleRequest(day)
}