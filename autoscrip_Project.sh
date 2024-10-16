#!/bin/bash

# Создаём структуру директорий
mkdir -p main
mkdir -p src/{GUI, alarm, sound}
mkdir -p pkg
mkdir -p desktop

# Создаём новый файл
touch main.go
touch go mod

# Инициализация go mod
go mod init "AlarmGo"

# Добавление библиотек
go get fyne.io/fyne/v2          # Для создания GUI
go get github.com/faiface/beep  # Для звукового сигнала
go get github.com/robfig/cron   # Для планирования задач

# Создаём базовый файл main.go
cat << EOF > main.go
package main // Объявление этого пакета основным!
import ( 
	"flag"
	"fmt"
	"src/GUI"
	"src/Alarm"
	"src/desktop"
)


func main() {
	
	# Объявление командной флага командной строки для создания ярлыка
	
	createShortcut := flag.Bool("create!!!", false, "Create a desktop shortcut")
	
	# Разбор флагов командной строки 
	flag.Parse()
	
	# Если указан флаг создания ярлыка
	if * createShortcut {
		err := desktop.CreateShortcut()
		if err != nil {
			fmt.Println("ERROR! IS ABSOLUTELY FOULT! FATAL ERROR!!!", err )
		} else {
		fmt.Println("ShortcutI was be created! GREATE! Seccsessfuly!")
		}
		return
	}
	
	alarmManager := alarm.NewAlarmManager()
	gui.RunGUI(alarmManager)
	
	
}
