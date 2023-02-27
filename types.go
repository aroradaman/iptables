package iptables

type IPTChain string

const kubeExternalServicesChain IPTChain = "KUBE-EXTERNAL-SERVICES"

type Chain []byte

var kubeExternalServicesNewChain Chain = []byte("KUBE-EXTERNAL-SERVICES")

type JumpTarget []byte

var JumpTargetReject JumpTarget = []byte("REJECT")

type AddrTypeSubType []byte

var AddrTypeDestinationType AddrTypeSubType = []byte("--dst-type")

type Protocol []byte

var ProtocolUDP = []byte("udp")

type AddressType []byte

var AddressTypeLocal = []byte("LOCAL")
