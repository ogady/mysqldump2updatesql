# mysqldump2updatesql

## install

```
go install github.com/ogady/mysqldump2updatesql@latest
export GOBIN=~/bin
export PATH=$PATH:$GOBIN
```

## usage

1. mysqldumpでCSVエクスポートするなどしてこんな感じのCSVを作る。

```input.csv:
"id","hoge_status","created_at","updated_at"
11111,1,"2022-01-01 11:11:11","2022-02-02 22:22:22"
11112,0,"2022-01-02 12:12:12","2023-03-03 23:23:23"
```

```sh
mysqldump2updatesql -input=${path to input csv} -output=${path to output sql}
```
