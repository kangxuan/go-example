package main

import (
	"fmt"
	"math"
)

func main() {
	l := int64(15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s\n", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d\n", l, i)
	}

	fmt.Println("-------------------------")

	l = int64(math.MaxInt32 + 15000)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s\n", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d\n", l, i)
	}
}

// ConvertInt64ToInt 将int64转换成int
func ConvertInt64ToInt(l int64) int {
	if math.MinInt32 <= l && math.MaxInt32 >= l {
		return int(l)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", l))
}

// IntFromInt64 将int类型转换成Int64
func IntFromInt64(l int64) (i int, err error) {
	// 捕获panic
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	i = ConvertInt64ToInt(l)
	return i, nil
}

//15000 converted to an int32 is 15000
//-------------------------
//The conversion of 2147498647 to an int32 resulted in an error: 2147498647 is out of the int32 range
