package xfloat64

import (
	"github.com/shopspring/decimal"
)

//float64加法
func Add(ff ...float64) float64 {
	df := decimal.NewFromFloat(ff[0])
	for i := 1; i < len(ff); i++ {
		df = df.Add(decimal.NewFromFloat(ff[i]))
	}
	return decimalToFloat64(df)
}

//string相加
func FromStringAdd(ss ...string) float64 {
	df := decimal.RequireFromString(ss[0])
	for i := 1; i < len(ss); i++ {
		df = df.Add(decimal.RequireFromString(ss[i]))
	}
	return decimalToFloat64(df)
}

//
func FromStringAddFloat(s string, f float64) float64 {
	d1 := decimal.RequireFromString(s)
	d2 := decimal.NewFromFloat(f)
	return decimalToFloat64(d1.Add(d2))
}

//float64减法
func Sub(ff ...float64) float64 {
	df := decimal.NewFromFloat(ff[0])
	for i := 1; i < len(ff); i++ {
		df = df.Sub(decimal.NewFromFloat(ff[i]))
	}
	return decimalToFloat64(df)
}

//string相减
func FromStringSub(ss ...string) float64 {
	df := decimal.RequireFromString(ss[0])
	for i := 1; i < len(ss); i++ {
		df = df.Sub(decimal.RequireFromString(ss[i]))
	}
	return decimalToFloat64(df)
}

//float64乘法
func Mul(ff ...float64) float64 {
	df := decimal.NewFromFloat(ff[0])
	for i := 1; i < len(ff); i++ {
		df = df.Mul(decimal.NewFromFloat(ff[i]))
	}
	return decimalToFloat64(df)
}

//string相乘
func FromStringMul(ss ...string) float64 {
	df := decimal.RequireFromString(ss[0])
	for i := 1; i < len(ss); i++ {
		df = df.Mul(decimal.RequireFromString(ss[i]))
	}
	return decimalToFloat64(df)
}

//float64除法
func Div(ff ...float64) float64 {
	df := decimal.NewFromFloat(ff[0])

	for i := 1; i < len(ff); i++ {
		df = df.Div(decimal.NewFromFloat(ff[i]))
	}
	return decimalToFloat64(df)
}

//string相除
func FromStringDiv(ss ...string) float64 {
	df := decimal.RequireFromString(ss[0])
	for i := 1; i < len(ss); i++ {
		df = df.Div(decimal.RequireFromString(ss[i]))
	}
	return decimalToFloat64(df)
}

//字符串浮点数比较
func FromStringCmp(s1, s2 string) int {
	d1 := decimal.RequireFromString(s1)
	d2 := decimal.RequireFromString(s2)
	return d1.Cmp(d2)
}

//浮点数比较
func FromFloatCmp(f1, f2 float64) int {
	d1 := decimal.NewFromFloat(f1)
	d2 := decimal.NewFromFloat(f2)
	return d1.Cmp(d2)
}

//从字符串浮点数比较浮点数
func FromStringCmpFloat(s string, f float64) int {
	d1 := decimal.RequireFromString(s)
	d2 := decimal.NewFromFloat(f)
	return d1.Cmp(d2)
}

//从字符串浮点数比较浮点数
func FromFloatCmpString(f float64, s string) int {
	d1 := decimal.NewFromFloat(f)
	d2 := decimal.RequireFromString(s)

	return d1.Cmp(d2)
}

//通过decimal转换成float64
func StrToFloat64(str string) float64 {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return 0
	}
	return decimalToFloat64(d)
}

//转成float64
func decimalToFloat64(d decimal.Decimal) float64 {
	f, _ := d.Float64()
	return f
}
