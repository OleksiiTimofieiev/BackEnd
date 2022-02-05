package regtest

import (
	"testing"

	"golang.org/x/tools/internal/lsp/tests"
)

const unformattedProgram = `
-- main.go --
package main
import "fmt"
func main(  ) {
	fmt.Println("Hello World.")
}
-- main.go.golden --
package main

import "fmt"

func main() {
	fmt.Println("Hello World.")
}
`

func TestFormatting(t *testing.T) {
	runner.Run(t, unformattedProgram, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		env.FormatBuffer("main.go")
		got := env.Editor.BufferText("main.go")
		want := env.ReadWorkspaceFile("main.go.golden")
		if got != want {
			t.Errorf("\n## got formatted file:\n%s\n## want:\n%s", got, want)
		}
	})
}

// Tests golang/go#36824.
func TestFormattingOneLine36824(t *testing.T) {

	const onelineProgram = `
-- a.go --
package main; func f() {}

-- a.go.formatted --
package main

func f() {}
`
	runner.Run(t, onelineProgram, func(t *testing.T, env *Env) {
		env.OpenFile("a.go")
		env.FormatBuffer("a.go")
		got := env.Editor.BufferText("a.go")
		want := env.ReadWorkspaceFile("a.go.formatted")
		if got != want {
			t.Errorf("got\n%q wanted\n%q", got, want)
		}
	})
}

// Tests golang/go#36824.
func TestFormattingOneLineImports36824(t *testing.T) {
	t.Skipf("golang/go#36824 has not been fixed yet")

	const onelineProgramA = `
-- a.go --
package x; func f() {fmt.Println()}

-- a.go.imported --
package x

import "fmt"

func f() { fmt.Println() }
`
	runner.Run(t, onelineProgramA, func(t *testing.T, env *Env) {
		env.OpenFile("a.go")
		env.OrganizeImports("a.go")
		got := env.Editor.BufferText("a.go")
		want := env.ReadWorkspaceFile("a.go.imported")
		if got != want {
			t.Errorf("OneLineImports:\n%s", tests.Diff(want, got))
		}
	})
}

const disorganizedProgram = `
-- main.go --
package main

import (
	"fmt"
	"errors"
)
func main(  ) {
	fmt.Println(errors.New("bad"))
}
-- main.go.organized --
package main

import (
	"errors"
	"fmt"
)
func main(  ) {
	fmt.Println(errors.New("bad"))
}
-- main.go.formatted --
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(errors.New("bad"))
}
`

func TestOrganizeImports(t *testing.T) {
	runner.Run(t, disorganizedProgram, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		env.OrganizeImports("main.go")
		got := env.Editor.BufferText("main.go")
		want := env.ReadWorkspaceFile("main.go.organized")
		if got != want {
			t.Errorf("\n## got formatted file:\n%s\n## want:\n%s", got, want)
		}
	})
}

func TestFormattingOnSave(t *testing.T) {
	runner.Run(t, disorganizedProgram, func(t *testing.T, env *Env) {
		env.OpenFile("main.go")
		env.SaveBuffer("main.go")
		got := env.Editor.BufferText("main.go")
		want := env.ReadWorkspaceFile("main.go.formatted")
		if got != want {
			t.Errorf("\n## got formatted file:\n%s\n## want:\n%s", got, want)
		}
	})
}
