package internal

func IsAdmin(userRole string) bool {
	return userRole == "admin"
}

func HasPermission(userRole string, requiredLevel int) bool {
	return requiredLevel <= 3 // Simplified permission check example
}
