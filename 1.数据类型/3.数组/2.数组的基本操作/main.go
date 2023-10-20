package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var couse = [3]string{"语文", "数学", "英语"}
	fmt.Println(reflect.TypeOf(couse))

	//(1) arr[索引]
	fmt.Println(couse[2])
	var x = couse[2]
	fmt.Println(x)

	//(2) 切片,数组[开始索引:结束索引]
	var s1 = couse[0:2]
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(s1)) //[]string
	//通过数组进行索引取值后就会生成切片，生成新的数据类型。不存在长度就是切片，会在长度就是数组

	//(3)数组的循环
	//方式1
	for i := 0; i < len(couse); i++ {
		fmt.Println(couse[i])
	}

	//方式2
	for i, v := range couse {
		//i就是遍历的每一个索引，v就是队以索引的值
		fmt.Println(i, v)
	}

	//添加判断，只想打印语文课程
	fmt.Println("===============")
	for i, v := range couse {
		if strings.HasPrefix(v, "语") {
			fmt.Println(i, v)
		}
	}

	//(4)细节问题
	//通过修改v是否能够修改数组值
	for i, v := range couse {
		//i就是遍历的每一个索引，v就是队以索引的值
		fmt.Println(i, v)
		if strings.HasPrefix(v, "语") {
			v = "物理"
		}
	}
	fmt.Println(couse)
	//事实证明,没法修改，这里的细节就是发生了值拷贝,所以直接修改v是没法修改数组的

}
