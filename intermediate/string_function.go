package intermediate

//string functions in Go are provided by the built-in "string" package. Here are some common string fucntions in go.
//1.len(s string) int: returns  the length of the string s in bytes.
//2. strings.Container(s, substr string) bool: returns true if the string s contains the substring substr.
//3. strings.HasPrefix(s, prefix string) bool: returns true if the string s starts with the prefix.
//4. strings.HasSuffix(s, suffix string) bool: returns true if the string s ends with the suffix.
//5. strings.ToUpper(s string) string: returns a copy of the string s with all letters converted to uppercase.
//6. strings.ToLower(s string) string: returns a copy of the string s with all letters converted to lowercase.
//7. strings.TrimSpace(s string) string: returns a copy of the string s with all leading and trailing white space removed.
//8. strings.Replace(s, old, new string, n int) string: returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
//9. strings.Split(s, sep string) []string: returns a slice of strings split from s by the separator sep.
//10. strings.Join(a []string, sep string) string: returns a string that is the concatenation of the elements of a, separated by the separator sep.

import "strings"

func main() {
	str := "Hello, World!"
	println("Length of string:", len(str))
	println("Contains 'World':", strings.Contains(str, "World"))
	println("Starts with 'Hellp':", strings.HasPrefix(str, "Hello"))
	println("End with world:", strings.HasSuffix(str, "world!"))
	println("To Upper:", strings.ToUpper(str))
	println("To Lower:", strings.ToLower(str))
	println("Trim space:", strings.TrimSpace("  Hello, world!  "))
	println("Replace '0' with 'X':", strings.Replace("1001", "0", "X", -1))
}