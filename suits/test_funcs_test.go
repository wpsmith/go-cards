package suits

/** Test Helper Functions **/
func getEmptyString() string {
    return "\"\""
}

func getEmptyStringMaybe(str string) string {
    if str == "" {
        str = getEmptyString()
    }
    return str
}

func getEmptyStrings(strs ...string) []string {
    for i, str := range strs {
        strs[i] = getEmptyStringMaybe(str)
    }
    return strs
}
