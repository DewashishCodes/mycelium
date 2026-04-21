package cmd

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mycelium/internal/resume"
	"net/http"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/spf13/cobra"
)

//go:embed templates/classic.html
var classicTemplateContent string

func init() {
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export your resume to a high-quality PDF",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Read Data
		res, err := resume.Read()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		fmt.Printf("[INFO] Generating professional PDF for %s...\n", res.Basics.Name)

		// 2. Setup Temporary HTTP Server for PDF Generation
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.New("resume").Parse(classicTemplateContent)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			tmpl.Execute(w, res)
		})

		server := &http.Server{Addr: ":7331", Handler: mux}
		go server.ListenAndServe()

		// 3. Use Rod to capture PDF
		// Wait briefly for server to start
		time.Sleep(500 * time.Millisecond)
		
		browser := rod.New().MustConnect()
		defer browser.MustClose()
		
		page := browser.MustPage("http://localhost:7331")
		page.MustWaitLoad()

		// High Quality PDF Config
		pdfStream, err := page.PDF(&proto.PagePrintToPDF{
			PrintBackground: true,
			PaperWidth:      toPtr(8.27), // A4
			PaperHeight:     toPtr(11.69),
		})

		if err != nil {
			fmt.Println("[ERROR] PDF Rendering failed:", err)
			server.Close()
			return
		}

		pdfData, _ := io.ReadAll(pdfStream)

		// 4. Cleanup & Save
		server.Close()

		filename := fmt.Sprintf("%s_Resume.pdf", resume.SafeFilename(res.Basics.Name))
		os.WriteFile(filename, pdfData, 0644)

		fmt.Printf("[SUCCESS] PDF Network Exported: %s\n", filename)
		fmt.Printf("[INFO] Timestamp: %s\n", time.Now().Format("15:04:05"))
	},
}

// Helper for Rod library
func toPtr(f float64) *float64 { return &f }



