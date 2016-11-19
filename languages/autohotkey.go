package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("autohotkey", `{"case_insensitive":true,"keywords":{"keyword":"Break Continue Else Gosub If Loop Return While","literal":"A|0 true false NOT AND OR","built_in":"ComSpec Clipboard ClipboardAll ErrorLevel"},"contains":[{"className":"built_in","begin":"A_[a-zA-Z0-9]+"},{"begin":"`+"`"+`[\\s\\S]"},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"`+"`"+`[\\s\\S]"}]},{"className":"comment","begin":";","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"variable","begin":"%","end":"%","illegal":"\\n","contains":[{"begin":"`+"`"+`[\\s\\S]"}]},{"className":"symbol","contains":[{"begin":"`+"`"+`[\\s\\S]"}],"variants":[{"begin":"^[^\\n\";]+::(?!=)"},{"begin":"^[^\\n\";]+:(?!=)","relevance":0}]},{"begin":",\\s*,"}]}`)
}