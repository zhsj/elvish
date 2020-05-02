// Package exc exposes an Elvish module containing functionalities for working
// with exceptions.
package exc

import "github.com/elves/elvish/pkg/eval"

//elvdoc:fn show
//
// ```elvish
// exc:show
// ```
//
// Prints the exception to the output, showing its cause and stacktrace using VT
// sequences.

var Ns = eval.Ns{}.
	AddGoFns("exc:", map[string]interface{}{
		"show": show,
	})

func show(fm *eval.Frame, e *eval.Exception) {
	fm.OutputFile().WriteString(e.Show(""))
}
