package urlmodifier

import (
	"testing"
)

func TestSubDomain(t *testing.T) {
	myUrl := "https://s2.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true"
	subdomain, err := SubDomain(myUrl)
	if err != nil {
		t.Errorf("Fail getting subdomain. Error: %v", err)
	}

	expectValue := "s2"
	if subdomain != expectValue {
		t.Errorf("Fail getting subdomain. Expect: %v, Got: %v", expectValue, subdomain)
	}
}