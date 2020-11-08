package fse

import (
	"path/filepath"
	"strings"
)

type Path struct {
	patterns []string
}

func newPath(path string) *Path {
	return &Path{patterns: strings.Split(path, string(filepath.Separator))}
}

func (p *Path) StepBack() {
	sz := len(p.patterns)
	if sz > 1 {
		p.patterns = p.patterns[:sz-1]
	}
}

func (p *Path) StepForward(pattern string) {
	p.patterns = append(p.patterns, pattern)
}

func (p *Path) String() string {
	sz := len(p.patterns)
	var buf strings.Builder
	for idx, pattern := range p.patterns {
		buf.WriteString(pattern)
		if idx == sz-1 {
			break
		}
		buf.WriteRune(filepath.Separator)
	}
	return buf.String()
}
