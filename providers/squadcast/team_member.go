package squadcast

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type TeamMemberGenerator struct {
	SquadcastService
}

type TeamMember struct {
	UserID string `json:"user_id"`
}

func (g *TeamMemberGenerator) createResources(team Team) []terraformutils.Resource {
	var teamMemberList []terraformutils.Resource
	for _, member := range team.Members {
		teamMemberList = append(teamMemberList, terraformutils.NewResource(
			member.UserID,
			fmt.Sprintf("squadcast_team_member_%s", member.UserID),
			"squadcast_team_member",
			g.GetProviderName(),
			map[string]string{
				"team_id": team.ID,
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return teamMemberList
}

func (g *TeamMemberGenerator) InitResources() error {
	teamName := g.Args["team_name"].(string)
	if len(teamName) == 0 {
		return errors.New("--team-name is required")
	}
	escapedTeamName := url.QueryEscape(teamName)
	getTeamURL := fmt.Sprintf("/v3/teams/by-name?name=%s", escapedTeamName)
	response, err := Request[Team](getTeamURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
