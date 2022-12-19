package pkg

import "fmt"

type SimpleVariable struct {
	var1 int64
	var2 string
}

func (v *SimpleVariable) SetVar1(c int64) {
	v.var1 = c
}

// StateCheck is a check that compares 2 Simple Variables and returns message about the difference
func (v *SimpleVariable) StateCheck(currentSimpleVariable SimpleVariable) (bool, string, error) {
	actuateMessage := ""
	if v.var1 != currentSimpleVariable.var1 {
		actuateMessage = actuateMessage + fmt.Sprintf("var1 is out of sync. Expected : %d ::: Desired : %d ;;;; ", v.var1, currentSimpleVariable.var1)
	}
	if v.var2 != currentSimpleVariable.var2 {
		actuateMessage = actuateMessage + fmt.Sprintf("var2 is out of sync. Expected : %s ::: Desired : %s ;;;; ", v.var2, currentSimpleVariable.var2)
	}
	return true, actuateMessage, nil
}
