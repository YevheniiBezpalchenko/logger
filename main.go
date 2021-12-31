package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Log struct {
	lvl   [4]bool
	msg   string
	time  string
	write *os.File
}

func (l *Log) Start(file *os.File, levels [4]bool) {
	currentTime := time.Now()
	l.time = currentTime.Format("2006-01-02 15:04:05")
	l.lvl = levels
	l.msg = ""
	l.write = file

}
func (l Log) Info(message string, a ...interface{}) {
	_, _, count, _ := runtime.Caller(1)
	cline := line(count - 2)
	l.msg = "Info | " + l.time + " | " + cline + message
	fmt.Fprintf(l.write, l.msg)
	for _, n := range a {
		fmt.Fprintf(l.write, " %v ", n)
	}
	fmt.Fprintf(l.write, "\n")
}
func (l Log) Error(message string, a ...interface{}) {
	_, _, count, _ := runtime.Caller(1)
	cline := line(count - 3)
	l.msg = "Error | " + l.time + " | " + cline + message
	fmt.Fprintf(l.write, l.msg)
	for _, n := range a {
		fmt.Fprintf(l.write, " %v ", n)
	}
	fmt.Fprintf(l.write, "\n")
}
func (l Log) Warning(message string, a ...interface{}) {
	_, _, count, _ := runtime.Caller(1)
	cline := line(count - 2)
	l.msg = "Warning | " + l.time + " | " + cline + message
	fmt.Fprintf(l.write, l.msg)
	for _, n := range a {
		fmt.Fprintf(l.write, " %v ", n)
	}
	fmt.Fprintf(l.write, "\n")
}

func (l Log) Debug(message string, a ...interface{}) {
	l.msg = "Debug | " + l.time + " | " + message
	fmt.Fprintf(l.write, l.msg)
	for _, n := range a {
		fmt.Fprintf(l.write, " %v ", n)
	}
	fmt.Fprintf(l.write, "\n")
}

func main() {
	log := Log{}

	file, err := os.OpenFile("logfile", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Start(file, [4]bool{true, true, true, true})

	defer file.Close()
	file2, err := os.OpenFile("logfil", os.O_RDONLY, 0755)
	if err != nil {
		errorMsg := "Неудалось открыть файл"
		log.Error(errorMsg, err)
	}
	fmt.Println("заказ оформлен")
	infomsg := "заказ оформлен"
	log.Info(infomsg)

	fmt.Println(file2, "Контекс ", "не разделенный на разные страки")
	warningmsg := "не правильное правописание слова строки/страки"
	log.Warning(warningmsg)

	debugmsg := "пример debug log"
	log.Debug(debugmsg)
}
func line(l int) string {
	var linestr string = "Строка " + strconv.Itoa(l) + ":"
	return linestr
}
