package helpers

import (
	"fmt"
	"strconv"
)

const runeMetricsURL string = "https://apps.runescape.com/runemetrics"
const runeMetricsProfileEndpoint = runeMetricsURL + "/profile/profile?user=%s&activities=%s"

// CreateRuneMetricsProfileEndpoint formats an endpoint with the given arguments
func CreateRuneMetricsProfileEndpoint(user string, activityCount int) string {
	count := strconv.FormatInt(int64(activityCount), 10)

	return fmt.Sprintf(runeMetricsProfileEndpoint, user, count)
}
