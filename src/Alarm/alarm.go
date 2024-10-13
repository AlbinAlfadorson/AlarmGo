package alarm

import (
	"AlarmGo/src/sound"
	"time"

	"github.com/robfig/cron"  // Пакет для палнирования задач
)
// Структура для управления будильниками
type AlarmManager struct {
  // Планировщик задач
	cron *cron.Cron
}

// Новый менеджер будильников
func NewAlarmManager() *AlarmManager {
	return &AlarmManager{
    // Новыйпланировщик задач
		cron: cron.New(),
	}
}

// Метод для установки будильника
func (am *AlarmManager) SetAlarm(t time.Time) {
  // Остановка текущего планировщика
	am.cron.Stop()
  // Создание нового планировщика
	am.cron = cron.New()

	am.cron.AddFunc(fmt.Sprintf("%d %d * * *", t.Minute(), t.Hour()), func() {
		sound.PlayAlarm()
	})

	am.cron.Start()
}
