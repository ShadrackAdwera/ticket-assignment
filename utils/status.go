package utils

const (
	ACTIVE   = "ACTIVE"
	INACTIVE = "INACTIVE"
)

func IsValidAgentStatus(status string) bool {
	switch status {
	case ACTIVE, INACTIVE:
		return true
	default:
		return false
	}
}
