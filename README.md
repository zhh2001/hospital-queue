# 医院叫号系统

```txt
hospital-queue/
├── main.go                 # 程序入口，Gin路由设置
├── go.mod                  # Go模块依赖
├── go.sum                  # 依赖校验文件
├── data/                   # 数据存储目录
│   └── queue.json          # 叫号数据的JSON文件
├── handlers/               # 处理器（控制器）
│   └── queue_handler.go    # 叫号逻辑处理
├── models/                 # 数据模型定义
│   └── queue.go            # 叫号数据结构定义
├── service/                # 业务逻辑层
│   └── queue_service.go    # 封装JSON文件的读写操作
├── static/                 # 静态资源
│   ├── css/
│   │   └── style.css       # 页面样式
│   ├── js/
│   │   └── main.js         # 前端交互逻辑
│   └── images/             # 图片资源
├── templates/              # HTML模板
│   └── index.html          # 主页面
├── README.md               # 项目说明
├── LICENSE                 # 许可证文件
└── .gitignore              # Git忽略文件
```
