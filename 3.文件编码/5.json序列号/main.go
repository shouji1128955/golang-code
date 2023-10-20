package main

//磁盘存储，网络传输
import (
	"encoding/json"
	"fmt"
	"os"
)

func example() {
	//对于非结构化数据，在golang中先要进行反序列化处理，根据数据格式规则，定义合适的数据类型
	//比如下面的就是类似map,通过定义map数据类型，然后进行写入数据
	s := `{"name": "zhangsan", "age": "30", "gender": "male"}`
	var data map[string]string
	json.Unmarshal([]byte(s), &data)
	fmt.Println(data["name"])
}

func main() {

	//1)序列化
	var stu01 = map[string]string{"name": "zhangsan", "age": "30", "gender": "male"}
	var stu02 = map[string]string{"name": "wangwu", "age": "22", "gender": "male"}
	var stu03 = map[string]string{"name": "zhaoliu", "age": "23", "gender": "male"}
	//切片管理
	var stus = []map[string]string{stu01, stu02, stu03}
	ret, _ := json.Marshal(stus)
	//fmt.Println(string(ret))

	//把json写入到文件中,输入的就是字节
	os.WriteFile("status.json", ret, 666)

	//2) 反序列化
	//根据json格式看在golang中哪种数据类型能够对应，然后进行序列化
	//从文件读取内容
	filel, _ := os.ReadFile("status.json")
	var data []map[string]string //定义了一个切片的数据类型，值为map类型
	json.Unmarshal(filel, &data) //然后写入数据  注意此处的&data就是通过地址来写数据
	fmt.Println(data)
	fmt.Println(data[1]["name"])
	example()
}
