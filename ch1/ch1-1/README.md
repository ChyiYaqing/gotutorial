# Go相对路径

Golang的相对路径是相对于执行命令时的目录

* go run
```
gotutorial/ch1/ch1-1 on  main [✘?] via 🐹 v1.21.5 
➜ go run main.go 
/var/folders/_q/5fhdnhsx0mg_52w806f3898m0000gn/T/go-build1552571369/b001/exe
```

go run 的输出结果时一个临时文件地址

* go build
```
gotutorial/ch1/ch1-1 on  main [✘?] via 🐹 v1.21.5 
➜ go build main.go

gotutorial/ch1/ch1-1 on  main [✘?] via 🐹 v1.21.5 
➜ ./main 
/Users/chyiyaqing/chyi/github.com/gotutorial/ch1/ch1-1
```