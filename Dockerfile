FROM scratch

go build -o /opt/rss-resource/in in/in.go
go build -o /opt/rss-resource/check check/check.go
go build -o /opt/rss-resource/out out/out.go
