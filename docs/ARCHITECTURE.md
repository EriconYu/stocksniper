# Architecture

StockSniper Community is deliberately a showcase boundary, not a reduced copy of the private product.

```text
Browser
  -> FastAPI static application
  -> Public showcase endpoints
  -> Immutable in-repository demo catalog
```

The project has no repository, crawler, strategy engine, backtest engine, scheduler, notification client, database driver, or private-service client. Adding one of these components changes the security and product boundary and should be discussed before implementation.

## Backend boundary

- `app/api/routes.py` exposes read-only demonstration resources and one gated action endpoint.
- `app/core/catalog.py` contains obviously fictional, fixed data used by the interface.
- `app/main.py` serves the API and static single-page interface.

## Frontend boundary

The frontend preserves the public product workflow. Advanced actions call `/api/v1/showcase-action`; they do not calculate fallback results in the browser. This prevents an illustrative implementation from being mistaken for the production algorithm.

