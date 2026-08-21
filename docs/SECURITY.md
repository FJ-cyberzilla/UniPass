# Security Policy

UniPass is committed to security and transparency. We follow a "zero-trust" design model.

## Reporting Vulnerabilities

If you believe you have found a security vulnerability in UniPass, please report it via [GitHub Issues](https://github.com/UniPass/UniPass/issues) using the "Security" label.

## Security Model

*   **Zero-Knowledge:** We do not store or transmit master seeds or credentials to any remote server. Everything remains local.
*   **Memory Hygiene:** We employ `ZeroMemory` techniques to wipe sensitive byte arrays after operations.
*   **EDR/AV Evasion:** Our design patterns minimize behavioral heuristic flags by decoupling operations and avoiding suspicious naming/polling.

## Dependencies

We aim to keep external dependencies to an absolute minimum to reduce the attack surface. All external dependencies are audited periodically.
