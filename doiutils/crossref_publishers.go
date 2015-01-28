package doiutils

import (
	"fmt"
)

type PublisherMeta struct{
	Names []string
	Id int
	DoiPrefixes []string
	TotalDoi int
}

var (
	IdToPublisher map[int16]PublisherMeta

	PublisherToId map[string]int16

	PrefixToId map[string]int16
)

func init() {
	IdToPublisher = map[int16]PublisherMeta{
		78: PublisherMeta{
			[]string{"Elsevier - Academic Press", "Elsevier", "Elsevier - WB Saunders", "Elsevier- Churchill Livingstone", "Elsevier - Mosby", "Elsevier - Urban & Fischer Verlag", "Elsevier - Ediciones Doyma", "Elsevier- Hanley and Belfus Inc.", "Elsevier - Institution of Chemical Engineers", "Elsevier- Spektrum Akademischer Verlag", "Elsevier - Ambulatory Pediatric Associates", "Elsevier - Medicine Publishing Company", "Elsevier - Biophysical Society", "Woodhead Publishing", "Elsevier - Wilderness Medical Society", "Elsevier - American Society for Experimental Neurotherapeutics", "Elsevier - American Society for Investigative Pathology", "Elsevier - International Federation of Automatic Control (IFAC)", "Elsevier - CIG Media Group LP", "Elsevier - Colegio Nacional de Opticos-Optometristas de Espana", "Elsevier - Mayo Clinic Proceedings", "Sociedad Argentina de Radiologia", "Elsevier BV"},
			78,
			[]string{"10.1006", "10.1016", "10.1053", "10.1054", "10.1067", "10.1078", "10.1157", "10.1197", "10.1205", "10.1240", "10.1367", "10.1383", "10.1529", "10.1533", "10.1580", "10.1602", "10.2353", "10.3182", "10.3816", "10.3921", "10.4065", "10.7811"},
			12363552	},
		311: PublisherMeta{
			[]string{"Wiley Blackwell (John Wiley & Sons)", "Wiley Blackwell (Blackwell Publishing - Munksgaard)", "Wiley Blackwell (Blackwell Publishing)", "Wiley Blackwell (Blackwell Publishing -The Physiological Society)", "Wiley Blackwell (New York Academy of Sciences)", "Wiley Blackwell (Royal Meteorological Society)", "Wiley Blackwell (International Life Sciences Institute)", "Wiley Blackwell (British Psychological Society)", "Wiley Blackwell (American Society Bone & Mineral Research)", "Wiley Blackwell (Canadian Academic Accounting Association)", "Wiley Blackwell (International Journal of Psychoanalysis)", "Wiley Blackwell (Rural Sociological Society)", "Wiley Blackwell (Robotic Publications)", "Pharmacotherapy", "Wiley Blackwell (American College of Veterinary Internal Medicine)", "Wiley Blackwell (Society of Environmental Toxicology and Chemistry)", "American Society of Andrology", "Wiley Blackwell (Equine Veterinary Journal)", "Wiley Blackwell (AHRC Research Centre)", "Wiley Blackwell (Comparative Legislative Research Center)", "Wiley Blackwell (Opulus Press)", "Wiley Blackwell (Production and Operations Management)", "Wiley Blackwell (New York Academy of Sciences E-Briefings)", "Wiley Blackwell (The Wildlife Society)", "Wiley Blackwell (Teachers of English to Speakers of Other Languages, Inc.)", "Wiley-Blackwell"},
			311,
			[]string{"10.1002", "10.1034", "10.1046", "10.1111", "10.1113", "10.1196", "10.1256", "10.1301", "10.1348", "10.1359", "10.1506", "10.1516", "10.1526", "10.1581", "10.1592", "10.1892", "10.1897", "10.2164", "10.2746", "10.2966", "10.3162", "10.3170", "10.3401", "10.3405", "10.4004", "10.5054", "10.14814"},
			6538697	},
		297: PublisherMeta{
			[]string{"Springer-Verlag", "Springer (Kluwer Academic Publishers)", "Springer - Ecomed Publishers", "Springer (Kluwer Academic Publishers - Biomedical Engineering Society (BMES))", "Springer (Biomed Central Ltd.)", "Springer - Society of Surgical Oncology", "Springer (Biological Procedures Online)", "Springer - ASM International", "Springer - Cell Stress Society International", "Springer - FD Communications", "Springer - Humana Press", "Springer - RILEM Publishing", "Springer - Psychonomic Society", "Wageningen Academic Publishers", "Springer - Mammal Research Institute", "Springer - The Korean Society of Pharmaceutical Sciences and Technology", "Springer - Real Academia de Ciencias Exactas, Fisicas y Naturales", "Springer - (backfiles)", "Island Press", "Springer - Global Science Journals", "Springer Science + Business Media"},
			297,
			[]string{"10.1007", "10.1023", "10.1065", "10.1114", "10.1140", "10.1186", "10.1245", "10.1251", "10.1361", "10.1379", "10.1381", "10.1385", "10.1617", "10.3758", "10.3920", "10.4098", "10.4333", "10.5052", "10.5819", "10.5822", "10.7603", "10.2165"},
			6205600	},
		301: PublisherMeta{
			[]string{"Informa UK (Taylor & Francis)", "Informa UK (Marcel Dekker)", "Fundacion Infancia y Aprendizaje", "Informa UK (American Statistical Association)", "CRC Press", "Informa UK (Lawrence Erlbaum Associates, Inc.)", "Informa UK (Taylor & Francis Books)", "Illuminating Engineering Society of North America", "Informa UK (IAHS Ltd.)", "Informa UK(IOP Books)", "Informa UK (Encyclopedia of Astronomy and Astrophysics)", "Informa UK (Multilingual Matters)", "Informa UK (Lawrence Erlbaum Associates, Inc)", "Informa UK (Human Frontier Science Program Publishing)", "Informa UK (National Inquiry Services Center)", "Informa UK (Regulatory Affairs)", "Informa UK (Canadian Meteorological and Oceanographic Society)", "Informa UK (Beech Tree Publishing)", "Informa UK (Air & Waste Management Association)", "Informa UK (National Inquiry Services Center/African Journals Online)", "Informa UK (Heldref Publications)", "Earthscan", "Informa UK (International Association of Hydraulic Engineering and Research (IAHR))", "Informa UK (Georg D. W. Callwey GambH and Co., KG)", "Informa UK (Journal of Maps)", "Informa UK (Routledge)", "Desalination Publications", "Nottingham University Press", "Informa UK Limited", "Council on Social Work Education"},
			301,
			[]string{"10.1080", "10.1081", "10.1174", "10.1198", "10.1201", "10.1207", "10.1531", "10.1582", "10.1623", "10.1887", "10.1888", "10.2167", "10.2513", "10.2976", "10.2989", "10.3110", "10.3137", "10.3152", "10.3155", "10.3187", "10.3200", "10.3763", "10.3826", "10.3939", "10.4113", "10.4324", "10.5004", "10.5661", "10.5175", "10.4161", "10.5370"},
			2636683	},
		1121: PublisherMeta{
			[]string{"JSTOR"},
			1121,
			[]string{"10.2307"},
			1978692	},
		276: PublisherMeta{
			[]string{"Ovid Technologies (Wolters Kluwer) - Lippincott Williams & Wilkins", "Ovid Technologies Wolters Kluwer -American Heart Association", "Ovid Technologies (Wolters Kluwer) - International Pediatric Research Foundation", "Ovid Technologies (Wolters Kluwer) - American Academy of Neurology", "Ovid Technologies (Wolters Kluwer) - Anesthesia & Analgesia", "Ovid Technologies (Wolters Kluwer) - Neurosurgery", "Ovid Technologies (Wolters-Kluwer) - American College of Sports Medicine", "Ovid Technologies (Wolters Kluwer) - Triological Society", "Ovid Technologies (Wolters Kluwer) - National Strength & Conditioning", "Ovid Technologies (Wolters-Kluwer) - Adis", "Ovid Technologies (Wolters Kluwer) - Italian Federation of Cardiology", "Ovid Technologies (Wolters Kluwer) - ENSTINET", "Ovid Technologies (Wolters Kluwer Health)"},
			276,
			[]string{"10.1097", "10.1161", "10.1203", "10.1212", "10.1213", "10.1227", "10.1249", "10.1288", "10.1519", "10.2165", "10.2459", "10.7123"},
			1663896	},
		286: PublisherMeta{
			[]string{"Oxford University Press", "Oxford University Press - London Mathematical Society", "European Association of Cardiothoracic Surgery", "Poultry Science Association", "Oxford University Press (OUP)"},
			286,
			[]string{"10.1093", "10.1112", "10.1510", "10.3382"},
			1431409	},
		179: PublisherMeta{
			[]string{"Sage Publications - Technomic Publishing Company", "SAGE Publications", "Arnold Publishers", "Research Publishing Services - Professional Engineering Publishing", "The Royal Society of Medicine", "Harvey Whitney Books Co.", "American College of Veterinary Pathologists", "SagePublications - National Association of School Nurses", "Sage Publications-International Institute for Environment and Development", "Canadian Association of Occupational Therapists", "Sage Publications - Bulletin of the Atomic Scientists", "Sage Publications (JRAAS Limited)", "Canadian Pharmacists Journal", "Sage Publications - Reference E-Books", "Sage Publications (Prufrock Press, Inc.)", "Association for Experiential Education", "SAGE PUBLICATIONS", "Research Publishing Services", "Harvey Whitney Book Co.", "Excellus Health Plan, Inc."},
			179,
			[]string{"10.1106", "10.1177", "10.1191", "10.1243", "10.1258", "10.1345", "10.1354", "10.1622", "10.1630", "10.2182", "10.2968", "10.3317", "10.3821", "10.4135", "10.4219", "10.5193", "10.5034", "10.5126", "10.2511", "10.17322"},
			1333082	},
		316: PublisherMeta{
			[]string{"American Chemical Society", "American Chemical Society (ACS)"},
			316,
			[]string{"10.1021"},
			1325443	},
		56: PublisherMeta{
			[]string{"Cambridge University Press", "Cambridge University Press (Australian Academic Press)", "Cambridge University Press (Materials Research Society)", "Cambridge University Press (Society for the Promotion of Roman Studies)", "Cambridge University Press (Entomological Society of Canada)", "Cambridge University Press (Mathematical Association of America)", "Cambridge University Press (Liverpool University Press)", "Cambridge University Press (Anthem Press)", "Cambridge University Press (Nottingham University Press)", "Cambridge University Press (CUP)"},
			56,
			[]string{"10.1017", "10.1375", "10.1557", "10.3815", "10.4039", "10.5948", "10.5949", "10.7135", "10.7313"},
			1036374	},
		339: PublisherMeta{
			[]string{"Nature Publishing Group", "Nature Publishing Group - Macmillan Publishers", "Korean Society for Biochemistry and Molecular Biology"},
			339,
			[]string{"10.1013", "10.1038", "10.1057", "10.3858"},
			935787	},
		263: PublisherMeta{
			[]string{"Institute of Electrical and Electronics Engineers", "Institute of Electrical & Electronics Engineers (IEEE)"},
			263,
			[]string{"10.1109", "10.1041", "10.15325"},
			842131	},
		239: PublisherMeta{
			[]string{"BMJ", "Faculty of Family Planning"},
			239,
			[]string{"10.1136", "10.1783"},
			746389	},
		266: PublisherMeta{
			[]string{"IOP Publishing", "IOP Publishing - Europhysics Letters"},
			266,
			[]string{"10.1088", "10.1209"},
			577472	},
		10: PublisherMeta{
			[]string{"American Medical Association", "American Medical Association (AMA)"},
			10,
			[]string{"10.1001"},
			574565	},
		16: PublisherMeta{
			[]string{"American Physical Society", "American Physical Society (APS)"},
			16,
			[]string{"10.1103"},
			560788	},
		317: PublisherMeta{
			[]string{"American Institute of Physics", "American Institute of Physics (AIP)", "AIP Publishing"},
			317,
			[]string{"10.1063"},
			504571	},
		194: PublisherMeta{
			[]string{"Thieme Publishing Group"},
			194,
			[]string{"10.1055"},
			435732	},
		374: PublisherMeta{
			[]string{"Walter de Gruyter GmbH", "Oldenbourg Wissenschaftsverlag", "Berkeley Electronic Press", "Walter de Gruyter GmbH (European Journal of Nanomedicine)", "Walter de Gruyter"},
			374,
			[]string{"10.1515", "10.1524", "10.2202", "10.3884", "10.2478"},
			410199	},
		292: PublisherMeta{
			[]string{"The Royal Society of Chemistry", "Royal Society of Chemistry (RSC)"},
			292,
			[]string{"10.1039"},
			409472	},
	}
	PublisherToId = map[string]int16{
		"Elsevier": 78,
		"Wiley": 311,
		"Springer": 297,
		"Informa": 301,
		"JSTOR": 1121,
		"Ovid": 276,
		"OUP": 286,
		"Sage": 179,
		"ACS": 316,
		"CUP": 56,
		"Nature": 339,
		"IEEE": 263,
		"BMJ": 239,
		"IOP": 266,
		"AMA": 10,
		"APS": 16,
		"AIP": 317,
		"Thieme": 194,
		"WalterdeGruyter": 374,
		"RSC": 292,
	}

	PrefixToId = map[string]int16{
		"10.1001": 10,
		"10.1002": 311,
		"10.1006": 78,
		"10.1007": 297,
		"10.1013": 339,
		"10.1016": 78,
		"10.1017": 56,
		"10.1021": 316,
		"10.1023": 297,
		"10.1034": 311,
		"10.1038": 339,
		"10.1039": 292,
		"10.1041": 263,
		"10.1046": 311,
		"10.1053": 78,
		"10.1054": 78,
		"10.1055": 194,
		"10.1057": 339,
		"10.1063": 317,
		"10.1065": 297,
		"10.1067": 78,
		"10.1078": 78,
		"10.1080": 301,
		"10.1081": 301,
		"10.1088": 266,
		"10.1093": 286,
		"10.1097": 276,
		"10.1103": 16,
		"10.1106": 179,
		"10.1109": 263,
		"10.1111": 311,
		"10.1112": 286,
		"10.1113": 311,
		"10.1114": 297,
		"10.1136": 239,
		"10.1140": 297,
		"10.1157": 78,
		"10.1161": 276,
		"10.1174": 301,
		"10.1177": 179,
		"10.1186": 297,
		"10.1191": 179,
		"10.1196": 311,
		"10.1197": 78,
		"10.1198": 301,
		"10.1201": 301,
		"10.1203": 276,
		"10.1205": 78,
		"10.1207": 301,
		"10.1209": 266,
		"10.1212": 276,
		"10.1213": 276,
		"10.1227": 276,
		"10.1240": 78,
		"10.1243": 179,
		"10.1245": 297,
		"10.1249": 276,
		"10.1251": 297,
		"10.1256": 311,
		"10.1258": 179,
		"10.1288": 276,
		"10.1301": 311,
		"10.1345": 179,
		"10.1348": 311,
		"10.1354": 179,
		"10.1359": 311,
		"10.1361": 297,
		"10.1367": 78,
		"10.1375": 56,
		"10.1379": 297,
		"10.1381": 297,
		"10.1383": 78,
		"10.1385": 297,
		"10.14814": 311,
		"10.1506": 311,
		"10.1510": 286,
		"10.1515": 374,
		"10.1516": 311,
		"10.1519": 276,
		"10.1524": 374,
		"10.1526": 311,
		"10.1529": 78,
		"10.1531": 301,
		"10.15325": 263,
		"10.1533": 78,
		"10.1557": 56,
		"10.1580": 78,
		"10.1581": 311,
		"10.1582": 301,
		"10.1592": 311,
		"10.1602": 78,
		"10.1617": 297,
		"10.1622": 179,
		"10.1623": 301,
		"10.1630": 179,
		"10.17322": 179,
		"10.1783": 239,
		"10.1887": 301,
		"10.1888": 301,
		"10.1892": 311,
		"10.1897": 311,
		"10.2164": 311,
		"10.2165": 276,
		"10.2167": 301,
		"10.2182": 179,
		"10.2202": 374,
		"10.2307": 1121,
		"10.2353": 78,
		"10.2459": 276,
		"10.2478": 374,
		"10.2511": 179,
		"10.2513": 301,
		"10.2746": 311,
		"10.2966": 311,
		"10.2968": 179,
		"10.2976": 301,
		"10.2989": 301,
		"10.3110": 301,
		"10.3137": 301,
		"10.3152": 301,
		"10.3155": 301,
		"10.3162": 311,
		"10.3170": 311,
		"10.3182": 78,
		"10.3187": 301,
		"10.3200": 301,
		"10.3317": 179,
		"10.3382": 286,
		"10.3401": 311,
		"10.3405": 311,
		"10.3758": 297,
		"10.3763": 301,
		"10.3815": 56,
		"10.3816": 78,
		"10.3821": 179,
		"10.3826": 301,
		"10.3858": 339,
		"10.3884": 374,
		"10.3920": 297,
		"10.3921": 78,
		"10.3939": 301,
		"10.4004": 311,
		"10.4039": 56,
		"10.4065": 78,
		"10.4098": 297,
		"10.4113": 301,
		"10.4135": 179,
		"10.4161": 301,
		"10.4219": 179,
		"10.4324": 301,
		"10.4333": 297,
		"10.5004": 301,
		"10.5034": 179,
		"10.5052": 297,
		"10.5054": 311,
		"10.5126": 179,
		"10.5175": 301,
		"10.5193": 179,
		"10.5370": 301,
		"10.5661": 301,
		"10.5819": 297,
		"10.5822": 297,
		"10.5948": 56,
		"10.5949": 56,
		"10.7123": 276,
		"10.7135": 56,
		"10.7313": 56,
		"10.7603": 297,
		"10.7811": 78,
	}

}

func GetPublisherByName(name string) (meta *PublisherMeta, err error) {
	var (
		m_id int16
		exists bool
		supportedNames []string
	)
	if m_id, exists = PublisherToId[name]; exists {
		metadata := IdToPublisher[m_id]
		return &metadata, nil
	} else {
		supportedNames = make([]string, 0, len(PublisherToId))
		for name, _ := range PublisherToId {
			supportedNames = append(supportedNames, name)
		}
		err = fmt.Errorf("Publisher '%s' not exist, options are: %v", name, supportedNames)
		return nil, err
	}
}
