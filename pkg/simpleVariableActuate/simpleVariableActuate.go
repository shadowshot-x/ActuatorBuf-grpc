package pkg

import "fmt"

type SimpleVariable struct {
	var1 int64
	var2 string
}

func (v *SimpleVariable) SetVar1(c int64) {
	v.var1 = c
}

func ConvertToSimpleVariable(v1 int64, v2 string) *SimpleVariable {
	return &SimpleVariable{
		var1: v1,
		var2: v2,
	}
}

// StateCheck is a check that compares 2 Simple Variables and returns message about the difference
func (v *SimpleVariable) StateCheck(currentSimpleVariable SimpleVariable) (bool, string, error) {
	actuateMessage := ""
	flag := true
	if v.var1 != currentSimpleVariable.var1 {
		actuateMessage = actuateMessage + fmt.Sprintf("var1 is out of sync. Expected : %d ::: Current : %d ;;;; ", v.var1, currentSimpleVariable.var1)
		flag = false
	}
	if v.var2 != currentSimpleVariable.var2 {
		actuateMessage = actuateMessage + fmt.Sprintf("var2 is out of sync. Expected : %s ::: Current : %s ;;;; ", v.var2, currentSimpleVariable.var2)
		flag = false
	}
	return flag, actuateMessage, nil
}
