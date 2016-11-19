package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("twig", `{"aliases":["craftcms"],"case_insensitive":true,"subLanguage":"xml","contains":[{"className":"comment","begin":"\\{#","end":"#}","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"template-tag","begin":"\\{%","end":"%}","contains":[{"className":"name","begin":"\\w+","keywords":"autoescape block do embed extends filter flush for if import include macro sandbox set spaceless use verbatim endautoescape endblock enddo endembed endextends endfilter endflush endfor endif endimport endinclude endmacro endsandbox endset endspaceless enduse endverbatim","starts":{"endsWithParent":true,"contains":[{"begin":"\\|[A-Za-z_]+:?","keywords":"abs batch capitalize convert_encoding date date_modify default escape first format join json_encode keys last length lower merge nl2br number_format raw replace reverse round slice sort split striptags title trim upper url_encode","contains":[{"beginKeywords":"attribute block constant cycle date dump include max min parent random range source template_from_string","keywords":{"name":"attribute block constant cycle date dump include max min parent random range source template_from_string"},"relevance":0,"contains":[{"className":"params","begin":"\\(","end":"\\)"}]}]},{"beginKeywords":"attribute block constant cycle date dump include max min parent random range source template_from_string","keywords":{"name":"attribute block constant cycle date dump include max min parent random range source template_from_string"},"relevance":0,"contains":[{"className":"params","begin":"\\(","end":"\\)"}]}],"relevance":0}}]},{"className":"template-variable","begin":"\\{\\{","end":"}}","contains":["self",{"begin":"\\|[A-Za-z_]+:?","keywords":"abs batch capitalize convert_encoding date date_modify default escape first format join json_encode keys last length lower merge nl2br number_format raw replace reverse round slice sort split striptags title trim upper url_encode","contains":[{"beginKeywords":"attribute block constant cycle date dump include max min parent random range source template_from_string","keywords":{"name":"attribute block constant cycle date dump include max min parent random range source template_from_string"},"relevance":0,"contains":[{"className":"params","begin":"\\(","end":"\\)"}]}]},{"beginKeywords":"attribute block constant cycle date dump include max min parent random range source template_from_string","keywords":{"name":"attribute block constant cycle date dump include max min parent random range source template_from_string"},"relevance":0,"contains":[{"className":"params","begin":"\\(","end":"\\)"}]}]}]}`)
}