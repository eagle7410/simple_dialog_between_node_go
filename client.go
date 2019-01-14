package main
import (
	"fmt"
	"net"
)
func main() {

	conn, err := net.Dial("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for{
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source));
			n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// получем ответ
		fmt.Print("Перевод:")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err !=nil{ break}
		fmt.Print(string(buff[0:n]))
		fmt.Println()
	}
}
