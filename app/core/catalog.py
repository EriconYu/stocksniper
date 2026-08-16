from __future__ import annotations

from copy import deepcopy
from typing import Any

PRODUCT = {
    "name": "StockSniper Community",
    "tagline": "A 股量化研究平台开源展示版",
    "edition": "community-showcase",
    "version": "2.0.0",
    "official_url": "https://facai.stackbang.com",
    "wechat": "sudo4fate",
    "disclaimer": "页面数据均为演示数据，不构成任何投资建议。",
}

NAVIGATION = [
    {"id": "dashboard", "label": "研究总览"},
    {"id": "market", "label": "市场环境"},
    {"id": "advice", "label": "今日观察"},
    {"id": "strategies", "label": "策略中心"},
    {"id": "ranking", "label": "综合评分"},
    {"id": "backtest", "label": "回测验证"},
    {"id": "boards", "label": "板块分析"},
    {"id": "signals", "label": "信号追踪"},
    {"id": "data", "label": "数据管理"},
    {"id": "logs", "label": "执行日志"},
]

OVERVIEW = {
    "as_of": "演示快照",
    "metrics": [
        {"label": "市场温度", "value": "62", "unit": "/ 100", "trend": "+4.2%"},
        {"label": "活跃策略", "value": "4", "unit": "个", "trend": "2 个观察中"},
        {"label": "候选标的", "value": "18", "unit": "只", "trend": "+6"},
        {"label": "数据覆盖", "value": "98.7", "unit": "%", "trend": "演示"},
    ],
    "market_curve": [
        58,
        61,
        60,
        64,
        63,
        67,
        65,
        66,
        70,
        68,
        69,
        72,
        71,
        67,
        66,
        62,
        64,
        63,
        62,
    ],
    "breadth": {"up": 2731, "flat": 186, "down": 2280},
    "watchlist": [
        {"symbol": "600000.SH", "name": "示例银行", "score": 86, "signal": "趋势延续", "change": 2.31},
        {"symbol": "000001.SZ", "name": "示例金融", "score": 81, "signal": "量价改善", "change": 1.47},
        {"symbol": "300001.SZ", "name": "示例科技", "score": 78, "signal": "回踩观察", "change": -0.62},
        {"symbol": "600519.SH", "name": "示例消费", "score": 74, "signal": "区间震荡", "change": 0.38},
    ],
}

MARKET = {
    "regime": "震荡偏强",
    "confidence": 0.72,
    "dimensions": [
        {"label": "趋势", "score": 68, "status": "改善"},
        {"label": "动量", "score": 61, "status": "中性"},
        {"label": "流动性", "score": 74, "status": "活跃"},
        {"label": "风险", "score": 43, "status": "可控"},
    ],
    "indices": [
        {"name": "上证指数", "value": "3,389.11", "change": 0.42},
        {"name": "深证成指", "value": "10,691.55", "change": 0.77},
        {"name": "创业板指", "value": "2,217.18", "change": 1.06},
    ],
}

STRATEGIES = [
    {
        "id": "trend-observer",
        "name": "趋势观察",
        "status": "active",
        "description": "识别结构完整、动能温和的趋势候选。社区版仅展示策略说明。",
        "last_run": "今日 15:18",
        "candidates": 7,
    },
    {
        "id": "volume-observer",
        "name": "量价观察",
        "status": "active",
        "description": "追踪量能变化与价格确认。阈值、基线和计算逻辑未开源。",
        "last_run": "今日 15:21",
        "candidates": 5,
    },
    {
        "id": "rotation-observer",
        "name": "轮动观察",
        "status": "watching",
        "description": "观察行业与概念的相对强弱变化，仅提供产品界面示例。",
        "last_run": "今日 15:25",
        "candidates": 4,
    },
    {
        "id": "risk-observer",
        "name": "风险过滤",
        "status": "active",
        "description": "对流动性、异常波动和市场环境进行综合过滤。",
        "last_run": "今日 15:27",
        "candidates": 2,
    },
]

RANKINGS = [
    {"rank": 1, "symbol": "600000.SH", "name": "示例银行", "score": 86, "trend": 89, "volume": 82, "risk": 87},
    {"rank": 2, "symbol": "000001.SZ", "name": "示例金融", "score": 81, "trend": 80, "volume": 88, "risk": 76},
    {"rank": 3, "symbol": "300001.SZ", "name": "示例科技", "score": 78, "trend": 84, "volume": 73, "risk": 75},
    {"rank": 4, "symbol": "600519.SH", "name": "示例消费", "score": 74, "trend": 71, "volume": 76, "risk": 79},
    {"rank": 5, "symbol": "000651.SZ", "name": "示例制造", "score": 72, "trend": 75, "volume": 68, "risk": 74},
]

BOARDS = [
    {"name": "数字经济", "heat": 87, "change": 2.84, "leaders": 9},
    {"name": "高端制造", "heat": 79, "change": 1.73, "leaders": 6},
    {"name": "创新药", "heat": 72, "change": 1.12, "leaders": 5},
    {"name": "大消费", "heat": 61, "change": 0.48, "leaders": 3},
    {"name": "金融服务", "heat": 57, "change": -0.21, "leaders": 2},
]

SIGNALS = [
    {"time": "15:27", "level": "focus", "title": "量能结构改善", "target": "000001.SZ", "source": "量价观察"},
    {"time": "15:21", "level": "normal", "title": "趋势结构保持", "target": "600000.SH", "source": "趋势观察"},
    {"time": "14:56", "level": "risk", "title": "波动率抬升", "target": "300001.SZ", "source": "风险过滤"},
    {"time": "14:32", "level": "normal", "title": "板块相对强度上行", "target": "数字经济", "source": "轮动观察"},
]

ACTIVITY = [
    {"time": "15:27:18", "module": "Strategy", "message": "Demo scan completed with 18 candidates"},
    {"time": "15:25:04", "module": "Sectors", "message": "Demo sector snapshot updated"},
    {"time": "15:21:42", "module": "Data", "message": "Demo data integrity check completed"},
    {"time": "15:18:09", "module": "Market", "message": "Demo market regime calculation completed"},
]


def public_payload(value: Any) -> Any:
    """Return isolated demo data so request handlers cannot mutate the catalog."""
    return deepcopy(value)
