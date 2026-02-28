# Mycelium
**The Version control system for you Resumes**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green)](./LICENSE)
[![Build Status](https://img.shields.io/badge/Status-Stable-success)](#)

Mycelium is a domain-specific version control system (VCS) designed to manage the lifecycle of professional resumes. It treats career data as structured infrastructure (JSON), enabling high-fidelity branching, semantic versioning, and automated PDF compilation.

---

## Overview
Modern technical recruitment requires extreme specialization. Maintaining separate files for "Machine Learning Engineer," "Backend Developer," and "Data Scientist" roles often leads to fragmented data where updates in one version are lost in others. 

Mycelium solves this by applying Git-based branching logic to a centralized career record. Users can maintain a master history while branching out specialized iterations for specific roles, ensuring all versions remain synchronized and traceable.

## Features
- **Branch-Based Specialization:** Use the `branch` command to create role-specific versions without polluting your master history.
- **Live Form Dashboard:** A built-in web-based editor with real-time browser preview. No manual JSON editing required.
- **Intelligence Layer:** Integrated Gemini-1.5-Flash AI to provide professional recruiter audits tailored to specific job titles.
- **Semantic Diff Engine:** Mycelium understands resume structures, reporting changes in specific fields like "Role," "Company," or "Skills" rather than raw line additions.
- **Headless PDF Production:** Automated PDF generation via a headless browser engine, ensuring a consistent and professional aesthetic.

---

## Quick Start

### 1. Installation
Install the Mycelium binary globally via Go:
```bash
go install github.com/DewashishCodes/mycelium@latest
```

### 2. Configuration (AI Layer)
Configure the intelligence layer using your Gemini API key:
```bash
mycelium config --key YOUR_GEMINI_API_KEY
```

### 3. Initialize a Network
```bash
mkdir my-career && cd my-career
mycelium init
```

### 4. Edit and Audit
```bash
mycelium edit                        # Launch the local dashboard
mycelium review --role="AI Engineer" # Get a targeted technical audit
mycelium export                      # Generate the production PDF
```

---

## Documentation
For detailed information regarding the architecture, VCS implementation, and the intelligence layer, please refer to the **[Technical Specification (TECHNICAL.md)](./TECHNICAL.md)**.

## License
Distributed under the MIT License. See `LICENSE` for more information.

---
**Maintained by [Dewashish Lambore](https://github.com/DewashishCodes)**
