# AlarmGo

План проекта:
	1) Исполняемый файл - func main()
	2) Граффический интерфейс для пользователя - fyne.io/fyne/v2
	3) Модуль работы со временем - time, os, github.com/robfig/cron
	4) Модуль воспроизведения звука - github.com/faiface/beep
	5) Модуль для создания ярлыка на рабочем столе
	6) Скрипт для автоматизации создания структуры проекта
	7) Папка с ассетами
	
Структура проекта:
	1) Главная директория - alarm-in-Go
	        |                  |
	        Каталоги:          Файлы:
	        0) gui/gui.go       1) main.go
	        1) alarm/alarm.go   2) go mod
	        2) sound/sound.go
	        3) desktop/desktop/go
	        4) assets/signal.mp3
	        5) git(Git.)
