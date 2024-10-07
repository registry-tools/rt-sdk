package token
type TokenPostRequestBody_grant_type int

const (
    AUTHORIZATION_CODE_TOKENPOSTREQUESTBODY_GRANT_TYPE TokenPostRequestBody_grant_type = iota
)

func (i TokenPostRequestBody_grant_type) String() string {
    return []string{"authorization_code"}[i]
}
func ParseTokenPostRequestBody_grant_type(v string) (any, error) {
    result := AUTHORIZATION_CODE_TOKENPOSTREQUESTBODY_GRANT_TYPE
    switch v {
        case "authorization_code":
            result = AUTHORIZATION_CODE_TOKENPOSTREQUESTBODY_GRANT_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTokenPostRequestBody_grant_type(values []TokenPostRequestBody_grant_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TokenPostRequestBody_grant_type) isMultiValue() bool {
    return false
}
