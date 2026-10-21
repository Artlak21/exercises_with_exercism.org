// https://golang.org/doc/effective_go.html#commentary
package twofer

// ShareWith use "you", if name is empty
func ShareWith(name string) string {
	if name == "" {
        name =  "you"
    }
	return  "One for " + name + ", one for me."
}
