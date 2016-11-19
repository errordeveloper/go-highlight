package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("haml", `{"case_insensitive":true,"contains":[{"className":"meta","begin":"^!!!( (5|1\\.1|Strict|Frameset|Basic|Mobile|RDFa|XML\\b.*))?$","relevance":10},{"className":"comment","begin":"^\\s*(!=#|=#|-#|/).*$","end":false,"contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"begin":"^\\s*(-|=|!=)(?!#)","starts":{"end":"\\n","subLanguage":"ruby"}},{"className":"tag","begin":"^\\s*%","contains":[{"className":"selector-tag","begin":"\\w+"},{"className":"selector-id","begin":"#[\\w-]+"},{"className":"selector-class","begin":"\\.[\\w-]+"},{"begin":"{\\s*","end":"\\s*}","contains":[{"begin":":\\w+\\s*=>","end":",\\s+","returnBegin":true,"endsWithParent":true,"contains":[{"className":"attr","begin":":\\w+"},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"\\w+","relevance":0}]}]},{"begin":"\\(\\s*","end":"\\s*\\)","excludeEnd":true,"contains":[{"begin":"\\w+\\s*=","end":"\\s+","returnBegin":true,"endsWithParent":true,"contains":[{"className":"attr","begin":"\\w+","relevance":0},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"\\w+","relevance":0}]}]}]},{"begin":"^\\s*[=~]\\s*"},{"begin":"#{","starts":{"end":"}","subLanguage":"ruby"}}]}`)
}