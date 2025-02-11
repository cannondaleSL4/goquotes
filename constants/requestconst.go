package constants

import (
	tinkoff "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
	"time"
)

type RequestData struct {
	From       time.Time
	To         time.Time
	Resolution tinkoff.CandleInterval
	FIGI       string
}

const (
	INDEX       = "./templates/index.html"
	INSTRUMENTS = "./templates/instr.html"
	FIGI        = "FIGI"
	Ticker      = "TICKER"
	Name        = "Name"
)

var QuotesMapDJ = []map[string]string{
	{FIGI: "BBG000B9XRY4", Ticker: "AAPL", Name: "Apple"},
	{FIGI: "BBG000BBQCY0", Ticker: "AMD", Name: "AMD"},
	{FIGI: "BBG000BSJK37", Ticker: "T", Name: "AT&T"},
	{FIGI: "BBG005P7Q881", Ticker: "AAL", Name: "American Airlines Group"},
	{FIGI: "BBG006G2JVL2", Ticker: "BABA", Name: "Alibaba M"},
	{FIGI: "BBG000BCSST7", Ticker: "BA", Name: "BOEING"},
	{FIGI: "BBG000BCTLF6", Ticker: "BAC", Name: "Bank of America Corp"},
	{FIGI: "BBG000K4ND22", Ticker: "CVX", Name: "Chevron"},
	{FIGI: "BBG00BN96922", Ticker: "DOW", Name: "Dow inc"},
	{FIGI: "BBG000GZQ728", Ticker: "XOM", Name: "Exxon Mobil Corporation"},
	{FIGI: "BBG000C6CFJ5", Ticker: "GS", Name: "Goldman Sachs"},
	{FIGI: "BBG000BPD168", Ticker: "MRK", Name: "Merck & Co."},
	{FIGI: "BBG000BJ81C1", Ticker: "TRV", Name: "Travelers Cos"},
	{FIGI: "BBG000BWXBC2", Ticker: "WMT", Name: "Wal-Mart Stores"},
	{FIGI: "BBG000BKZB36", Ticker: "HD", Name: "The Home Depot"},
	{FIGI: "BBG000HS77T5", Ticker: "VZ", Name: "Verizon Communications"},
	{FIGI: "BBG000C3J3C9", Ticker: "CSCO", Name: "Cisco"},
	{FIGI: "BBG000BBJQV0", Ticker: "NVDA", Name: "NVIDIA MinPriceIncrement"},
	{FIGI: "BBG000BLNNH6", Ticker: "IBM", Name: "IBM"},
	{FIGI: "BBG000BLZRJ2", Ticker: "MS", Name: "Morgan Stanley"},
	{FIGI: "BBG000C43RR5", Ticker: "EBAY", Name: "eBay"},
	{FIGI: "BBG000BR2B91", Ticker: "PFE", Name: "Pfizer"},
	{FIGI: "BBG000PSKYX7", Ticker: "V", Name: "Visa"},
	{FIGI: "BBG000BP52R2", Ticker: "MMM", Name: "3M Company"},
	{FIGI: "BBG000BMX289", Ticker: "KO", Name: "COCA-COLA"},
	{FIGI: "BBG000FY4S11", Ticker: "C", Name: "Citigroup"},
	{FIGI: "BBG000BLBVN4", Ticker: "BKNG", Name: "Booking"},
	{FIGI: "BBG000BD8PN9", Ticker: "BK", Name: "Bank of New York Mellon"},
	{FIGI: "BBG000BWQFY7", Ticker: "WFC", Name: "Wells Fargo & Company"},
	{FIGI: "BBG000C5HS04", Ticker: "NKE", Name: "NIKE"},
	{FIGI: "BBG00BN961G4", Ticker: "DD", Name: "DuPont de Nemours Inc"},
	{FIGI: "BBG000BR2TH3", Ticker: "PG", Name: "Procter & Gamble"},
	{FIGI: "BBG000BQLTW7", Ticker: "ORCL", Name: "Oracle"},
	{FIGI: "BBG000BMHYD1", Ticker: "JNJ", Name: "Johnson & Johnson"},
	{FIGI: "BBG000F1ZSQ2", Ticker: "MA", Name: "Mastercard"},
	{FIGI: "BBG000BPH459", Ticker: "MSFT", Name: "Microsoft Corporation"},
	{FIGI: "BBG000C0G1D1", Ticker: "INTC", Name: "Intel Corporation"},
	{FIGI: "BBG000DMBXR2", Ticker: "JPM", Name: "JPMorgan"},
	{FIGI: "BBG000BCQZS4", Ticker: "AXP", Name: "American Express"},
	{FIGI: "BBG000BNSZP1", Ticker: "MCD", Name: "Mc'DONALDS"},
	{FIGI: "BBG000D9KWL9", Ticker: "MTB", Name: "M&T Bank"},
	{FIGI: "BBG000BF0K17", Ticker: "CAT", Name: "Caterpillar"},
	{FIGI: "BBG000BH4R78", Ticker: "DIS", Name: "Walt Disney"},
	{FIGI: "BBG000DH7JK6", Ticker: "PEP", Name: "PepsiCo"},
	{FIGI: "BBG000CH5208", Ticker: "UNH", Name: "UnitedHealth"},
	{FIGI: "BBG000BWLMJ4", Ticker: "WBA", Name: "Walgreens Boots Alliance"},
	{FIGI: "BBG000BB5373", Ticker: "WU", Name: "Western Union"},
	{FIGI: "BBG000BGPTV6", Ticker: "CMI", Name: "Cummins"},
	{FIGI: "BBG000R7Z112", Ticker: "DAL", Name: "Delta Air Lines"},

	//{FIGI: "BBG000C4LN67", Ticker: "GRMN", Name: "Garmin"},
	//{FIGI: "BBG000BKZTP3", Ticker: "HOG", Name: "Harley-Davidson"},
	//{FIGI: "BBG000BRXP69", Ticker: "EW", Name: "Edwards Lifesciences"},
	//{FIGI: "BBG000BJF1Z8", Ticker: "FDX", Name: "FedEx"},
	//{FIGI: "BBG000BCTQ65", Ticker: "XEL", Name: "Xcel Energy"},
	//{FIGI: "BBG000BD2NY8", Ticker: "BF.B", Name: "Brown-Forman"},
	//{FIGI: "BBG000BSLZY7", Ticker: "SCHW", Name: "The Charles Schwab Corporation"},
	//{FIGI: "BBG000M65M61", Ticker: "UAL", Name: "United Airlines Holdings"},
	//{FIGI: "BBG000BN84F3", Ticker: "DGX", Name: "Quest Diagnostics"},
	//{FIGI: "BBG000BGVW60", Ticker: "D", Name: "Dominion Energy Inc"},
	//{FIGI: "BBG000CGJMB9", Ticker: "ESS", Name: "Essex Property Trust"},
	//{FIGI: "BBG000BJX8C8", Ticker: "NOV", Name: "National Oilwell Varco"},
	//{FIGI: "BBG000BQ5DS5", Ticker: "NSC", Name: "Norfolk Southern"},
	//{FIGI: "BBG000BRJ809", Ticker: "PPG", Name: "PPG Industries"},
	//{FIGI: "BBG000BRD0D8", Ticker: "PNC", Name: "PNC Financial Services"},
	//{FIGI: "BBG000BR14K5", Ticker: "CB", Name: "Chubb"},
	//{FIGI: "BBG000BLHRS2", Ticker: "HSY", Name: "Hershey"},
	//{FIGI: "BBG000BQ8KV2", Ticker: "NUE", Name: "Nucor"},
	//{FIGI: "BBG000BX3BL3", Ticker: "WY", Name: "Weyerhaeuser"},
	//{FIGI: "BBG000BH3GZ2", Ticker: "YUM", Name: "Yum!"},
	//{FIGI: "BBG000BRNLL2", Ticker: "ALGN", Name: "Align Technology"},
	//{FIGI: "BBG005CPNTQ2", Ticker: "KHC", Name: "Kraft Heinz"},
	//{FIGI: "BBG000G8N9C6", Ticker: "JWN", Name: "Nordstrom"},
	//{FIGI: "BBG000BR37X2", Ticker: "PGR", Name: "Progressive"},
	//{FIGI: "BBG000BK2F42", Ticker: "FIS", Name: "Fidelity National Information"},
	//{FIGI: "BBG000BGRY34", Ticker: "CVS", Name: "CVS Health Corporation"},
	//{FIGI: "BBG000KHWT55", Ticker: "HPQ", Name: "HP"},
	//{FIGI: "BBG000BNNKG9", Ticker: "MAS", Name: "Masco"},
	//{FIGI: "BBG000BS0ZF1", Ticker: "RL", Name: "Ralph Lauren"},
	//{FIGI: "BBG000BLXZN1", Ticker: "TSCO", Name: "Tractor Supply"},
	//{FIGI: "BBG000BLPBL5", Ticker: "AVB", Name: "AvalonBay Communities"},
	//{FIGI: "BBG000B9ZXB4", Ticker: "ABT", Name: "Abbott"},
	//{FIGI: "BBG000BJPDZ1", Ticker: "ISRG", Name: "Intuitive Surgical"},
	//{FIGI: "BBG000BWNFZ9", Ticker: "WDC", Name: "Western Digital"},
	//{FIGI: "BBG000D8RG11", Ticker: "NRG", Name: "NRG Energy"},
	//{FIGI: "BBG000B9Z0J8", Ticker: "PLD", Name: "Prologis"},
	//{FIGI: "BBG000BNJHS8", Ticker: "LUV", Name: "Southwest Airlines"},
	//{FIGI: "BBG000CKGBP2", Ticker: "GILD", Name: "GILEAD"},
	//{FIGI: "BBG000RGM5P1", Ticker: "TEL", Name: "TE Connectivity"},
	//{FIGI: "BBG000BZX1N5", Ticker: "AIZ", Name: "Assurant"},
	//{FIGI: "BBG000C90DH9", Ticker: "IPG", Name: "IPG"},
	//{FIGI: "BBG000BBDZG3", Ticker: "AIG", Name: "American International Group"},
	//{FIGI: "BBG000B9YJ35", Ticker: "APH", Name: "Amphenol"},
	//{FIGI: "BBG000BXMFC3", Ticker: "URI", Name: "United Rentals"},
	//{FIGI: "BBG000CB8Q50", Ticker: "UHS", Name: "Universal Health Services"},
	//{FIGI: "BBG000C1S2X2", Ticker: "VRTX", Name: "Vertex Pharmaceuticals"},
	//{FIGI: "BBG0088CYPX8", Ticker: "NLSN", Name: "Nielsen Holdings plc"},
	//{FIGI: "BBG000BCVJ77", Ticker: "BAX", Name: "Baxter International"},
	//{FIGI: "BBG000D898T9", Ticker: "CAH", Name: "Cardinal Health"},
	//{FIGI: "BBG000BDXCJ5", Ticker: "CNC", Name: "Centene Corporation"},
	//{FIGI: "BBG000BS7KS3", Ticker: "TAP", Name: "Molson Coors Brewing"},

	//{FIGI: "BBG003PS7JV1", Ticker: "ALLE", Name: "Allegion"},
	//{FIGI: "BBG001D8R5D0", Ticker: "XYL", Name: "Xylem"},
	//{FIGI: "BBG000N2HDY5", Ticker: "ANET", Name: "Arista Networks Inc"},
	//{FIGI: "BBG000BBDV81", Ticker: "CTSH", Name: "Cognizant Technology Solutions"},
	//{FIGI: "BBG000C45984", Ticker: "L", Name: "Loews"},
	//{FIGI: "BBG0035M2ZB7", Ticker: "NWS", Name: "News Corp"},
	//{FIGI: "BBG000BWGYF8", Ticker: "VMC", Name: "Vulcan Materials"},
	//{FIGI: "BBG000BLMZK6", Ticker: "KMX", Name: "CarMax"},
	//{FIGI: "BBG00JHNKJY8", Ticker: "FOX", Name: "Twenty-First Century Fox"},
	//{FIGI: "BBG000PV27K3", Ticker: "EXR", Name: "Extra Space Storage"},
	//{FIGI: "BBG000BMW2Z0", Ticker: "KMB", Name: "Kimberly-Clark"},
	//{FIGI: "BBG000L9CV04", Ticker: "UPS", Name: "United Parcel Service"},
	//{FIGI: "BBG000BWVSR1", Ticker: "WM", Name: "Waste Management"},
	//{FIGI: "BBG000FQRVM3", Ticker: "WAT", Name: "Waters"},
	//{FIGI: "BBG000D9D830", Ticker: "ACN", Name: "Accenture"},
	//{FIGI: "BBG000BM7HL0", Ticker: "ADSK", Name: "Autodesk"},
	//{FIGI: "BBG000BG4202", Ticker: "CPB", Name: "Campbell Soup"},
	//{FIGI: "BBG000BLW530", Ticker: "EMN", Name: "Eastman Chemical Company"},

	//{FIGI: "BBG000BHPL78", Ticker: "EFX", Name: "Equifax"},
	//{FIGI: "BBG000BM6788", Ticker: "IR", Name: "Ingersoll-Rand"},
	//{FIGI: "BBG000BZCKH3", Ticker: "MTD", Name: "Mettler Toledo"},
	//{FIGI: "BBG000BB8SW7", Ticker: "PKG", Name: "Packaging Corp of America"},
	//{FIGI: "BBG000BPPN67", Ticker: "PSA", Name: "Public Storage"},
	//{FIGI: "BBG000BB5792", Ticker: "RCL", Name: "Royal Caribbean Cruises"},
	//{FIGI: "BBG000BQQH30", Ticker: "COP", Name: "ConocoPhillips"},
	//{FIGI: "BBG000CHWP52", Ticker: "DISCA", Name: "Discovery Communications (A)"},
	//{FIGI: "BBG006Q0HY77", Ticker: "CFG", Name: "Citizens Financial Group"},
	//{FIGI: "BBG000BN53G7", Ticker: "LEG", Name: "Leggett & Platt"},
	//{FIGI: "BBG00658F3P3", Ticker: "SYF", Name: "Synchrony Financial"},
	//{FIGI: "BBG00FWQ4VD6", Ticker: "ULTA", Name: "Ulta Beauty"},
	//{FIGI: "BBG000BCZYD3", Ticker: "BDX", Name: "Becton, Dickinson and Company"},
	//{FIGI: "BBG000BBD070", Ticker: "HES", Name: "Hess"},
	//{FIGI: "BBG000C1FB75", Ticker: "ICE", Name: "Intercontinental Exchange"},
	//{FIGI: "BBG000WCFV84", Ticker: "LYB", Name: "LyondellBasell"},
	//{FIGI: "BBG000BLDV98", Ticker: "HRB", Name: "H&R Block"},
	//{FIGI: "BBG000BT9DW0", Ticker: "SO", Name: "Southern"},
	//{FIGI: "BBG000Q3JN03", Ticker: "RF", Name: "Regions Financial"},
	//{FIGI: "BBG000BKL348", Ticker: "GPC", Name: "Genuine Parts"},
	//{FIGI: "BBG000BR3KL6", Ticker: "PH", Name: "Parker-Hannifin"},
	//{FIGI: "BBG000BTVJ25", Ticker: "SYY", Name: "Sysco"},
	//{FIGI: "BBG000BVVQQ8", Ticker: "TXT", Name: "Textron"},
	//{FIGI: "BBG001QD41M9", Ticker: "APTV", Name: "Aptiv"},
	//{FIGI: "BBG000PPFKQ7", Ticker: "BR", Name: "Broadridge Financial Solutions"},
	//{FIGI: "BBG000BQVTF5", Ticker: "PCAR", Name: "PACCAR"},
	//{FIGI: "BBG000BJL3N0", Ticker: "FITB", Name: "Fifth Third Bancorp"},
	//{FIGI: "BBG000BB5373", Ticker: "WU", Name: "Western Union"},
	//{FIGI: "BBG000BFPK65", Ticker: "CINF", Name: "Cincinnati Financial Corporation"},
	//{FIGI: "BBG000BPXVJ6", Ticker: "RSG", Name: "Republic Services"},
	//{FIGI: "BBG006G063F9", Ticker: "INFO", Name: "IHS Markit"},
	//{FIGI: "BBG000RTDY25", Ticker: "MSCI", Name: "MSCI Inc"},
	//{FIGI: "BBG000BKFBD7", Ticker: "STT", Name: "State Street"},
	//{FIGI: "BBG000BV8DN6", Ticker: "TJX", Name: "TJX"},
	//{FIGI: "BBG000BJ8YN7", Ticker: "FAST", Name: "Fastenal Company"},
	//{FIGI: "BBG000C13CD9", Ticker: "CDNS", Name: "Cadence Design Systems"},
	//{FIGI: "BBG000BP0KQ8", Ticker: "EA", Name: "Electronic Arts"},
	//{FIGI: "BBG000BCMBG4", Ticker: "FTNT", Name: "Fortinet Inc"},
	//{FIGI: "BBG000DCGRL8", Ticker: "IPGP", Name: "IPG Photonics Corporation"},
	//{FIGI: "BBG000BB6KF5", Ticker: "MET", Name: "MetLife"},
	//{FIGI: "BBG000BQC9V2", Ticker: "NWL", Name: "Newell Brands"},
	//{FIGI: "BBG000Q5ZRM7", Ticker: "DLR", Name: "Digital Realty"},
	//{FIGI: "BBG000BQ2C28", Ticker: "NOC", Name: "Northrop Grumman"},
	//{FIGI: "BBG000MBDGM6", Ticker: "EQIX", Name: "Equinix"},
	//{FIGI: "BBG000BJKPG0", Ticker: "FISV", Name: "Fiserv"},
	//{FIGI: "BBG000F7RCJ1", Ticker: "AAP", Name: "Advance Auto Parts"},
	//{FIGI: "BBG000C0LW92", Ticker: "BSX", Name: "Boston Scientific Corporation"},
	//{FIGI: "BBG000BF6LY3", Ticker: "CCL", Name: "Carnival"},
	//{FIGI: "BBG000QBR5J5", Ticker: "DFS", Name: "Discover Financial Services"},
	//{FIGI: "BBG000BW3299", Ticker: "UNP", Name: "Union Pacific Corporation"},

	//{FIGI: "BBG00FN64XT9", Ticker: "DXC", Name: "DXC Technology"},

	//{FIGI: "BBG000C3GN47", Ticker: "COG", Name: "Cabot Oil & Gas"},
	//{FIGI: "BBG000BNPSQ9", Ticker: "INCY", Name: "Incyte"},
	//{FIGI: "BBG000L4M7F1", Ticker: "RMD", Name: "ResMed Inc"},
	//{FIGI: "BBG000BY29C7", Ticker: "TPR", Name: "Tapestry"},
	//{FIGI: "BBG000BJQWD2", Ticker: "AKAM", Name: "Akamai Technologies"},
	//{FIGI: "BBG000DWG505", Ticker: "BRK.B", Name: "Berkshire Hathaway"},
	//{FIGI: "BBG000BFJT36", Ticker: "CHD", Name: "Church & Dwight"},
	//{FIGI: "BBG000BHKYH4", Ticker: "ECL", Name: "Ecolab"},
	//{FIGI: "BBG000BJ2D31", Ticker: "SPG", Name: "Simon Property Group"},
	//{FIGI: "BBG001M8HHB7", Ticker: "TRIP", Name: "TripAdvisor"},
	//{FIGI: "BBG000H3YXF8", Ticker: "CTAS", Name: "Cintas"},
	//{FIGI: "BBG000BKTFN2", Ticker: "HAL", Name: "Halliburton"},
	//{FIGI: "BBG000BHCP19", Ticker: "MCHP", Name: "Microchip Technology"},
	//{FIGI: "BBG000BV59Y6", Ticker: "TFX", Name: "Teleflex Inc"},
	//{FIGI: "BBG000BX57K1", Ticker: "XRAY", Name: "Dentsply Sirona"},
	//{FIGI: "BBG000C2ZCH8", Ticker: "SRE", Name: "Sempra Energy"},
	//{FIGI: "BBG00DL8NMV2", Ticker: "FTI", Name: "TechnipFMC"},
	//{FIGI: "BBG000BV75B7", Ticker: "TIF", Name: "Tiffany & Co"},
	//{FIGI: "BBG000BNGTQ7", Ticker: "LB", Name: "L Brands"},
	//{FIGI: "BBG000BJBZ23", Ticker: "MKTX", Name: "MarketAxess Holdings Inc"},
	//{FIGI: "BBG000BGKTF9", Ticker: "COF", Name: "Capital One"},
	//{FIGI: "BBG000BJP882", Ticker: "FMC", Name: "FMC"},
	//{FIGI: "BBG000C46HM9", Ticker: "M", Name: "Macy's"},
	//{FIGI: "BBG000BVPV84", Ticker: "AMZN", Name: "Amazon.com"},
	//{FIGI: "BBG000BS5CM9", Ticker: "BXP", Name: "Boston Properties"},
	//{FIGI: "BBG000FKJRC5", Ticker: "EL", Name: "The Estee Lauder Companies"},
	//{FIGI: "BBG000G0Z878", Ticker: "HIG", Name: "Hartford Financial Services"},
	//{FIGI: "BBG0059FN811", Ticker: "KEYS", Name: "Keysight Technologies Inc"},
	//{FIGI: "BBG000BNFLM9", Ticker: "LRCX", Name: "Lam Research"},
	//{FIGI: "BBG008NVB1C0", Ticker: "MNST", Name: "Monster Beverage"},
	//{FIGI: "BBG000B9XYV2", Ticker: "AMT", Name: "American Tower"},
	//{FIGI: "BBG000C2PW58", Ticker: "BLK", Name: "BlackRock"},
	//{FIGI: "BBG000BB6G37", Ticker: "ADI", Name: "Analog Devices"},
	//{FIGI: "BBG0035LY913", Ticker: "NWSA", Name: "News Corp (A)"},
	//{FIGI: "BBG000C060M4", Ticker: "AMG", Name: "Affiliated Managers Group"},
	//{FIGI: "BBG000BVMGF2", Ticker: "ALL", Name: "The Allstate Corporation"},
	//{FIGI: "BBG000BP6LJ8", Ticker: "MO", Name: "Altria Group"},
	//{FIGI: "BBG000BPQD31", Ticker: "MYL", Name: "Mylan NV"},
	//{FIGI: "BBG000BL9JQ1", Ticker: "HFC", Name: "HollyFrontier"},
	//{FIGI: "BBG000BPNP00", Ticker: "MXIM", Name: "Maxim Integrated Products Inc"},
	//{FIGI: "BBG000BT0CM2", Ticker: "SIVB", Name: "SVB Financial Group"},

	//{FIGI: "BBG000BFVXX0", Ticker: "CMS", Name: "CMS Energy Corporation"},
	//{FIGI: "BBG000NV1KK7", Ticker: "DG", Name: "Dollar General"},
	//{FIGI: "BBG00BLVZ228", Ticker: "FTV", Name: "Fortive"},
	//{FIGI: "BBG000BMY992", Ticker: "KR", Name: "Kroger"},
	//{FIGI: "BBG000C1XVK6", Ticker: "RE", Name: "Everest Re Group"},
	//{FIGI: "BBG00333FYS2", Ticker: "IQV", Name: "IQVIA Holdings Inc"},
	//{FIGI: "BBG000BMTFR4", Ticker: "KLAC", Name: "KLA-Tencor"},
	//{FIGI: "BBG000C41023", Ticker: "UDR", Name: "UDR"},
	//{FIGI: "BBG000R7RDB4", Ticker: "CXO", Name: "Concho Resources"},
	//{FIGI: "BBG000DQTXY6", Ticker: "DHI", Name: "D.R. Horton"},
	//{FIGI: "BBG000QY3XZ2", Ticker: "EXPE", Name: "Expedia"},
	//{FIGI: "BBG000J2XL74", Ticker: "PM", Name: "Philip Morris"},
	//{FIGI: "BBG000VPGNR2", Ticker: "CHTR", Name: "Charter Communications"},
	//{FIGI: "BBG000D7RKJ5", Ticker: "EIX", Name: "Edison International"},
	//{FIGI: "BBG000BXM6V2", Ticker: "UAA", Name: "Under Armour"},
	//{FIGI: "BBG000BL8804", Ticker: "HST", Name: "Host Hotels & Resorts"},
	//{FIGI: "BBG000BN5HF7", Ticker: "LEN", Name: "Lennar"},
	//{FIGI: "BBG00KHY5S69", Ticker: "AVGO", Name: "Broadcom"},
	//{FIGI: "BBG000C1BW00", Ticker: "LMT", Name: "Lockheed Martin"},
	//{FIGI: "BBG000F0KF42", Ticker: "STX", Name: "Seagate"},
	//{FIGI: "BBG000BVP5P2", Ticker: "SLG", Name: "SL Green Realty"},
	//{FIGI: "BBG000BMQPL1", Ticker: "KEY", Name: "KeyCorp"},
	//{FIGI: "BBG000C8H633", Ticker: "MRO", Name: "Marathon Oil"},
	//{FIGI: "BBG000BP1Q11", Ticker: "SPGI", Name: "S&P Global"},
	//{FIGI: "BBG000BWBZN1", Ticker: "VAR", Name: "Varian Medical Systems"},
	//{FIGI: "BBG000BHLYS1", Ticker: "ED", Name: "Consolidated Edison"},
	//{FIGI: "BBG000CXYSZ6", Ticker: "FFIV", Name: "F5 Networks"},
	//{FIGI: "BBG000NDYB67", Ticker: "GM", Name: "General Motors"},
	//{FIGI: "BBG000BLSL58", Ticker: "IFF", Name: "International Flavors & Fragrances"},
	//{FIGI: "BBG000BQZMH4", Ticker: "PEG", Name: "Public Service Enterprise Group"},
	//{FIGI: "BBG000C23KJ3", Ticker: "AES", Name: "The AES Corporation"},
	//{FIGI: "BBG000BDDNH5", Ticker: "BLL", Name: "Ball Corporation"},

	//{FIGI: "BBG000BMDBZ1", Ticker: "JBHT", Name: "JB Hunt"},
	//{FIGI: "BBG000BT41Q8", Ticker: "SLB", Name: "Schlumberger"},
	//{FIGI: "BBG000B9XG87", Ticker: "AME", Name: "AMETEK"},
	//{FIGI: "BBG000GXZ4W7", Ticker: "ANSS", Name: "ANSYS"},
	//{FIGI: "BBG000BKLH74", Ticker: "GPS", Name: "Gap"},
	//{FIGI: "BBG000BC2C10", Ticker: "APA", Name: "Apache"},
	//{FIGI: "BBG000BKVJK4", Ticker: "HAS", Name: "Hasbro"},
	//{FIGI: "BBG000BY2Y78", Ticker: "IVZ", Name: "Invesco"},
	//{FIGI: "BBG00BN969C1", Ticker: "CTVA", Name: "Corteva Inc"},
	//{FIGI: "BBG000QW7VC1", Ticker: "HCA", Name: "HCA Healthcare"},
	//{FIGI: "BBG000C7LMS8", Ticker: "AZO", Name: "AutoZone"},
	//{FIGI: "BBG000BBCDZ2", Ticker: "ROK", Name: "Rockwell Automation"},
	//{FIGI: "BBG000C3NTN5", Ticker: "HOLX", Name: "Hologic"},
	//{FIGI: "BBG000KLB4Q1", Ticker: "SWKS", Name: "Skyworks Solutions"},
	//{FIGI: "BBG000BWCKB6", Ticker: "VFC", Name: "VF"},
	//{FIGI: "BBG000DB3KT1", Ticker: "WLTW", Name: "Willis Towers Watson"},
	//{FIGI: "BBG000BC1L02", Ticker: "AOS", Name: "A. O. Smith Corporation"},
	//{FIGI: "BBG000BGYWY6", Ticker: "ORLY", Name: "O'Reilly Automotive"},
	//{FIGI: "BBG000BN2DC2", Ticker: "CRM", Name: "Salesforce"},

	//{FIGI: "BBG000BC15S7", Ticker: "AON", Name: "Aon PLC"},
	//{FIGI: "BBG000GPXKX9", Ticker: "FLT", Name: "FleetCor Technologies Inc"},
	//{FIGI: "BBG000CN3S73", Ticker: "KIM", Name: "Kimco Realty"},
	//{FIGI: "BBG000BSXQV7", Ticker: "SHW", Name: "Sherwin-Williams"},
	//{FIGI: "BBG000BWP7D9", Ticker: "WEC", Name: "WEC Energy Group"},
	//{FIGI: "BBG000BLCPY4", Ticker: "HP", Name: "Helmerich & Payne"},
	//{FIGI: "BBG0058KMH30", Ticker: "HLT", Name: "Hilton"},
	//{FIGI: "BBG00GVR8YQ9", Ticker: "LIN", Name: "Linde"},
	//{FIGI: "BBG000BVDLH9", Ticker: "TMO", Name: "Thermo Fisher Scientific"},
	//{FIGI: "BBG000DKCC19", Ticker: "TSN", Name: "Tyson Foods"},
	//{FIGI: "BBG000BWVCP8", Ticker: "WMB", Name: "The Williams Companies"},
	//{FIGI: "BBG009S39JX6", Ticker: "GOOGL", Name: "Alphabet Class A"},
	//{FIGI: "BBG009S3NB30", Ticker: "GOOG", Name: "Google"},
	//{FIGI: "BBG000D6L294", Ticker: "XEC", Name: "Cimarex Energy"},
	//{FIGI: "BBG009DTD8H2", Ticker: "UA", Name: "Under Armour (C)"},
	//{FIGI: "BBG000DD3510", Ticker: "AIV", Name: "Apartment Investment & Management"},
	//{FIGI: "BBG000BJ49H3", Ticker: "BWA", Name: "BorgWarner"},
	//{FIGI: "BBG000BFNR17", Ticker: "ADS", Name: "Alliance Data Systems"},
	//{FIGI: "BBG000BRJL00", Ticker: "PPL", Name: "PPL"},
	//{FIGI: "BBG000H556T9", Ticker: "HON", Name: "Honeywell"},
	//{FIGI: "BBG000PXDL44", Ticker: "LKQ", Name: "LKQ Corporation"},
	//{FIGI: "BBG00JHNJW99", Ticker: "FOXA", Name: "Twenty-First Century Fox (A)"},
	//{FIGI: "BBG000CX0P89", Ticker: "GPN", Name: "Global Payments"},
	//{FIGI: "BBG000BSC0K9", Ticker: "DLTR", Name: "Dollar Tree"},
	//{FIGI: "BBG000BHLYP4", Ticker: "CME", Name: "CME GROUP"},
	//{FIGI: "BBG000BG3445", Ticker: "COO", Name: "The Cooper Companies"},
	//{FIGI: "BBG000F86GP6", Ticker: "MCO", Name: "Moody's"},
	//{FIGI: "BBG000BCQ4P6", Ticker: "AVY", Name: "Avery Dennison"},
	//{FIGI: "BBG000F5VVB6", Ticker: "NDAQ", Name: "Nasdaq"},
	//{FIGI: "BBG000BL9C59", Ticker: "MAC", Name: "Macerich"},
	//{FIGI: "BBG000BFXHL6", Ticker: "MOS", Name: "Mosaic"},
	//{FIGI: "BBG000G6Y5W4", Ticker: "MKC", Name: "McCormick & Co"},
	//{FIGI: "BBG000BCZL41", Ticker: "VRSK", Name: "Verisk Analytics"},
	//{FIGI: "BBG000BVV7G1", Ticker: "TXN", Name: "Texas Instruments"},
}

var QuotesMapRUS = []map[string]string{
	{FIGI: "BBG004730N88", Ticker: "SBER", Name: "Сбербанк России"},
	{FIGI: "BBG004730RP0", Ticker: "GAZP", Name: "Газпром"},
	{FIGI: "BBG004731489", Ticker: "GMKN", Name: "Норильский никель"},
	{FIGI: "BBG004731354", Ticker: "ROSN", Name: "Роснефть"},
	{FIGI: "BBG004731032", Ticker: "LKOH", Name: "ЛУКОЙЛ"},
	{FIGI: "BBG00475KKY8", Ticker: "NVTK", Name: "НОВАТЭК"},
	{FIGI: "BBG004730ZJ9", Ticker: "VTBR", Name: "Банк ВТБ"},
	{FIGI: "BBG004S684M6", Ticker: "SIBN", Name: "Газпром нефть"},
	{FIGI: "BBG004S681M2", Ticker: "SNGSP", Name: "Сургутнефтегаз - привилегированные акции"},
	{FIGI: "BBG000R607Y3", Ticker: "PLZL", Name: "Полюс Золото"},
	{FIGI: "BBG004RVFFC0", Ticker: "TATN", Name: "Татнефть"},
	{FIGI: "BBG00475K6C3", Ticker: "CHMF", Name: "Северсталь"},
	{FIGI: "BBG004S681B4", Ticker: "NLMK", Name: "НЛМК"},
	{FIGI: "BBG004S681W1", Ticker: "MTSS", Name: "МТС"},
	{FIGI: "BBG004PYF2N3", Ticker: "POLY", Name: "Polymetal"},
	{FIGI: "BBG00JXPFBN0", Ticker: "FIVE", Name: "ГДР X5 RetailGroup"},
	{FIGI: "BBG004S68B31", Ticker: "ALRS", Name: "АЛРОСА"},
	{FIGI: "BBG008F2T3T2", Ticker: "RUAL", Name: "РУСАЛ"},
	{FIGI: "BBG004S68507", Ticker: "MAGN", Name: "Магнитогорский металлургический комбинат"},
}

func GetQuotesDJ() []string {
	var quotes = make([]string, 0, len(QuotesMapDJ))
	for _, k := range QuotesMapDJ {
		quotes = append(quotes, k[FIGI])
	}
	return quotes
}

func GetQuotesRus() []string {
	var quotes = make([]string, 0, len(QuotesMapRUS))
	for _, k := range QuotesMapRUS {
		quotes = append(quotes, k[FIGI])
	}
	return quotes
}

func GetInstrNamesDJ() []string {
	names := make([]string, 0)
	for index, _ := range GetQuotesDJ() {
		names = append(names, QuotesMapDJ[index]["Name"])
	}
	return names
}

func GetInstrNamesRUS() []string {
	names := make([]string, 0)
	for index, _ := range GetQuotesRus() {
		names = append(names, QuotesMapRUS[index]["Name"])
	}
	return names
}

func GetQuoteNameByFigi(figi string) string {

	for _, element := range QuotesMapDJ {
		if element[FIGI] == figi {
			return element["Name"]
		}
	}

	for _, element := range QuotesMapRUS {
		if element[FIGI] == figi {
			return element["Name"]
		}
	}

	log.Printf("Unknown element figi: %s", figi)
	return ""
}

func GetFigiByName(name string) string {
	for _, element := range QuotesMapDJ {
		if element[Name] == name {
			return element[Ticker]
		}
	}

	for _, element := range QuotesMapRUS {
		if element[Name] == name {
			return element[Ticker]
		}
	}

	log.Printf("Unknown element name: %s", name)
	return ""
}
