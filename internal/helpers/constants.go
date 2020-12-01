package helpers

import (
	"fmt"
	"strconv"
)

const runeMetricsURL string = "https://apps.runescape.com/runemetrics"
const runeMetricsProfileEndpoint = runeMetricsURL + "/profile/profile?user=%s&activities=%s"

// CreateRuneMetricsProfileEndpoint formats an endpoint with the given arguments
func CreateRuneMetricsProfileEndpoint(user string, activityCount uint32) string {
	count := strconv.FormatUint(uint64(activityCount), 10)

	return fmt.Sprintf(runeMetricsProfileEndpoint, user, count)
}
