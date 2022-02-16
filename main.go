package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	command := "powershell"
	timeUntilBreak := 15
	countdownInterval := 1
	breakCount := 1

	fmt.Println("Starting the script")

	for {
		fmt.Printf("This is your break number %v\n", breakCount)

		playBeepSounds(command)

		for i := timeUntilBreak; i > 0; i = i - countdownInterval {
			printCountdownMessage(i)

			time.Sleep(time.Duration(countdownInterval) * time.Minute)
		}
		breakCount++
	}
}

func playBeepSounds(command string) {
	exec.Command(command, makeBeepFromFrequency(300)).Run()
	exec.Command(command, makeBeepFromFrequency(400)).Run()
	exec.Command(command, makeBeepFromFrequency(600)).Run()
	exec.Command(command, makeBeepFromFrequency(800)).Run()
}

func makeBeepFromFrequency(frequency int) string {
	return fmt.Sprintf("[console]::beep(%v,50)", frequency)
}

func printCountdownMessage(minutes int) {
	fmt.Printf("%v minute(s) left for your next break\n", minutes)
}
