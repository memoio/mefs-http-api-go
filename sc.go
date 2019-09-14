package shell

import (
	"context"
	"fmt"
)

// TestLocalinfo 获取节点信息的操作
func (s *Shell) TestLocalinfo() {
	var sl StringList
	rb := s.Request("test/localinfo")
	rb.Exec(context.Background(), &sl)
	fmt.Println(sl.String())
}

// ResultSummary keeper计算时空值命令，用于测试，返回计算好的时空值
func (s *Shell) ResultSummary() int {
	var il IntList
	rb := s.Request("test/resultsummary")
	rb.Exec(context.Background(), &il)
	if len(il.ChildLists) < 1 {
		fmt.Println("计算失败")
		return 0
	}
	return il.ChildLists[0]
}
