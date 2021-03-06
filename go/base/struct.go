package main

/**
1. 创建结构
		type 结构名 struct{
		field1 type1
		field2 type2
	}
2. 创建机构对象
	2.1   var t T
	2.2  t := new(T)   等同于     t *T    t = new (T)
	2.3  t := &T{field:xxx}  结构体字面量


*/

import (
	"fmt"
)

type Books struct {
	title  string
	author string
	id     int
}


type SuperBooks struct {
	title string
	book Books
}


func main() {

	var book1 Books

	book1.title = "Go struct"
	book1.author = "shenjun"

	fmt.Println(book1.title)

	//*************
	book2 := new(Books) // t := new(T)  其实创建的是结构指针，通过用test方法可以知道，而且，的的确确的修改了book 的id
	fmt.Println(book2.id)

	test(book2)
	fmt.Println(book2.id)
	//*************

	//*************
	// 字面量生成时，还可以使用变量， 比如  {title:"123",id :13, author : "3123"}
	book3 := &Books{"10", "10", 1} //这里需要注意的是 字面量创建结构，生成不是指针，除非使用&{} 方法生成
	fmt.Println(book3.id)
	//*************

	//*************
	//使用工厂方法创建
	//book5 :=  NewBook(0,"123","123")   // 返回的nil
	book5 :=  NewBook(123,"123","123")
	fmt.Println(book5.title)
	//*************



//*************
// 结构体的组合  碰到同名的参数也没关系，注意调用层级就行
sb := &SuperBooks{title:"sb",book:Books{"title","author",12}}

fmt.Println(sb.title)
fmt.Println(sb.book.title)
//*************





}

func NewBook(id int, title string, author string) *Books {
	if id < 1 {
		return nil
	}
	return &Books{title: title, id: id, author: author}
}

func test(book *Books) {
	book.id = 123
	fmt.Println(book.id)
}
