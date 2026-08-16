from __future__ import annotations

import os
from pathlib import Path

import uvicorn
from fastapi import FastAPI, Request
from fastapi.responses import FileResponse, JSONResponse
from fastapi.staticfiles import StaticFiles

from app import __version__
from app.api.routes import router

BASE_DIR = Path(__file__).resolve().parent
STATIC_DIR = BASE_DIR / "static"


def create_app() -> FastAPI:
    application = FastAPI(
        title="StockSniper Community API",
        description="Public showcase API. No market crawler, strategy engine, or credentials included.",
        version=__version__,
        docs_url="/api/docs",
        redoc_url=None,
        openapi_url="/api/openapi.json",
    )
    application.include_router(router)
    application.mount("/assets", StaticFiles(directory=STATIC_DIR / "assets"), name="assets")

    @application.middleware("http")
    async def security_headers(request: Request, call_next):
        response = await call_next(request)
        response.headers["X-Content-Type-Options"] = "nosniff"
        response.headers["X-Frame-Options"] = "DENY"
        response.headers["Referrer-Policy"] = "strict-origin-when-cross-origin"
        response.headers["Permissions-Policy"] = "camera=(), microphone=(), geolocation=()"
        return response

    @application.get("/api/health", include_in_schema=False)
    def health() -> dict:
        return {"status": "ok", "edition": "community-showcase", "version": __version__}

    @application.exception_handler(404)
    async def not_found(request: Request, _exc: Exception):
        if request.url.path.startswith("/api/"):
            return JSONResponse({"detail": "API endpoint not found"}, status_code=404)
        return FileResponse(STATIC_DIR / "index.html")

    @application.get("/", include_in_schema=False)
    def index() -> FileResponse:
        return FileResponse(STATIC_DIR / "index.html")

    return application


app = create_app()


def run() -> None:
    host = os.getenv("STOCKSNIPER_HOST", "127.0.0.1")
    port = int(os.getenv("STOCKSNIPER_PORT", "8000"))
    reload_enabled = os.getenv("STOCKSNIPER_RELOAD", "false").lower() == "true"
    uvicorn.run("app.main:app", host=host, port=port, reload=reload_enabled)


if __name__ == "__main__":
    run()
