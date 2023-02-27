package iptables

import (
	"bytes"
	"strconv"
)

type RuleBuilder func(*bytes.Buffer)

func ForChain(chain Chain) RuleBuilder {
	return func(buffer *bytes.Buffer) {
		buffer.Write(appendPrefix)
		buffer.Write(chain)
		buffer.WriteByte(' ')
	}
}

func MatchProtocol(protocol Protocol) RuleBuilder {
	return func(buffer *bytes.Buffer) {
		buffer.Write(protocolPrefix)
		buffer.Write(protocol)
		buffer.WriteByte(' ')
	}
}

func MatchAddrType(subType AddrTypeSubType, val AddressType) RuleBuilder {
	return func(buffer *bytes.Buffer) {
		buffer.Write(matchModuleAddressTypePrefix)
		buffer.Write(subType)
		buffer.WriteByte(' ')
		buffer.Write(val)
		buffer.WriteByte(' ')
	}
}

func MatchDestinationPort(port int) RuleBuilder {
	return func(buffer *bytes.Buffer) {
		buffer.Write(destinationPortPrefix)
		buffer.WriteString(strconv.Itoa(port))
		buffer.WriteByte(' ')
	}
}

func AddComment(comment string) RuleBuilder {
	return func(buffer *bytes.Buffer) {
		buffer.Write(matchModuleCommentPrefix)
		buffer.WriteString(comment)
		buffer.WriteByte(' ')
	}
}

func JumpTo(jumpTarget JumpTarget) RuleBuilder {
	return func(buffer *bytes.Buffer) {
		buffer.Write(jumpPrefix)
		buffer.Write(jumpTarget)
	}
}
