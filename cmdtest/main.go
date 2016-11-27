package main
import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c", "lss -cal ~/")
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err := cmd.Run()
	log.Println("Done running")
	if err != nil {
		log.Printf("Error: %s", errOut.String())
		log.Fatal(err)
	}
	log.Printf("StdOut %s", out.String())
	log.Printf("StdErr %s", errOut.String())
}
