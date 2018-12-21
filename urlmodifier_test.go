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

func TestRandomizeSubDomain(t *testing.T) {
	var found bool

	myUrl := "https://www.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true"
	subDomains := []string{"s1", "s2", "s3"}
	newUrl, err := RandomizeSubDomain(myUrl, subDomains)
	if err != nil {
		t.Errorf("Fail randomizing subdomain. Error: %v", err)
	}

	newUrls := []string{
		"https://s1.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true",
		"https://s2.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true",
		"https://s3.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true",
	}

	for _, v := range newUrls {
		if v == newUrl {
			found = true
			break
		}
	}

	if !found {
		t.Error("Fail randomizing subdomain. Got: ", newUrl)
	}
}

func TestReplacePathComponent(t *testing.T) {
	myUrl := "https://s1.bukalapak.com/avt/7700597/medium/02420bf25da2677d984420e4c8479c61.jpg?go=true"
	newUrl, err := ReplacePathComponent(myUrl, "small", 3)
	if err != nil {
		t.Error("ReplacePathComponent is Fail. Error: ", err)
	}

	expectValue := "https://s1.bukalapak.com/avt/7700597/small/02420bf25da2677d984420e4c8479c61.jpg?go=true"
	if newUrl != expectValue {
		t.Errorf("ReplacePathComponent is Fail. Expect: %v, Got: %v", expectValue, newUrl)
	}
}
