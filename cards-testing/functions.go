package cards_testing

/** Test Helper Functions **/
func GetEmptyString() string {
    return "\"\""
}

func GetEmptyStringMaybe(str string) string {
    if str == "" {
        str = GetEmptyString()
    }
    return str
}

func GetEmptyStrings(strs ...string) []string {
    for i, str := range strs {
        strs[i] = GetEmptyStringMaybe(str)
    }
    return strs
}
