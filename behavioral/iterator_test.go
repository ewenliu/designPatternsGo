package behavioral

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T){
	nr := &NameRepo{}
	for iterator := nr.GetIterator(); iterator.HasNext();{
		fmt.Println(iterator.Next())
	}

}
