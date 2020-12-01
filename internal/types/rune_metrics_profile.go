package types

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	rs3pb "github.com/alog-rs/proto/rs3"
	"github.com/alog-rs/shared-packages/pkg/types"
)

// RuneMetricsPlayerProfileError represents an error returned form the PlayerProfile endpoint
type RuneMetricsProfileError int

const (
	// ProfileErrorNotFound is returned when the user does not exist
	ProfileErrorNotFound RuneMetricsProfileError = iota
	// ProfileErrorPrivate is returned when the users profile is private
	ProfileErrorPrivate
	// ProfileErrorUnknown is returned when an unknown error is found
	ProfileErrorUnknown
	// ProfileErrorNone is returned when no error is returned
	ProfileErrorNone
)

// RuneMetricsActivity represents a players recent activity
type RuneMetricsActivity struct {
	// Date is the date string from when the activity occured
	Date string `json:"date"`
	// Details is the long form description of the activity
	Details string `json:"details"`
	// Text is the short form description of the activity
	Text string `json:"text"`
}

// RuneMetricsSkillValue represents a players statistics for a single skill
type RuneMetricsSkillValue struct {
	// Level is the users level in this skill
	Level uint32 `json:"level"`
	// XP is the users xp in this skill
	XP uint64 `json:"xp"`
	// Rank is the users rank in this skill
	Rank uint32 `json:"rank"`
	// ID is the ID for the skill
	ID types.Skill `json:"id"`
}

// RuneMetricsProfile represents a players profile directly from the
// runemetrics/profile/profile?users=<name>&activities=<count> endpoint
type RuneMetricsProfile struct {
	// Name is the users username
	Name string `json:"name"`
	// LoggedIn
	LoggedIn string `json:"loggedIn"`
	// Rank is the users overall rank
	Rank string `json:"rank"`
	// TotalSkill is the users total level
	TotalSkill uint32 `json:"totalskill"`
	// TotalXP is the users total XP across all skills
	TotalXP uint64 `json:"totalxp"`
	// CombatLevel is the users combat level
	CombatLevel uint32 `json:"combatlevel"`
	// Melee is the users total XP in all melee skills
	Melee uint64 `json:"melee"`
	// Ranged is the users total XP in the ranged skill
	Ranged uint64 `json:"ranged"`
	// Magic is the users total XP in the magic skill
	Magic uint64 `json:"magic"`
	// QuestsStarted are the number of quests the user has started but not completed
	QuestsStarted uint32 `json:"questsstarted"`
	// QuestsComplete is the number of quests the user has completed
	QuestsComplete uint32 `json:"questscomplete"`
	// QuestsNotStarted is the number of quests the user has yet to start
	QuestsNotStarted uint32 `json:"questsnotstarted"`
	// Activities is a list of the users recent activity
	Activities []RuneMetricsActivity `json:"activities"`
	// SkillValues is a list of the users skills
	SkillValues []RuneMetricsSkillValue `json:"skillvalues"`
	// Error is populated if an error occured
	Error string `json:"error,omitempty"`
}

func activitiesToPb(activities []RuneMetricsActivity) ([]*rs3pb.PlayerActivityItem, error) {
	items := make([]*rs3pb.PlayerActivityItem, len(activities))

	for i, activity := range activities {
		ts, err := time.Parse("02-Jan-2006 15:04", activity.Date)

		if err != nil {
			return nil, err
		}

		items[i] = &rs3pb.PlayerActivityItem{
			Timestamp: ts.Unix(),
			Long:      activity.Details,
			Short:     activity.Text,
		}
	}

	return items, nil
}

func skillsToPb(skills []RuneMetricsSkillValue) ([]*rs3pb.SkillData, error) {
	items := make([]*rs3pb.SkillData, len(skills))

	for i, skill := range skills {
		s := types.Skill(skill.ID)
		// JAGEX why do you do stupid things like this?
		xp := uint64(float64(skill.XP) * 0.1)

		items[i] = &rs3pb.SkillData{
			Rank:         skill.Rank,
			Level:        skill.Level,
			VirtualLevel: s.GetVirtualLevel(xp),
			Xp:           xp,
		}
	}

	return items, nil
}

// ProfileErrorFromString parses a string into a RuneMetricsPlayerProfileError
func ProfileErrorFromString(str string) RuneMetricsProfileError {
	switch str {
	case "NOT_FOUND":
		return ProfileErrorNotFound
	case "PROFILE_PRIVATE":
		return ProfileErrorPrivate
	case "":
		return ProfileErrorNone
	default:
		return ProfileErrorUnknown
	}
}

// NewRuneMetricsPlayerProfile takes JSON as input and attempts to create a
// RuneMetricsPlayerProfile
func NewRuneMetricsPlayerProfile(input []byte) (*RuneMetricsProfile, error) {
	var profile RuneMetricsProfile

	if err := json.Unmarshal(input, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

// GetError returns an error if it exists from a PlayerProfile
func (p *RuneMetricsProfile) GetError() RuneMetricsProfileError {
	return ProfileErrorFromString(p.Error)
}

// ConvertToPB converts a RuneMetricsProfile into our protobuf PlayerProfile
func (p *RuneMetricsProfile) ConvertToPB() (*rs3pb.PlayerProfile, error) {
	rank, rankErr := strconv.ParseUint(strings.Replace(p.Rank, ",", "", -1), 10, 32)

	if rankErr != nil {
		return nil, errors.New("Failed to parse rank from RuneMetricsProfile")
	}

	activities, activitiesErr := activitiesToPb(p.Activities)

	if activitiesErr != nil {
		return nil, errors.New("Failed to parse activies from RuneMetricsProfile")
	}

	skills, skillsErr := skillsToPb(p.SkillValues)

	if skillsErr != nil {
		return nil, errors.New("Failed to parse skills from RuneMetricsProfile")
	}

	return &rs3pb.PlayerProfile{
		Name:        p.Name,
		Rank:        rank,
		TotalLevel:  p.TotalSkill,
		TotalXp:     p.TotalXP,
		CombatLevel: &p.CombatLevel,
		QuestInfo: &rs3pb.QuestData{
			Completed:  p.QuestsComplete,
			Started:    p.QuestsStarted,
			NotStarted: p.QuestsNotStarted,
		},
		Activity: activities,
		Skills:   skills,
	}, nil
}
