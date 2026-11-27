// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cilium/cilium/pkg/common"
	"github.com/cilium/cilium/pkg/maps/ctmap"
	"github.com/cilium/cilium/pkg/maps/timestamp"
	"github.com/cilium/stream"
	"github.com/spf13/cobra"
)

var bpfCtGcCmd = cobra.Command{
	Use:   "gc global",
	Short: "Trigger safe connection tracking GC",
	Run: func(cmd *cobra.Command, args []string) {
		common.RequireRootPrivilege("cilium bpf ct gc")
		gcCt()
	},
}

func init() {
	BPFCtCmd.AddCommand(&bpfCtGcCmd)
}

func gcCt() {

	clockSource, err := timestamp.GetClockSourceFromAgent(client.Daemon)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get clock source from agent: %s. Falling back to runtime config.\n", err)
		clockSource, err = timestamp.GetClockSourceFromRuntimeConfig()
		if err != nil {
			Fatalf("Unable to determine clock source: %s", err)
		}
	}

	ipv4, ipv6 := getIpEnableStatuses()

	observable4, next4, complete4 := stream.Multicast[ctmap.GCEvent]()
	observable6, next6, complete6 := stream.Multicast[ctmap.GCEvent]()
	observable4.Observe(context.Background(), ctmap.NatMapNext4, func(error) {})
	observable6.Observe(context.Background(), ctmap.NatMapNext6, func(error) {})

	for _, m := range ctmap.GlobalMaps(ipv4, ipv6) {
		path, err := ctmap.OpenCTMap(m)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Unable to open %s: %s. Skipping.\n", path, err)
				continue
			}
			Fatalf("Unable to open %s: %s", path, err)
		}
		defer m.Close()

		filter := ctmap.GCFilter{
			RemoveExpired: true,
		}
		if _, err := ctmap.GCWithClockSource(m, filter, clockSource, next4, next6); err != nil {
			Fatalf("GC failed on %s: %s", m.Name(), err)
		}
	}
	complete4(nil)
	complete6(nil)

}
