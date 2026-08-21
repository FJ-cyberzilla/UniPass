# UniPass

UniPass is a zero-knowledge, cross-platform CLI cryptographic utility designed with a strict "no-passive-naming" and "zero-beaconing" security model. It derives high-precision entropy locally using astronomical and system-level constants, ensuring the application behaves entirely like a benign mathematical utility.

---

## Key Features

*   **Anti-Heuristic Architecture:** Designed to evade detection by EDR/AV solutions.
*   **Local Entropy Generation:** No network requests during credential generation; utilizes local system state and astronomical calculations.
*   **Zero-Beaconing:** Does not poll sensors or network endpoints.
*   **Secure Storage:** AES-GCM encrypted vault with memory hygiene.

## Getting Started

### Prerequisites
*   Go (v1.20+)
*   Make

### Building
```bash
make build
```

The binary will be located in `dist/unipass`.

## Documentation
*   [User Guide](USERGUIDE.md)
*   [Engineering Documents](docs/DOCUMENTS.md)
*   [Security Policy](SECURITY.md)

---
*UniPass is intended for security-conscious users. Always review the source and use it responsibly.*
