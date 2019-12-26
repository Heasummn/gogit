package ui

import (
	"fmt"
	"sort"
//	"path/filepath"

	"github.com/rivo/tview"
	"github.com/heasummn/gogit/util"
	"gopkg.in/src-d/go-git.v4"
)

type MainPanel struct {
	git *util.GitInfo
	text string
	*tview.TextView
}

// NewMainPanel creates main panel
func NewMainPanel() *MainPanel {
	m := &MainPanel{}
	m.text = ""
	m.git = &util.GitInfo{}
	

	if !util.InitGitInfo(m.git) {
		m.text = "No Git Repo Here! Init Repository below."
	} else {
		m.text = GitInfoToString(m.git)
	}
	
	m.TextView = tview.NewTextView().SetDynamicColors(true).SetText(m.text)
	m.SetBorder(true).SetTitle("Info")
	return m
}

func (m *MainPanel) Refresh() {
	m.text = GitInfoToString(m.git)
	m.TextView.SetText(m.text)

}

type pair struct {
	Key string
	Value *git.FileStatus
}
  
type pairList []pair

func sortByValue(dict git.Status) pairList {
	pl := make(pairList, len(dict))
	i := 0
	for k, v := range dict {
	  pl[i] = pair{k, v}
	  i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}
func (p pairList) Len() int { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value.Staging < p[j].Value.Staging }
func (p pairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

// GitInfoToString converts info to displayable string
func GitInfoToString(g *util.GitInfo) string {
	status := g.GetStaging()
	ret := ""

	flippedStatus := sortByValue(status)

	for _, pair := range flippedStatus {
		file := pair.Key
		val := pair.Value
		if val.Staging == git.Added {
			ret += fmt.Sprintf("[green]added:    \t%s[white]\n", file)
		} else if val.Staging == git.Modified {
			ret += fmt.Sprintf("[green]modified: \t%s[white]\n", file)
		} else if val.Worktree == git.Modified {
			ret += fmt.Sprintf("[red]modified: \t%s[white]\n", file)
		} else if val.Worktree == git.Untracked {
			ret += fmt.Sprintf("[red]untracked:\t%s[white]\n", file)
		} 
	}
	return ret
}