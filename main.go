package main // 包的声明
import "unicode/utf8"

// 1. main方法必须在main包里面
// 2. go run main.go
// 3. 同一个文件夹下声明一致
// 4. 匿名引用可以加_不报错
// 5. len算的长度是字节数
// 6. rune类型 不是byte 本质是int32 一个rune四个字节 非常不常用 一年用不了几次
func main() {
	print("hello Go!")                     // 输出的语句
	println(`fdsfsaf`)                     // 反引号通常用于大段文字输入的时候
	println(utf8.RuneCountInString("我来了")) //3
}
