package models

type PDTeam struct {
	ID   string `json:"id"`
	Name string `json:"summary"`
}

type PDUser struct {
	ID    string   `json:"id"`
	Name  string   `json:"summary"`
	Email string   `json: "email"`
	Teams []PDTeam `json: "teams"`
}

func (u PDUser) GetID() string {
	return u.ID
}

func (u PDUser) GetName() string {
	return u.Name
}

func (u PDUser) GetEmail() string {
	return u.Email
}

func (u PDUser) GetTeams() []string {
	teams := []string{}

	for _, team := range u.Teams {
		teams = append(teams, team.ID)
	}
	return teams
}
