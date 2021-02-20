package main

import "fmt"

// Go语言设计模式之函数式选项模式
// https://www.liwenzhou.com/posts/Go/functional_options_pattern/

type Option struct {
	A string
	B string
	C int
}

type OptionFunc func(*Option)

func NewOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

func WithA(a string) OptionFunc {
	return func(option *Option) {
		option.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(option *Option) {
		option.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(option *Option) {
		option.C = c
	}
}

var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)

func NewOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}

	return
}

func main() {
	x := NewOption("luochao", "dashage", 10)
	fmt.Println(x)

	x = NewOption2()
	fmt.Println(x)

	x = NewOption2(
		WithA("xiaoxue"),
		WithC(24),
	)
	fmt.Println(x)
}
