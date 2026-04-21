package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"mycelium/internal/resume"
	"net/http"

	"github.com/spf13/cobra"
)

//go:embed templates/editor.html
var editorTemplateContent string

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the Mycelium Live Form Editor",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := resume.ReadRaw()
			if err != nil {
				http.Error(w, "[ERROR] resume.json not found. Run 'mycelium init' first.", 404)
				return
			}
			tmpl, _ := template.New("editor").Parse(editorTemplateContent)
			// Pass as template.HTML to preserve JSON quotes
			tmpl.Execute(w, template.HTML(data))
		})

		http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				return
			}
			body, _ := io.ReadAll(r.Body)
			
			// Try to unmarshal to validate before saving
			var temp resume.Resume
			if err := json.Unmarshal(body, &temp); err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Invalid JSON: %v", err)
				return
			}

			err := resume.Write(&temp)
			if err != nil {
				w.WriteHeader(500)
				return
			}
			w.WriteHeader(200)
		})

		fmt.Println("[INFO] Mycelium Editor started.")
		fmt.Println("[INFO] Local Network Link: http://localhost:9090")
		fmt.Println("[INFO] Press Ctrl+C to disconnect from the network.")
		http.ListenAndServe(":9090", nil)
	},
}
