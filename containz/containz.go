package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	res, err := Run("example")
	if err != nil {
		log.Fatalf("epic fail %s", err)
	}
	log.Print(res)
}

// returns the output logs
func Run(code string) (string, error) {
	// if err := os.WriteFile("./docker/main.go", []byte(code), 0644); err != nil {
	// 	return "", errors.New("could not write to file")
	// }
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fp := filepath.Join(cwd, "docker")
	err = os.Chdir(fp)

	t := fmt.Sprint(time.Now().Unix())

	// cmd := exec.Command("docker", "run", "--rm", "-it", "$(docker", "build", "-q", ".")
	cmd := exec.Command("docker", "build", ".", "-t", t)
	if err = cmd.Run(); err != nil {
		log.Println("build")
		return "", err
	}

	cmd = exec.Command("docker", "run", "--name", t, t)
	if err = cmd.Run(); err != nil {
		log.Println("run")
		return "", err
	}

	out, err := exec.Command("docker", "logs", t).Output()
	if err != nil {
		log.Println("logs")
		return "", err
	}
	return string(out), nil
}
