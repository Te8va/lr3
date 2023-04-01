package main

import (
	"fmt"
	"net"
)

const (
	host = "localhost:8080"
	Type = "tcp"
)

func main() {
	// установить соединение с сервером на порту 8080
	conn, err := net.Dial(Type, host)
	if err != nil {
		fmt.Println("Server startup error:", err)
		return
	}
	defer conn.Close()
	// Бесконечный цикл, который просит пользователя угадать число
	fmt.Print("Guess a number between 0 and 99: ")
	for {
		var guess int
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}

		if guess <= 0 || guess >= 99 {
			fmt.Println("Incorrect data.Try again.")
			continue
		}

		// Отправляем угаданное число на сервер
		_, err = conn.Write([]byte(fmt.Sprintf("%d", guess)))
		if err != nil {
			fmt.Println("Error when sending data to the server", err)
			continue
		}

		// Читаем ответ от сервера и выводим его на экран
		buffer := make([]byte, 1024)
		numberByte, err := conn.Read(buffer)
		if numberByte == 0 || err != nil {
			fmt.Println("Error reading the response from the server:", err)
			continue
		}
		fmt.Println(string(buffer[:numberByte]))

		// Закрываем соединение, если число было угадано
		if string(buffer[:numberByte]) == "Congratulations, you guessed it!" {
			conn.Close()
			break
		}
	}

	fmt.Println("Game over")
}
