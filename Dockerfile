FROM python:3.12-slim AS runtime

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    STOCKSNIPER_HOST=0.0.0.0 \
    STOCKSNIPER_PORT=8000

WORKDIR /app

RUN addgroup --system stocksniper && adduser --system --ingroup stocksniper stocksniper

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY app ./app
COPY pyproject.toml README.md LICENSE ./

USER stocksniper

EXPOSE 8000
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD python -c "import urllib.request; urllib.request.urlopen('http://127.0.0.1:8000/api/health', timeout=2)"

CMD ["python", "-m", "app.main"]
