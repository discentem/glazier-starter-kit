package unattend

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Unattend struct {
	XMLName   xml.Name `xml:"unattend"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Servicing struct {
		Text    string `xml:",chardata"`
		Package struct {
			Text             string `xml:",chardata"`
			Action           string `xml:"action,attr"`
			AssemblyIdentity struct {
				Text                  string `xml:",chardata"`
				Name                  string `xml:"name,attr"`
				Version               string `xml:"version,attr"`
				ProcessorArchitecture string `xml:"processorArchitecture,attr"`
				PublicKeyToken        string `xml:"publicKeyToken,attr"`
				Language              string `xml:"language,attr"`
			} `xml:"assemblyIdentity"`
		} `xml:"package"`
	} `xml:"servicing"`
	Settings []struct {
		Text      string `xml:",chardata"`
		Pass      string `xml:"pass,attr"`
		Component []struct {
			Text                   string `xml:",chardata"`
			Name                   string `xml:"name,attr"`
			ProcessorArchitecture  string `xml:"processorArchitecture,attr"`
			PublicKeyToken         string `xml:"publicKeyToken,attr"`
			Language               string `xml:"language,attr"`
			VersionScope           string `xml:"versionScope,attr"`
			Wcm                    string `xml:"wcm,attr"`
			Xsi                    string `xml:"xsi,attr"`
			DisableAccelerators    string `xml:"DisableAccelerators"`
			DisableFirstRunWizard  string `xml:"DisableFirstRunWizard"`
			DisableOOBAccelerators string `xml:"DisableOOBAccelerators"`
			HomePage               string `xml:"Home_Page"`
			SuggestedSitesEnabled  string `xml:"SuggestedSitesEnabled"`
			SearchScopes           struct {
				Text  string `xml:",chardata"`
				Scope struct {
					Text               string `xml:",chardata"`
					Action             string `xml:"action,attr"`
					FaviconURL         string `xml:"FaviconURL"`
					ScopeDisplayName   string `xml:"ScopeDisplayName"`
					ScopeDefault       string `xml:"ScopeDefault"`
					SuggestionsURL     string `xml:"SuggestionsURL"`
					SuggestionsURLJSON string `xml:"SuggestionsURL_JSON"`
					ScopeKey           string `xml:"ScopeKey"`
					ScopeUrl           string `xml:"ScopeUrl"`
				} `xml:"Scope"`
			} `xml:"SearchScopes"`
			OEMInformation struct {
				Text           string `xml:",chardata"`
				HelpCustomized string `xml:"HelpCustomized"`
				SupportURL     string `xml:"SupportURL"`
			} `xml:"OEMInformation"`
			ComputerName           string `xml:"ComputerName"`
			CopyProfile            string `xml:"CopyProfile"`
			RegisteredOrganization string `xml:"RegisteredOrganization"`
			RegisteredOwner        string `xml:"RegisteredOwner"`
			TimeZone               string `xml:"TimeZone"`
			TaskbarLinks           struct {
				Text  string `xml:",chardata"`
				Link0 string `xml:"Link0"`
				Link1 string `xml:"Link1"`
			} `xml:"TaskbarLinks"`
			RunAsynchronous struct {
				Text                   string `xml:",chardata"`
				RunAsynchronousCommand []struct {
					Text        string `xml:",chardata"`
					Action      string `xml:"action,attr"`
					Order       string `xml:"Order"`
					Description string `xml:"Description"`
					Path        string `xml:"Path"`
				} `xml:"RunAsynchronousCommand"`
			} `xml:"RunAsynchronous"`
			IcmpRedirectsEnabled string `xml:"IcmpRedirectsEnabled"`
			InputLocale          string `xml:"InputLocale"`
			SystemLocale         string `xml:"SystemLocale"`
			UILanguage           string `xml:"UILanguage"`
			UserLocale           string `xml:"UserLocale"`
			AutoLogon            struct {
				Text       string `xml:",chardata"`
				Enabled    string `xml:"Enabled"`
				LogonCount string `xml:"LogonCount"`
				Username   string `xml:"Username"`
				Password   struct {
					Text      string `xml:",chardata"`
					Value     string `xml:"Value"`
					PlainText string `xml:"PlainText"`
				} `xml:"Password"`
			} `xml:"AutoLogon"`
			FirstLogonCommands struct {
				Text               string `xml:",chardata"`
				SynchronousCommand []struct {
					Text        string `xml:",chardata"`
					Action      string `xml:"action,attr"`
					CommandLine string `xml:"CommandLine"`
					Description string `xml:"Description"`
					Order       string `xml:"Order"`
				} `xml:"SynchronousCommand"`
			} `xml:"FirstLogonCommands"`
			OOBE struct {
				Text                      string `xml:",chardata"`
				HideEULAPage              string `xml:"HideEULAPage"`
				HideWirelessSetupInOOBE   string `xml:"HideWirelessSetupInOOBE"`
				NetworkLocation           string `xml:"NetworkLocation"`
				ProtectYourPC             string `xml:"ProtectYourPC"`
				HideOnlineAccountScreens  string `xml:"HideOnlineAccountScreens"`
				HideOEMRegistrationScreen string `xml:"HideOEMRegistrationScreen"`
				HideLocalAccountScreen    string `xml:"HideLocalAccountScreen"`
				UnattendEnableRetailDemo  string `xml:"UnattendEnableRetailDemo"`
			} `xml:"OOBE"`
			UserAccounts struct {
				Text                  string `xml:",chardata"`
				AdministratorPassword struct {
					Text      string `xml:",chardata"`
					Value     string `xml:"Value"`
					PlainText string `xml:"PlainText"`
				} `xml:"AdministratorPassword"`
			} `xml:"UserAccounts"`
			DesktopOptimization struct {
				Text                          string `xml:",chardata"`
				ShowWindowsStoreAppsOnTaskbar string `xml:"ShowWindowsStoreAppsOnTaskbar"`
				GoToDesktopOnSignIn           string `xml:"GoToDesktopOnSignIn"`
			} `xml:"DesktopOptimization"`
			ComplianceCheck struct {
				Text          string `xml:",chardata"`
				DisplayReport string `xml:"DisplayReport"`
			} `xml:"ComplianceCheck"`
			DiskConfiguration struct {
				Text                             string `xml:",chardata"`
				DisableEncryptedDiskProvisioning string `xml:"DisableEncryptedDiskProvisioning"`
				WillShowUI                       string `xml:"WillShowUI"`
				Disk                             struct {
					Text             string `xml:",chardata"`
					Action           string `xml:"action,attr"`
					CreatePartitions struct {
						Text            string `xml:",chardata"`
						CreatePartition []struct {
							Text   string `xml:",chardata"`
							Action string `xml:"action,attr"`
							Order  string `xml:"Order"`
							Size   string `xml:"Size"`
							Type   string `xml:"Type"`
							Extend string `xml:"Extend"`
						} `xml:"CreatePartition"`
					} `xml:"CreatePartitions"`
					ModifyPartitions struct {
						Text            string `xml:",chardata"`
						ModifyPartition []struct {
							Text        string `xml:",chardata"`
							Action      string `xml:"action,attr"`
							PartitionID string `xml:"PartitionID"`
							Format      string `xml:"Format"`
							Label       string `xml:"Label"`
							Order       string `xml:"Order"`
							TypeID      string `xml:"TypeID"`
						} `xml:"ModifyPartition"`
					} `xml:"ModifyPartitions"`
					DiskID       string `xml:"DiskID"`
					WillWipeDisk string `xml:"WillWipeDisk"`
				} `xml:"Disk"`
			} `xml:"DiskConfiguration"`
			ImageInstall struct {
				Text    string `xml:",chardata"`
				OSImage struct {
					Text      string `xml:",chardata"`
					InstallTo struct {
						Text        string `xml:",chardata"`
						DiskID      string `xml:"DiskID"`
						PartitionID string `xml:"PartitionID"`
					} `xml:"InstallTo"`
				} `xml:"OSImage"`
			} `xml:"ImageInstall"`
			UserData struct {
				Text       string `xml:",chardata"`
				ProductKey struct {
					Text       string `xml:",chardata"`
					Key        string `xml:"Key"`
					WillShowUI string `xml:"WillShowUI"`
				} `xml:"ProductKey"`
			} `xml:"UserData"`
			EnableFirewall string `xml:"EnableFirewall"`
			EnableNetwork  string `xml:"EnableNetwork"`
		} `xml:"component"`
	} `xml:"settings"`
	OfflineImage struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Cpi    string `xml:"cpi,attr"`
	} `xml:"offlineImage"`
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	up := flag.String("unattend_path", pwd, "Path to unattend.xml")
	byt, err := ioutil.ReadFile(fmt.Sprintf("%s/unattend.xml", *up))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byt))

}
