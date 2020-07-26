package utils

import (
	"fmt"
	"github.com/json-iterator/go"
)

func test01(){
	fmt.Println("test01")
}

//----------
func test02(){
	fmt.Println("test02")
}

//----------------------------------------
type User struct {
	Id      int    `json:"id,string"`
	Name    string `json:"username"`
	Age     int    `json:"age,omitempty"` //各种类型的零值，当字段为零值的时候，不会被编码到json中
	Address string `json:"-"`
}


func  test03()  {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	u := User{
		Id:      12,
		Name:    "wendell",
		Age:     1,
		Address: "成都高新区",
	}
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	//data:=[]byte(`{"id":"12","username":"wendell","age":1,"Address":"北京"}`)
	u2 := &User{}
	err = json.Unmarshal(data, u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", u2)


}