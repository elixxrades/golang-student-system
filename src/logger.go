package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TwiN/go-color"
)

func NewLogger(name string) *Logger {
	w := &Logger{name}
	w.Info("Logger loaded. Name : "+name, nil)

	return w
}

func (l *Logger) Info(str string, data any) {
	l.log("INFO", str, data)
}

func (l *Logger) Warn(str string, data any) {
	l.log("WARN", str, data)
}

func (l *Logger) Error(str string, data any) {
	l.log("ERROR", str, data)
}

func (l *Logger) Debug(str string, data any) {
	l.log("DEBUG", str, data)
}

func (l *Logger) log(ty string, str string, data any) {
	head := ""
	switch ty {
	case "INFO":
		head = color.InGreen(ty + " / " + l.Name)
	case "ERROR":
		head = color.InRed(ty + " / " + l.Name)
	case "WARN":
		head = color.InYellow(ty + " / " + l.Name)
	case "DEBUG":
		head = color.InWhite(ty + " / " + l.Name)
	}

	if data == nil {
		wz := ""
		wz += head + " - "
		wz += color.InBold(time.Now().Format("2006.01.02 - 15:04:05") + " - ")
		wz += str + "\n"

		log.Println(wz)

		writeLog("./logs/"+ty+".log", str)
	} else {
		wz := ""
		wz += head + " - "
		wz += color.InBold(time.Now().Format("2006.01.02 - 15:04:05") + " - ")
		wz += str

		log.Println(wz)
		log.Println(data)
		out, err := json.Marshal(data)
		checkErr(err)
		writeLog("./logs/"+ty+".log", str+"\n	- "+string(out))
	}
}

func writeLog(path string, data string) {
	files, err := os.ReadDir("logs")
	if len(files) < 1 {
		os.Mkdir("logs", 0755)
	}

	if !FileExists(path) {
		CreateFile(path)
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	nl := "\n" + data
	_, err = fmt.Fprintln(file, nl)
	if err != nil {
		log.Fatal()
	}
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
