package iptables

import (
	"strconv"
)

func OldWriter(n int) []byte {
	rules := LineBuffer{}
	for i := 0; i < n; i++ {
		protocol := "udp"
		rules.Write("-A", string(kubeExternalServicesChain),
			"-m", "comment", "--comment", `"service has no endpoints"`,
			"-m", "addrtype", "--dst-type", "LOCAL",
			"-p", protocol,
			"--dport", strconv.Itoa(32320),
			"-j", "REJECT")
	}
	return rules.Bytes()
}

func NewWriter(n int) []byte {
	rules := NewRules()

	for i := 0; i < n; i++ {
		//rules.Add(
		//	ForChain(kubeExternalServicesNewChain),
		//	AddComment(`"service has no endpoints"`),
		//	MatchAddrType(AddrTypeDestinationType, AddressTypeLocal),
		//	MatchProtocol(ProtocolUDP),
		//	MatchDestinationPort(32320),
		//	JumpTo(JumpTargetReject),
		//)

		rule := rules.NewRule()
		rule = rule.ForChain(kubeExternalServicesNewChain)
		rule = rule.AddComment(`"service has no endpoints"`)
		rule = rule.MatchAddrType(AddrTypeDestinationType, AddressTypeLocal)
		rule = rule.MatchProtocol(ProtocolUDP)
		rule = rule.MatchDestinationPort(32320)
		rule = rule.JumpTo(JumpTargetReject)
		rule.Add()
	}

	return rules.Bytes()
}
