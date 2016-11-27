package main
import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//"strings"
)

func main() {
	cmd := exec.Command("bash", "-c", "lss -cal ~/")
	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut
	err := cmd.Run()
	fmt.Printf("%s\n", errOut.String())
	if err != nil {
		fmt.Printf("H1\n")
		log.Fatal(err)
		fmt.Printf("H2")
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
