package cmd

import (
	"context"
	"fmt"
	"mycelium/internal/resume"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.Flags().StringP("role", "r", "Software Engineer", "The role you are targeting")
	writeCmd.Flags().StringP("company", "c", "", "The company name")
	writeCmd.Flags().StringP("for", "f", "experience", "Section to generate for (experience, projects)")
}

var writeCmd = &cobra.Command{
	Use:   "write [raw-bullet]",
	Short: "AI-powered bullet point generator to strengthen your resume",
	Args:  cobra.MinimumArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := getAPIKey()
		if apiKey == "" {
			fmt.Println("[ERROR] API Key not found. Run: mycelium config --key YOUR_KEY")
			return
		}

		rawBullet := strings.Join(args, " ")
		targetRole, _ := cmd.Flags().GetString("role")
		company, _ := cmd.Flags().GetString("company")
		targetSection, _ := cmd.Flags().GetString("for")

		fmt.Printf("[AI] Strengthening bullet for [%s] role...\n", targetRole)

		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			fmt.Println("[ERROR] AI Connection failed:", err)
			return
		}
		defer client.Close()

		model := client.GenerativeModel("gemini-1.5-flash")

		prompt := fmt.Sprintf(`
			Act as a world-class resume writer.
			I have a raw bullet point: "%s"
			
			Please rewrite this bullet point to be extremely professional, high-impact, and metrics-driven.
			Target Role: %s
			Target Company: %s
			Section: %s
			
			Guidelines:
			- Use strong action verbs (e.g., Spearheaded, Engineered, Optimized).
			- Quantify the impact wherever possible (use placeholders like "[X]%%" or "$[Y]M" if not provided).
			- Keep it concise and professional.
			
			Provide exactly 3 variations of the strengthened bullet point, each on a new line.
			Do not add any other text.
		`, rawBullet, targetRole, company, targetSection)

		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			fmt.Println("[ERROR] AI Generation failed:", err)
			return
		}

		fmt.Println("\n✨ STRENGTHENED VARIATIONS:")
		fmt.Println("---------------------------")
		for _, cand := range resp.Candidates {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
		fmt.Println("---------------------------")
		fmt.Println("[INFO] Copy your favorite and add it to your resume using 'mycelium edit'.")
	},
}
