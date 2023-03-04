#!/bin/bash
go test -v
#GOMAXPROCS=1

WRITER=OLD go test -bench=. --count=25 -run=^#  > old.report
WRITER=NEW go test -bench=. --count=25 -run=^#  > new.report

benchstat old.report new.report
rm old.report new.report