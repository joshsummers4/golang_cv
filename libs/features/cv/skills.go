package cv

import (
	"context"
	"fmt"

	"github.com/joshsummers4/golang_cv/libs/utils/database"
	"github.com/joshsummers4/golang_cv/libs/utils/logger"
)

type Skill struct {
	Skill string "json:skill"
	Type  string "json:type"
}

func GetSkills(ctx context.Context) []Skill {
	db, err := database.Open(fmt.Sprintf("%s/%s", routePath, skillsTableName))
	if err != nil {
		logger.Error(ctx, "error opening skills database", err, []string{"server"}, nil)
		return []Skill{}
	}

	defer db.Close()
	query := "SELECT skill, type FROM skills"
	rows, err := db.Query(query)
	if err != nil {
		logger.Error(ctx, "error querying skills database", err, []string{"server"}, nil)
		return []Skill{}
	}

	var skills []Skill
	for rows.Next() {
		var skill Skill
		err := rows.Scan(&skill.Skill, &skill.Type)
		if err != nil {
			logger.Error(ctx, "error scanning skills database", err, []string{"server"}, nil)
			return []Skill{}
		}
		skills = append(skills, skill)
	}
	logger.Trace(ctx, "get skills", []string{"server"}, map[string]any{"skills": len(skills)})
	return skills
}
