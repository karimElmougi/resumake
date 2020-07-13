package templates_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"resumake/internal/resume/templates"
)

func TestPlaintextTemplate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() { templates.Plaintext() }).ToNot(Panic())
	tmpl := templates.Plaintext()

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
Senior Software Engineer
Microsoft, Seattle, WA, Jan. 2004 - Current
* did a thing
* did another thing
Technologies used: C#, C++

Software Engineer
IBM, Seattle, WA, Jan. 2004 - Dec. 2006
* did a thing
* did another thing
Technologies used: Java

Software Engineer Intern
SAP, Seattle, WA, Winter 2004
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
