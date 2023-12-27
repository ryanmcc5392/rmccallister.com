package jsonHandlers

import (
	"encoding/json"
	"net/http"
	"rmccallister/pkg/jsonHandlers/types"
)

func ServeJson(w http.ResponseWriter, r *http.Request) {

	data := CreateData()

	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func CreateData() types.Data {
	data := types.Data{
		About:      CreateDataAbout(),
		Experience: CreateExperience(),
	}
	return data
}

func CreateExperience() types.Experience {
	experience := types.Experience{
		Experience: []types.Job{
			{
				Company:   "DailyWire",
				Position:  "Junior Front End Developer",
				StartDate: "March 2023",
				EndDate:   "Present",
				Technologies: []string{
					"React",
					"JavaScript",
					"GO",
					"HTML",
					"CSS",
					"Tailwind-CSS",
					"Git",
					"GitHub",
					"Jira",
					"Confluence",
					"Figma",
					"Agile",
				},
				Description: "I am currently employed with the DailyWire as a Junior Front End Developer. I work on the front end of the website, and I am responsible for creating new features and fixing bugs. I work with React, JavaScript, GO, HTML, CSS, and Tailwind-CSS.",
			},
		},
	}
	return experience
}

func CreateDataAbout() types.About {
	about := types.About{
		Name: "Ryan McCallister",
	}
	return about
}
