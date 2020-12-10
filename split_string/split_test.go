package split_string

import (
	"reflect"
	"testing"
)

// 测试用列

// 测试组
func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	testMap := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		//{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	for _, ts := range testMap {
		got := Split(ts.input, ts.sep)
		if !reflect.DeepEqual(got, ts.want) {
			t.Errorf("excepted:%#v, got:%#v", ts.want, got)
		}
	}
}

// 子测试
func TestSplit2(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":    {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":  {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		//"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
		}
	}
}

// 测试覆盖率
// go test -cover

// 测试覆盖率相关记录输出到一个文件
//  go test -cover -coverprofile=c.out

// 测试覆盖率记录信息，生成一个html报告
// go tool cover-html=c.out

// 基准测试
//  go test -bench=Split
// go test -bench=Split -benchmem
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:dfa:afa", ":")
	}
}

// 性能比较
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
