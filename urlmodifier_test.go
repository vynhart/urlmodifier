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

func TestReplaceSubDomain(t *testing.T) {
	myUrl := "https://www.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true"
	newUrl, err := ReplaceSubDomain(myUrl, "s9")
	if err != nil {
		t.Errorf("Fail replacing subdomain. Error: %v", err)
	}

	expectValue := "https://s9.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true"
	if newUrl != expectValue {
		t.Errorf("Fail replacing subdomain. Expect: %v, Got: %v", expectValue, newUrl)
	}
}
