package helpers

const runemetricsURL string = "https://apps.runescape.com/runemetrics"

// RunemetricsProfileEndpoint provides a formatted string for use in calling the
// Runemetrics profile endpoint
const RunemetricsProfileEndpoint = runemetricsURL + "/profile/profile?user=%s&activities=%s"
