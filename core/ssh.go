package core

import (
	"fmt"
	"log"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/fatih/color"
	"github.com/j3ssie/sfleet/utils"
	"github.com/melbahja/goph"
)

func SampleConnect(host string) error {
	sshPrivateKey := utils.NormalizePath("~/.ssh/id_rsa")
	passPhrase := ""
	auth, err := goph.Key(sshPrivateKey, passPhrase)
	if err != nil {
		// handle error
		log.Fatal(err)
		return err
	}

	client, err := goph.New("root", host, auth)
	if err != nil {
		// handle error
		log.Fatal(err)
		return err
	}

	out, err := client.Run("ls -lat /")

	fmt.Println(string(out))
	return err
}

func (r *Runner) Connect() error {
	auth, err := goph.Key(r.SSHPrivateKey, r.PassPhrase)
	if err != nil {
		return err
	}
	utils.DebugF("connecting to %v -- %v", r.SSHUser, r.Host)

	client, err := goph.New(r.SSHUser, r.Host, auth)
	if err != nil {
		return err
	}

	r.Client = client
	return nil
}

func (r *Runner) RunCommnad(cmd string) (err error) {
	operation := func() error {
		utils.DebugF("running command '%v' on %v", cmd, r.Host)

		out, err := r.Client.Run(cmd)
		if err != nil {
			return err
		}

		prefixOutput(r.Host, string(out))
		fmt.Println("")
		return nil
	}
	err = backoff.Retry(operation, r.BackOff)
	if err != nil {
		utils.WarnF("error create instance action %v", r.Host)
		return err
	}
	return nil
}

func prefixOutput(prefix string, out string) (content string) {
	c := color.New(color.FgGreen, color.Bold)
	whiteBackground := c.Add(color.BgBlack)

	prefix = whiteBackground.Sprintf(fmt.Sprintf("%v >>", prefix))
	newOut := strings.Split(out, "\n")

	for _, line := range newOut {
		data := prefix + " " + line
		fmt.Println(data)
		content += data
	}

	return content
}
