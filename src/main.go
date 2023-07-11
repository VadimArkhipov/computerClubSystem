package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]

	// Читаем файл
	data := readFile(path)
	// Проверяем данные из файла на корректность
	validate(data)

	// Создаем и инициализируем объект компьютерного клуба
	club := Club{}
	club.init(data)

	// Обрабатываем информацию из прочитанного файла
	club.processing()

	// Выводим полученные данные
	for _, line := range club.output {
		fmt.Println(line)
	}
}
