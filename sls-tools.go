package main

import "sls-tools/core"

func main() {

	ic := core.NewIndexClient()

	for _, logstoreName := range ic.DumpLogstore {
		ic.DumpConfiguration(ic.DumpProject, logstoreName)
	}

	for _, logstoreName := range ic.ReceiveLogstore {
		ic.ApplyConfiguration(ic.ReceiveProject, logstoreName)
	}

}
