<div align='center'>
<img width="100" height="100" alt="WhatsApp_Image_2026-02-28_at_22 40 34-removebg-preview" src="https://github.com/user-attachments/assets/af959cbf-a738-4d9e-955e-3ff0d0e6baf5" />

# 🍄 Mycelium: The Resume Versioning Network

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Release](https://img.shields.io/github/v/release/DewashishCodes/mycelium)
![License](https://img.shields.io/badge/License-MIT-green)

Mycelium is a version control system designed to manage the lifecycle of professional resumes. Instead of tracking binary files, Mycelium tracks structured career data (JSON), allowing for semantic versioning, branching for specific job roles, and automated high-fidelity PDF generation.
</div>

## 🚀 Installation

### 1. The Quickest Way (Download Binary)
No coding knowledge required.
1. Go to the [Latest Releases](https://github.com/DewashishCodes/mycelium/releases/latest).
2. Download the `.zip` or `.tar.gz` file for your OS (e.g., `mycelium_Windows_x86_64.zip`).
3. Extract the `mycelium.exe` and move it to a folder in your **System PATH**.
4. Open your terminal and type `mycelium version`.

### 2. For Developers (Go installed)
```bash
go install github.com/DewashishCodes/mycelium@latest
```

---

## 🛠️ Quick Start

1. **Seed the Network**: Create a new folder and run `mycelium init`. This creates a professional John Doe template.
2. **Launch the Dashboard**: Run `mycelium edit`. A live form-based editor will open in your browser.
3. **Audit Your Profile**: Run `mycelium review --role="AI Engineer"`. Our integrated Gemini AI will critique your resume.
4. **Tailor with Branches**: Use `mycelium branch create company-name` to create specific versions of your CV.
5. **Deploy**: Run `mycelium export` to generate a perfectly formatted PDF.

---

## 🧠 Core Features

- **Semantic Diff**: Run `mycelium diff` to see human-readable changes between versions (e.g., "Changed Role from X to Y") instead of messy code lines.
- **Time Travel**: Use `mycelium restore <hash>` to instantly revert your resume to any previous state in your history.
- **Role-Specific Intelligence**: Deep integration with Google Gemini-1.5-Flash to provide specialized technical audits.

---

**1. Getting Started (The Onboarding)**
- **Installation (Non-Devs):** Step-by-step for Windows/Mac/Linux. Download from Releases -> Extract -> Add to PATH.
- **Installation (Go Devs):** Use `go install github.com/DewashishCodes/mycelium@latest`.
- **First Run:** Running `mycelium init` to bootstrap the network with the 'John Doe' professional template.

**2. The Management Suite (Editing & Exporting)**
- **`mycelium edit`:** Explain the visual dashboard. Mention it runs on `localhost:9090`, features a live split-screen preview, and auto-syncs with `resume.json`.
- **`mycelium export`:** Explain the headless browser orchestration. Mention it generates a pixel-perfect PDF using Jake's Resume format.

**3. Version Control (The Time-Machine)**
- **`mycelium commit -m "msg"`:** Saving a snapshot of the current state.
- **`mycelium status` & `list`:** Tracking the active branch and viewing the 40-character commit history.
- **`mycelium branch [create/switch]`:** The logic of specializing resumes. Explain the use case: creating a 'frontend-role' branch vs a 'backend-role' branch.
- **`mycelium diff`:** Explain the 'Semantic Diff' engine. Contrast it with raw Git diffs—show how Mycelium understands that a 'Role' changed, not just a line of text.
- **`mycelium restore [hash] --force`:** The 'Time Travel' command. Explain how to revert a resume to any point in the history.
- **`mycelium sync [branch]`:** The rebase logic. Explain how to pull updates from 'main' into a specialized branch.

**4. Intelligence Layer (AI Audit)**
- **`mycelium config --key [key]`:** Guide on obtaining a Google Gemini API key and storing it locally in `~/.cvvc_config.json`.
- **`mycelium review --role="[title]"`:** Detail the AI Audit feature. Explain how it provides a Role-Match score, missing keywords, and bullet point strengthening.

**5. Advanced Configuration**
- **Manual JSON Editing:** Explaining the JSON schema (Basics, Education, Experience, Skills, Projects).
- **The .gitignore File:** Why Mycelium ignores generated PDFs to keep the history clean.

---

## 🤝 Contributing & Feedback

The Mycelium network grows through community input! 

- **Found a Bug?** Open an [Issue](https://github.com/DewashishCodes/mycelium/issues). I will review and work on them actively.
- **Have a Feature Idea?** We would love to hear about new templates or AI features.
- **Contributing Code**: 
    1. Fork the repo.
    2. Create your feature branch.
    3. Open a Pull Request.
    *Assigning issues is not compulsory—feel free to pick up any open issue and start building!*

---

## ⚖️ License
Distributed under the MIT License. 

*Built with ❤️ by [Dewashish Lambore](https://github.com/DewashishCodes)*
**This shows recruiters you aren't just a coder—you are a Project Lead.** 

**Ready to update the README and open the network?** 🍄🚀
