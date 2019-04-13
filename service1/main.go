package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"

	"github.com/harveywangdao/micro/service1/test"
)

func write() {
	phone1 := &test.Phone{
		Type:   test.PhoneType_HOME,
		Number: "111111111",
	}

	phone2 := &test.Phone{
		Type:   test.PhoneType_WORK,
		Number: "222222222",
	}

	phone3 := &test.Phone{
		Type:   test.PhoneType_HOME,
		Number: "333333333",
	}

	phone4 := &test.Phone{
		Type:   test.PhoneType_WORK,
		Number: "444444444",
	}

	p1 := &test.Person{
		Id:     1,
		Name:   "小张",
		Phones: []*test.Phone{phone1, phone2},
	}

	p2 := &test.Person{
		Id:     2,
		Name:   "小王",
		Phones: []*test.Phone{phone3, phone4},
	}

	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	data, _ := proto.Marshal(book)

	ioutil.WriteFile("./test.txt", data, os.ModePerm)
}

func read() {
	data, _ := ioutil.ReadFile("./test.txt")
	book := &test.ContactBook{}

	proto.Unmarshal(data, book)

	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)

		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	write()
	read()
}
