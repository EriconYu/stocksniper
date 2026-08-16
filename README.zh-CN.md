# StockSniper Community

[English](README.md) | 简体中文

面向 A 股量化研究场景的 **Python 3 开源展示版**。项目提供一个可直接运行的产品壳，用于展示市场环境、今日观察、策略管理、综合评分、回测、板块、信号、数据任务和执行日志等工作流。

> 本仓库是社区展示项目，不包含正式版的行情采集、私有数据源、策略算法、参数阈值、回测引擎、任务调度、交易辅助和数据库实现。页面中的股票、行情、评分和信号均为固定演示数据，不构成任何投资建议。

[在线体验：facai.stackbang.com](https://facai.stackbang.com)

## 项目定位

StockSniper 最初是一个 Go 股票爬虫。社区版 2.0 已彻底重构为 Python 3 + FastAPI 项目，目标不是公开生产系统源码，而是提供：

- 清晰的量化产品信息架构与前端工作台示例；
- 可运行、可二次开发的 FastAPI 项目骨架；
- 不需要数据库、行情账号或任何 API Key 的本地体验；
- 正式产品能力的公开说明和体验入口；
- 安全的开源边界，避免私有算法、凭据和数据被误提交。

## 界面功能

| 模块 | 社区展示版 | 正式版 |
| --- | --- | --- |
| 研究总览 | 固定演示快照、市场温度图、观察列表 | 实时任务状态与研究结果 |
| 市场环境 | 演示指数、趋势、动量、流动性、风险 | 多数据源市场状态计算 |
| 今日观察 | 演示候选与盘后摘要 | 完整筛选、过滤与研究建议 |
| 策略中心 | 策略方向、状态和候选数量展示 | 私有策略执行、参数与准入管理 |
| 综合评分 | 演示排行和因子维度 | 全市场因子计算与排名 |
| 回测验证 | 参数表单与完整工作流入口 | 成交约束、费用、滑点与绩效分析 |
| 板块分析 | 演示行业/概念热度 | 板块同步、成分股与轮动分析 |
| 信号追踪 | 演示信号时间线 | 实时信号、通知与状态跟踪 |
| 数据管理 | 任务表单与受限提示 | 增量同步、质量检查与多源回退 |
| 执行日志 | 固定英文任务日志 | 调度、进度、异常与审计日志 |

回测、策略执行、数据同步、导出和交易助手按钮会调用受限展示接口，明确提示该能力未包含，并引导用户访问正式版或联系作者。仓库内没有这些功能的占位算法或弱化实现，避免对生产能力产生误导。

## 技术结构

```text
stocksniper/
├── app/
│   ├── api/                 # 公开演示 API
│   ├── core/                # 产品元信息与固定演示数据
│   ├── static/              # 无构建依赖的响应式研究台
│   └── main.py              # FastAPI 应用入口
├── docs/                    # 架构与安全边界说明
├── scripts/                 # 提交前密钥检查
├── tests/                   # API、安全边界与静态资源测试
├── Dockerfile
└── pyproject.toml
```

前端使用原生 HTML、CSS 和 JavaScript，无 Node.js 构建步骤；后端仅依赖 FastAPI 与 Uvicorn。默认不读写数据库、不访问第三方行情接口、不启动后台任务。

## 快速开始

需要 Python 3.10 或更高版本。

```bash
git clone https://github.com/EriconYu/stocksniper.git
cd stocksniper

python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python -m app.main
```

浏览器访问 [http://127.0.0.1:8000](http://127.0.0.1:8000)，API 文档位于 [http://127.0.0.1:8000/api/docs](http://127.0.0.1:8000/api/docs)。

开发模式：

```bash
pip install -r requirements-dev.txt
make dev
```

Docker：

```bash
docker compose up --build
```

## 公开 API

| 方法 | 路径 | 用途 |
| --- | --- | --- |
| `GET` | `/api/health` | 健康检查 |
| `GET` | `/api/v1/meta` | 产品和导航元信息 |
| `GET` | `/api/v1/overview` | 总览演示数据 |
| `GET` | `/api/v1/market` | 市场环境演示数据 |
| `GET` | `/api/v1/strategies` | 策略元信息 |
| `GET` | `/api/v1/rankings` | 综合评分演示数据 |
| `GET` | `/api/v1/boards` | 板块演示数据 |
| `GET` | `/api/v1/signals` | 信号演示数据 |
| `POST` | `/api/v1/showcase-action` | 受限能力的统一产品引导 |

API 只返回仓库内可见的固定演示内容，不会隐式调用正式版服务。

## 安全边界

- 不提交 `.env`、数据库地址、用户名、密码、Cookie、Token、API Key 或私钥；
- 不复制正式版数据源实现、策略源码、参数、回测逻辑和内部接口；
- `.gitignore` 默认排除配置、数据、结果、日志和数据库文件；
- `python scripts/security_check.py` 会检查危险文件名和疑似硬编码凭据；
- CI 在每次提交和 Pull Request 中执行代码检查、测试和密钥检查。

发现安全问题请不要公开提交 Issue，处理方式见 [SECURITY.md](SECURITY.md)。

## 开发检查

```bash
pip install -r requirements-dev.txt
make check
```

## 正式版与联系

正式产品地址：[https://facai.stackbang.com](https://facai.stackbang.com)

需要完整能力、部署合作或定制开发，请联系作者，微信：**sudo4fate**。

<p align="center">
  <img src="app/static/assets/wechat-qr.jpg" width="360" alt="作者微信二维码" />
</p>

## 免责声明

本项目仅用于软件展示、技术交流和研究，不提供证券投资咨询，不承诺数据准确性、完整性或及时性。任何页面、接口和示例均不构成买入、卖出或持有任何证券的建议。使用者应独立判断并自行承担风险。

## License

[GNU General Public License v3.0](LICENSE)
