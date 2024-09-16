func backspaceCompare(s string, t string) bool {
    sr := processStr(s)
    tr := processStr(t)
    return string(sr) == string(tr)
}

func processStr(s string) []rune{
    r := make([]rune, 0)
    for _, c := range s {
        if c != '#' {
            r = append(r, c)
        }else{
            if len(r) > 0 {
                r = r[:len(r)-1]
            }
        }
    }
    return r
}

