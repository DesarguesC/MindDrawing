package controller

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const sd_path = "~/autodl-tmp"

func Txt2Img(prompt string, path string, name string) error {
	// status: user, path has been calculated before cin
	var cmd *exec.Cmd
	cmd = exec.Command("conda", "activate", "ldm")
	outdir := path
	if strings.LastIndex(name, ".jpg") == -1 || strings.LastIndex(name, ".jpeg") == -1 {
		name = name + ".jpg"
	}
	cmd.Dir = sd_path
	args := []string{"--prompt", prompt, "--outdir", outdir, "--name", name}

	cmd.Stdout = os.Stdout
	cmd = exec.Command("python", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("command output: %q", out.String())
	outs := out.String()
	if strings.LastIndex(outs, "CUDA") != -1 {
		return errors.New("cuda out of memory, please wait for queue")
	} else {
		return nil
	}
}
