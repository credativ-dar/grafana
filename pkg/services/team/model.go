package team

import (
	"errors"
	"time"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/user"
)

// Typed errors
var (
	ErrTeamNotFound                         = errors.New("team not found")
	ErrTeamNameTaken                        = errors.New("team name is taken")
	ErrTeamMemberNotFound                   = errors.New("team member not found")
	ErrLastTeamAdmin                        = errors.New("not allowed to remove last admin")
	ErrNotAllowedToUpdateTeam               = errors.New("user not allowed to update team")
	ErrNotAllowedToUpdateTeamInDifferentOrg = errors.New("user not allowed to update team in another org")

	ErrTeamMemberAlreadyAdded = errors.New("user is already added to this team")
)

// Team model
type Team struct {
	ID    int64  `json:"id" xorm:"pk autoincr 'id'"`
	OrgID int64  `json:"orgId" xorm:"org_id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// ---------------------
// COMMANDS

type CreateTeamCommand struct {
	Name  string `json:"name" binding:"Required"`
	Email string `json:"email"`
	OrgID int64  `json:"-"`
}

type UpdateTeamCommand struct {
	ID    int64
	Name  string
	Email string
	OrgID int64 `json:"-"`
}

type DeleteTeamCommand struct {
	OrgID int64
	ID    int64
}

type GetTeamByIDQuery struct {
	OrgID        int64
	ID           int64
	SignedInUser *user.SignedInUser
	HiddenUsers  map[string]struct{}
	UserIdFilter int64
}

// FilterIgnoreUser is used in a get / search teams query when the caller does not want to filter teams by user ID / membership
const FilterIgnoreUser int64 = 0

type GetTeamsByUserQuery struct {
	OrgID        int64
	UserID       int64 `json:"userId"`
	SignedInUser *user.SignedInUser
}

type SearchTeamsQuery struct {
	Query        string
	Name         string
	Limit        int
	Page         int
	OrgID        int64 `xorm:"org_id"`
	UserIDFilter int64 `xorm:"user_id_filter"`
	SignedInUser *user.SignedInUser
	HiddenUsers  map[string]struct{}
}

type TeamDTO struct {
	ID            int64                 `json:"id" xorm:"id"`
	OrgID         int64                 `json:"orgId" xorm:"org_id"`
	Name          string                `json:"name"`
	Email         string                `json:"email"`
	AvatarURL     string                `json:"avatarUrl"`
	MemberCount   int64                 `json:"memberCount"`
	Permission    models.PermissionType `json:"permission"`
	AccessControl map[string]bool       `json:"accessControl"`
}

type SearchTeamQueryResult struct {
	TotalCount int64      `json:"totalCount"`
	Teams      []*TeamDTO `json:"teams"`
	Page       int        `json:"page"`
	PerPage    int        `json:"perPage"`
}

type IsAdminOfTeamsQuery struct {
	SignedInUser *user.SignedInUser
}

// TeamMember model
type TeamMember struct {
	ID         int64 `xorm:"pk autoincr 'id'"`
	OrgID      int64 `xorm:"org_id"`
	TeamID     int64 `xorm:"team_id"`
	UserID     int64 `xorm:"user_id"`
	External   bool  // Signals that the membership has been created by an external systems, such as LDAP
	Permission models.PermissionType

	Created time.Time
	Updated time.Time
}

// ---------------------
// COMMANDS

type AddTeamMemberCommand struct {
	UserID     int64                 `json:"userId" binding:"Required"`
	OrgID      int64                 `json:"-"`
	TeamID     int64                 `json:"-"`
	External   bool                  `json:"-"`
	Permission models.PermissionType `json:"-"`
}

type UpdateTeamMemberCommand struct {
	UserID     int64                 `json:"-"`
	OrgID      int64                 `json:"-"`
	TeamID     int64                 `json:"-"`
	Permission models.PermissionType `json:"permission"`
}

type RemoveTeamMemberCommand struct {
	OrgID  int64 `json:"-"`
	UserID int64
	TeamID int64
}

// ----------------------
// QUERIES

type GetTeamMembersQuery struct {
	OrgID        int64
	TeamID       int64
	UserID       int64
	External     bool
	SignedInUser *user.SignedInUser
}

// ----------------------
// Projections and DTOs

type TeamMemberDTO struct {
	OrgID      int64                 `json:"orgId" xorm:"org_id"`
	TeamID     int64                 `json:"teamId" xorm:"team_id"`
	UserID     int64                 `json:"userId" xorm:"user_id"`
	External   bool                  `json:"-"`
	AuthModule string                `json:"auth_module"`
	Email      string                `json:"email"`
	Name       string                `json:"name"`
	Login      string                `json:"login"`
	AvatarURL  string                `json:"avatarUrl" xorm:"avatar_url"`
	Labels     []string              `json:"labels"`
	Permission models.PermissionType `json:"permission"`
}
