# 医院叫号系统

```txt
hospital-queue/
├── main.go                 # 程序入口，Gin路由设置
├── go.mod                  # Go模块依赖
├── go.sum                  # 依赖校验文件
├── database/               # 数据库相关
│   ├── sqlite.go           # SQLite连接和初始化
│   └── models/             # 数据模型
│       └── queue.go        # 叫号相关模型定义
├── handlers/               # 处理器（控制器）
│   └── queue_handler.go    # 叫号逻辑处理
├── static/                 # 静态资源
│   ├── css/
│   │   └── style.css       # 页面样式
│   ├── js/
│   │   └── main.js         # 前端交互逻辑
│   └── images/             # 图片资源
└── templates/              # HTML模板
    └── index.html          # 主页面
```
