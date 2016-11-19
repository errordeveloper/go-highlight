package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("coffeescript", `{"aliases":[0,1,2],"keywords":{"keyword":"in if for while finally new do return else break catch instanceof throw try this switch continue typeof delete debugger super then unless until loop of by when and or is isnt not","literal":"true false null undefined yes no on off","built_in":"npm require console print module global window document"},"illegal":"/\\/\\*/","contains":[0,1,2,3,4,5,6,7,8,9,10,11]}`)
}