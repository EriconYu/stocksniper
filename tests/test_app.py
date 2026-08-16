from fastapi.testclient import TestClient

from app.main import create_app

client = TestClient(create_app())


def test_health_identifies_showcase_edition() -> None:
    response = client.get("/api/health")
    assert response.status_code == 200
    assert response.json() == {
        "status": "ok",
        "edition": "community-showcase",
        "version": "2.0.0",
    }


def test_meta_exposes_product_and_navigation() -> None:
    response = client.get("/api/v1/meta")
    assert response.status_code == 200
    payload = response.json()
    assert payload["product"]["official_url"] == "https://facai.stackbang.com"
    assert payload["product"]["wechat"] == "sudo4fate"
    assert len(payload["navigation"]) == 10


def test_public_catalog_endpoints_return_demo_data() -> None:
    for endpoint in ("overview", "market", "strategies", "rankings", "boards", "signals", "activity"):
        response = client.get(f"/api/v1/{endpoint}")
        assert response.status_code == 200
        assert response.json()


def test_activity_logs_are_in_english() -> None:
    activity = client.get("/api/v1/activity").json()
    assert activity
    assert all(item["module"].isascii() for item in activity)
    assert all(item["message"].isascii() for item in activity)


def test_advanced_action_is_explicitly_gated() -> None:
    response = client.post(
        "/api/v1/showcase-action",
        json={"feature": "backtest", "symbol": "600000.SH"},
    )
    assert response.status_code == 200
    payload = response.json()
    assert payload["available"] is False
    assert payload["feature"] == "backtest"
    assert payload["official_url"] == "https://facai.stackbang.com"


def test_advanced_action_rejects_unknown_feature() -> None:
    response = client.post("/api/v1/showcase-action", json={"feature": "private-engine"})
    assert response.status_code == 422


def test_static_app_and_qr_asset_are_served() -> None:
    index = client.get("/")
    assert index.status_code == 200
    assert "StockSniper Community" in index.text
    assert "sudo4fate" in index.text

    qr_code = client.get("/assets/wechat-qr.jpg")
    assert qr_code.status_code == 200
    assert qr_code.headers["content-type"] == "image/jpeg"
    assert len(qr_code.content) > 10_000


def test_unknown_api_path_stays_json_404() -> None:
    response = client.get("/api/v1/not-found")
    assert response.status_code == 404
    assert response.json() == {"detail": "API endpoint not found"}


def test_spa_path_returns_index() -> None:
    response = client.get("/some/future/page")
    assert response.status_code == 200
    assert "StockSniper Community" in response.text
