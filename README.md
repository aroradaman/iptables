#### Benchmark
Use `./compare.sh` to run benchmarking

#### Benchmark Results
```
$> ./compare.sh

goos: darwin
goarch: amd64
pkg: github.com/daman1807/iptables
cpu: VirtualApple @ 2.50GHz
                    │    old.report     │                new.report                │
                    │      sec/op       │      sec/op       vs base                │
Writer/Rules-100-8    0.00003030n ± 12%   0.00002210n ± 8%  -27.06% (p=0.000 n=25)
Writer/Rules-200-8    0.00004890n ±  5%   0.00003740n ± 7%  -23.52% (p=0.000 n=25)
Writer/Rules-300-8    0.00006990n ±  6%   0.00005530n ± 5%  -20.89% (p=0.000 n=25)
Writer/Rules-400-8    0.00008860n ±  3%   0.00006700n ± 3%  -24.38% (p=0.000 n=25)
Writer/Rules-500-8    0.00011290n ±  2%   0.00008750n ± 3%  -22.50% (p=0.000 n=25)
Writer/Rules-600-8    0.00012860n ±  3%   0.00009900n ± 1%  -23.02% (p=0.000 n=25)
Writer/Rules-700-8     0.0001554n ±  4%    0.0001117n ± 2%  -28.12% (p=0.000 n=25)
Writer/Rules-800-8     0.0001795n ±  2%    0.0001264n ± 3%  -29.58% (p=0.000 n=25)
Writer/Rules-900-8     0.0001944n ±  4%    0.0001535n ± 3%  -21.04% (p=0.000 n=25)
Writer/Rules-1000-8    0.0002239n ±  3%    0.0001628n ± 1%  -27.29% (p=0.000 n=25)
Writer/Rules-1100-8    0.0002336n ±  2%    0.0001762n ± 1%  -24.57% (p=0.000 n=25)
Writer/Rules-1200-8    0.0002489n ±  3%    0.0001910n ± 1%  -23.26% (p=0.000 n=25)
Writer/Rules-1300-8    0.0002679n ±  1%    0.0002025n ± 1%  -24.41% (p=0.000 n=25)
Writer/Rules-1400-8    0.0002845n ±  3%    0.0002145n ± 1%  -24.60% (p=0.000 n=25)
Writer/Rules-1500-8    0.0003116n ±  1%    0.0002305n ± 3%  -26.03% (p=0.000 n=25)
Writer/Rules-1600-8    0.0003213n ±  1%    0.0002453n ± 2%  -23.65% (p=0.000 n=25)
Writer/Rules-1700-8    0.0003500n ±  3%    0.0002642n ± 3%  -24.51% (p=0.000 n=25)
Writer/Rules-1800-8    0.0003916n ±  3%    0.0002896n ± 3%  -26.05% (p=0.000 n=25)
Writer/Rules-1900-8    0.0004197n ±  2%    0.0003112n ± 2%  -25.85% (p=0.000 n=25)
Writer/Rules-2000-8    0.0004365n ±  2%    0.0003192n ± 2%  -26.87% (p=0.000 n=25)
geomean                0.0001831n          0.0001376n       -24.89%
```