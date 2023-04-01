package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const (
	port = "8080"
	TYPE = "tcp"
)

func main() {

	// Слушаем входящие соединения на заданном порту
	listener, err := net.Listen(TYPE, ":"+port)
	if err != nil {
		fmt.Println("Server startup error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started")

	// Бесконечно принимаем входящие соединения и запускаем для каждого свой go-поток
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Server startup error:", err)
			conn.Close()
			continue
		}
		// Запускаем горутину для обработки запроса
		go handleConnection(conn)
	}
}

// Обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New client connected")

	// Создаем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())
	// Генерируем случайное число от 0 до 99
	value := rand.Intn(100)
	fmt.Println(value)

	for {
		// Читаем данные из соединения
		buffer := make([]byte, 1024)
		numberByte, err := conn.Read(buffer)
		if numberByte == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		// Преобразовываем входящее сообщение в число
		guess, err := strconv.Atoi(string(buffer[:numberByte]))
		if err != nil {
			fmt.Println("Conversion error ", err)
			continue
		}

		// Сравниваем угаданное число с загаданным
		if guess == value {
			conn.Write([]byte("Congratulations, you guessed it!"))
			conn.Close()
			fmt.Println("Client disconnected")
			break
		} else if guess < value {
			conn.Write([]byte("Your guess is too low. Try again:"))
		} else {
			conn.Write([]byte("Your guess is too high. Try again:"))
		}
	}
}
