from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]


def test_private_capability_modules_do_not_exist() -> None:
    forbidden = {
        "datafeed",
        "strategies",
        "backtest",
        "models",
        "db",
        "scheduler",
        "repositories",
    }
    app_children = {path.name for path in (ROOT / "app").iterdir()}
    assert forbidden.isdisjoint(app_children)


def test_local_environment_is_not_committed() -> None:
    assert not (ROOT / ".env").exists()


def test_frontend_labels_all_catalog_data_as_demo() -> None:
    html = (ROOT / "app/static/index.html").read_text(encoding="utf-8")
    javascript = (ROOT / "app/static/assets/app.js").read_text(encoding="utf-8")
    assert "DEMO" in html
    assert "演示" in javascript
    assert "不构成投资建议" in javascript


def test_readme_language_switches_are_bidirectional() -> None:
    english = (ROOT / "README.md").read_text(encoding="utf-8")
    chinese = (ROOT / "README.zh-CN.md").read_text(encoding="utf-8")
    assert "[简体中文](README.zh-CN.md)" in english
    assert "[English](README.md)" in chinese
