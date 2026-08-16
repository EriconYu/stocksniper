# Security Policy

## Supported version

Security updates are applied to the latest version on the default branch.

## Reporting a vulnerability

Do not disclose credentials or security issues in a public Issue. Contact the author through WeChat `sudo4fate` and include the affected version, reproduction steps, and impact.

## Repository boundary

This repository must never contain production credentials, private market-data integrations, strategy parameters, customer data, database dumps, or internal service URLs. Run `python scripts/security_check.py` before every release.

