sass --no-source-map handlers/static/styles.scss:handlers/static/styles.min.css --style compressed
pandoc -f markdown -t html ./handlers/static/resources.md -o ./handlers/static/resources.html
go run ./cmd/server/main.go
