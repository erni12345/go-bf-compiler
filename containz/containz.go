package containz

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// returns the output logs
func Run(code string) (string, error) {
	// if err := os.WriteFile("./docker/main.go", []byte(code), 0644); err != nil {
	// 	return "", err
	// }

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fp := filepath.Join(cwd, "docker")
	err = os.Chdir(fp)

	t := fmt.Sprint(time.Now().Unix())

	if err = build(t); err != nil {
		return "", err
	}

	out, err := run(t)
	if err != nil {
		return "", err
	}

	return out, nil
}

func build(tag string) error {
	// TODO do it with this command
	// cmd := exec.Command("docker", "run", "--rm", "-it", "$(docker", "build", "-q", ".")
	cmd := exec.Command("docker", "build", ".", "-t", tag)
	if err := cmd.Run(); err != nil {
		log.Println("build")
		return err
	}
	return nil
}

func run(tag string) (string, error) {
	out, err := exec.Command("docker", "run", tag).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
