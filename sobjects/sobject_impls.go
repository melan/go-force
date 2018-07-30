
package sobjects

var SObjectsImplementations = map[string]SObject {
	"Account": &Account{},
	"Lead": &Lead{},
	"Opportunity": &Opportunity{},
	"Profile": &Profile{},
	"Task": &Task{},
	"User": &User{},
}
