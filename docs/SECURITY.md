# Security Policy

UniPass is committed to security and transparency. We follow a "zero-trust" design model.

## Reporting Vulnerabilities

If you believe you have found a security vulnerability in UniPass, please report it via [GitHub Issues](https://github.com/UniPass/UniPass/issues) using the "Security" label.

## Security Model

*   **Zero-Knowledge:** We do not store or transmit master seeds or credentials to any remote server. Everything remains local.
*   **Memory Hygiene:** We employ `ZeroMemory` techniques to wipe sensitive byte arrays after operations.
*   **EDR/AV Evasion:** Our design patterns minimize behavioral heuristic flags by decoupling operations and avoiding suspicious naming/polling.

## Geolocation Fallbacks

UniPass employs a robust, multi-tier geolocation fallback strategy to ensure high availability and accuracy without compromising local security:

1.  **Native Hardware GPS (Primary):** Accesses built-in hardware GPS via `termux-location` (Android/Termux).
2.  **HTML5 Geolocation (Web Fallback):** Intended for browser-based environments, providing direct client-side coordinate access.
3.  **IP-Based Geo (Tertiary Fallback):** Used only if hardware or HTML5 methods are unavailable, with mandatory trust validation.

