package main

import (
	"testing"
)
//测试add方法
func TestAdd(t *testing.T) {
	//列出要测试的参数和返回结果
	cases := []struct {
		first    int
		second   int
		excepted int
	}{
		{1, 2, 3},
		{1, 3, 4},
		{2, 2, 4},
	}

	//遍历，进行测试
	for _, c := range cases {
		result := add(c.first, c.second)
		if result != c.excepted {
			t.Fatalf("add function failed, first: %d, second:%d, execpted:%d, result:%d", c.first, c.second, c.excepted, result)
		}
		/*
		else {
			fmt.Printf("success:first: %d, second:%d, execpted:%d, result:%d \n", c.first, c.second, c.excepted, result)
		}
		 */
	}
}
