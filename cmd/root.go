package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/j3ssie/sfleet/core"
	"github.com/j3ssie/sfleet/libs"
	"github.com/j3ssie/sfleet/utils"
	"github.com/spf13/cobra"
)

var options = libs.Options{}

var RootCmd = &cobra.Command{
	Use:   libs.BINARY,
	Short: libs.DESC,
}

// Execute main function
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&options.SSHPrivateKey, "key", "~/.ssh/id_rsa", "Default SSH Key")
	RootCmd.PersistentFlags().StringVar(&options.Credentials, "cred", "", "Password or Passphrase for SSHKey")
	RootCmd.PersistentFlags().StringVar(&options.ConfigFile, "config", "~/.sfleet/config", "Config file to store host and keypairs")

	RootCmd.PersistentFlags().StringVar(&options.Port, "port", "22", "Port for SSH")
	RootCmd.PersistentFlags().StringVar(&options.User, "user", "root", "User for SSH")

	RootCmd.PersistentFlags().StringSliceVarP(&options.Inputs, "target", "t", []string{}, "host to run")
	RootCmd.PersistentFlags().StringVarP(&options.InputFile, "targets", "T", "", "List of hosts to run")

	RootCmd.PersistentFlags().IntVarP(&options.Concurrency, "concurrency", "c", 5, "Concurrency")
	RootCmd.PersistentFlags().BoolVar(&options.Debug, "debug", false, "Debug")
	RootCmd.PersistentFlags().IntVar(&options.Retry, "retry", 8, "Number of retry when command is error")

	// RootCmd.SetHelpFunc(RootMessage)
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	logLevel := "info"
	if options.Debug {
		logLevel = "debug"
	}
	utils.InitLog(logLevel)
	core.InitConfig(&options)

	// detect if anything came from stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			target := strings.TrimSpace(sc.Text())
			if err := sc.Err(); err == nil && target != "" {
				options.Inputs = append(options.Inputs, target)
			}
		}
	}

	//options.InputFile, _ = cmd.Flags().GetString("targets")
	if options.InputFile != "" {
		if utils.FileExists(options.InputFile) {
			options.Inputs = append(options.Inputs, utils.ReadingLines(options.InputFile)...)
		}
	}

}
