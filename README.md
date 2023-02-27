#### Simple and Efficient interface for generating iptables rules.

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
Writer/Rules-1000-8   0.0002131n ± 5%   0.0001312n ± 9%  -38.42% (p=0.000 n=10)
Writer/Rules-1100-8   0.0002254n ± 3%   0.0001411n ± 8%  -37.42% (p=0.000 n=10)
Writer/Rules-1200-8   0.0002399n ± 4%   0.0001460n ± 6%  -39.15% (p=0.000 n=10)
Writer/Rules-1300-8   0.0002573n ± 8%   0.0001551n ± 4%  -39.75% (p=0.000 n=10)
Writer/Rules-1400-8   0.0002741n ± 4%   0.0001690n ± 8%  -38.35% (p=0.000 n=10)
Writer/Rules-1500-8   0.0003004n ± 3%   0.0001734n ± 4%  -42.28% (p=0.000 n=10)
Writer/Rules-1600-8   0.0003108n ± 4%   0.0001815n ± 2%  -41.61% (p=0.000 n=10)
Writer/Rules-1700-8   0.0003290n ± 6%   0.0001926n ± 3%  -41.44% (p=0.000 n=10)
Writer/Rules-1800-8   0.0003454n ± 7%   0.0001993n ± 1%  -42.29% (p=0.000 n=10)
Writer/Rules-1900-8   0.0003636n ± 3%   0.0002101n ± 3%  -42.22% (p=0.000 n=10)
Writer/Rules-2000-8   0.0004211n ± 5%   0.0002475n ± 2%  -41.24% (p=0.000 n=10)
geomean               0.0002921n        0.0001741n       -40.40%

```