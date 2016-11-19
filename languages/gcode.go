package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("gcode", `{"aliases":["nc"],"case_insensitive":true,"lexemes":"[A-Z_][A-Z0-9_.]*","keywords":"IF DO WHILE ENDWHILE CALL ENDIF SUB ENDSUB GOTO REPEAT ENDREPEAT EQ LT GT NE GE LE OR XOR","contains":[{"className":"meta","begin":"\\%"},{"className":"meta","begin":"([O])([0-9]+)"},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"\\(","end":"\\)","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"([-+]?([0-9]*\\.?[0-9]+\\.?))|(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"string","begin":"'","end":"'","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"name","begin":"([G])([0-9]+\\.?[0-9]?)"},{"className":"name","begin":"([M])([0-9]+\\.?[0-9]?)"},{"className":"attr","begin":"(VC|VS|#)","end":"(\\d+)"},{"className":"attr","begin":"(VZOFX|VZOFY|VZOFZ)"},{"className":"built_in","begin":"(ATAN|ABS|ACOS|ASIN|SIN|COS|EXP|FIX|FUP|ROUND|LN|TAN)(\\[)","end":"([-+]?([0-9]*\\.?[0-9]+\\.?))(\\])"},{"className":"symbol","variants":[{"begin":"N","end":"\\d+","illegal":"\\W"}]}]}`)
}