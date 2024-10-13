package sound // Объявление пакета sound

import (
    "os"    // Пакет для работы с операционной системой
    "time"  // Пакет для работы со временем

    "github.com/faiface/beep"          // Пакет для работы со звуком
    "github.com/faiface/beep/mp3"      // Пакет для декодирования MP3
    "github.com/faiface/beep/speaker"  // Пакет для воспроизведения звука
)

// Функция для воспроизведения звука будильника
func PlayAlarm() {
    // Открытие файла со звуком будильника
    f, err := os.Open("src/assets/Signal.mp3")
    if err != nil {
        // В настоящих приложениях здесь должно быть событие обработки ошибок, может, когда - нибудь оно здесь и будет!
        return
    }
    // Отложенное закрытие файла
    defer f.Close()

    // Декодирование MP3 файла
    streamer, format, err := mp3.Decode(f)
    if err != nil {
        // Обработка ошибки (в реальном приложении здесь должен быть код обработки ошибки)
        return
    }
    // Отложенное закрытие стримера
    defer streamer.Close()

    // Инициализация динамика с параметрами аудио
    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)) // функция, почему то не всегда отробатывает, хотя строго повторяет документацию!
    // Воспроизведение звука
    speaker.Play(streamer)

    // Ожидание 30 секунд (длительность воспроизведения)
    time.Sleep(30 * time.Second)
}
