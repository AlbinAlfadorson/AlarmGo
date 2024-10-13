package gui // Объявление пакета gui

import (
	"AlarmGo/src/Alarm" // Импорт нашего пакета alarm
	"time"              // Пакет для работы со временем

	// Основной пакет Fyne для GUI
	"fyne.io/fyne/v2/app"       // Пакет Fyne для создания приложения
	"fyne.io/fyne/v2/container" // Пакет Fyne для контейнеров элементов GUI
	"fyne.io/fyne/v2/widget"    // Пакет Fyne для виджетов GUI
)

// Функция для запуска графического интерфейса
func RunGUI(alarmManager *Alarm.AlarmManager) {
	// Создание нового приложения Fyne
	a := app.New()
	// Создание нового окна с заголовком "Go Alarm"
	w := a.NewWindow("Go Alarm")

	// Создание поля ввода для времени будильника
	timeEntry := widget.NewEntry()
	// Установка подсказки в поле ввода
	timeEntry.SetPlaceHolder("HH:MM")

	// Создание кнопки "Set Alarm"
	setButton := widget.NewButton("Set Alarm", func() {
		// Парсинг введенного времени
		t, err := time.Parse("15:04", timeEntry.Text)
		if err != nil {
			timeEntry.SetPlaceHolder("Absolutely is not right!")
			return
		}
		// Установка будильника на указанное время
		alarmManager.SetAlarm(t)
	})

	// Создание вертикального контейнера с полем ввода и кнопкой
	content := container.NewVBox(
		timeEntry,
		setButton,
	)

	// Установка содержимого окна
	w.SetContent(content)
	// Отображение окна и запуск главного цикла приложения
	w.ShowAndRun()
}
