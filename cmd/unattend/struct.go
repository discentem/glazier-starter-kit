package unattend

import "encoding/xml"

type Unattend struct {
	XMLName   xml.Name `xml:"unattend,omitempty"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,omitempty,attr"`
	Servicing struct {
		Text    string `xml:",chardata"`
		Package struct {
			Text             string `xml:",chardata"`
			Action           string `xml:"action,omitempty,attr"`
			AssemblyIdentity struct {
				Text                  string `xml:",chardata"`
				Name                  string `xml:"name,omitempty,attr"`
				Version               string `xml:"version,omitempty,attr"`
				ProcessorArchitecture string `xml:"processorArchitecture,omitempty,attr"`
				PublicKeyToken        string `xml:"publicKeyToken,omitempty,attr"`
				Language              string `xml:"language,omitempty,attr"`
			} `xml:"assemblyIdentity,omitempty"`
		} `xml:"package,omitempty"`
	} `xml:"servicing,omitempty"`
	Settings []struct {
		Text      string `xml:",chardata"`
		Pass      string `xml:"pass,omitempty,attr"`
		Component []struct {
			Text                   string `xml:",chardata"`
			Name                   string `xml:"name,omitempty,attr"`
			ProcessorArchitecture  string `xml:"processorArchitecture,omitempty,attr"`
			PublicKeyToken         string `xml:"publicKeyToken,omitempty,attr"`
			Language               string `xml:"language,omitempty,attr"`
			VersionScope           string `xml:"versionScope,omitempty,attr"`
			Wcm                    string `xml:"wcm,omitempty,attr"`
			Xsi                    string `xml:"xsi,omitempty,attr"`
			DisableAccelerators    string `xml:"DisableAccelerators,omitempty"`
			DisableFirstRunWizard  string `xml:"DisableFirstRunWizard,omitempty"`
			DisableOOBAccelerators string `xml:"DisableOOBAccelerators,omitempty"`
			HomePage               string `xml:"Home_Page,omitempty"`
			SuggestedSitesEnabled  string `xml:"SuggestedSitesEnabled,omitempty"`
			SearchScopes           struct {
				Text  string `xml:",chardata"`
				Scope struct {
					Text               string `xml:",chardata"`
					Action             string `xml:"action,omitempty,attr"`
					FaviconURL         string `xml:"FaviconURL,omitempty"`
					ScopeDisplayName   string `xml:"ScopeDisplayName,omitempty"`
					ScopeDefault       string `xml:"ScopeDefault,omitempty"`
					SuggestionsURL     string `xml:"SuggestionsURL,omitempty"`
					SuggestionsURLJSON string `xml:"SuggestionsURL_JSON,omitempty"`
					ScopeKey           string `xml:"ScopeKey,omitempty"`
					ScopeUrl           string `xml:"ScopeUrl,omitempty"`
				} `xml:"Scope,omitempty"`
			} `xml:"SearchScopes,omitempty"`
			OEMInformation struct {
				Text           string `xml:",chardata"`
				HelpCustomized string `xml:"HelpCustomized,omitempty"`
				SupportURL     string `xml:"SupportURL,omitempty"`
			} `xml:"OEMInformation,omitempty"`
			ComputerName           string `xml:"ComputerName,omitempty"`
			CopyProfile            string `xml:"CopyProfile,omitempty"`
			RegisteredOrganization string `xml:"RegisteredOrganization,omitempty"`
			RegisteredOwner        string `xml:"RegisteredOwner,omitempty"`
			TimeZone               string `xml:"TimeZone,omitempty"`
			TaskbarLinks           struct {
				Text  string `xml:",chardata"`
				Link0 string `xml:"Link0,omitempty"`
				Link1 string `xml:"Link1,omitempty"`
			} `xml:"TaskbarLinks,omitempty"`
			RunAsynchronous struct {
				Text                   string `xml:",chardata"`
				RunAsynchronousCommand []struct {
					Text        string `xml:",chardata"`
					Action      string `xml:"action,omitempty,attr"`
					Order       string `xml:"Order,omitempty"`
					Description string `xml:"Description,omitempty"`
					Path        string `xml:"Path,omitempty"`
				} `xml:"RunAsynchronousCommand,omitempty"`
			} `xml:"RunAsynchronous,omitempty"`
			IcmpRedirectsEnabled string `xml:"IcmpRedirectsEnabled,omitempty"`
			InputLocale          string `xml:"InputLocale,omitempty"`
			SystemLocale         string `xml:"SystemLocale,omitempty"`
			UILanguage           string `xml:"UILanguage,omitempty"`
			UserLocale           string `xml:"UserLocale,omitempty"`
			AutoLogon            struct {
				Text       string `xml:",chardata"`
				Enabled    string `xml:"Enabled,omitempty"`
				LogonCount string `xml:"LogonCount,omitempty"`
				Username   string `xml:"Username,omitempty"`
				Password   struct {
					Text      string `xml:",chardata"`
					Value     string `xml:"Value,omitempty"`
					PlainText string `xml:"PlainText,omitempty"`
				} `xml:"Password,omitempty"`
			} `xml:"AutoLogon,omitempty"`
			FirstLogonCommands struct {
				Text               string `xml:",chardata"`
				SynchronousCommand []struct {
					Text        string `xml:",chardata"`
					Action      string `xml:"action,omitempty,attr"`
					CommandLine string `xml:"CommandLine,omitempty"`
					Description string `xml:"Description,omitempty"`
					Order       string `xml:"Order,omitempty"`
				} `xml:"SynchronousCommand,omitempty"`
			} `xml:"FirstLogonCommands,omitempty"`
			OOBE struct {
				Text                      string `xml:",chardata"`
				HideEULAPage              string `xml:"HideEULAPage,omitempty"`
				HideWirelessSetupInOOBE   string `xml:"HideWirelessSetupInOOBE,omitempty"`
				NetworkLocation           string `xml:"NetworkLocation,omitempty"`
				ProtectYourPC             string `xml:"ProtectYourPC,omitempty"`
				HideOnlineAccountScreens  string `xml:"HideOnlineAccountScreens,omitempty"`
				HideOEMRegistrationScreen string `xml:"HideOEMRegistrationScreen,omitempty"`
				HideLocalAccountScreen    string `xml:"HideLocalAccountScreen,omitempty"`
				UnattendEnableRetailDemo  string `xml:"UnattendEnableRetailDemo,omitempty"`
			} `xml:"OOBE,omitempty"`
			UserAccounts struct {
				Text                  string `xml:",chardata"`
				AdministratorPassword struct {
					Text      string `xml:",chardata"`
					Value     string `xml:"Value,omitempty"`
					PlainText string `xml:"PlainText,omitempty"`
				} `xml:"AdministratorPassword,omitempty"`
			} `xml:"UserAccounts,omitempty"`
			DesktopOptimization struct {
				Text                          string `xml:",chardata"`
				ShowWindowsStoreAppsOnTaskbar string `xml:"ShowWindowsStoreAppsOnTaskbar,omitempty"`
				GoToDesktopOnSignIn           string `xml:"GoToDesktopOnSignIn,omitempty"`
			} `xml:"DesktopOptimization,omitempty"`
			ComplianceCheck struct {
				Text          string `xml:",chardata"`
				DisplayReport string `xml:"DisplayReport,omitempty"`
			} `xml:"ComplianceCheck,omitempty"`
			DiskConfiguration struct {
				Text                             string `xml:",chardata"`
				DisableEncryptedDiskProvisioning string `xml:"DisableEncryptedDiskProvisioning,omitempty"`
				WillShowUI                       string `xml:"WillShowUI,omitempty"`
				Disk                             struct {
					Text             string `xml:",chardata"`
					Action           string `xml:"action,omitempty,attr"`
					CreatePartitions struct {
						Text            string `xml:",chardata"`
						CreatePartition []struct {
							Text   string `xml:",chardata"`
							Action string `xml:"action,omitempty,attr"`
							Order  string `xml:"Order,omitempty"`
							Size   string `xml:"Size,omitempty"`
							Type   string `xml:"Type,omitempty"`
							Extend string `xml:"Extend,omitempty"`
						} `xml:"CreatePartition,omitempty"`
					} `xml:"CreatePartitions,omitempty"`
					ModifyPartitions struct {
						Text            string `xml:",chardata"`
						ModifyPartition []struct {
							Text        string `xml:",chardata"`
							Action      string `xml:"action,omitempty,attr"`
							PartitionID string `xml:"PartitionID,omitempty"`
							Format      string `xml:"Format,omitempty"`
							Label       string `xml:"Label,omitempty"`
							Order       string `xml:"Order,omitempty"`
							TypeID      string `xml:"TypeID,omitempty"`
						} `xml:"ModifyPartition,omitempty"`
					} `xml:"ModifyPartitions,omitempty"`
					DiskID       string `xml:"DiskID,omitempty"`
					WillWipeDisk string `xml:"WillWipeDisk,omitempty"`
				} `xml:"Disk,omitempty"`
			} `xml:"DiskConfiguration,omitempty"`
			ImageInstall struct {
				Text    string `xml:",chardata"`
				OSImage struct {
					Text      string `xml:",chardata"`
					InstallTo struct {
						Text        string `xml:",chardata"`
						DiskID      string `xml:"DiskID,omitempty"`
						PartitionID string `xml:"PartitionID,omitempty"`
					} `xml:"InstallTo,omitempty"`
				} `xml:"OSImage,omitempty"`
			} `xml:"ImageInstall,omitempty"`
			UserData struct {
				Text       string `xml:",chardata"`
				ProductKey struct {
					Text       string `xml:",chardata"`
					Key        string `xml:"Key,omitempty"`
					WillShowUI string `xml:"WillShowUI,omitempty"`
				} `xml:"ProductKey,omitempty"`
			} `xml:"UserData,omitempty"`
			EnableFirewall string `xml:"EnableFirewall,omitempty"`
			EnableNetwork  string `xml:"EnableNetwork,omitempty"`
		} `xml:"component,omitempty"`
	} `xml:"settings,omitempty"`
	OfflineImage struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,omitempty,attr"`
		Cpi    string `xml:"cpi,omitempty,attr"`
	} `xml:"offlineImage,omitempty"`
}
