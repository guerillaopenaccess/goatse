package doiutils

var (
    crossRefLicenses map[string]bool
)

func LicenseIsOpen(licenseUrl string) bool {
	if license_openness, license_known := crossRefLicenses[licenseUrl]; license_openness && license_known {
		return true
	} else {
		return false
	}
}

func init() {
	crossRefLicenses = map[string]bool{
		"http://creativecommons.org/licenses/by-nc-nd/3.0/de" : true,
		"http://creativecommons.org/licenses/by-nc/3.0" :  true,
		"http://creativecommons.org/licenses/by-nc/3.0/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.1/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.10/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.11/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.12/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.13/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.14/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.15/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.16/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.17/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.18/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.19/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.2/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.20/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.21/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.22/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.23/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.24/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.25/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.26/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.27/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.28/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.29/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.3/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.30/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.31/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.32/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.33/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.34/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.35/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.4/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.5/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.6/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.7/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.8/" :  true,
		"http://creativecommons.org/licenses/by-nc/3.9/" :  true,
		"http://creativecommons.org/licenses/by-nc/4.0/" :  true,
		"http://creativecommons.org/licenses/by/3.0/" :  true,
		"http://creativecommons.org/licenses/by/3.0/deed.en_US" :  true,
		"http://creativecommons.org/licenses/by/4.0/" :  true,
		"http://creativecommons.org/publicdomain/zero/1.0/" :  true,
		"http://link.aps.org/licenses/aps-default-accepted-manuscript-license" : false,
		"http://link.aps.org/licenses/aps-default-license" : false,
		"http://link.aps.org/licenses/aps-default-text-mining-license" : false,
		"http://olabout.wiley.com/WileyCDA/Section/id-815641.html" : false,
		"http://pubs.acs.org/page/policy/authorchoice_ccby_termsofuse.html" :  true,
		"http://pubs.acs.org/page/policy/authorchoice_ccbyncnd_termsofuse.html" :  true,
		"http://pubs.acs.org/page/policy/authorchoice_termsofuse.html" : false,
		"http://pubs.acs.org/userimages/ContentEditor/1388526979973/authorchoice_form.pdf" : false,
		"http://www.acm.org/publications/policies/copyright_policy#Background" : false,
		"http://www.acs.org/content/acs/en/copyright.html" : false,
		"http://www.crossref.org/license" : false,
		"http://www.developers.elsevier.com/cms/content/text-mining-elsevier-publications" : false,
		"http://www.elsevier.com/open-access/userlicense/1.0/" :  true,
		"http://www.elsevier.com/tdm/userlicense/1.0/" : false,
		"http://www.ieee.org/publications_standards/publications/subscriptions/info/licensing.html" : false,
		"http://www.sciencemag.org/site/feature/contribinfo/prep/license.xhtml" : false,
		"http://www.tandfonline.com/page/terms-and-conditions#link2" : false,
	}
}
