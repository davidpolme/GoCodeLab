package two_pointers

// ReverseString reverses a slice of bytes in place.
func ReverseString(s []byte) {
    for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
        s[left], s[right] = s[right], s[left]
    }
}
