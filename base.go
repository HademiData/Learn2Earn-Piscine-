package piscine

import "github.com/01-edu/z01"

// PrintNbrBase prints the integer nbr in the numeral system defined by the base string.
// If the base is invalid, it prints "NV".
func PrintNbrBase(nbr int, base string) {
    // Step 1: Validate the base string before doing anything
    if !isValidBase(base) {
        printNV()
        return
    }

    // Step 2: Handle negative numbers
    if nbr < 0 {
        // Print the minus sign to indicate negativity
        z01.PrintRune('-')
        // Convert nbr to positive to simplify the conversion logic
        // This works because we already printed the '-' sign
        nbr = -nbr
    }

    // Step 3: Convert and print the number in the new base
    printNumberBase(nbr, base)
}

// isValidBase checks if the base string is valid according to these rules:
// - length >= 2
// - does not contain '+' or '-'
// - all characters are unique
func isValidBase(base string) bool {
    if len(base) < 2 {
        return false // base too short to represent numbers
    }

    for i, c := range base {
        if c == '+' || c == '-' {
            return false // invalid characters in base
        }

        // Check that this character is unique in the base string
        for j := i + 1; j < len(base); j++ {
            if base[j] == byte(c) {
                return false // duplicate character found
            }
        }
    }

    return true // base is valid
}

// printNV prints "NV" to indicate an invalid base
func printNV() {
    z01.PrintRune('N')
    z01.PrintRune('V')
}

// printNumberBase prints nbr in the base system defined by base string.
// This function uses recursion to print digits from most significant to least significant.
func printNumberBase(nbr int, base string) {
    baseLen := len(base) // the length of the base string defines the base

    // If nbr is greater or equal to baseLen, there are more digits to print before the current digit
    if nbr >= baseLen {
        // Recursive call with the "quotient" of nbr divided by baseLen
        // This effectively "shifts" the number right, moving towards the most significant digit
        printNumberBase(nbr/baseLen, base)
    }

    // Calculate remainder of nbr divided by baseLen to find the current digit's index in the base string
    // This gives the rightmost digit for the current recursion level
    digitIndex := nbr % baseLen

    // Print the digit corresponding to the remainder
    z01.PrintRune(rune(base[digitIndex]))
}
