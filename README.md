# project-layout
Go Project Layout

目錄說明
```
.
├── initialize  <- 初始化集中地 (手動初始的地方，為了能控制初始的時機
│   └── initialize.go <- 在 main 初始時要注意順序
├── internal <- package 的地方
│   └── user
│       ├── user.go
│       └── tools.go <- 局部 package 共用的函示可以建立 tools.go
├── obj
│   └── obj.go
├── procedure   <- 背景程式目錄
│   ├── task       <- 任務型程式
│   └── worker     <- 背景型程式
├── repository
│   ├── cache
│   └── models
├── server <- http Server 入口目錄 
│   ├── controller <- (這一層的檔案名稱與 package 名稱會使用底線)
│   │   └── api            <- 一個資料夾為一個 group      
│   │     └── v1
│   │       └── auth
│   │         └── post_login.go <- 一個 api 一個檔案(前面為 method, 後面為 api 名稱)
│   ├── middleware
│   ├── protocol  <- api相關協定，包含可以回傳的error code訊息
│   └── router   <- 管理 group
├── wsserver
├── grpcserver
├── service  <-  other Server api 
├── test <- 測項(test api)
│   ├── controller
│   │   └── api
│   │     └── v1
│   │       └── auth
│   │         └── post_login_test.go
│   └── itest <- 共用 test 工具
│   └── testingmock
└── utils  <- 共用函式目錄
    ├── consts
    ├── env
    └── utils.go
```
