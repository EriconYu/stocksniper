from __future__ import annotations

from typing import Literal

from fastapi import APIRouter
from pydantic import BaseModel, Field

from app.core.catalog import (
    ACTIVITY,
    BOARDS,
    MARKET,
    NAVIGATION,
    OVERVIEW,
    PRODUCT,
    RANKINGS,
    SIGNALS,
    STRATEGIES,
    public_payload,
)

router = APIRouter(prefix="/api/v1")


class ShowcaseAction(BaseModel):
    feature: Literal["backtest", "data-sync", "strategy-run", "export", "trade-assistant"]
    symbol: str | None = Field(default=None, max_length=32)


@router.get("/meta")
def meta() -> dict:
    return {"product": public_payload(PRODUCT), "navigation": public_payload(NAVIGATION)}


@router.get("/overview")
def overview() -> dict:
    return public_payload(OVERVIEW)


@router.get("/market")
def market() -> dict:
    return public_payload(MARKET)


@router.get("/strategies")
def strategies() -> list[dict]:
    return public_payload(STRATEGIES)


@router.get("/rankings")
def rankings() -> list[dict]:
    return public_payload(RANKINGS)


@router.get("/boards")
def boards() -> list[dict]:
    return public_payload(BOARDS)


@router.get("/signals")
def signals() -> list[dict]:
    return public_payload(SIGNALS)


@router.get("/activity")
def activity() -> list[dict]:
    return public_payload(ACTIVITY)


@router.post("/showcase-action")
def showcase_action(action: ShowcaseAction) -> dict:
    return {
        "available": False,
        "feature": action.feature,
        "message": "该能力属于完整产品，社区展示版不包含核心实现。",
        "official_url": PRODUCT["official_url"],
        "wechat": PRODUCT["wechat"],
    }
