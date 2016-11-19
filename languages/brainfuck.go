package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("brainfuck", `{"aliases":["bf"],"contains":[{"className":"comment","begin":"[^\\[\\]\\.,\\+\\-<> \r\n]","end":"[\\[\\]\\.,\\+\\-<> \r\n]","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"returnEnd":true,"relevance":0},{"className":"title","begin":"[\\[\\]]","relevance":0},{"className":"string","begin":"[\\.,]","relevance":0},{"begin":"\\+\\+|\\-\\-","returnBegin":true,"contains":[{"className":"literal","begin":"[\\+\\-]","relevance":0}]},{"className":"literal","begin":"[\\+\\-]","relevance":0}]}`)
}