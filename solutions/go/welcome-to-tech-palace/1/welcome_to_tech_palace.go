package techpalace

import ("strings")

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer);
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var stars string = strings.Repeat("*", numStarsPerLine);
	return stars + "\n" + welcomeMsg + "\n" + stars;
}

// CleanupMessage cleans up an old marketing message.
/*
func CleanupMessage(oldMsg string) string {
    var s = strings.Replace(oldMsg, "*", " ", -1);
    s = strings.Replace(s, "\n", " ", -1);
    s = strings.TrimLeft(s, " ");
    s = strings.TrimRight(s, " ");
	return s;
}
*/
func CleanupMessage(oldMsg string) string {
    r := strings.NewReplacer("*", " ", "\n", " ",);
    s := r.Replace(oldMsg);
    return strings.TrimSpace(s);
}
