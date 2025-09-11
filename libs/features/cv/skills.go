package cv

import (
	"context"
	"fmt"

	"github.com/joshsummers4/golang_cv/libs/utils/database"
)

type Skill struct {
	Skill string "json:skill"
	Type  string "json:type"
}

func GetSkills(ctx context.Context) []Skill {
	db, err := database.Open(fmt.Sprintf("%s/%s", routePath, skillsTableName))
	if err != nil {
		fmt.Printf("error opening skills database: %v\n", err)
		return []Skill{}
	}

	defer db.Close()
	query := "SELECT skill, type FROM skills"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("error querying skills database: %v\n", err)
		return []Skill{}
	}

	var skills []Skill
	for rows.Next() {
		var skill Skill
		err := rows.Scan(&skill.Skill, &skill.Type)
		if err != nil {
			fmt.Printf("error scanning skill row: %v\n", err)
			return []Skill{}
		}
		skills = append(skills, skill)
	}
	fmt.Printf("get skills: %d\n", len(skills))
	return skills
}
