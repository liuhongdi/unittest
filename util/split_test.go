package util

import (
	"reflect"
	"testing"
)

//测试字符串切分
func TestSplit(t *testing.T) {
	// test struct
	type testcase struct {
		input string
		sep   string
		excepted  []string
	}
	// 用例: 用map存储
	cases := map[string]testcase{
		"冒号分隔":      {input: "a:b:c", sep: ":", excepted: []string{"a", "b", "c"}},
		"逗号分隔":   {input: "a:b:c", sep: ",", excepted: []string{"a:b:c"}},
		"字符串分割":    {input: "abcd", sep: "bc", excepted: []string{"a", "d"}},
		"中文分隔符位于开始":    {input: "百度员工百度不到时用google", sep: "百度", excepted: []string{"","员工", "不到时用google"}},
		"分隔符last":        {input: "abcd", sep: "d", excepted: []string{"abc", ""}},
	}
	for name, tc := range cases {
		// t.Run():执行子测试
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.excepted) {
				t.Errorf("excepted:%#v, got:%#v", tc.excepted, got)
			}
		})
	}
}
