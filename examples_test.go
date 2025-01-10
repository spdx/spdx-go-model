package spdx_go_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/spdx/spdx-go-model/spdx_v3_0_1"
)

func TestImport(t *testing.T) {
	spdx_v3_0_1.MakePerson()
}

func TestExample(t *testing.T) {
	// Create a new Person
	p := spdx_v3_0_1.MakePerson()

	// Set the unique identifier (IRI) for the person
	p.ID().Set("http://spdx.org/spdxdocs/person-example")

	// Set the name for the person
	p.Name().Set("Joshua")

	// Create a creation info object
	ci := spdx_v3_0_1.MakeCreationInfo()
	ci.SpecVersion().Set("3.0.1")
	ci.Created().Set(time.Now())

	// Set the person as the creator of the SPDX data
	ci.CreatedBy().AppendObj(p)

	// Set the creation info as the creation info of the Person
	p.CreationInfo().SetObj(ci)

	// Add an e-mail address for the person
	email := spdx_v3_0_1.MakeExternalIdentifier()
	email.ExternalIdentifierType().Set(spdx_v3_0_1.ExternalIdentifierTypeEmail)
	email.Identifier().Set("noreply@noreply.com")
	p.ExternalIdentifier().AppendObj(email)

	// Create a SPDX Document object
	doc := spdx_v3_0_1.MakeSpdxDocument()
	doc.ID().Set("http://spdx.org/spdxdocs/doc-example")
	doc.CreationInfo().SetObj(ci)

	// Add the document and person to an object set and serialize
	objset := spdx_v3_0_1.NewSHACLObjectSet()
	objset.AddObject(p)
	objset.AddObject(doc)

	file, err := os.Create("out.json")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := objset.Encode(encoder); err != nil {
		t.Error(err)
		return
	}
}
