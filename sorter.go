package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	//пользователь вводит путь к файлам
	fmt.Print("Введите путь к папке:")
	var filesPath string
	fmt.Scanln(&filesPath)

	//проверяем файлы
	files, err := ioutil.ReadDir(filesPath)
	if err != nil {
		fmt.Println("Ошибка при чтении файлов: ", err)
		os.Exit(1)
	}

	//пробигаемся по файлам
	for _, file := range files {
		filename := file.Name()

		//разделяем строки и находим префикс
		prefix := strings.Split(filename, ".")[0]

		//создание папки если она не создана
		folderPath := filepath.Join(filesPath, prefix)
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			err := os.Mkdir(folderPath, 0755)
			if err != nil {
				fmt.Println("Ошибка при создании папки: ", err)
				os.Exit(1)
			}
		}

		//перемещение файлов в папку
		oldPath := filepath.Join(filesPath, filename)
		newPath := filepath.Join(folderPath, filename)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("Ошибка при перемещении файлов!", err)
			os.Exit(1)
		}
	}
}
