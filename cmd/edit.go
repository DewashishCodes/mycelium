package cmd

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the Comprehensive Resume Editor",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data, err := os.ReadFile("resume.json")
			if err != nil {
				http.Error(w, "resume.json not found.", 404)
				return
			}
			tmpl, _ := template.New("editor").Parse(editorHTML)
			tmpl.Execute(w, template.HTML(data))
		})

		http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			os.WriteFile("resume.json", body, 0644)
			w.WriteHeader(200)
		})

		fmt.Println("🚀 CVVC Editor: http://localhost:9090")
		http.ListenAndServe(":9090", nil)
	},
}

const editorHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>CVVC Master Editor</title>
    <style>
        :root { --bg: #f8f9fa; --sidebar: #212529; --primary: #0d6efd; --text: #f8f9fa; }
        body { margin: 0; display: flex; height: 100vh; font-family: 'Inter', sans-serif; background: var(--bg); overflow: hidden; }
        .nav-sidebar { width: 80px; background: var(--sidebar); display: flex; flex-direction: column; align-items: center; padding-top: 20px; gap: 20px; border-right: 1px solid #333; }
        .nav-item { color: #888; cursor: pointer; font-size: 20px; transition: 0.2s; }
        .nav-item.active { color: white; }
        .form-panel { width: 450px; background: white; border-right: 1px solid #dee2e6; display: flex; flex-direction: column; }
        .form-header { padding: 20px; border-bottom: 1px solid #eee; display: flex; justify-content: space-between; align-items: center; }
        .form-content { flex: 1; overflow-y: auto; padding: 20px; }
        .input-group { margin-bottom: 15px; }
        label { display: block; font-size: 11px; font-weight: bold; text-transform: uppercase; color: #666; margin-bottom: 5px; }
        input, textarea { width: 100%; padding: 10px; border: 1px solid #ddd; border-radius: 4px; box-sizing: border-box; font-size: 14px; margin-bottom: 10px; }
        .card { border: 1px solid #eee; padding: 15px; border-radius: 8px; margin-bottom: 15px; background: #fafafa; }
        .controls { display: flex; gap: 5px; margin-top: 10px; }
        .btn-sm { padding: 4px 8px; font-size: 11px; border: 1px solid #ccc; background: white; cursor: pointer; border-radius: 3px; }
        .btn-add { background: #e7f3ff; color: #007bff; border: 1px dashed #007bff; width: 100%; padding: 10px; cursor: pointer; margin-top: 10px; border-radius: 5px;}
        .save-btn { background: #198754; color: white; border: none; padding: 8px 20px; border-radius: 4px; cursor: pointer; font-weight: bold; }
        .preview-panel { flex: 1; background: #525659; overflow-y: auto; display: flex; justify-content: center; padding: 40px 0; }
        .paper { background: white; width: 210mm; min-height: 297mm; padding: 50px; box-shadow: 0 0 20px rgba(0,0,0,0.5); font-family: 'Times New Roman', Times, serif; }
        .res-name { text-align: center; font-size: 26pt; border-bottom: 1.5px solid black; padding-bottom: 5px; margin-bottom: 10px; }
        .res-contact { text-align: center; font-size: 11pt; margin-bottom: 10px; }
        .res-sec { font-weight: bold; text-transform: uppercase; border-bottom: 1px solid black; margin-top: 15px; font-size: 12pt; }
        .res-row { display: flex; justify-content: space-between; font-weight: bold; margin-top: 4px; font-size: 11pt; }
        ul { margin: 5px 0; padding-left: 20px; }
        li { font-size: 10.5pt; margin-bottom: 2px; }
    </style>
</head>
<body>
    <div class="nav-sidebar">
        <div class="nav-item active" onclick="tab('basics', this)">👤</div>
        <div class="nav-item" onclick="tab('education', this)">🎓</div>
        <div class="nav-item" onclick="tab('experience', this)">💼</div>
        <div class="nav-item" onclick="tab('projects', this)">🚀</div>
        <div class="nav-item" onclick="tab('skills', this)">🛠️</div>
        <div class="nav-item" onclick="tab('order', this)">🔃</div>
    </div>
    <div class="form-panel">
        <div class="form-header">
            <h3 id="tab-title">Basics</h3>
            <button class="save-btn" onclick="save()">SAVE</button>
        </div>
        <div class="form-content" id="form-area"></div>
    </div>
    <div class="preview-panel">
        <div id="capture-area" class="paper"></div>
    </div>

    <script id="data-raw" type="text/plain">{{.}}</script>

    <script>
        let resume = JSON.parse(document.getElementById('data-raw').textContent);
        if (!resume.sectionOrder) resume.sectionOrder = ['education', 'skills', 'experience', 'projects'];
        let currentTab = 'basics';

        function tab(t, el) {
            currentTab = t;
            document.querySelectorAll('.nav-item').forEach(i => i.classList.remove('active'));
            el.classList.add('active');
            document.getElementById('tab-title').innerText = t.charAt(0).toUpperCase() + t.slice(1);
            renderForm();
        }

        function renderForm() {
            const area = document.getElementById('form-area');
            area.innerHTML = '';

            if (currentTab === 'basics') {
                area.innerHTML = '<label>Full Name</label><input id="inp-name" value="' + resume.basics.name + '">' +
                                 '<label>Email</label><input id="inp-email" value="' + resume.basics.email + '">' +
                                 '<label>Phone</label><input id="inp-phone" value="' + resume.basics.phone + '">';
                
                document.getElementById('inp-name').oninput = (e) => { resume.basics.name = e.target.value; render(); };
                document.getElementById('inp-email').oninput = (e) => { resume.basics.email = e.target.value; render(); };
                document.getElementById('inp-phone').oninput = (e) => { resume.basics.phone = e.target.value; render(); };

            } else if (currentTab === 'experience') {
                resume.experience.forEach((exp, i) => {
                    let card = document.createElement('div');
                    card.className = 'card';
                    card.innerHTML = '<label>Company</label><input class="c-comp" value="' + exp.company + '">' +
                                     '<label>Role</label><input class="c-role" value="' + exp.role + '">' +
                                     '<label>Points (Enter for new line)</label><textarea class="c-pts">' + exp.points.join('\n') + '</textarea>' +
                                     '<div class="controls"><button onclick="remove(\'experience\','+i+')">Delete</button></div>';
                    
                    card.querySelector('.c-comp').oninput = (e) => { resume.experience[i].company = e.target.value; render(); };
                    card.querySelector('.c-role').oninput = (e) => { resume.experience[i].role = e.target.value; render(); };
                    card.querySelector('.c-pts').oninput = (e) => { resume.experience[i].points = e.target.value.split('\n'); render(); };
                    area.appendChild(card);
                });
                let addBtn = document.createElement('button');
                addBtn.className = 'btn-add';
                addBtn.innerText = '+ Add Experience';
                addBtn.onclick = () => { resume.experience.push({company:'', role:'', date:'', points:[]}); renderForm(); render(); };
                area.appendChild(addBtn);

            } else if (currentTab === 'skills') {
                for (let k in resume.skills) {
                    let div = document.createElement('div');
                    div.innerHTML = '<label>'+k+'</label><textarea id="sk-'+k+'">'+resume.skills[k]+'</textarea>';
                    div.querySelector('textarea').oninput = (e) => { resume.skills[k] = e.target.value; render(); };
                    area.appendChild(div);
                }
            } else if (currentTab === 'order') {
                resume.sectionOrder.forEach((sec, i) => {
                    let div = document.createElement('div');
                    div.className = 'card';
                    div.style.display = 'flex';
                    div.style.justifyContent = 'space-between';
                    div.innerHTML = '<span>'+sec.toUpperCase()+'</span>' +
                                    '<div><button onclick="moveOrder('+i+', -1)">Up</button><button onclick="moveOrder('+i+', 1)">Down</button></div>';
                    area.appendChild(div);
                });
            } else {
                area.innerHTML = '<p>Section coming soon or use resume.json</p>';
            }
        }

        function moveOrder(i, dir) {
            let target = i + dir;
            if (target < 0 || target >= resume.sectionOrder.length) return;
            let temp = resume.sectionOrder[i];
            resume.sectionOrder[i] = resume.sectionOrder[target];
            resume.sectionOrder[target] = temp;
            renderForm(); render();
        }

        function remove(arr, i) {
            resume[arr].splice(i, 1);
            renderForm(); render();
        }

        function render() {
            const paper = document.getElementById('capture-area');
            let h = '<div class="res-name">' + resume.basics.name + '</div>';
            h += '<div class="res-contact">' + resume.basics.phone + ' | ' + resume.basics.email + '</div>';

            resume.sectionOrder.forEach(sec => {
                if (sec === 'education' && resume.education.length) {
                    h += '<div class="res-sec">Education</div>';
                    resume.education.forEach(e => { h += '<div class="res-row"><span>'+e.school+'</span><span>'+e.date+'</span></div><div>'+e.degree+'</div>'; });
                } else if (sec === 'skills') {
                    h += '<div class="res-sec">Technical Skills</div>';
                    for (let k in resume.skills) { h += '<div><strong>'+k+':</strong> '+resume.skills[k]+'</div>'; }
                } else if (sec === 'experience') {
                    h += '<div class="res-sec">Experience</div>';
                    resume.experience.forEach(exp => {
                        h += '<div class="res-row"><span>'+exp.company+'</span><span>'+exp.date+'</span></div><div>'+exp.role+'</div>';
                        h += '<ul>' + exp.points.map(p => p ? '<li>'+p+'</li>' : '').join('') + '</ul>';
                    });
                }
            });
            paper.innerHTML = h;
        }

        async function save() {
            await fetch('/save', { method: 'POST', body: JSON.stringify(resume) });
            alert("Saved!");
        }

        renderForm();
        render();
    </script>
</body>
</html>
`
