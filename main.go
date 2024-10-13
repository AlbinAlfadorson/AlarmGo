package main

import (

  "AlarmGo/src/Alarm"
  "AlarmGo/src/GUI"
  "AlarmGo/src/desktop"

  "fmt"
  "flag"
)

func main() {

// Парсинг аргументов командной строки
// Объявляем флаг командной строки для создания ярлыка
createShortcut := flag.Bool("create-shortcut", false, "Create a desktop shortcut")
// Разбор флагов командной строки
flag.Parse()

  // Если у нас указан флаг создания ярлыка
	if *createShortcut {
    // Вызов функции создания ярлыка
		err := desktop.CreateShortcut()
		if err != nil {
			fmt.Println("Error creating shortcut:", err)
		} else {
			fmt.Println("Shortcut created successfully")
		}
		return
	}

	// Запуск GUI

  // Создаём новый экземпляр менеджера будильника
	alarmManager := alarm.NewAlarmManager()
  // Вызываем для запуска графический интерфейс и передаём ему менеджер будильника.. примерно так
	gui.RunGUI(alarmManager)


}
