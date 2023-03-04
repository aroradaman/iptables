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
Writer/Rules-1000-8   0.0002208n ± 4%   0.0001650n ± 2%  -25.27% (p=0.000 n=25)
Writer/Rules-1100-8   0.0002326n ± 2%   0.0001772n ± 3%  -23.82% (p=0.000 n=25)
Writer/Rules-1200-8   0.0002489n ± 1%   0.0001890n ± 1%  -24.07% (p=0.000 n=25)
Writer/Rules-1300-8   0.0002676n ± 1%   0.0002027n ± 1%  -24.25% (p=0.000 n=25)
Writer/Rules-1400-8   0.0002845n ± 1%   0.0002158n ± 1%  -24.15% (p=0.000 n=25)
Writer/Rules-1500-8   0.0003085n ± 3%   0.0002300n ± 3%  -25.45% (p=0.000 n=25)
Writer/Rules-1600-8   0.0003407n ± 3%   0.0002434n ± 3%  -28.56% (p=0.000 n=25)
Writer/Rules-1700-8   0.0003570n ± 3%   0.0002546n ± 1%  -28.68% (p=0.000 n=25)
Writer/Rules-1800-8   0.0003978n ± 2%   0.0002909n ± 3%  -26.87% (p=0.000 n=25)
Writer/Rules-1900-8   0.0004162n ± 2%   0.0003065n ± 2%  -26.36% (p=0.000 n=25)
Writer/Rules-2000-8   0.0004327n ± 1%   0.0003197n ± 4%  -26.12% (p=0.000 n=25)
geomean               0.0003108n        0.0002306n       -25.80%
```