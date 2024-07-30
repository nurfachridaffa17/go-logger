package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Membuat instance logger baru
	log := logrus.New()

	// Mengatur format output log
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Menentukan level logging
	log.SetLevel(logrus.InfoLevel)

	// Mengatur log ke file dengan rotasi harian
	log.SetOutput(&lumberjack.Logger{
		Filename:   "log/app.log", // Nama file log
		MaxSize:    10,            // Ukuran maksimum file log (dalam MB)
		MaxBackups: 30,            // Jumlah backup file yang disimpan
		MaxAge:     7,             // Usia maksimum file backup (dalam hari)
		Compress:   true,          // Mengompresi file backup
	})

	// Logging dengan berbagai level
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	// Logging dengan fields
	log.WithFields(logrus.Fields{
		"user": "JohnDoe",
		"id":   1234,
	}).Info("User logged in")
}
