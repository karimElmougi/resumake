package templates_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/karimElmougi/resumake/internal/resume/templates"
)

func TestPlaintextTemplate(t *testing.T) {
	g := NewGomegaWithT(t)

	censor := false
	g.Expect(func() { templates.Plaintext(&censor) }).ToNot(Panic())
	tmpl := templates.Plaintext(&censor)

	b := &strings.Builder{}
	err := tmpl.Execute(b, testResume)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(b.String()).To(Equal(plaintextResume))
}

var plaintextResume = `John Smith
john.smith@gmail.com
==============================

EDUCATION
==============================
M.S. in Computer Science, Georgia Institute of Technology, Jan. 2004 - Current
GPA: 3.9

B.S. in Computer Science, University of Philadelphia, Jan. 2004 - Dec. 2006

PROFESSIONAL EXPERIENCE
===============================
Microsoft, Seattle, WA, Jan. 2004 - Current
Senior Software Engineer
* did a thing
* did another thing
Technologies used: C#, C++

IBM, Seattle, WA, Jan. 2004 - Dec. 2006
Software Engineer
* did a thing
* did another thing
Technologies used: Java

SAP, Seattle, WA, Winter 2004
Software Engineer Intern
* did a thing
* did another thing
Technologies used: ABAP

SKILLS
==============================
Languages: C++, Java, C#
Technologies: git, Docker

PROJECTS
==============================
Compiler
Compiles stuff
Technologies used: C#, ANTLR, LLVM

Gameboy Emulator
Emulates stuff
Technologies used: C++
`
