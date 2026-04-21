package cmd

import (
	"fmt"
	"html/template"
	"io"
	"mycelium/internal/resume"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Generate a professional PDF of your resume",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := resume.Read()
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}

		// 1. Start temporary server for the PDF engine
		// Use a local mux to avoid conflicts with other commands
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl, _ := template.New("p").Parse(pdfHTML)
			tmpl.Execute(w, res)
		})

		// Find an available port if 9091 is taken, or just use 9091 with a listener
		listener, err := net.Listen("tcp", ":9091")
		if err != nil {
			fmt.Println("[ERROR] Could not start PDF engine server on :9091:", err)
			return
		}
		defer listener.Close()

		go http.Serve(listener, mux)

		fmt.Println("⏳ Starting PDF generation...")
		time.Sleep(1 * time.Second)

		// 2. Launch Local Browser
		path, _ := launcher.LookPath()
		if path == "" {
			fmt.Println("[ERROR] Error: Could not find Chrome or Edge.")
			return
		}

		u := launcher.New().Bin(path).Leakless(false).MustLaunch()
		browser := rod.New().ControlURL(u).MustConnect()
		defer browser.MustClose()

		page := browser.MustPage("http://localhost:9091")
		page.MustWaitLoad()

		// 3. Capture PDF Stream
		fmt.Println("[INFO] Rendering PDF...")
		pdfStream, err := page.PDF(&proto.PagePrintToPDF{
			PrintBackground: true,
			PaperWidth:      toPtr(8.27),
			PaperHeight:     toPtr(11.69),
			MarginTop:       toPtr(0),
			MarginBottom:    toPtr(0),
			MarginLeft:      toPtr(0),
			MarginRight:     toPtr(0),
		})

		if err != nil {
			fmt.Println("[ERROR] Error rendering:", err)
			return
		}

		// 4. Convert Stream to Bytes
		pdfBytes, err := io.ReadAll(pdfStream)
		if err != nil {
			fmt.Println("[ERROR] Error reading stream:", err)
			return
		}

		// 5. Save file
		outputName := fmt.Sprintf("%s_Resume.pdf", resume.SafeFilename(res.Basics.Name))
		err = os.WriteFile(outputName, pdfBytes, 0644)
		if err != nil {
			fmt.Println("[ERROR] Error saving file:", err)
			return
		}

		fmt.Printf("[SUCCESS] Success! Exported to %s\n", outputName)
	},
}

// Helper for Rod library
func toPtr(f float64) *float64 { return &f }

// --- THE CSS/HTML DESIGN ---
const pdfHTML = `
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: "Times New Roman", serif; padding: 45px; line-height: 1.15; color: black; }
        .name { text-align: center; font-size: 28pt; margin: 0; }
        .contact { text-align: center; font-size: 11pt; border-bottom: 1.5px solid black; padding-bottom: 6px; margin-bottom: 10px; }
        .section { font-weight: bold; text-transform: uppercase; border-bottom: 1.5px solid black; margin-top: 15px; font-size: 12.5pt; }
        .row { display: flex; justify-content: space-between; font-weight: bold; margin-top: 5px; font-size: 11pt; }
        .sub-row { display: flex; justify-content: space-between; font-style: italic; font-size: 10.5pt; }
        ul { margin: 4px 0; padding-left: 18px; }
        li { margin-bottom: 1.5px; font-size: 10.5pt; text-align: justify; }
    </style>
</head>
<body>
    <div class="name">{{.Basics.Name}}</div>
    <div class="contact">
        {{if .Basics.Phone}}{{.Basics.Phone}} | {{end}}
        {{if .Basics.Email}}{{.Basics.Email}} | {{end}}
        {{if .Basics.LinkedIn}}{{.Basics.LinkedIn}} | {{end}}
        {{if .Basics.GitHub}}{{.Basics.GitHub}}{{end}}
    </div>
    <div class="section">Education</div>
    {{range .Education}}<div class="row"><span>{{.School}}</span><span>{{.Date}}</span></div><div class="sub-row"><span>{{.Degree}}</span><span>{{if .CGPA}}CGPA: {{.CGPA}}{{end}}</span></div>{{end}}
    <div class="section">Technical Skills</div>
    <div style="font-size: 10.5pt; margin-top: 5px;">{{range $cat, $val := .Skills}}<strong>{{$cat}}:</strong> {{$val}}<br>{{end}}</div>
    <div class="section">Experience</div>
    {{range .Experience}}<div class="row"><span>{{.Company}}</span><span>{{.Date}}</span></div><div class="sub-row"><span>{{.Role}}</span></div><ul>{{range .Points}}<li>{{.}}</li>{{end}}</ul>{{end}}
    <div class="section">Projects</div>
    {{range .Projects}}<div class="row"><span>{{.Name}} | <span style="font-weight:normal; font-style:italic;">{{.Tech}}</span></span></div><ul>{{range .Points}}<li>{{.}}</li>{{end}}</ul>{{end}}
</body>
</html>`

