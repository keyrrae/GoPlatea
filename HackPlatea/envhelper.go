package main

import (
	"log"
	"os"
	"os/exec"
	"fmt"
)

type EnvHelper struct {
	Filename []string
}

func NewEnvHelper(filename []string) EnvHelper {
	res := EnvHelper{
		Filename: filename,
	}

	return res
}

func (eh EnvHelper) GetCurrDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/"
}

func (eh EnvHelper) ClearWorkspace() (string, error) {
	for _, file := range eh.Filename {
		os.Remove(file)
	}
	cmd := exec.Command("touch", ".hhconfig")
	output, err := cmd.CombinedOutput()
	res := string(output)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (eh EnvHelper) WriteProgToSystem(prog string) error {
	f, err := os.Create(eh.Filename[0])
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("prog"+prog)

	_, err = f.Write([]byte(prog))
	if err != nil {
		return err
	}

	f.Sync()
	return nil
}
