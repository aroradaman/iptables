package iptables

import (
	"fmt"
	"strconv"
)

func OldWriter(n int) []byte {
	rules := LineBuffer{}
	args := []string{
		"-A", "kubeExternalServicesChain",
		"-m", "comment", "--comment", fmt.Sprintf(`"service %s has no endpoints"`, "test/test"),
		"-p", "udp",
	}
	for i := 0; i < n; i++ {
		rules.Write(args,
			"-m", "addrtype", "--dst-type", "LOCAL",
			"--dport", strconv.Itoa(32320),
			"-j", "REJECT")
	}
	return rules.Bytes()
}

func pack(args ...string) []string {
	return args
}

func NewWriter(n int) []byte {
	rules := LineBuffer{}
	args := []string{
		"-A", "kubeExternalServicesChain",
		"-m", "comment", "--comment", fmt.Sprintf(`"service %s has no endpoints"`, "test/test"),
		"-p", "udp",
	}
	for i := 0; i < n; i++ {
		rules.NewWrite(args,
			pack("-m", "addrtype", "--dst-type", "LOCAL",
				"--dport", strconv.Itoa(32320),
				"-j", "REJECT"))
	}
	return rules.Bytes()
}
