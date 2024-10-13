package desktop // Объявление пакета desktop

import (
    "os"               // Пакет для работы с операционной системой
    "path/filepath"    // Пакет для работы с путями файлов
    "text/template"    // Пакет для работы с шаблонами
)

// Шаблон файла .desktop для создания ярлыка
const desktopEntry = `
[Desktop Entry]
Type=Application
Name=Go Alarm
Exec={{.ExecPath}}
Icon={{.IconPath}}
`

// Функция для создания ярлыка на рабочем столе
func CreateShortcut() error {
    // Получение домашней директории пользователя
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return err
    }

    // Формирование пути к директории рабочего стола
    desktopDir := filepath.Join(homeDir, "Desktop")
    // Формирование пути к файлу ярлыка
    shortcutPath := filepath.Join(desktopDir, "GoAlarm.desktop")

    // Создание шаблона из строки desktopEntry
    tmpl, err := template.New("desktop").Parse(desktopEntry)
    if err != nil {
        return err
    }

    // Создание файла ярлыка
    f, err := os.Create(shortcutPath)
    if err != nil {
        return err
    }
    // Отложенное закрытие файла
    defer f.Close()

    // Получение пути к исполняемому файлу приложения
    execPath, err := os.Executable()
    if err != nil {
        return err
    }

    // Структура с данными для шаблона
    data := struct {
        ExecPath string
        IconPath string
    }{
        ExecPath: execPath,
        IconPath: "/home/constantine/Desktop/AlarmGo/alarm.png", // Здесь нужно указать путь к иконке приложения
    }

    // Заполнение шаблона данными и запись в файл
    return tmpl.Execute(f, data)
}
