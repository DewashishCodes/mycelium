package cmd

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize CVVC with Dewashish's AI/ML Template",
	Run: func(cmd *cobra.Command, args []string) {
		git.PlainInit(".", false)

		content := `{
  "basics": {
    "name": "Dewashish Lambore",
    "email": "dewashish.lambore@gmail.com",
    "phone": "+91 9307059152",
    "linkedin": "linkedin.com/in/dewashish",
    "github": "github.com/dewashish"
  },
  "education": [
    {
      "school": "Symbiosis Institute of Technology, Pune",
      "degree": "Bachelor of Technology in Electronics and Telecommunication",
      "date": "2024-2028",
      "cgpa": "8.40"
    }
  ],
  "skills": {
    "Languages": "Python, JavaScript, C++, C, SQL",
    "AI Tools": "Fine tuning, RAG, Supervised/Unsupervised Learning",
    "Libraries": "Numpy, Pandas, Scikitlearn, React, Tailwind, PyTorch",
    "Development": "Docker, Kubernetes, Github Actions"
  },
  "certifications": [
    "Oracle Cloud Infrastructure 2025 Certified Generative AI Professional",
    "Advanced Learning Algorithms – DeepLearning.AI",
    "MATLAB Machine Learning Onramp – MathWorks"
  ],
  "experience": [
    {
      "company": "SCAAI Pune",
      "role": "AI Intern",
      "date": "July 2025-Present",
      "points": [
        "Developing AI-driven FinTech predictive modeling solutions",
        "Collaborating to deploy compliant (AML, KYC) production-ready models"
      ]
    }
  ],
  "projects": [
    {
      "name": "Resilient Multi-Modal Agentic RAG System",
      "tech": "Python, LangChain, FastAPI",
      "points": [
        "Engineered RAG pipeline with Hybrid Search + GPU reranking, achieving 90% QA accuracy",
        "Cut processing time by 86% (3 mins to 25 seconds)"
      ]
    }
  ],
  "achievements": [
    "Top 180 out of 46,178 in Bajaj Finserv Hackrx 6.0 Runner-up",
    "2Fast2Hack Hackathon Runner-up (600+ participants)"
  ]
}`
		os.WriteFile("resume.json", []byte(content), 0644)
		os.WriteFile(".gitignore", []byte("*.pdf\ncvvc.exe\n"), 0644)
		fmt.Println("✅ CVVC Initialized with your AI/ML profile.")
	},
}
