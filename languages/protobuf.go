package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("protobuf", `{"keywords":{"keyword":"package import option optional required repeated group","built_in":"double float int32 int64 uint32 uint64 sint32 sint64 fixed32 fixed64 sfixed32 sfixed64 bool string bytes","literal":"true false"},"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"class","beginKeywords":"message enum service","end":"\\{","illegal":"\\n","contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0,"starts":{"endsWithParent":true,"excludeEnd":true}}]},{"className":"function","beginKeywords":"rpc","end":";","excludeEnd":true,"keywords":"rpc returns"},{"begin":"^\\s*[A-Z_]+","end":"\\s*=","excludeEnd":true}]}`)
}