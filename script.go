package main

import (
	"fmt"
	"os/exec"
)

func createCronJob() {
	cmd := exec.Command("crontab", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error fetching current crontab:", err)
		return
	}

	// Add a new cron job
	newJob := "0 * * * * /usr/bin/my-script.sh"
	cmd = exec.Command("crontab", "-")
	cmd.Stdin = append(output, []byte("\n"+newJob)...)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error updating crontab:", err)
	}
	fmt.Println("Cron job created successfully.")
}

func main() {
	createCronJob()
}
