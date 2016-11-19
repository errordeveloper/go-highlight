package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("http", `{"aliases":["https"],"illegal":"\\S","contains":[{"begin":"^HTTP/[0-9\\.]+","end":"$","contains":[{"className":"number","begin":"\\b\\d{3}\\b"}]},{"begin":"^[A-Z]+ (.*?) HTTP/[0-9\\.]+$","returnBegin":true,"end":"$","contains":[{"className":"string","begin":" ","end":" ","excludeBegin":true,"excludeEnd":true},{"begin":"HTTP/[0-9\\.]+"},{"className":"keyword","begin":"[A-Z]+"}]},{"className":"attribute","begin":"^\\w","end":": ","excludeEnd":true,"illegal":"\\n|\\s|=","starts":{"end":"$","relevance":0}},{"begin":"\\n\\n","starts":{"subLanguage":[],"endsWithParent":true}}]}`)
}