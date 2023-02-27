package iptables

import (
	"bytes"
	"strconv"
)

var appendPrefix = []byte("-A ")

var matchModuleCommentPrefix = []byte("-m comment --comment ")
var matchModuleAddressTypePrefix = []byte("-m addrtype ")

var destinationPortPrefix = []byte("--dport ")

var jumpPrefix = []byte("-j ")

var protocolPrefix = []byte("-p ")

type Rules struct {
	buffer *bytes.Buffer
}

func NewRules() *Rules {
	return &Rules{buffer: new(bytes.Buffer)}
}

func (r *Rules) Add(builder ...RuleBuilder) {
	for i := 0; i < len(builder); i++ {
		builder[i](r.buffer)
	}
	r.buffer.WriteByte('\n')
}

func (r *Rules) Bytes() []byte {
	return r.buffer.Bytes()
}

func (r *Rules) Write(i []byte) {
	r.buffer.Write(i)
}

type Rule struct {
	buffer *bytes.Buffer
}

func (r *Rules) NewRule() *Rule {
	return &Rule{buffer: r.buffer}
}

func (r *Rule) ForChain(chain Chain) *Rule {
	r.buffer.Write(appendPrefix)
	r.buffer.Write(chain)
	r.buffer.WriteByte(' ')
	return r
}

func (r *Rule) MatchProtocol(protocol Protocol) *Rule {
	r.buffer.Write(protocolPrefix)
	r.buffer.Write(protocol)
	r.buffer.WriteByte(' ')
	return r
}

func (r *Rule) MatchAddrType(subType AddrTypeSubType, val AddressType) *Rule {
	r.buffer.Write(matchModuleAddressTypePrefix)

	r.buffer.Write(subType)
	r.buffer.WriteByte(' ')
	r.buffer.Write(val)
	r.buffer.WriteByte(' ')
	return r
}

func (r *Rule) MatchDestinationPort(port int) *Rule {

	r.buffer.Write(destinationPortPrefix)
	r.buffer.WriteString(strconv.Itoa(port))
	r.buffer.WriteByte(' ')
	return r
}

func (r *Rule) AddComment(comment string) *Rule {
	r.buffer.Write(matchModuleCommentPrefix)
	r.buffer.WriteString(comment)
	r.buffer.WriteByte(' ')
	return r
}

func (r *Rule) JumpTo(jumpTarget JumpTarget) *Rule {
	r.buffer.Write(jumpPrefix)
	r.buffer.Write(jumpTarget)
	return r
}

func (r *Rule) Add() {
	r.buffer.WriteByte('\n')
}
