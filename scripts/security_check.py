from __future__ import annotations

import re
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SELF = Path(__file__).resolve()
BLOCKED_NAMES = {".env", "id_rsa", "id_ed25519"}
BLOCKED_SUFFIXES = {".pem", ".p12", ".pfx"}
TEXT_SUFFIXES = {
    "",
    ".css",
    ".dockerignore",
    ".html",
    ".ini",
    ".js",
    ".json",
    ".md",
    ".py",
    ".toml",
    ".txt",
    ".yaml",
    ".yml",
}
PATTERNS = [
    re.compile(r"(?i)(api[_-]?key|secret|password|token)\s*[:=]\s*['\"][^'\"]{8,}['\"]"),
    re.compile(r"-----BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY-----"),
    re.compile(r"(?i)mysql(?:\+\w+)?://[^\s:/]+:[^\s@]+@"),
]


def tracked_files() -> list[Path]:
    result = subprocess.run(
        ["git", "ls-files", "--cached", "--others", "--exclude-standard"],
        cwd=ROOT,
        check=True,
        capture_output=True,
        text=True,
    )
    return [ROOT / line for line in result.stdout.splitlines() if line]


def main() -> int:
    findings: list[str] = []
    for path in tracked_files():
        if path.resolve() == SELF or not path.is_file():
            continue
        if path.name in BLOCKED_NAMES or path.suffix.lower() in BLOCKED_SUFFIXES:
            findings.append(f"blocked sensitive filename: {path.relative_to(ROOT)}")
            continue
        if path.suffix.lower() not in TEXT_SUFFIXES:
            continue
        try:
            text = path.read_text(encoding="utf-8")
        except UnicodeDecodeError:
            continue
        for pattern in PATTERNS:
            if pattern.search(text):
                findings.append(f"possible hard-coded credential: {path.relative_to(ROOT)}")
                break

    if findings:
        print("Security check failed:")
        for finding in findings:
            print(f"- {finding}")
        return 1

    print("Security check passed: no blocked files or credential patterns found.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
