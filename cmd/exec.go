package cmd

import (
	"sync"

	"github.com/j3ssie/sfleet/core"
	"github.com/panjf2000/ants"
	"github.com/spf13/cobra"
)

func init() {
	var execCmd = &cobra.Command{
		Use:   "exec",
		Short: "Exec command on remote host",
		//Long:  core.Banner(),
		RunE: runExec,
	}

	execCmd.Flags().StringSliceVarP(&options.Commands, "cmd", "s", []string{}, "Target to running")
	RootCmd.AddCommand(execCmd)
}

var workers = []core.Worker{}

func runExec(cmd *cobra.Command, _ []string) error {
	workerPool, err := core.GetConfig(options)
	if err == nil {
		workers = workerPool.Workers
		for _, worker := range workerPool.Workers {
			options.Inputs = append(options.Inputs, worker.Host)
		}
	}

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(options.Concurrency, func(i interface{}) {
		defer wg.Done()
		host := i.(string)
		runner, err := core.InitProvider(host, options)
		if err != nil {
			return
		}
		runner.Start()
	}, ants.WithPreAlloc(true))
	defer p.Release()

	for _, host := range options.Inputs {
		wg.Add(1)
		_ = p.Invoke(host)
	}
	wg.Wait()
	return nil
}
