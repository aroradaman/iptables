#### Simple and Efficient interface for generating iptables rules.
```
rule := rules.NewRule()
rule = rule.ForChain(kubeExternalServicesNewChain)
rule = rule.AddComment(fmt.Sprintf(`"service %s has no endpoints"`, testServiceName))
rule = rule.MatchAddrType(AddrTypeDestinationType, AddressTypeLocal)
rule = rule.MatchProtocol(ProtocolUDP)
rule = rule.MatchDestinationPort(32320)
rule = rule.JumpTo(JumpTargetReject)
rule.Add()
```
#### Benchmark
Use `./compare.sh` to run benchmarking

#### Benchmark Results
```
./compare.sh

goos: darwin
goarch: amd64
pkg: github.com/daman1807/iptables
cpu: VirtualApple @ 2.50GHz
                    │   old.report    │               new.report                │
                    │     sec/op      │     sec/op       vs base                │
Writer/Rules-1000-8   0.0003046n ± 6%   0.0001967n ± 4%  -35.44% (p=0.000 n=10)
Writer/Rules-1100-8   0.0003152n ± 3%   0.0002102n ± 3%  -33.30% (p=0.000 n=10)
Writer/Rules-1200-8   0.0003409n ± 4%   0.0002266n ± 4%  -33.54% (p=0.000 n=10)
Writer/Rules-1300-8   0.0003680n ± 3%   0.0002419n ± 1%  -34.25% (p=0.000 n=10)
Writer/Rules-1400-8   0.0004062n ± 4%   0.0002595n ± 2%  -36.14% (p=0.000 n=10)
Writer/Rules-1500-8   0.0004335n ± 5%   0.0002749n ± 1%  -36.57% (p=0.000 n=10)
Writer/Rules-1600-8   0.0004533n ± 4%   0.0002932n ± 3%  -35.33% (p=0.000 n=10)
Writer/Rules-1700-8   0.0004706n ± 2%   0.0003073n ± 5%  -34.71% (p=0.000 n=10)
Writer/Rules-1800-8   0.0005230n ± 5%   0.0003220n ± 3%  -38.42% (p=0.000 n=10)
Writer/Rules-1900-8   0.0005682n ± 3%   0.0003640n ± 1%  -35.94% (p=0.000 n=10)
Writer/Rules-2000-8   0.0005877n ± 3%   0.0003791n ± 3%  -35.49% (p=0.000 n=10)
geomean               0.0004237n        0.0002738n       -35.39%


```
