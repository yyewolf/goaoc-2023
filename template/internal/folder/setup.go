package folder

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateDay(year string, day int) {
	// Create ./year/day/
	// Copy what's in ./template/go ./year/day
	os.MkdirAll("./"+year+"/day"+fmt.Sprintf("%02d", day), 0755)
	c := "cp -r ./template/go/* ./" + year + "/day" + fmt.Sprintf("%02d", day)
	cmd := exec.Command("sh", "-c", c)
	cmd.Run()
}

func CreateDayInput(year string, day int, input string) error {
	// Write to ./year/day/input.txt
	f, err := os.OpenFile("./"+year+"/day"+fmt.Sprintf("%02d", day)+"/input.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(input)
	return err
}
