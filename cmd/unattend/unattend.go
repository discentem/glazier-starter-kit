package unattend

import (
	"encoding/xml"
	"io/ioutil"
)

type Settings struct {
	Source       string
	Destination  string
	HomePage     string
	OEMSupport   string
	RegisteredCo string
	Pass         string
}

func (s Settings) Execute() error {
	byt, err := ioutil.ReadFile(s.Source)
	if err != nil {
		return err
	}
	u := &Unattend{}

	if err := xml.Unmarshal(byt, u); err != nil {
		return err
	}

	u.Settings[0].Component[0].HomePage = s.HomePage
	u.Settings[0].Component[1].OEMInformation.SupportURL = s.OEMSupport
	u.Settings[0].Component[1].RegisteredOrganization = s.RegisteredCo
	u.Settings[0].Component[1].RegisteredOwner = s.RegisteredCo
	u.Settings[0].Component[3].HomePage = s.HomePage
	u.Settings[1].Component[1].AutoLogon.Password.Value = s.Pass
	u.Settings[1].Component[1].UserAccounts.AdministratorPassword.Value = s.Pass
	u.Settings[1].Component[1].RegisteredOrganization = s.RegisteredCo
	u.Settings[1].Component[1].RegisteredOwner = s.RegisteredCo
	//fmt.Printf("%+v\n", u.Settings[1].Component[1].AutoLogon.Password.Value)
	//fmt.Printf("%+v\n", u)
	byt, err = xml.MarshalIndent(u, "", " ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(s.Destination, byt, 0644); err != nil {
		return err
	}

	return nil

}
