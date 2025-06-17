# Validation

By default, data checking is supported for structs by `Check(structData any, msgForTag MsgForTagFunc) (map[string]any, error)`

### Usage

Install
```bash
go get -u github.com/gflydev/validation@v1.0.3
```

Quick usage
```go
import "github.com/gflydev/validation"

if errorData, err := validation.Check(loginDto); err != nil {
    return c.BadRequest(errorData)
}
```

Customize error message by yourself
```go
if errorData, err := validation.Check(loginDto, func(fe validator.FieldError) string {
    switch fe.Tag() {
    case "required":
        return "This field is required"
    case "gte":
        return fmt.Sprintf("This field is gte %s characters", fe.Param())
    case "email":
        return "Invalid email"
    }

    return fe.Error()
}); err != nil {
    return c.BadRequest(errorData)
}
```

### Message for Tag

Review and add more code in file `message_for_tag.go`.

### Add custom Validator `uuid_rule.go`

```go
package validation

import (
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
)

// UuidRule Custom validation for uuid fields. Use `validate:"uuid"`
type UuidRule string

func (v UuidRule) GetTag() string {
    return string(v)
}

func (v UuidRule) Handler() validator.Func {
    return func(fl validator.FieldLevel) bool {
        field := fl.Field().String()

        if _, err := uuid.Parse(field); err != nil {
            return false
        }
        return true
    }
}
```
Add rule to validator
```go
import "github.com/vinhio/gfly-modules/validation"

// Add rule before validation.Validate()
validation.AddRule(uuidValidator("uuid"))
```

## Rules

Default rules about `Database`, `Fields`, `Network`, `Strings`, `Format`, `Comparisons`, `Other`, `Aliases`

### Database:

| Tag | Description                                                                                                      |
| - |------------------------------------------------------------------------------------------------------------------|
| db_exists | Value is equal to a specific column in a table. Example: `db_exists=ads.id` with table name `ads` and column `id` |

### Fields:

| Tag | Description |
| - | - |
| eqcsfield | Field Equals Another Field (relative)|
| eqfield | Field Equals Another Field |
| fieldcontains | Check the indicated characters are present in the Field |
| fieldexcludes | Check the indicated characters are not present in the field |
| gtcsfield | Field Greater Than Another Relative Field |
| gtecsfield | Field Greater Than or Equal To Another Relative Field |
| gtefield | Field Greater Than or Equal To Another Field |
| gtfield | Field Greater Than Another Field |
| ltcsfield | Less Than Another Relative Field |
| ltecsfield | Less Than or Equal To Another Relative Field |
| ltefield | Less Than or Equal To Another Field |
| ltfield | Less Than Another Field |
| necsfield | Field Does Not Equal Another Field (relative) |
| nefield | Field Does Not Equal Another Field |

### Network:

| Tag | Description |
| - | - |
| cidr | Classless Inter-Domain Routing CIDR |
| cidrv4 | Classless Inter-Domain Routing CIDRv4 |
| cidrv6 | Classless Inter-Domain Routing CIDRv6 |
| datauri | Data URL |
| fqdn | Full Qualified Domain Name (FQDN) |
| hostname | Hostname RFC 952 |
| hostname_port | HostPort |
| hostname_rfc1123 | Hostname RFC 1123 |
| ip | Internet Protocol Address IP |
| ip4_addr | Internet Protocol Address IPv4 |
| ip6_addr | Internet Protocol Address IPv6 |
| ip_addr | Internet Protocol Address IP |
| ipv4 | Internet Protocol Address IPv4 |
| ipv6 | Internet Protocol Address IPv6 |
| mac | Media Access Control Address MAC |
| tcp4_addr | Transmission Control Protocol Address TCPv4 |
| tcp6_addr | Transmission Control Protocol Address TCPv6 |
| tcp_addr | Transmission Control Protocol Address TCP |
| udp4_addr | User Datagram Protocol Address UDPv4 |
| udp6_addr | User Datagram Protocol Address UDPv6 |
| udp_addr | User Datagram Protocol Address UDP |
| unix_addr | Unix domain socket end point Address |
| uri | URI String |
| url | URL String |
| http_url | HTTP URL String |
| url_encoded | URL Encoded |
| urn_rfc2141 | Urn RFC 2141 String |

### Strings:

| Tag | Description |
| - | - |
| alpha | Alpha Only |
| alphanum | Alphanumeric |
| alphanumunicode | Alphanumeric Unicode |
| alphaunicode | Alpha Unicode |
| ascii | ASCII |
| boolean | Boolean |
| contains | Contains |
| containsany | Contains Any |
| containsrune | Contains Rune |
| endsnotwith | Ends Not With |
| endswith | Ends With |
| excludes | Excludes |
| excludesall | Excludes All |
| excludesrune | Excludes Rune |
| lowercase | Lowercase |
| multibyte | Multi-Byte Characters |
| number | Number |
| numeric | Numeric |
| printascii | Printable ASCII |
| startsnotwith | Starts Not With |
| startswith | Starts With |
| uppercase | Uppercase |

### Format:
| Tag | Description |
| - | - |
| base64 | Base64 String |
| base64url | Base64URL String |
| base64rawurl | Base64RawURL String |
| bic | Business Identifier Code (ISO 9362) |
| bcp47_language_tag | Language tag (BCP 47) |
| btc_addr | Bitcoin Address |
| btc_addr_bech32 | Bitcoin Bech32 Address (segwit) |
| credit_card | Credit Card Number |
| mongodb | MongoDB ObjectID |
| cron | Cron |
| spicedb | SpiceDb ObjectID/Permission/Type |
| datetime | Datetime |
| e164 | e164 formatted phone number |
| email | E-mail String
| eth_addr | Ethereum Address |
| hexadecimal | Hexadecimal String |
| hexcolor | Hexcolor String |
| hsl | HSL String |
| hsla | HSLA String |
| html | HTML Tags |
| html_encoded | HTML Encoded |
| isbn | International Standard Book Number |
| isbn10 | International Standard Book Number 10 |
| isbn13 | International Standard Book Number 13 |
| issn | International Standard Serial Number |
| iso3166_1_alpha2 | Two-letter country code (ISO 3166-1 alpha-2) |
| iso3166_1_alpha3 | Three-letter country code (ISO 3166-1 alpha-3) |
| iso3166_1_alpha_numeric | Numeric country code (ISO 3166-1 numeric) |
| iso3166_2 | Country subdivision code (ISO 3166-2) |
| iso4217 | Currency code (ISO 4217) |
| json | JSON |
| jwt | JSON Web Token (JWT) |
| latitude | Latitude |
| longitude | Longitude |
| luhn_checksum | Luhn Algorithm Checksum (for strings and (u)int) |
| postcode_iso3166_alpha2 | Postcode |
| postcode_iso3166_alpha2_field | Postcode |
| rgb | RGB String |
| rgba | RGBA String |
| ssn | Social Security Number SSN |
| timezone | Timezone |
| uuid | Universally Unique Identifier UUID |
| uuid3 | Universally Unique Identifier UUID v3 |
| uuid3_rfc4122 | Universally Unique Identifier UUID v3 RFC4122 |
| uuid4 | Universally Unique Identifier UUID v4 |
| uuid4_rfc4122 | Universally Unique Identifier UUID v4 RFC4122 |
| uuid5 | Universally Unique Identifier UUID v5 |
| uuid5_rfc4122 | Universally Unique Identifier UUID v5 RFC4122 |
| uuid_rfc4122 | Universally Unique Identifier UUID RFC4122 |
| md4 | MD4 hash |
| md5 | MD5 hash |
| sha256 | SHA256 hash |
| sha384 | SHA384 hash |
| sha512 | SHA512 hash |
| ripemd128 | RIPEMD-128 hash |
| ripemd128 | RIPEMD-160 hash |
| tiger128 | TIGER128 hash |
| tiger160 | TIGER160 hash |
| tiger192 | TIGER192 hash |
| semver | Semantic Versioning 2.0.0 |
| ulid | Universally Unique Lexicographically Sortable Identifier ULID |
| cve | Common Vulnerabilities and Exposures Identifier (CVE id) |

### Comparisons:
| Tag | Description |
| - | - |
| eq | Equals |
| eq_ignore_case | Equals ignoring case |
| gt | Greater than|
| gte | Greater than or equal |
| lt | Less Than |
| lte | Less Than or Equal |
| ne | Not Equal |
| ne_ignore_case | Not Equal ignoring case |

### Other:
| Tag | Description |
| - | - |
| dir | Existing Directory |
| dirpath | Directory Path |
| file | Existing File |
| filepath | File Path |
| image | Image |
| isdefault | Is Default |
| len | Length |
| max | Maximum |
| min | Minimum |
| oneof | One Of |
| required | Required |
| required_if | Required If |
| required_unless | Required Unless |
| required_with | Required With |
| required_with_all | Required With All |
| required_without | Required Without |
| required_without_all | Required Without All |
| excluded_if | Excluded If |
| excluded_unless | Excluded Unless |
| excluded_with | Excluded With |
| excluded_with_all | Excluded With All |
| excluded_without | Excluded Without |
| excluded_without_all | Excluded Without All |
| unique | Unique |

#### Aliases:
| Tag | Description |
| - | - |
| iscolor | hexcolor\|rgb\|rgba\|hsl\|hsla |
| country_code | iso3166_1_alpha2\|iso3166_1_alpha3\|iso3166_1_alpha_numeric |


## Release

- Date 2025-06-17 09:55:32: Add rule `db_exists`
