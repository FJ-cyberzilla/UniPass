# UniPass User Guide

Welcome to the **UniPass** documentation portal. This guide provides comprehensive instructions on installation, usage, project architecture, and developer workflows for the UniPass CLI utility.

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [Project Documentation & Resources](#2-project-documentation--resources)
3. [Prerequisites](#3-prerequisites)
4. [Installation & Building](#4-installation--building)
5. [Usage Overview](#5-usage-overview)
6. [Makefile Commands](#6-makefile-commands)
7. [Security & Compliance](#7-security--compliance)
8. [Architecture & Data Flow](#8-architecture--data-flow)

---

## 1. Introduction
UniPass is a zero-knowledge, cross-platform CLI cryptographic utility designed with a strict security model that avoids passive naming and beaconing. It generates entropy locally using astronomical and system-level constants, avoiding external network dependencies during credential generation.

## 2. Project Documentation & Resources

The repository contains several key files to ensure transparency, security, and maintainability:

*   **`README.md`**: The project's primary entry point. Provides a high-level summary, installation instructions, and feature overview.
*   **`docs/DOCUMENTS.md`**: Deep dive into the "Anti-Heuristic" architecture, cryptographic engineering, and the astronomical entropy generation sources.
*   **`SECURITY.md`**: Outlines our vulnerability disclosure policy and security model commitments.
*   **`CONTRIBUTORS.md`**: Credits and guidelines for contributing to UniPass.
*   **`LICENSE`**: Details the project's legal usage rights.
*   **`CHANGELOG`**: A historical record of changes, fixes, and feature updates.
*   **`DISCLAIMER`**: Essential warnings regarding the use of this software, including "as-is" clauses.

## 3. Prerequisites
To build or develop UniPass, ensure you have the following installed:
*   **Go** (v1.20 or later)
*   **Make**
*   **Git**

## 4. Installation & Building
Clone the repository and use the included `Makefile` to compile the application:

```bash
git clone <repository_url>
cd UniPass
make build
```
The resulting binary will be located in the `dist/` directory.

## 5. Usage Overview
Once built, you can execute the utility directly:

```bash
./dist/unipass
```
*(Refer to individual command help outputs via `--help` for specific flags and subcommands.)*

## 6. Makefile Commands
The `Makefile` simplifies the development lifecycle. Use the following commands:

| Command | Description |
| :--- | :--- |
| `make build` | Compiles the UniPass binary into the `dist/` directory. |
| `make test` | Executes the full unit test suite across all packages. |
| `make check` | Runs `vet`, `fmt`, and `test` sequentially (Static Analysis). |
| `make run-dev` | Builds and immediately launches the CLI for local testing. |
| `make clean` | Removes build artifacts and cleans the `dist/` directory. |
| `make release` | Runs the full `build.sh` release script. |

## 7. Security & Compliance
UniPass is designed with a "zero-trust" approach to its own environment. Please review `SECURITY.md` for details on:
*   **Memory Hygiene:** How we prevent sensitive data leakage.
*   **EDR/AV Evasion:** Our design patterns for preventing false-positive flagging.
*   **Vulnerability Reporting:** Instructions on how to securely report potential issues.

## 8. Architecture & Data Flow
The `docs/` folder contains several Mermaid (`.mmd`) diagrams mapping the internal logic:

*   **`ARCHITECTURE.mmd`**: High-level component relationship map.
*   **`DATAFLOW.mmd`**: How entropy moves from sources to the vault.
*   **`ENTROPY_GENERATION.mmd`**: The logic behind astronomical seed generation.
*   **`INTERACTION_SEQUENCE.mmd`**: The sequence of events during a standard command execution.

---
*For further assistance, please consult the project's `README.md` or contact the maintainers.*
