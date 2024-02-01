package piscine

func SplitWhiteSpaces(s string) []string {
    var result []string
    var word string
    for i, c := range s {
        if c != ' ' && c != '\t' && c != '\n' {
            word += string(c)
        }
        if ((c == ' ' || c == '\t' || c == '\n') && word != "") || i == len(s)-1 {
            result = append(result, word)
            word = ""
        }
    }
    return result
}

func Split(s, sep string) []string {
    ln := 0

    for i := range sep {
        ln = i + 1
    }

    ln2 := 0
    for i := range s {
        ln2 = i + 1
    }

    for i := 0; i < ln2-ln; i++ {
        if s[i:i+ln] == sep {
            s = s[:i] + " " + s[i+ln:]
            ln2 -= ln
        }
    }

    return SplitWhiteSpaces(s)
}