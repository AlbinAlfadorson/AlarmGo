module AlarmGo

go 1.22.5

Связывание библиотек с директорией, но при запуске на локальном компьютере придётся создавать этот файл по новой,
с помощью команды go mod init из репозитория AlarmGo в терминале!
Важно чтоб в идентичном приложение он выглядел так
require (
	fyne.io/fyne/v2 v2.5.1 // indirect
	github.com/faiface/beep v1.1.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
)
