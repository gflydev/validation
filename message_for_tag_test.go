package validation

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

// TestStruct is a test struct with validation tags
type TestStructForMsg struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Age    int    `json:"age" validate:"required,gt=0"`
	URL    string `json:"url" validate:"url"`
	IP     string `json:"ip" validate:"ip"`
	Length string `json:"length" validate:"len=10"`
	Min    string `json:"min" validate:"min=5"`
	Max    string `json:"max" validate:"max=10"`
	Alpha  string `json:"alpha" validate:"alpha"`
	Equal  string `json:"equal" validate:"eq=test"`
	// Additional fields for more validation tags
	AlphaNum    string   `json:"alphanum" validate:"alphanum"`
	ASCII       string   `json:"ascii" validate:"ascii"`
	Boolean     string   `json:"boolean" validate:"boolean"`
	Contains    string   `json:"contains" validate:"contains=test"`
	ContainsAny string   `json:"containsany" validate:"containsany=!@#"`
	EndsWith    string   `json:"endswith" validate:"endswith=test"`
	ExcludesAll string   `json:"excludesall" validate:"excludesall=!@#"`
	Lowercase   string   `json:"lowercase" validate:"lowercase"`
	Uppercase   string   `json:"uppercase" validate:"uppercase"`
	Base64      string   `json:"base64" validate:"base64"`
	CreditCard  string   `json:"credit_card" validate:"credit_card"`
	JSON        string   `json:"json" validate:"json"`
	Latitude    string   `json:"latitude" validate:"latitude"`
	Longitude   string   `json:"longitude" validate:"longitude"`
	UUID        string   `json:"uuid" validate:"uuid"`
	CIDR        string   `json:"cidr" validate:"cidr"`
	MAC         string   `json:"mac" validate:"mac"`
	IPV4        string   `json:"ipv4" validate:"ipv4"`
	IPV6        string   `json:"ipv6" validate:"ipv6"`
	FQDN        string   `json:"fqdn" validate:"fqdn"`
	Unique      []string `json:"unique" validate:"unique"`
	OnOf        string   `json:"oneof" validate:"oneof=red green blue"`
	// Removed problematic validation tags
	GreaterThanField int    `json:"greater_than_field" validate:"gtfield=Age"`
	LessThanField    int    `json:"less_than_field" validate:"ltfield=Age"`
	EqualField       string `json:"equal_field" validate:"eqfield=Name"`
	NotEqualField    string `json:"not_equal_field" validate:"nefield=Name"`
	// Additional fields for more validation tags
	//EqualIgnoreCase    string `json:"equal_ignore_case" validate:"eq_ignore_case=TEST"`
	GreaterThan      string `json:"greater_than" validate:"gt=5"`
	GreaterThanEqual string `json:"greater_than_equal" validate:"gte=5"`
	//LessThan           string `json:"less_than" validate:"lt=5"`
	//LessThanEqual      string `json:"less_than_equal" validate:"lte=5"`
	NotEqual           string `json:"not_equal" validate:"ne=test"`
	NotEqualIgnoreCase string `json:"not_equal_ignore_case" validate:"ne_ignore_case=TEST"`
	AlphaUnicode       string `json:"alpha_unicode" validate:"alphaunicode"`
	AlphaNumUnicode    string `json:"alpha_num_unicode" validate:"alphanumunicode"`
	ContainsRune       string `json:"contains_rune" validate:"containsrune=☺"`
	EndsNotWith        string `json:"ends_not_with" validate:"endsnotwith=test"`
	Excludes           string `json:"excludes" validate:"excludes=test"`
	ExcludesRune       string `json:"excludes_rune" validate:"excludesrune=☺"`
	MultiBytes         string `json:"multi_bytes" validate:"multibyte"`
	Number             string `json:"number" validate:"number"`
	Numeric            string `json:"numeric" validate:"numeric"`
	PrintASCII         string `json:"print_ascii" validate:"printascii"`
	StartsNotWith      string `json:"starts_not_with" validate:"startsnotwith=test"`
	StartsWith         string `json:"starts_with" validate:"startswith=test"`
	Base64URL          string `json:"base64_url" validate:"base64url"`
	Base64RawURL       string `json:"base64_raw_url" validate:"base64rawurl"`
	BIC                string `json:"bic" validate:"bic"`
	BCP47LanguageTag   string `json:"bcp47_language_tag" validate:"bcp47_language_tag"`
	BTCAddr            string `json:"btc_addr" validate:"btc_addr"`
	BTCAddrBech32      string `json:"btc_addr_bech32" validate:"btc_addr_bech32"`
	MongoDB            string `json:"mongodb" validate:"mongodb"`
	Cron               string `json:"cron" validate:"cron"`
	//SpiceDB                    string `json:"spicedb" validate:"spicedb"`
	DateTime                   string `json:"datetime" validate:"datetime"`
	E164                       string `json:"e164" validate:"e164"`
	EthAddr                    string `json:"eth_addr" validate:"eth_addr"`
	Hexadecimal                string `json:"hexadecimal" validate:"hexadecimal"`
	HexColor                   string `json:"hex_color" validate:"hexcolor"`
	HSL                        string `json:"hsl" validate:"hsl"`
	HSLA                       string `json:"hsla" validate:"hsla"`
	HTML                       string `json:"html" validate:"html"`
	HTMLEncoded                string `json:"html_encoded" validate:"html_encoded"`
	ISBN                       string `json:"isbn" validate:"isbn"`
	ISBN10                     string `json:"isbn10" validate:"isbn10"`
	ISBN13                     string `json:"isbn13" validate:"isbn13"`
	ISSN                       string `json:"issn" validate:"issn"`
	ISO3166_1_Alpha2           string `json:"iso3166_1_alpha2" validate:"iso3166_1_alpha2"`
	ISO3166_1_Alpha3           string `json:"iso3166_1_alpha3" validate:"iso3166_1_alpha3"`
	ISO3166_1_AlphaNumeric     string `json:"iso3166_1_alpha_numeric" validate:"iso3166_1_alpha_numeric"`
	ISO3166_2                  string `json:"iso3166_2" validate:"iso3166_2"`
	ISO4217                    string `json:"iso4217" validate:"iso4217"`
	JWT                        string `json:"jwt" validate:"jwt"`
	LuhnChecksum               string `json:"luhn_checksum" validate:"luhn_checksum"`
	PostcodeISO3166Alpha2      string `json:"postcode_iso3166_alpha2" validate:"postcode_iso3166_alpha2"`
	PostcodeISO3166Alpha2Field string `json:"postcode_iso3166_alpha2_field" validate:"postcode_iso3166_alpha2_field"`
	RGB                        string `json:"rgb" validate:"rgb"`
	RGBA                       string `json:"rgba" validate:"rgba"`
	SSN                        string `json:"ssn" validate:"ssn"`
	Timezone                   string `json:"timezone" validate:"timezone"`
	UUID3                      string `json:"uuid3" validate:"uuid3"`
	UUID3RFC4122               string `json:"uuid3_rfc4122" validate:"uuid3_rfc4122"`
	UUID4                      string `json:"uuid4" validate:"uuid4"`
	UUID4RFC4122               string `json:"uuid4_rfc4122" validate:"uuid4_rfc4122"`
	UUID5                      string `json:"uuid5" validate:"uuid5"`
	UUID5RFC4122               string `json:"uuid5_rfc4122" validate:"uuid5_rfc4122"`
	UUIDRFC4122                string `json:"uuid_rfc4122" validate:"uuid_rfc4122"`
	MD4                        string `json:"md4" validate:"md4"`
	MD5                        string `json:"md5" validate:"md5"`
	SHA256                     string `json:"sha256" validate:"sha256"`
	SHA384                     string `json:"sha384" validate:"sha384"`
	SHA512                     string `json:"sha512" validate:"sha512"`
	RIPEMD128                  string `json:"ripemd128" validate:"ripemd128"`
	RIPEMD160                  string `json:"ripemd160" validate:"ripemd160"`
	TIGER128                   string `json:"tiger128" validate:"tiger128"`
	TIGER160                   string `json:"tiger160" validate:"tiger160"`
	TIGER192                   string `json:"tiger192" validate:"tiger192"`
	Semver                     string `json:"semver" validate:"semver"`
	ULID                       string `json:"ulid" validate:"ulid"`
	CVE                        string `json:"cve" validate:"cve"`
	EqCSField                  string `json:"eq_cs_field" validate:"eqcsfield=Name"`
	FieldContains              string `json:"field_contains" validate:"fieldcontains=test"`
	FieldExcludes              string `json:"field_excludes" validate:"fieldexcludes=test"`
	GtCSField                  string `json:"gt_cs_field" validate:"gtcsfield=Age"`
	GteCSField                 string `json:"gte_cs_field" validate:"gtecsfield=Age"`
	GteField                   string `json:"gte_field" validate:"gtefield=Age"`
	LtCSField                  string `json:"lt_cs_field" validate:"ltcsfield=Age"`
	LteCSField                 string `json:"lte_cs_field" validate:"ltecsfield=Age"`
	LteField                   string `json:"lte_field" validate:"ltefield=Age"`
	NeCSField                  string `json:"ne_cs_field" validate:"necsfield=Name"`
	CIDRv4                     string `json:"cidrv4" validate:"cidrv4"`
	CIDRv6                     string `json:"cidrv6" validate:"cidrv6"`
	DataURI                    string `json:"data_uri" validate:"datauri"`
	HostnamePort               string `json:"hostname_port" validate:"hostname_port"`
	//HostnameRFC1123            string `json:"hostname_rfc1123" validate:"hostname_rfc1123"`
	IP4Addr  string `json:"ip4_addr" validate:"ip4_addr"`
	IP6Addr  string `json:"ip6_addr" validate:"ip6_addr"`
	IPAddr   string `json:"ip_addr" validate:"ip_addr"`
	TCP4Addr string `json:"tcp4_addr" validate:"tcp4_addr"`
	TCP6Addr string `json:"tcp6_addr" validate:"tcp6_addr"`
	TCPAddr  string `json:"tcp_addr" validate:"tcp_addr"`
	UDP4Addr string `json:"udp4_addr" validate:"udp4_addr"`
	UDP6Addr string `json:"udp6_addr" validate:"udp6_addr"`
	UDPAddr  string `json:"udp_addr" validate:"udp_addr"`
	//UnixAddr           string `json:"unix_addr" validate:"unix_addr"`
	URI     string `json:"uri" validate:"uri"`
	HTTPURL string `json:"http_url" validate:"http_url"`
	//URLEncoded         string `json:"url_encoded" validate:"url_encoded"`
	URNRFC2141 string `json:"urn_rfc2141" validate:"urn_rfc2141"`
	Dir        string `json:"dir" validate:"dir"`
	DirPath    string `json:"dir_path" validate:"dirpath"`
	File       string `json:"file" validate:"file"`
	//FilePath           string `json:"file_path" validate:"filepath"`
	Image          string `json:"image" validate:"image"`
	IsDefault      string `json:"is_default" validate:"isdefault"`
	RequiredIf     string `json:"required_if" validate:"required_if"`
	RequiredUnless string `json:"required_unless" validate:"required_unless"`
	//RequiredWith       string `json:"required_with" validate:"required_with"`
	RequiredWithAll    string `json:"required_with_all" validate:"required_with_all"`
	RequiredWithout    string `json:"required_without" validate:"required_without"`
	RequiredWithoutAll string `json:"required_without_all" validate:"required_without_all"`
	ExcludedIf         string `json:"excluded_if" validate:"excluded_if"`
	//ExcludedUnless     string `json:"excluded_unless" validate:"excluded_unless"`
	//ExcludedWith       string `json:"excluded_with" validate:"excluded_with"`
	ExcludedWithAll string `json:"excluded_with_all" validate:"excluded_with_all"`
	//ExcludedWithout    string `json:"excluded_without" validate:"excluded_without"`
	ExcludedWithoutAll string `json:"excluded_without_all" validate:"excluded_without_all"`
	IsColor            string `json:"is_color" validate:"iscolor"`
	CountryCode        string `json:"country_code" validate:"country_code"`
}

func TestMsgForTag(t *testing.T) {
	// Create a validator instance
	v := validator.New()

	// Test cases
	tests := []struct {
		name          string
		data          interface{}
		fieldName     string
		expectedError string
	}{
		// Original test cases
		{
			name: "Required field",
			data: TestStructForMsg{
				// Name is missing (required)
				Email: "test@example.com",
				Age:   30,
			},
			fieldName:     "Name",
			expectedError: "is required",
		},
		{
			name: "Email validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "invalid-email", // Invalid email
				Age:   30,
			},
			fieldName:     "Email",
			expectedError: "invalid email",
		},
		{
			name: "Greater than validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   0, // Should be greater than 0
			},
			fieldName:     "Age",
			expectedError: "is required", // The validator reports this as "required" first
		},
		{
			name: "URL validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				URL:   "not-a-url", // Invalid URL
			},
			fieldName:     "URL",
			expectedError: "invalid URL string",
		},
		{
			name: "IP validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				IP:    "not-an-ip", // Invalid IP
			},
			fieldName:     "IP",
			expectedError: "invalid Internet Protocol Address IP",
		},
		{
			name: "Length validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				Length: "12345", // Not length 10
			},
			fieldName:     "Length",
			expectedError: "length 10",
		},
		{
			name: "Min validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Min:   "123", // Less than min 5
			},
			fieldName:     "Min",
			expectedError: "minimum 5",
		},
		{
			name: "Max validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Max:   "12345678901", // More than max 10
			},
			fieldName:     "Max",
			expectedError: "maximum 10",
		},
		{
			name: "Alpha validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Alpha: "123abc", // Not alpha
			},
			fieldName:     "Alpha",
			expectedError: "invalid alpha",
		},
		{
			name: "Equal validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Equal: "not-test", // Not equal to "test"
			},
			fieldName:     "Equal",
			expectedError: "equals test",
		},

		// Additional test cases for new validation tags
		{
			name: "AlphaNum validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				AlphaNum: "abc123!@#", // Not alphanumeric
			},
			fieldName:     "AlphaNum",
			expectedError: "invalid alphanumeric",
		},
		{
			name: "ASCII validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				ASCII: "ñáéíóú", // Non-ASCII characters
			},
			fieldName:     "ASCII",
			expectedError: "invalid ASCII",
		},
		{
			name: "Boolean validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				Boolean: "notbool", // Not a boolean
			},
			fieldName:     "Boolean",
			expectedError: "invalid boolean type",
		},
		{
			name: "Contains validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				Contains: "nothing", // Doesn't contain "test"
			},
			fieldName:     "Contains",
			expectedError: "contains test",
		},
		{
			name: "ContainsAny validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				ContainsAny: "abcdef", // Doesn't contain any of "!@#"
			},
			fieldName:     "ContainsAny",
			expectedError: "contains any",
		},
		{
			name: "EndsWith validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				EndsWith: "nothing", // Doesn't end with "test"
			},
			fieldName:     "EndsWith",
			expectedError: "ends with test",
		},
		{
			name: "ExcludesAll validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				ExcludesAll: "abc!def", // Contains "!"
			},
			fieldName:     "ExcludesAll",
			expectedError: "excludes all !@#",
		},
		{
			name: "Lowercase validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				Lowercase: "abcDEF", // Not all lowercase
			},
			fieldName:     "Lowercase",
			expectedError: "lowercase only",
		},
		{
			name: "Uppercase validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				Uppercase: "ABCdef", // Not all uppercase
			},
			fieldName:     "Uppercase",
			expectedError: "uppercase only",
		},
		{
			name: "Base64 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				Base64: "not-base64!", // Not valid base64
			},
			fieldName:     "Base64",
			expectedError: "invalid Base64",
		},
		{
			name: "CreditCard validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				CreditCard: "1234-5678-9012-3456", // Not a valid credit card
			},
			fieldName:     "CreditCard",
			expectedError: "invalid Credit Card Number",
		},
		{
			name: "JSON validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				JSON:  "not-json", // Not valid JSON
			},
			fieldName:     "JSON",
			expectedError: "invalid JSON",
		},
		{
			name: "Latitude validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				Latitude: "200", // Not a valid latitude
			},
			fieldName:     "Latitude",
			expectedError: "invalid Latitude",
		},
		{
			name: "Longitude validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				Longitude: "200", // Not a valid longitude
			},
			fieldName:     "Longitude",
			expectedError: "invalid Longitude",
		},
		{
			name: "UUID validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				UUID:  "not-uuid", // Not a valid UUID
			},
			fieldName:     "UUID",
			expectedError: "invalid Universally Unique Identifier UUID",
		},
		{
			name: "CIDR validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				CIDR:  "not-cidr", // Not a valid CIDR
			},
			fieldName:     "CIDR",
			expectedError: "invalid Classless Inter-Domain Routing CIDR",
		},
		{
			name: "MAC validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				MAC:   "not-mac", // Not a valid MAC address
			},
			fieldName:     "MAC",
			expectedError: "invalid Media Access Control Address MAC",
		},
		{
			name: "IPV4 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				IPV4:  "not-ipv4", // Not a valid IPv4
			},
			fieldName:     "IPV4",
			expectedError: "invalid Internet Protocol Address IPv4",
		},
		{
			name: "IPV6 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				IPV6:  "not-ipv6", // Not a valid IPv6
			},
			fieldName:     "IPV6",
			expectedError: "invalid Internet Protocol Address IPv6",
		},
		{
			name: "FQDN validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				FQDN:  "not-fqdn", // Not a valid FQDN
			},
			fieldName:     "FQDN",
			expectedError: "invalid Full Qualified Domain Name (FQDN)",
		},
		{
			name: "Unique validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				Unique: []string{"a", "b", "a"}, // Not unique
			},
			fieldName:     "Unique",
			expectedError: "not unique",
		},
		{
			name: "OneOf validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				OnOf:  "yellow", // Not one of red, green, blue
			},
			fieldName:     "OnOf",
			expectedError: "one of red green blue",
		},
		// Removed problematic validation test cases
		{
			name: "GreaterThanField validation",
			data: TestStructForMsg{
				Name:             "John Doe",
				Email:            "test@example.com",
				Age:              30,
				GreaterThanField: 20, // Not greater than Age (30)
			},
			fieldName:     "GreaterThanField",
			expectedError: "field greater than another field Age",
		},
		{
			name: "LessThanField validation",
			data: TestStructForMsg{
				Name:          "John Doe",
				Email:         "test@example.com",
				Age:           30,
				LessThanField: 40, // Not less than Age (30)
			},
			fieldName:     "LessThanField",
			expectedError: "field less than another field Age",
		},
		{
			name: "EqualField validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				EqualField: "Different", // Not equal to Name
			},
			fieldName:     "EqualField",
			expectedError: "field equals another field Name",
		},
		{
			name: "NotEqualField validation",
			data: TestStructForMsg{
				Name:          "John Doe",
				Email:         "test@example.com",
				Age:           30,
				NotEqualField: "John Doe", // Equal to Name but should not be
			},
			fieldName:     "NotEqualField",
			expectedError: "field does not equal another field Name",
		},
		// Additional test cases for more validation tags
		//{
		//	name: "EqualIgnoreCase validation",
		//	data: TestStructForMsg{
		//		Name:            "John Doe",
		//		Email:           "test@example.com",
		//		Age:             30,
		//		EqualIgnoreCase: "test", // Not equal to "TEST" ignoring case
		//	},
		//	fieldName:     "EqualIgnoreCase",
		//	expectedError: "equal ignoring case TEST",
		//},
		{
			name: "GreaterThan validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				GreaterThan: "3", // Not greater than 5
			},
			fieldName:     "GreaterThan",
			expectedError: "greater than 5",
		},
		{
			name: "GreaterThanEqual validation",
			data: TestStructForMsg{
				Name:             "John Doe",
				Email:            "test@example.com",
				Age:              30,
				GreaterThanEqual: "3", // Not greater than or equal to 5
			},
			fieldName:     "GreaterThanEqual",
			expectedError: "greater than or equal 5",
		},
		//{
		//	name: "LessThan validation",
		//	data: TestStructForMsg{
		//		Name:     "John Doe",
		//		Email:    "test@example.com",
		//		Age:      30,
		//		LessThan: "7", // Not less than 5
		//	},
		//	fieldName:     "LessThan",
		//	expectedError: "less than 5",
		//},
		//{
		//	name: "LessThanEqual validation",
		//	data: TestStructForMsg{
		//		Name:          "John Doe",
		//		Email:         "test@example.com",
		//		Age:           30,
		//		LessThanEqual: "7", // Not less than or equal to 5
		//	},
		//	fieldName:     "LessThanEqual",
		//	expectedError: "less than or equal 5",
		//},
		{
			name: "NotEqual validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				NotEqual: "test", // Equal to "test" but should not be
			},
			fieldName:     "NotEqual",
			expectedError: "not equal test",
		},
		{
			name: "NotEqualIgnoreCase validation",
			data: TestStructForMsg{
				Name:               "John Doe",
				Email:              "test@example.com",
				Age:                30,
				NotEqualIgnoreCase: "test", // Equal to "TEST" ignoring case but should not be
			},
			fieldName:     "NotEqualIgnoreCase",
			expectedError: "not equal ignoring case TEST",
		},
		{
			name: "AlphaUnicode validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				AlphaUnicode: "123abc", // Not alpha unicode
			},
			fieldName:     "AlphaUnicode",
			expectedError: "invalid alpha unicode",
		},
		{
			name: "AlphaNumUnicode validation",
			data: TestStructForMsg{
				Name:            "John Doe",
				Email:           "test@example.com",
				Age:             30,
				AlphaNumUnicode: "abc123!@#", // Not alphanumeric unicode
			},
			fieldName:     "AlphaNumUnicode",
			expectedError: "invalid alphanumeric unicode",
		},
		{
			name: "ContainsRune validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				ContainsRune: "abc", // Doesn't contain the rune ☺
			},
			fieldName:     "ContainsRune",
			expectedError: "contains rune",
		},
		{
			name: "EndsNotWith validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				EndsNotWith: "nothing_test", // Ends with "test" but should not
			},
			fieldName:     "EndsNotWith",
			expectedError: "ends not with test",
		},
		{
			name: "Excludes validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				Excludes: "contains_test_string", // Contains "test" but should not
			},
			fieldName:     "Excludes",
			expectedError: "excludes test",
		},
		{
			name: "ExcludesRune validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				ExcludesRune: "abc☺def", // Contains the rune ☺ but should not
			},
			fieldName:     "ExcludesRune",
			expectedError: "excludes rune ☺",
		},
		{
			name: "MultiBytes validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				MultiBytes: "abc", // Not multi-byte
			},
			fieldName:     "MultiBytes",
			expectedError: "invalid multi-byte",
		},
		{
			name: "Number validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				Number: "abc", // Not a number
			},
			fieldName:     "Number",
			expectedError: "invalid number",
		},
		{
			name: "Numeric validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				Numeric: "abc", // Not numeric
			},
			fieldName:     "Numeric",
			expectedError: "invalid numeric",
		},
		{
			name: "PrintASCII validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				PrintASCII: "ñáéíóú", // Not printable ASCII
			},
			fieldName:     "PrintASCII",
			expectedError: "invalid a printable ASCII",
		},
		{
			name: "StartsNotWith validation",
			data: TestStructForMsg{
				Name:          "John Doe",
				Email:         "test@example.com",
				Age:           30,
				StartsNotWith: "test_string", // Starts with "test" but should not
			},
			fieldName:     "StartsNotWith",
			expectedError: "starts not with test",
		},
		{
			name: "StartsWith validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				StartsWith: "nothing", // Doesn't start with "test"
			},
			fieldName:     "StartsWith",
			expectedError: "starts with test",
		},
		{
			name: "Base64URL validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				Base64URL: "not-base64-url!", // Not valid base64 URL
			},
			fieldName:     "Base64URL",
			expectedError: "invalid Base64 URL",
		},
		{
			name: "Base64RawURL validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				Base64RawURL: "not-base64-raw-url!", // Not valid base64 raw URL
			},
			fieldName:     "Base64RawURL",
			expectedError: "invalid Base64 raw URL",
		},
		{
			name: "BIC validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				BIC:   "not-bic", // Not a valid BIC
			},
			fieldName:     "BIC",
			expectedError: "invalid Business Identifier Code (ISO 9362)",
		},
		{
			name: "BCP47LanguageTag validation",
			data: TestStructForMsg{
				Name:             "John Doe",
				Email:            "test@example.com",
				Age:              30,
				BCP47LanguageTag: "not-bcp47", // Not a valid BCP47 language tag
			},
			fieldName:     "BCP47LanguageTag",
			expectedError: "invalid Language tag (BCP 47)",
		},
		{
			name: "BTCAddr validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				BTCAddr: "not-btc-addr", // Not a valid Bitcoin address
			},
			fieldName:     "BTCAddr",
			expectedError: "invalid Bitcoin address",
		},
		{
			name: "BTCAddrBech32 validation",
			data: TestStructForMsg{
				Name:          "John Doe",
				Email:         "test@example.com",
				Age:           30,
				BTCAddrBech32: "not-btc-addr-bech32", // Not a valid Bitcoin Bech32 address
			},
			fieldName:     "BTCAddrBech32",
			expectedError: "invalid Bitcoin Bech32 Address (segwit)",
		},
		{
			name: "MongoDB validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				MongoDB: "not-mongodb", // Not a valid MongoDB ObjectID
			},
			fieldName:     "MongoDB",
			expectedError: "invalid MongoDB ObjectID",
		},
		{
			name: "Cron validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Cron:  "not-cron", // Not a valid cron expression
			},
			fieldName:     "Cron",
			expectedError: "invalid cron",
		},
		//{
		//	name: "SpiceDB validation",
		//	data: TestStructForMsg{
		//		Name:    "John Doe",
		//		Email:   "test@example.com",
		//		Age:     30,
		//		SpiceDB: "not-spicedb", // Not a valid SpiceDb ObjectID/Permission/Type
		//	},
		//	fieldName:     "SpiceDB",
		//	expectedError: "invalid SpiceDb ObjectID/Permission/Type",
		//},
		{
			name: "DateTime validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				DateTime: "not-datetime", // Not a valid datetime
			},
			fieldName:     "DateTime",
			expectedError: "invalid Datetime",
		},
		{
			name: "E164 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				E164:  "not-e164", // Not a valid E.164 phone number
			},
			fieldName:     "E164",
			expectedError: "invalid e164 formatted phone number",
		},
		{
			name: "EthAddr validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				EthAddr: "not-eth-addr", // Not a valid Ethereum address
			},
			fieldName:     "EthAddr",
			expectedError: "invalid Ethereum address",
		},
		{
			name: "Hexadecimal validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				Hexadecimal: "not-hex", // Not a valid hexadecimal
			},
			fieldName:     "Hexadecimal",
			expectedError: "invalid Hexadecimal string",
		},
		{
			name: "HexColor validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				HexColor: "not-hex-color", // Not a valid hex color
			},
			fieldName:     "HexColor",
			expectedError: "invalid Hexcolor string",
		},
		{
			name: "HSL validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				HSL:   "not-hsl", // Not a valid HSL color
			},
			fieldName:     "HSL",
			expectedError: "invalid HSL string",
		},
		{
			name: "HSLA validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				HSLA:  "not-hsla", // Not a valid HSLA color
			},
			fieldName:     "HSLA",
			expectedError: "invalid HSLA string",
		},
		{
			name: "HTML validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				HTML:  "not-html", // Not valid HTML
			},
			fieldName:     "HTML",
			expectedError: "invalid HTTP tags",
		},
		{
			name: "HTMLEncoded validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				HTMLEncoded: "not-html-encoded", // Not valid HTML encoded
			},
			fieldName:     "HTMLEncoded",
			expectedError: "invalid HTML Encoded",
		},
		{
			name: "ISBN validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				ISBN:  "not-isbn", // Not a valid ISBN
			},
			fieldName:     "ISBN",
			expectedError: "invalid International Standard Book Number",
		},
		{
			name: "ISBN10 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				ISBN10: "not-isbn10", // Not a valid ISBN-10
			},
			fieldName:     "ISBN10",
			expectedError: "invalid International Standard Book Number 10",
		},
		{
			name: "ISBN13 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				ISBN13: "not-isbn13", // Not a valid ISBN-13
			},
			fieldName:     "ISBN13",
			expectedError: "invalid International Standard Book Number 13",
		},
		{
			name: "ISSN validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				ISSN:  "not-issn", // Not a valid ISSN
			},
			fieldName:     "ISSN",
			expectedError: "invalid International Standard Serial Number",
		},
		{
			name: "ISO3166_1_Alpha2 validation",
			data: TestStructForMsg{
				Name:             "John Doe",
				Email:            "test@example.com",
				Age:              30,
				ISO3166_1_Alpha2: "not-iso3166-1-alpha2", // Not a valid ISO 3166-1 alpha-2 code
			},
			fieldName:     "ISO3166_1_Alpha2",
			expectedError: "invalid Two-letter country code (ISO 3166-1 alpha-2)",
		},
		{
			name: "ISO3166_1_Alpha3 validation",
			data: TestStructForMsg{
				Name:             "John Doe",
				Email:            "test@example.com",
				Age:              30,
				ISO3166_1_Alpha3: "not-iso3166-1-alpha3", // Not a valid ISO 3166-1 alpha-3 code
			},
			fieldName:     "ISO3166_1_Alpha3",
			expectedError: "invalid Three-letter country code (ISO 3166-1 alpha-3)",
		},
		{
			name: "ISO3166_1_AlphaNumeric validation",
			data: TestStructForMsg{
				Name:                   "John Doe",
				Email:                  "test@example.com",
				Age:                    30,
				ISO3166_1_AlphaNumeric: "not-iso3166-1-alpha-numeric", // Not a valid ISO 3166-1 numeric code
			},
			fieldName:     "ISO3166_1_AlphaNumeric",
			expectedError: "invalid Numeric country code (ISO 3166-1 numeric)",
		},
		{
			name: "ISO3166_2 validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				ISO3166_2: "not-iso3166-2", // Not a valid ISO 3166-2 code
			},
			fieldName:     "ISO3166_2",
			expectedError: "invalid Country subdivision code (ISO 3166-2)",
		},
		{
			name: "ISO4217 validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				ISO4217: "not-iso4217", // Not a valid ISO 4217 code
			},
			fieldName:     "ISO4217",
			expectedError: "invalid Currency code (ISO 4217)",
		},
		{
			name: "JWT validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				JWT:   "not-jwt", // Not a valid JWT
			},
			fieldName:     "JWT",
			expectedError: "invalid JSON Web Token (JWT)",
		},
		{
			name: "LuhnChecksum validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				LuhnChecksum: "not-luhn", // Not a valid Luhn checksum
			},
			fieldName:     "LuhnChecksum",
			expectedError: "invalid Luhn Algorithm Checksum (for strings and (u)int)",
		},
		{
			name: "PostcodeISO3166Alpha2 validation",
			data: TestStructForMsg{
				Name:                  "John Doe",
				Email:                 "test@example.com",
				Age:                   30,
				PostcodeISO3166Alpha2: "not-postcode", // Not a valid postcode
			},
			fieldName:     "PostcodeISO3166Alpha2",
			expectedError: "invalid Postcode",
		},
		{
			name: "PostcodeISO3166Alpha2Field validation",
			data: TestStructForMsg{
				Name:                       "John Doe",
				Email:                      "test@example.com",
				Age:                        30,
				PostcodeISO3166Alpha2Field: "not-postcode", // Not a valid postcode
			},
			fieldName:     "PostcodeISO3166Alpha2Field",
			expectedError: "invalid Postcode",
		},
		{
			name: "RGB validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				RGB:   "not-rgb", // Not a valid RGB color
			},
			fieldName:     "RGB",
			expectedError: "invalid RGB string",
		},
		{
			name: "RGBA validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				RGBA:  "not-rgba", // Not a valid RGBA color
			},
			fieldName:     "RGBA",
			expectedError: "invalid RGBA string",
		},
		{
			name: "SSN validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				SSN:   "not-ssn", // Not a valid SSN
			},
			fieldName:     "SSN",
			expectedError: "invalid Social Security Number SSN",
		},
		{
			name: "Timezone validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				Timezone: "not-timezone", // Not a valid timezone
			},
			fieldName:     "Timezone",
			expectedError: "invalid Timezone",
		},
		{
			name: "UUID3 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				UUID3: "not-uuid3", // Not a valid UUID v3
			},
			fieldName:     "UUID3",
			expectedError: "invalid Universally Unique Identifier UUID v3",
		},
		{
			name: "UUID3RFC4122 validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				UUID3RFC4122: "not-uuid3-rfc4122", // Not a valid UUID v3 RFC4122
			},
			fieldName:     "UUID3RFC4122",
			expectedError: "invalid Universally Unique Identifier UUID v3 RFC4122",
		},
		{
			name: "UUID4 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				UUID4: "not-uuid4", // Not a valid UUID v4
			},
			fieldName:     "UUID4",
			expectedError: "invalid Universally Unique Identifier UUID v4",
		},
		{
			name: "UUID4RFC4122 validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				UUID4RFC4122: "not-uuid4-rfc4122", // Not a valid UUID v4 RFC4122
			},
			fieldName:     "UUID4RFC4122",
			expectedError: "invalid Universally Unique Identifier UUID v4 RFC4122",
		},
		{
			name: "UUID5 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				UUID5: "not-uuid5", // Not a valid UUID v5
			},
			fieldName:     "UUID5",
			expectedError: "invalid Universally Unique Identifier UUID v5",
		},
		{
			name: "UUID5RFC4122 validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				UUID5RFC4122: "not-uuid5-rfc4122", // Not a valid UUID v5 RFC4122
			},
			fieldName:     "UUID5RFC4122",
			expectedError: "invalid Universally Unique Identifier UUID v5 RFC4122",
		},
		{
			name: "UUIDRFC4122 validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				UUIDRFC4122: "not-uuid-rfc4122", // Not a valid UUID RFC4122
			},
			fieldName:     "UUIDRFC4122",
			expectedError: "invalid Universally Unique Identifier UUID RFC4122",
		},
		{
			name: "MD4 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				MD4:   "not-md4", // Not a valid MD4 hash
			},
			fieldName:     "MD4",
			expectedError: "invalid MD4 hash",
		},
		{
			name: "MD5 validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				MD5:   "not-md5", // Not a valid MD5 hash
			},
			fieldName:     "MD5",
			expectedError: "invalid MD5 hash",
		},
		{
			name: "SHA256 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				SHA256: "not-sha256", // Not a valid SHA256 hash
			},
			fieldName:     "SHA256",
			expectedError: "invalid SHA256 hash",
		},
		{
			name: "SHA384 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				SHA384: "not-sha384", // Not a valid SHA384 hash
			},
			fieldName:     "SHA384",
			expectedError: "invalid SHA384 hash",
		},
		{
			name: "SHA512 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				SHA512: "not-sha512", // Not a valid SHA512 hash
			},
			fieldName:     "SHA512",
			expectedError: "invalid SHA512 hash",
		},
		{
			name: "RIPEMD128 validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				RIPEMD128: "not-ripemd128", // Not a valid RIPEMD-128 hash
			},
			fieldName:     "RIPEMD128",
			expectedError: "invalid RIPEMD-128 hash",
		},
		{
			name: "RIPEMD160 validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				RIPEMD160: "not-ripemd160", // Not a valid RIPEMD-160 hash
			},
			fieldName:     "RIPEMD160",
			expectedError: "invalid RIPEMD-160 hash",
		},
		{
			name: "TIGER128 validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				TIGER128: "not-tiger128", // Not a valid TIGER128 hash
			},
			fieldName:     "TIGER128",
			expectedError: "invalid TIGER128 hash",
		},
		{
			name: "TIGER160 validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				TIGER160: "not-tiger160", // Not a valid TIGER160 hash
			},
			fieldName:     "TIGER160",
			expectedError: "invalid TIGER160 hash",
		},
		{
			name: "TIGER192 validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				TIGER192: "not-tiger192", // Not a valid TIGER192 hash
			},
			fieldName:     "TIGER192",
			expectedError: "invalid TIGER192 hash",
		},
		{
			name: "Semver validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				Semver: "not-semver", // Not a valid semantic version
			},
			fieldName:     "Semver",
			expectedError: "invalid Semantic Versioning 2.0.0",
		},
		{
			name: "ULID validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				ULID:  "not-ulid", // Not a valid ULID
			},
			fieldName:     "ULID",
			expectedError: "invalid Universally Unique Lexicographically Sortable Identifier ULID",
		},
		{
			name: "CVE validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				CVE:   "not-cve", // Not a valid CVE ID
			},
			fieldName:     "CVE",
			expectedError: "Common Vulnerabilities and Exposures Identifier (CVE id)",
		},
		{
			name: "EqCSField validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				EqCSField: "Different", // Not equal to Name (case sensitive)
			},
			fieldName:     "EqCSField",
			expectedError: "field equals another field (relative) Name",
		},
		{
			name: "FieldContains validation",
			data: TestStructForMsg{
				Name:          "John Doe",
				Email:         "test@example.com",
				Age:           30,
				FieldContains: "nothing", // Doesn't contain "test"
			},
			fieldName:     "FieldContains",
			expectedError: "field contains test",
		},
		//{
		//	name: "FieldExcludes validation",
		//	data: TestStructForMsg{
		//		Name:          "John Doe",
		//		Email:         "test@example.com",
		//		Age:           30,
		//		FieldExcludes: "contains_test_string", // Contains "test" but should not
		//	},
		//	fieldName:     "FieldExcludes",
		//	expectedError: "field excludes content test",
		//},
		{
			name: "GtCSField validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				GtCSField: "20", // Not greater than Age (30)
			},
			fieldName:     "GtCSField",
			expectedError: "field greater than another relative field Age",
		},
		{
			name: "GteCSField validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				GteCSField: "20", // Not greater than or equal to Age (30)
			},
			fieldName:     "GteCSField",
			expectedError: "field greater than or equal to another relative field Age",
		},
		{
			name: "GteField validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				GteField: "20", // Not greater than or equal to Age (30)
			},
			fieldName:     "GteField",
			expectedError: "field greater than or equal to another field Age",
		},
		{
			name: "LtCSField validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				LtCSField: "40", // Not less than Age (30)
			},
			fieldName:     "LtCSField",
			expectedError: "field less than another relative field Age",
		},
		{
			name: "LteCSField validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				LteCSField: "40", // Not less than or equal to Age (30)
			},
			fieldName:     "LteCSField",
			expectedError: "field less than or equal to another relative field Age",
		},
		{
			name: "LteField validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				LteField: "40", // Not less than or equal to Age (30)
			},
			fieldName:     "LteField",
			expectedError: "field less than another relative field Age",
		},
		{
			name: "NeCSField validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				NeCSField: "John Doe", // Equal to Name but should not be (case sensitive)
			},
			fieldName:     "NeCSField",
			expectedError: "field does not equal another field (relative) Name",
		},
		{
			name: "CIDRv4 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				CIDRv4: "not-cidrv4", // Not a valid CIDRv4
			},
			fieldName:     "CIDRv4",
			expectedError: "invalid Classless Inter-Domain Routing CIDRv4",
		},
		{
			name: "CIDRv6 validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				CIDRv6: "not-cidrv6", // Not a valid CIDRv6
			},
			fieldName:     "CIDRv6",
			expectedError: "invalid Classless Inter-Domain Routing CIDRv6",
		},
		{
			name: "DataURI validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				DataURI: "not-datauri", // Not a valid Data URI
			},
			fieldName:     "DataURI",
			expectedError: "invalid Data URL",
		},
		{
			name: "HostnamePort validation",
			data: TestStructForMsg{
				Name:         "John Doe",
				Email:        "test@example.com",
				Age:          30,
				HostnamePort: "not-hostname-port", // Not a valid hostname:port
			},
			fieldName:     "HostnamePort",
			expectedError: "invalid HostPort",
		},
		//{
		//	name: "HostnameRFC1123 validation",
		//	data: TestStructForMsg{
		//		Name:            "John Doe",
		//		Email:           "test@example.com",
		//		Age:             30,
		//		HostnameRFC1123: "not-hostname-rfc1123", // Not a valid hostname RFC 1123
		//	},
		//	fieldName:     "HostnameRFC1123",
		//	expectedError: "invalid Hostname RFC 1123",
		//},
		{
			name: "IP4Addr validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				IP4Addr: "not-ip4-addr", // Not a valid IPv4 address
			},
			fieldName:     "IP4Addr",
			expectedError: "invalid Internet Protocol Address IPv4",
		},
		{
			name: "IP6Addr validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				IP6Addr: "not-ip6-addr", // Not a valid IPv6 address
			},
			fieldName:     "IP6Addr",
			expectedError: "invalid Internet Protocol Address IPv6",
		},
		{
			name: "IPAddr validation",
			data: TestStructForMsg{
				Name:   "John Doe",
				Email:  "test@example.com",
				Age:    30,
				IPAddr: "not-ip-addr", // Not a valid IP address
			},
			fieldName:     "IPAddr",
			expectedError: "invalid Internet Protocol Address IP",
		},
		{
			name: "TCP4Addr validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				TCP4Addr: "not-tcp4-addr", // Not a valid TCPv4 address
			},
			fieldName:     "TCP4Addr",
			expectedError: "invalid Transmission Control Protocol Address TCPv4",
		},
		{
			name: "TCP6Addr validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				TCP6Addr: "not-tcp6-addr", // Not a valid TCPv6 address
			},
			fieldName:     "TCP6Addr",
			expectedError: "Transmission Control Protocol Address TCPv6",
		},
		{
			name: "TCPAddr validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				TCPAddr: "not-tcp-addr", // Not a valid TCP address
			},
			fieldName:     "TCPAddr",
			expectedError: "invalid Transmission Control Protocol Address TCP",
		},
		{
			name: "UDP4Addr validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				UDP4Addr: "not-udp4-addr", // Not a valid UDPv4 address
			},
			fieldName:     "UDP4Addr",
			expectedError: "invalid User Datagram Protocol Address UDPv4",
		},
		{
			name: "UDP6Addr validation",
			data: TestStructForMsg{
				Name:     "John Doe",
				Email:    "test@example.com",
				Age:      30,
				UDP6Addr: "not-udp6-addr", // Not a valid UDPv6 address
			},
			fieldName:     "UDP6Addr",
			expectedError: "invalid User Datagram Protocol Address UDPv6",
		},
		{
			name: "UDPAddr validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				UDPAddr: "not-udp-addr", // Not a valid UDP address
			},
			fieldName:     "UDPAddr",
			expectedError: "invalid User Datagram Protocol Address UDP",
		},
		//{
		//	name: "UnixAddr validation",
		//	data: TestStructForMsg{
		//		Name:     "John Doe",
		//		Email:    "test@example.com",
		//		Age:      30,
		//		UnixAddr: "not-unix-addr", // Not a valid Unix domain socket address
		//	},
		//	fieldName:     "UnixAddr",
		//	expectedError: "invalid Unix domain socket end point Address",
		//},
		{
			name: "URI validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				URI:   "not-uri", // Not a valid URI
			},
			fieldName:     "URI",
			expectedError: "invalid URI string",
		},
		{
			name: "HTTPURL validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				HTTPURL: "not-http-url", // Not a valid HTTP URL
			},
			fieldName:     "HTTPURL",
			expectedError: "invalid HTTP URL string",
		},
		//{
		//	name: "URLEncoded validation",
		//	data: TestStructForMsg{
		//		Name:       "John Doe",
		//		Email:      "test@example.com",
		//		Age:        30,
		//		URLEncoded: "not-url-encoded", // Not a valid URL encoded string
		//	},
		//	fieldName:     "URLEncoded",
		//	expectedError: "invalid URL encoded",
		//},
		{
			name: "URNRFC2141 validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				URNRFC2141: "not-urn-rfc2141", // Not a valid URN RFC 2141
			},
			fieldName:     "URNRFC2141",
			expectedError: "invalid Urn RFC 2141 string",
		},
		{
			name: "Dir validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Dir:   "/path/to/nonexistent/dir", // Not an existing directory
			},
			fieldName:     "Dir",
			expectedError: "existing directory",
		},
		{
			name: "DirPath validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				DirPath: "invalid/dir/path", // Not a valid directory path
			},
			fieldName:     "DirPath",
			expectedError: "invalid directory path",
		},
		{
			name: "File validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				File:  "/path/to/nonexistent/file", // Not an existing file
			},
			fieldName:     "File",
			expectedError: "existing file",
		},
		//{
		//	name: "FilePath validation",
		//	data: TestStructForMsg{
		//		Name:     "John Doe",
		//		Email:    "test@example.com",
		//		Age:      30,
		//		FilePath: "invalid/file/path", // Not a valid file path
		//	},
		//	fieldName:     "FilePath",
		//	expectedError: "invalid file path",
		//},
		{
			name: "Image validation",
			data: TestStructForMsg{
				Name:  "John Doe",
				Email: "test@example.com",
				Age:   30,
				Image: "not-an-image", // Not a valid image
			},
			fieldName:     "Image",
			expectedError: "invalid image",
		},
		{
			name: "IsDefault validation",
			data: TestStructForMsg{
				Name:      "John Doe",
				Email:     "test@example.com",
				Age:       30,
				IsDefault: "not-default", // Not the default value
			},
			fieldName:     "IsDefault",
			expectedError: "is default",
		},
		{
			name: "RequiredIf validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				RequiredIf: "", // Required if some condition is met
			},
			fieldName:     "RequiredIf",
			expectedError: "required if",
		},
		{
			name: "RequiredUnless validation",
			data: TestStructForMsg{
				Name:           "John Doe",
				Email:          "test@example.com",
				Age:            30,
				RequiredUnless: "", // Required unless some condition is met
			},
			fieldName:     "RequiredUnless",
			expectedError: "required unless",
		},
		//{
		//	name: "RequiredWith validation",
		//	data: TestStructForMsg{
		//		Name:         "John Doe",
		//		Email:        "test@example.com",
		//		Age:          30,
		//		RequiredWith: "", // Required with some other field
		//	},
		//	fieldName:     "RequiredWith",
		//	expectedError: "required with",
		//},
		{
			name: "RequiredWithAll validation",
			data: TestStructForMsg{
				Name:            "John Doe",
				Email:           "test@example.com",
				Age:             30,
				RequiredWithAll: "", // Required with all other fields
			},
			fieldName:     "RequiredWithAll",
			expectedError: "required with all",
		},
		{
			name: "RequiredWithout validation",
			data: TestStructForMsg{
				Name:            "John Doe",
				Email:           "test@example.com",
				Age:             30,
				RequiredWithout: "", // Required without some other field
			},
			fieldName:     "RequiredWithout",
			expectedError: "required without",
		},
		{
			name: "RequiredWithoutAll validation",
			data: TestStructForMsg{
				Name:               "John Doe",
				Email:              "test@example.com",
				Age:                30,
				RequiredWithoutAll: "", // Required without all other fields
			},
			fieldName:     "RequiredWithoutAll",
			expectedError: "required without all",
		},
		{
			name: "ExcludedIf validation",
			data: TestStructForMsg{
				Name:       "John Doe",
				Email:      "test@example.com",
				Age:        30,
				ExcludedIf: "should-be-excluded", // Should be excluded if some condition is met
			},
			fieldName:     "ExcludedIf",
			expectedError: "excluded if",
		},
		//{
		//	name: "ExcludedUnless validation",
		//	data: TestStructForMsg{
		//		Name:           "John Doe",
		//		Email:          "test@example.com",
		//		Age:            30,
		//		ExcludedUnless: "should-be-excluded", // Should be excluded unless some condition is met
		//	},
		//	fieldName:     "ExcludedUnless",
		//	expectedError: "excluded unless",
		//},
		//{
		//	name: "ExcludedWith validation",
		//	data: TestStructForMsg{
		//		Name:         "John Doe",
		//		Email:        "test@example.com",
		//		Age:          30,
		//		ExcludedWith: "should-be-excluded", // Should be excluded with some other field
		//	},
		//	fieldName:     "ExcludedWith",
		//	expectedError: "excluded with",
		//},
		{
			name: "ExcludedWithAll validation",
			data: TestStructForMsg{
				Name:            "John Doe",
				Email:           "test@example.com",
				Age:             30,
				ExcludedWithAll: "should-be-excluded", // Should be excluded with all other fields
			},
			fieldName:     "ExcludedWithAll",
			expectedError: "excluded with all",
		},
		//{
		//	name: "ExcludedWithout validation",
		//	data: TestStructForMsg{
		//		Name:            "John Doe",
		//		Email:           "test@example.com",
		//		Age:             30,
		//		ExcludedWithout: "should-be-excluded", // Should be excluded without some other field
		//	},
		//	fieldName:     "ExcludedWithout",
		//	expectedError: "excluded without",
		//},
		{
			name: "ExcludedWithoutAll validation",
			data: TestStructForMsg{
				Name:               "John Doe",
				Email:              "test@example.com",
				Age:                30,
				ExcludedWithoutAll: "should-be-excluded", // Should be excluded without all other fields
			},
			fieldName:     "ExcludedWithoutAll",
			expectedError: "excluded without all",
		},
		{
			name: "IsColor validation",
			data: TestStructForMsg{
				Name:    "John Doe",
				Email:   "test@example.com",
				Age:     30,
				IsColor: "not-a-color", // Not a valid color
			},
			fieldName:     "IsColor",
			expectedError: "color format hexcolor|rgb|rgba|hsl|hsla",
		},
		{
			name: "CountryCode validation",
			data: TestStructForMsg{
				Name:        "John Doe",
				Email:       "test@example.com",
				Age:         30,
				CountryCode: "not-a-country-code", // Not a valid country code
			},
			fieldName:     "CountryCode",
			expectedError: "country format iso3166_1_alpha2|iso3166_1_alpha3|iso3166_1_alpha_numeric",
		},
		// Default error case
		{
			name: "Default error case",
			data: TestStructForMsg{
				Name:  "",
				Email: "test@example.com",
				Age:   30,
				// Use a field with a validation tag that doesn't have a specific message in MsgForTag
			},
			fieldName:     "Name",
			expectedError: "is required", // This is the default error message
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Validate the struct
			err := v.Struct(tt.data)
			if err == nil {
				t.Fatalf("Expected validation error for %s, got nil", tt.fieldName)
			}

			// Check if it's a validation error
			var validationErrors validator.ValidationErrors
			var valErrs validator.ValidationErrors
			if errors.As(err, &valErrs) {
				validationErrors = valErrs
			}

			// Find the error for the specific field
			var fieldError validator.FieldError
			for _, fe := range validationErrors {
				if fe.Field() == tt.fieldName {
					fieldError = fe
					break
				}
			}

			if fieldError == nil {
				t.Fatalf("No validation error found for field %s", tt.fieldName)
			}

			// Test the MsgForTag function
			message := MsgForTag(fieldError)
			if message != tt.expectedError {
				t.Errorf("Expected error message '%s', got '%s'", tt.expectedError, message)
			}
		})
	}
}

// MockFieldError is a mock implementation of validator.FieldError for testing
type MockFieldError struct {
	tag    string
	param  string
	field  string
	errMsg string
}

func (m MockFieldError) Tag() string {
	return m.tag
}

func (m MockFieldError) ActualTag() string {
	return m.tag
}

func (m MockFieldError) Namespace() string {
	return ""
}

func (m MockFieldError) StructNamespace() string {
	return ""
}

func (m MockFieldError) Field() string {
	return m.field
}

func (m MockFieldError) StructField() string {
	return m.field
}

func (m MockFieldError) Value() interface{} {
	return nil
}

func (m MockFieldError) Param() string {
	return m.param
}

func (m MockFieldError) Kind() reflect.Kind {
	return reflect.String
}

func (m MockFieldError) Type() reflect.Type {
	return reflect.TypeOf("")
}

func (m MockFieldError) Error() string {
	return m.errMsg
}

func (m MockFieldError) Translate(ut ut.Translator) string {
	return ""
}

// TestMsgForTagDirect tests the MsgForTag function directly with mock field errors
func TestMsgForTagDirect(t *testing.T) {
	tests := []struct {
		name          string
		mockError     MockFieldError
		expectedError string
	}{
		{
			name: "eq_ignore_case tag",
			mockError: MockFieldError{
				tag:    "eq_ignore_case",
				param:  "TEST",
				field:  "EqualIgnoreCase",
				errMsg: "default error",
			},
			expectedError: "equal ignoring case TEST",
		},
		{
			name: "lt tag",
			mockError: MockFieldError{
				tag:    "lt",
				param:  "5",
				field:  "LessThan",
				errMsg: "default error",
			},
			expectedError: "less than 5",
		},
		{
			name: "lte tag",
			mockError: MockFieldError{
				tag:    "lte",
				param:  "5",
				field:  "LessThanEqual",
				errMsg: "default error",
			},
			expectedError: "less than or equal 5",
		},
		{
			name: "spicedb tag",
			mockError: MockFieldError{
				tag:    "spicedb",
				param:  "",
				field:  "SpiceDB",
				errMsg: "default error",
			},
			expectedError: "invalid SpiceDb ObjectID/Permission/Type",
		},
		{
			name: "fieldexcludes tag",
			mockError: MockFieldError{
				tag:    "fieldexcludes",
				param:  "test",
				field:  "FieldExcludes",
				errMsg: "default error",
			},
			expectedError: "field excludes content test",
		},
		{
			name: "hostname_rfc1123 tag",
			mockError: MockFieldError{
				tag:    "hostname_rfc1123",
				param:  "",
				field:  "HostnameRFC1123",
				errMsg: "default error",
			},
			expectedError: "invalid Hostname RFC 1123",
		},
		{
			name: "hostname_rfc1123 tag (duplicate to ensure coverage)",
			mockError: MockFieldError{
				tag:    "hostname_rfc1123",
				param:  "",
				field:  "HostnameRFC1123Duplicate",
				errMsg: "default error",
			},
			expectedError: "invalid Hostname RFC 1123",
		},
		{
			name: "hostname_port tag",
			mockError: MockFieldError{
				tag:    "hostname_port",
				param:  "",
				field:  "HostnamePort",
				errMsg: "default error",
			},
			expectedError: "invalid HostPort",
		},
		{
			name: "unix_addr tag",
			mockError: MockFieldError{
				tag:    "unix_addr",
				param:  "",
				field:  "UnixAddr",
				errMsg: "default error",
			},
			expectedError: "invalid Unix domain socket end point Address",
		},
		{
			name: "url_encoded tag",
			mockError: MockFieldError{
				tag:    "url_encoded",
				param:  "",
				field:  "URLEncoded",
				errMsg: "default error",
			},
			expectedError: "invalid URL encoded",
		},
		{
			name: "filepath tag",
			mockError: MockFieldError{
				tag:    "filepath",
				param:  "",
				field:  "FilePath",
				errMsg: "default error",
			},
			expectedError: "invalid file path",
		},
		{
			name: "required_with tag",
			mockError: MockFieldError{
				tag:    "required_with",
				param:  "",
				field:  "RequiredWith",
				errMsg: "default error",
			},
			expectedError: "required with",
		},
		{
			name: "excluded_unless tag",
			mockError: MockFieldError{
				tag:    "excluded_unless",
				param:  "",
				field:  "ExcludedUnless",
				errMsg: "default error",
			},
			expectedError: "excluded unless",
		},
		{
			name: "excluded_with tag",
			mockError: MockFieldError{
				tag:    "excluded_with",
				param:  "",
				field:  "ExcludedWith",
				errMsg: "default error",
			},
			expectedError: "excluded with",
		},
		{
			name: "excluded_without tag",
			mockError: MockFieldError{
				tag:    "excluded_without",
				param:  "",
				field:  "ExcludedWithout",
				errMsg: "default error",
			},
			expectedError: "excluded without",
		},
		{
			name: "default error case",
			mockError: MockFieldError{
				tag:    "unknown_tag",
				param:  "",
				field:  "UnknownField",
				errMsg: "default error message",
			},
			expectedError: "default error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := MsgForTag(tt.mockError)
			if message != tt.expectedError {
				t.Errorf("Expected error message '%s', got '%s'", tt.expectedError, message)
			}
		})
	}

	// Direct test for hostname_rfc1123 tag to ensure coverage
	t.Run("hostname_rfc1123 direct test", func(t *testing.T) {
		mockError := MockFieldError{
			tag:    "hostname_rfc1123",
			param:  "",
			field:  "HostnameRFC1123Direct",
			errMsg: "default error",
		}

		message := MsgForTag(mockError)
		expected := "invalid Hostname RFC 1123"
		if message != expected {
			t.Errorf("Expected error message '%s', got '%s'", expected, message)
		}
	})

	// Direct test for hostname tag to ensure coverage
	t.Run("hostname direct test", func(t *testing.T) {
		mockError := MockFieldError{
			tag:    "hostname",
			param:  "",
			field:  "HostnameDirect",
			errMsg: "default error",
		}

		message := MsgForTag(mockError)
		expected := "invalid Hostname RFC 952"
		if message != expected {
			t.Errorf("Expected error message '%s', got '%s'", expected, message)
		}
	})
}
