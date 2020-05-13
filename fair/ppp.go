package fair

// https://www.cia.gov/library/publications/the-world-factbook/rankorder/2004rank.html
var pppGDPperCapita = map[string]float64{
	"LIECHTENSTEIN":                    139100,
	"QATAR":                            124500,
	"MONACO":                           115700,
	"MACAU":                            111600,
	"LUXEMBOURG":                       106300,
	"BERMUDA":                          99400,
	"SINGAPORE":                        93900,
	"ISLE OF MAN":                      84600,
	"BRUNEI":                           78200,
	"IRELAND":                          75500,
	"NORWAY":                           71800,
	"FALKLAND ISLANDS":                 70800,
	"UNITED ARAB EMIRATES":             67700,
	"SINT MAARTEN":                     66800,
	"KUWAIT":                           66200,
	"GIBRALTAR":                        61700,
	"HONG KONG":                        61400,
	"SWITZERLAND":                      61400,
	"UNITED STATES OF AMERICA":         59500,
	"SAN MARINO":                       58600,
	"JERSEY":                           56600,
	"SAUDI ARABIA":                     54800,
	"NETHERLANDS":                      53600,
	"GUERNSEY":                         52500,
	"ICELAND":                          51800,
	"SWEDEN":                           51500,
	"GERMANY":                          50400,
	"TAIWAN":                           50300,
	"AUSTRALIA":                        50300,
	"AUSTRIA":                          49900,
	"DENMARK":                          49900,
	"ANDORRA":                          49900,
	"BAHRAIN":                          48500,
	"CANADA":                           48300,
	"BELGIUM":                          46600,
	"SAINT PIERRE AND MIQUELON":        46200,
	"OMAN":                             45200,
	"FINLAND":                          44300,
	"UNITED KINGDOM":                   44100,
	"CAYMAN ISLANDS":                   43800,
	"FRANCE":                           43800,
	"JAPAN":                            42800,
	"MALTA":                            42000,
	"GREENLAND":                        41800,
	"EUROPEAN UNION":                   40900,
	"FAROE ISLANDS":                    40000,
	"SOUTH KOREA":                      39400,
	"NEW ZEALAND":                      38900,
	"SPAIN":                            38300,
	"ITALY":                            38100,
	"PUERTO RICO":                      37300,
	"VIRGIN ISLANDS (U.S.)":            37000,
	"CYPRUS":                           37000,
	"ISRAEL":                           36300,
	"EQUATORIAL GUINEA":                36000,
	"GUAM":                             35600,
	"CZECHIA":                          35500,
	"SLOVENIA":                         34400,
	"VIRGIN ISLANDS (BRITISH)":         34200,
	"MONTSERRAT":                       34000,
	"SLOVAKIA":                         33000,
	"LITHUANIA":                        32300,
	"ESTONIA":                          31800,
	"TRINIDAD AND TOBAGO":              31400,
	"BAHAMAS":                          31200,
	"NEW CALEDONIA":                    31100,
	"PORTUGAL":                         30400,
	"HUNGARY":                          29500,
	"POLAND":                           29500,
	"TURKS AND CAICOS ISLANDS":         29100,
	"MALAYSIA":                         29000,
	"SEYCHELLES":                       28900,
	"RUSSIA":                           27800,
	"GREECE":                           27700,
	"LATVIA":                           27600,
	"TURKEY":                           26900,
	"SAINT KITTS AND NEVIS":            26800,
	"KAZAKHSTAN":                       26300,
	"ANTIGUA AND BARBUDA":              26300,
	"PANAMA":                           25400,
	"ARUBA":                            25300,
	"NORTHERN MARIANA ISLANDS":         24500,
	"CHILE":                            24500,
	"ROMANIA":                          24500,
	"CROATIA":                          24400,
	"URUGUAY":                          22400,
	"BULGARIA":                         21700,
	"MAURITIUS":                        21600,
	"ARGENTINA":                        20900,
	"IRAN":                             20200,
	"MEXICO":                           19900,
	"LEBANON":                          19400,
	"SAINT MARTIN":                     19300,
	"GABON":                            19200,
	"MALDIVES":                         19100,
	"BELARUS":                          18900,
	"BARBADOS":                         18700,
	"TURKMENISTAN":                     18100,
	"THAILAND":                         17900,
	"BOTSWANA":                         17800,
	"MONTENEGRO":                       17700,
	"AZERBAIJAN":                       17500,
	"FRENCH POLYNESIA":                 17000,
	"IRAQ":                             17000,
	"COSTA RICA":                       16900,
	"DOMINICAN REPUBLIC":               16900,
	"COOK ISLANDS":                     16700,
	"CHINA":                            16700,
	"PALAU":                            16200,
	"BRAZIL":                           15600,
	"ALGERIA":                          15200,
	"SERBIA":                           15000,
	"CURACAO":                          15000,
	"MACEDONIA":                        14900,
	"GRENADA":                          14900,
	"SURINAME":                         14600,
	"COLOMBIA":                         14500,
	"SAINT LUCIA":                      14400,
	"SOUTH AFRICA":                     13500,
	"PERU":                             13300,
	"MONGOLIA":                         13000,
	"SRI LANKA":                        12800,
	"EGYPT":                            12700,
	"BOSNIA AND HERZEGOVINA":           12700,
	"ALBANIA":                          12500,
	"JORDAN":                           12500,
	"INDONESIA":                        12400,
	"CUBA":                             12300,
	"NAURU":                            12200,
	"ANGUILLA":                         12200,
	"VENEZUELA":                        12100,
	"TUNISIA":                          11800,
	"ECUADOR":                          11500,
	"SAINT VINCENT AND THE GRENADINES": 11500,
	"NAMIBIA":                          11300,
	"AMERICAN SAMOA":                   11200,
	"DOMINICA":                         11100,
	"GEORGIA":                          10700,
	"KOSOVO":                           10500,
	"LIBYA":                            10000,
	"ESWATINI":                         9900,
	"PARAGUAY":                         9800,
	"FIJI":                             9800,
	"ARMENIA":                          9500,
	"JAMAICA":                          9200,
	"EL SALVADOR":                      8900,
	"BHUTAN":                           8700,
	"UKRAINE":                          8700,
	"MOROCCO":                          8600,
	"BELIZE":                           8300,
	"PHILIPPINES":                      8300,
	"GUYANA":                           8200,
	"GUATEMALA":                        8100,
	"SAINT HELENA, ASCENSION, AND TRISTAN DA CUNHA": 7800,
	"BOLIVIA":                           7500,
	"LAOS":                              7400,
	"INDIA":                             7200,
	"CABO VERDE":                        6900,
	"VIETNAM":                           6900,
	"UZBEKISTAN":                        6900,
	"ANGOLA":                            6800,
	"CONGO":                             6600,
	"BURMA":                             6200,
	"NIGERIA":                           5900,
	"NICARAGUA":                         5800,
	"NIUE":                              5800,
	"SAMOA":                             5700,
	"MOLDOVA":                           5700,
	"TONGA":                             5600,
	"HONDURAS":                          5600,
	"TIMOR-LESTE":                       5400,
	"PAKISTAN":                          5400,
	"GHANA":                             4700,
	"SUDAN":                             4600,
	"MAURITANIA":                        4400,
	"WEST BANK":                         4300,
	"BANGLADESH":                        4200,
	"ZAMBIA":                            4000,
	"CAMBODIA":                          4000,
	"COTE D'IVOIRE":                     3900,
	"WALLIS AND FUTUNA":                 3800,
	"TUVALU":                            3800,
	"CAMEROON":                          3700,
	"PAPUA NEW GUINEA":                  3700,
	"KYRGYZSTAN":                        3700,
	"LESOTHO":                           3600,
	"DJIBOUTI":                          3600,
	"KENYA":                             3500,
	"MARSHALL ISLANDS":                  3400,
	"MICRONESIA":                        3400,
	"TAJIKISTAN":                        3200,
	"TANZANIA":                          3200,
	"SAO TOME AND PRINCIPE":             3200,
	"SYRIA":                             2900,
	"VANUATU":                           2700,
	"SENEGAL":                           2700,
	"NEPAL":                             2700,
	"WESTERN SAHARA":                    2500,
	"UGANDA":                            2400,
	"BENIN":                             2300,
	"ZIMBABWE":                          2300,
	"CHAD":                              2300,
	"SOLOMON ISLANDS":                   2200,
	"MALI":                              2200,
	"ETHIOPIA":                          2200,
	"RWANDA":                            2100,
	"AFGHANISTAN":                       2000,
	"KIRIBATI":                          2000,
	"GUINEA":                            2000,
	"BURKINA FASO":                      1900,
	"GUINEA-BISSAU":                     1800,
	"HAITI":                             1800,
	"GAMBIA":                            1700,
	"TOGO":                              1700,
	"NORTH KOREA":                       1700,
	"COMOROS":                           1600,
	"SIERRA LEONE":                      1600,
	"MADAGASCAR":                        1600,
	"ERITREA":                           1600,
	"SOUTH SUDAN":                       1500,
	"LIBERIA":                           1400,
	"YEMEN":                             1300,
	"MALAWI":                            1200,
	"NIGER":                             1200,
	"MOZAMBIQUE":                        1200,
	"TOKELAU":                           1000,
	"CONGO, DEMOCRATIC REPUBLIC OF THE": 800,
	"CENTRAL AFRICAN REPUBLIC":          700,
	"BURUNDI":                           700,
}

type C struct {
	Name     string
	Currency string // upper case
}

// Currency codes taken from: https://www.iban.com/currency-codes
var Ccodes = map[string]C{
	"ad": C{"andorra", "EUR"},
	"ae": C{"united arab emirates", "AED"},
	"af": C{"afghanistan", "AFN"},
	"ag": C{"antigua and barbuda", "XCD"},
	"ai": C{"anguilla", "XCD"},
	"al": C{"albania", "ALL"},
	"am": C{"armenia", "AMD"},
	"ao": C{"angola", "AOA"},
	"aq": C{"antarctica", ""},
	"ar": C{"argentina", "ARS"},
	"as": C{"american samoa", "USD"},
	"at": C{"austria", "EUR"},
	"au": C{"australia", "AUD"},
	"aw": C{"aruba", "AWG"},
	"ax": C{"åland islands", "EUR"},
	"az": C{"azerbaijan", "AZN"},
	"ba": C{"bosnia and herzegovina", "BAM"},
	"bb": C{"barbados", "BBD"},
	"bd": C{"bangladesh", "BDT"},
	"be": C{"belgium", "EUR"},
	"bf": C{"burkina faso", "XOF"},
	"bg": C{"bulgaria", "BGN"},
	"bh": C{"bahrain", "BHD"},
	"bi": C{"burundi", "BIF"},
	"bj": C{"benin", "XOF"},
	"bl": C{"saint barthélemy", "EUR"},
	"bm": C{"bermuda", "BMD"},
	"bn": C{"brunei", "BND"},
	"bo": C{"bolivia", ""}, // ?
	"bq": C{"bonaire, sint eustatius and saba", "USD"},
	"br": C{"brazil", "BRL"},
	"bs": C{"bahamas", "BSD"},
	"bt": C{"bhutan", ""}, // ?
	"bv": C{"bouvet island", "NOK"},
	"bw": C{"botswana", "BWP"},
	"by": C{"belarus", "BYN"},
	"bz": C{"belize", "BZD"},
	"ca": C{"canada", "CAD"},
	"cc": C{"cocos (keeling) islands", "AUD"},
	"cd": C{"congo, democratic republic of the", "CDF"},
	"cf": C{"central african republic", "XAF"},
	"cg": C{"congo", "XAF"},
	"ch": C{"switzerland", "CHF"},
	"ci": C{"cote d'ivoire", "XOF"},
	"ck": C{"cook islands", "NZD"},
	"cl": C{"chile", ""}, // ?
	"cm": C{"cameroon", "XAF"},
	"cn": C{"china", "CNY"},
	"co": C{"colombia", ""}, // ?
	"cr": C{"costa rica", "CRC"},
	"cu": C{"cuba", ""}, // ?
	"cv": C{"cabo verde", "CVE"},
	"cw": C{"curaçao", "ANG"},
	"cx": C{"christmas island", "AUD"},
	"cy": C{"cyprus", "EUR"},
	"cz": C{"czechia", "CZK"},
	"de": C{"germany", "EUR"},
	"dj": C{"djibouti", "DJF"},
	"dk": C{"denmark", "DKK"},
	"dm": C{"dominica", "XCD"},
	"do": C{"dominican republic", "DOP"},
	"dz": C{"algeria", "DZD"},
	"ec": C{"ecuador", "USD"},
	"ee": C{"estonia", "EUR"},
	"eg": C{"egypt", "EGP"},
	"eh": C{"western sahara", "MAD"},
	"er": C{"eritrea", "ERN"},
	"es": C{"spain", "EUR"},
	"et": C{"ethiopia", "ETB"},
	"fi": C{"finland", "EUR"},
	"fj": C{"fiji", "FJD"},
	"fk": C{"falkland islands", "FKP"},
	"fm": C{"micronesia", "USD"},
	"fo": C{"faroe islands", "DKK"},
	"fr": C{"france", "EUR"},
	"ga": C{"gabon", "XAF"},
	"gb": C{"united kingdom", "GBP"},
	"gd": C{"grenada", "XCD"},
	"ge": C{"georgia", "GEL"},
	"gf": C{"french guiana", "EUR"},
	"gg": C{"guernsey", "GBP"},
	"gh": C{"ghana", "GHS"},
	"gi": C{"gibraltar", "GIP"},
	"gl": C{"greenland", "DKK"},
	"gm": C{"gambia", "GMD"},
	"gn": C{"guinea", "GNF"},
	"gp": C{"guadeloupe", "EUR"},
	"gq": C{"equatorial guinea", "XAF"},
	"gr": C{"greece", "EUR"},
	"gs": C{"south georgia and the south sandwich islands", ""},
	"gt": C{"guatemala", "GTQ"},
	"gu": C{"guam", "USD"},
	"gw": C{"guinea-bissau", "XOF"},
	"gy": C{"guyana", "GYD"},
	"hk": C{"hong kong", "HKD"},
	"hm": C{"heard island and mcdonald islands", "AUD"},
	"hn": C{"honduras", "HNL"},
	"hr": C{"croatia", "HRK"},
	"ht": C{"haiti", ""}, // ?
	"hu": C{"hungary", "HUF"},
	"id": C{"indonesia", "IDR"},
	"ie": C{"ireland", "EUR"},
	"il": C{"israel", "ILS"},
	"im": C{"isle of man", "GBP"},
	"in": C{"india", "INR"},
	"io": C{"british indian ocean territory", "USD"},
	"iq": C{"iraq", "IQD"},
	"ir": C{"iran", "IRR"},
	"is": C{"iceland", "ISK"},
	"it": C{"italy", "EUR"},
	"je": C{"jersey", "GBP"},
	"jm": C{"jamaica", "JMD"},
	"jo": C{"jordan", "JOD"},
	"jp": C{"japan", "JPY"},
	"ke": C{"kenya", "KES"},
	"kg": C{"kyrgyzstan", "KGS"},
	"kh": C{"cambodia", "KHR"},
	"ki": C{"kiribati", "AUD"},
	"km": C{"comoros", "KMF"},
	"kn": C{"saint kitts and nevis", "XCD"},
	"kp": C{"north korea", "KPW"},
	"kr": C{"south korea", "KRW"},
	"kw": C{"kuwait", "KWD"},
	"ky": C{"cayman islands", "KYD"},
	"kz": C{"kazakhstan", "KZT"},
	"la": C{"laos", "LAK"},
	"lb": C{"lebanon", "LBP"},
	"lc": C{"saint lucia", "XCD"},
	"li": C{"liechtenstein", "CHF"},
	"lk": C{"sri lanka", "LKR"},
	"lr": C{"liberia", "LRD"},
	"ls": C{"lesotho", ""}, // ?
	"lt": C{"lithuania", "EUR"},
	"lu": C{"luxembourg", "EUR"},
	"lv": C{"latvia", "EUR"},
	"ly": C{"libya", "LYD"},
	"ma": C{"morocco", "MAD"},
	"mc": C{"monaco", "EUR"},
	"md": C{"moldova", ""}, // ?
	"me": C{"montenegro", "EUR"},
	"mf": C{"saint martin (french part)", "EUR"},
	"mg": C{"madagascar", "MGA"},
	"mh": C{"marshall islands", "USD"},
	"mk": C{"north macedonia", "MKD"},
	"ml": C{"mali", "XOF"},
	"mm": C{"myanmar", "MMK"},
	"mn": C{"mongolia", "MNT"},
	"mo": C{"macao", "MOP"},
	"mp": C{"northern mariana islands", "USD"},
	"mq": C{"martinique", "EUR"},
	"mr": C{"mauritania", "MRU"},
	"ms": C{"montserrat", "XCD"},
	"mt": C{"malta", "EUR"},
	"mu": C{"mauritius", "MUR"},
	"mv": C{"maldives", "MVR"},
	"mw": C{"malawi", "MWK"},
	"mx": C{"mexico", "MXN"},
	"my": C{"malaysia", "MYR"},
	"mz": C{"mozambique", "MZN"},
	"na": C{"namibia", ""}, // ?
	"nc": C{"new caledonia", "XPF"},
	"ne": C{"niger", "XOF"},
	"nf": C{"norfolk island", "AUD"},
	"ng": C{"nigeria", "NGN"},
	"ni": C{"nicaragua", "NIO"},
	"nl": C{"netherlands", "EUR"},
	"no": C{"norway", "NOK"},
	"np": C{"nepal", "NPR"},
	"nr": C{"nauru", "AUD"},
	"nu": C{"niue", "NZD"},
	"nz": C{"new zealand", "NZD"},
	"om": C{"oman", "OMR"},
	"pa": C{"panama", ""}, // ?
	"pe": C{"peru", "PEN"},
	"pf": C{"french polynesia", "XPF"},
	"pg": C{"papua new guinea", "PGK"},
	"ph": C{"philippines", "PHP"},
	"pk": C{"pakistan", "PKR"},
	"pl": C{"poland", "PLN"},
	"pm": C{"saint pierre and miquelon", "EUR"},
	"pn": C{"pitcairn", "NZD"},
	"pr": C{"puerto rico", "USD"},
	"ps": C{"palestine", ""},
	"pt": C{"portugal", "EUR"},
	"pw": C{"palau", "USD"},
	"py": C{"paraguay", "PYG"},
	"qa": C{"qatar", "QAR"},
	"re": C{"réunion", "EUR"},
	"ro": C{"romania", "RON"},
	"rs": C{"serbia", "RSD"},
	"ru": C{"russia", "RUB"},
	"rw": C{"rwanda", "RWF"},
	"sa": C{"saudi arabia", "SAR"},
	"sb": C{"solomon islands", "SBD"},
	"sc": C{"seychelles", "SCR"},
	"sd": C{"sudan", "SDG"},
	"se": C{"sweden", "SEK"},
	"sg": C{"singapore", "SGD"},
	"sh": C{"saint helena, ascension and tristan da cunha", "SHP"},
	"si": C{"slovenia", "EUR"},
	"sj": C{"svalbard and jan mayen", "NOK"},
	"sk": C{"slovakia", "EUR"},
	"sl": C{"sierra leone", "SLL"},
	"sm": C{"san marino", "EUR"},
	"sn": C{"senegal", "XOF"},
	"so": C{"somalia", "SOS"},
	"sr": C{"suriname", "SRD"},
	"ss": C{"south sudan", "SSP"},
	"st": C{"sao tome and principe", "STN"},
	"sv": C{"el salvador", ""}, // ?
	"sx": C{"sint maarten (dutch part)", "ANG"},
	"sy": C{"syria", "SYP"},
	"sz": C{"eswatini", ""}, // ?
	"tc": C{"turks and caicos islands", "USD"},
	"td": C{"chad", "XAF"},
	"tf": C{"french southern territories", "EUR"},
	"tg": C{"togo", "XOF"},
	"th": C{"thailand", "THB"},
	"tj": C{"tajikistan", "TJS"},
	"tk": C{"tokelau", "NZD"},
	"tl": C{"timor-leste", "USD"},
	"tm": C{"turkmenistan", "TMT"},
	"tn": C{"tunisia", "TND"},
	"to": C{"tonga", "TOP"},
	"tr": C{"turkey", "TRY"},
	"tt": C{"trinidad and tobago", "TTD"},
	"tv": C{"tuvalu", "AUD"},
	"tw": C{"taiwan", "TWD"},
	"tz": C{"tanzania", "TZS"},
	"ua": C{"ukraine", "UAH"},
	"ug": C{"uganda", "UGX"},
	"um": C{"united states minor outlying islands", "USD"},
	"us": C{"united states of america", "USD"},
	"uy": C{"uruguay", ""}, // ?
	"uz": C{"uzbekistan", "UZS"},
	"va": C{"holy see", "EUR"},
	"vc": C{"saint vincent and the grenadines", "XCD"},
	"ve": C{"venezuela", "VEF"},
	"vg": C{"virgin islands (british)", "USD"},
	"vi": C{"virgin islands (u.s.)", "USD"},
	"vn": C{"vietnam", "VND"},
	"vu": C{"vanuatu", "VUV"},
	"wf": C{"wallis and futuna", "XPF"},
	"ws": C{"samoa", "WST"},
	"ye": C{"yemen", "YER"},
	"yt": C{"mayotte", "EUR"},
	"za": C{"south africa", "ZAR"},
	"zm": C{"zambia", "ZMW"},
	"zw": C{"zimbabwe", "ZWL"},
}
