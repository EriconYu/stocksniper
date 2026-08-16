# StockSniper Community

English | [简体中文](README.zh-CN.md)

A **Python 3 open-source showcase** for an A-share quantitative research platform. The project provides a runnable product shell covering market regime analysis, daily observations, strategy management, composite rankings, backtesting, sector analysis, signals, data tasks, and execution logs.

> This repository is a community showcase. It does not include production market-data collectors, private data providers, strategy algorithms, parameter thresholds, the backtest engine, task scheduling, trading assistance, or database implementations. All stocks, prices, scores, and signals shown in the interface are fixed demonstration data and do not constitute investment advice.

[Live product: facai.stackbang.com](https://facai.stackbang.com)

## Positioning

StockSniper started as a Go-based stock crawler. Community Edition 2.0 has been completely rebuilt with Python 3 and FastAPI. Its purpose is to provide:

- a clear information architecture and polished workstation for a quantitative research product;
- a runnable FastAPI foundation that is easy to inspect and extend;
- a local experience that requires no database, market-data account, or API key;
- a public overview of the complete product and a direct path to try it;
- an explicit open-source boundary that keeps private algorithms, credentials, and data out of the repository.

## Interface

| Module | Community showcase | Full product |
| --- | --- | --- |
| Research overview | Fixed snapshot, market temperature chart, watchlist | Live task status and research results |
| Market regime | Demo indices, trend, momentum, liquidity, and risk | Multi-source market regime calculation |
| Daily observations | Demo candidates and closing summary | Full screening, filtering, and research guidance |
| Strategy center | Strategy themes, status, and demo candidate counts | Private strategy execution, parameters, and admission controls |
| Composite ranking | Demo ranking and factor dimensions | Full-market factor calculation and ranking |
| Backtesting | Parameter form and complete workflow entry | Trading constraints, fees, slippage, and performance analysis |
| Sector analysis | Demo industry and concept heat | Sector synchronization, constituents, and rotation analysis |
| Signal tracking | Demo signal timeline | Live signals, notifications, and state tracking |
| Data management | Task form and gated product prompt | Incremental sync, quality checks, and provider fallback |
| Execution logs | Fixed English task logs | Scheduling, progress, exceptions, and audit logs |

Backtest, strategy execution, data sync, export, and trading-assistant controls call a single gated showcase endpoint. The interface clearly identifies those capabilities as unavailable in the community edition and directs users to the full product or the author. The repository intentionally contains no placeholder algorithms or weakened copies of production logic.

## Architecture

```text
stocksniper/
├── app/
│   ├── api/                 # Public showcase API
│   ├── core/                # Product metadata and fixed demo data
│   ├── static/              # Responsive workstation with no build step
│   └── main.py              # FastAPI application entry point
├── docs/                    # Architecture and security-boundary notes
├── scripts/                 # Pre-commit credential check
├── tests/                   # API, public-boundary, and static-asset tests
├── Dockerfile
└── pyproject.toml
```

The frontend uses plain HTML, CSS, and JavaScript, so no Node.js build is required. The backend depends only on FastAPI and Uvicorn. It does not read or write a database, contact a market-data provider, or start background jobs.

## Quick Start

Python 3.10 or newer is required.

```bash
git clone https://github.com/EriconYu/stocksniper.git
cd stocksniper

python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python -m app.main
```

Open [http://127.0.0.1:8000](http://127.0.0.1:8000). OpenAPI documentation is available at [http://127.0.0.1:8000/api/docs](http://127.0.0.1:8000/api/docs).

Development mode:

```bash
pip install -r requirements-dev.txt
make dev
```

Docker:

```bash
docker compose up --build
```

## Public API

| Method | Path | Purpose |
| --- | --- | --- |
| `GET` | `/api/health` | Health check |
| `GET` | `/api/v1/meta` | Product and navigation metadata |
| `GET` | `/api/v1/overview` | Overview demonstration data |
| `GET` | `/api/v1/market` | Market-regime demonstration data |
| `GET` | `/api/v1/strategies` | Public strategy metadata |
| `GET` | `/api/v1/rankings` | Composite-ranking demonstration data |
| `GET` | `/api/v1/boards` | Sector demonstration data |
| `GET` | `/api/v1/signals` | Signal demonstration data |
| `POST` | `/api/v1/showcase-action` | Unified product prompt for gated capabilities |

The API returns only fixed content visible in this repository. It never calls the full product implicitly.

## Security Boundary

- Never commit `.env` files, database addresses, usernames, passwords, cookies, tokens, API keys, or private keys.
- Do not copy production data-provider code, strategy source, parameters, backtest logic, or internal endpoints.
- `.gitignore` excludes configuration, market data, results, logs, and database files by default.
- `python scripts/security_check.py` checks dangerous filenames and likely hard-coded credentials.
- CI runs linting, tests, and the credential check for every push and pull request.

Do not report security issues in a public issue. See [SECURITY.md](SECURITY.md) for the private reporting process.

## Development Checks

```bash
pip install -r requirements-dev.txt
make check
```

## Full Product and Contact

Full product: [https://facai.stackbang.com](https://facai.stackbang.com)

For complete capabilities, deployment cooperation, or custom development, contact the author on WeChat: **sudo4fate**.

<p align="center">
  <img src="app/static/assets/wechat-qr.jpg" width="360" alt="Author WeChat QR code" />
</p>

## Disclaimer

This project is provided solely for software demonstration, technical exchange, and research. It is not investment advisory software and makes no guarantee regarding data accuracy, completeness, or timeliness. Nothing in the interface, API, or examples is a recommendation to buy, sell, or hold any security. Users are responsible for their own decisions and risks.

## License

[GNU General Public License v3.0](LICENSE)
