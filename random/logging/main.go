package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	// допишите код
	mylog := log.New(&buf, `serv `, 0)
	mylog.Print("yes")

	// 1) создайте переменную типа *log.Logger
	// 2) запишите в неё нужные строки

	// ...

	fmt.Print(&buf)
	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
