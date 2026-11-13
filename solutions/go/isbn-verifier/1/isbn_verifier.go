package isbnverifier

import "strings"

// IsValidISBN check if the provided string is a valid ISBN-10
func IsValidISBN(isbn string) bool {
    result := 0
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
        return false
    }
	// first 9 digits
    for i := 0; i < 9; i++ {
        if (isbn[i] < '0' || isbn[i] > '9') {
            return false
        }
        digit := int(isbn[i] - '0')
        result += digit * (10 - i)
    }
    // last isbn digit
    if isbn[9] == 'X' {
        result += 10
    } else if isbn[9] >= '0' && isbn[9] <= '9'{
        result += int(isbn[9] - '0')
    } else {
        return false
    }

	return result % 11 == 0
}
