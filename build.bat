@echo off
git describe --all --long | pipe run "in=text.cut(-8, -1)=$buildId&in.file('version.tmp')=text.replace('{buildId}', $buildId)=out.file('cmd/version.go')"
go install