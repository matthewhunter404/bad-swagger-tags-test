package bad_swagger_tags

import "fmt"

//
// Computes the set of conflicting or ambiguous alternatives from a
// configuration set, if that information was not already provided by the
// parser.
//
// @param ReportedAlts The set of conflicting or ambiguous alternatives, as
// Reported by the parser.
// @param configs The conflicting or ambiguous configuration set.
// @return Returns {@code ReportedAlts} if it is not {@code nil}, otherwise
// returns the set of alternatives represented in {@code configs}.
//
func PrintSomething() {
	fmt.Println("Something")
}
