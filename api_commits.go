package secondAssignment

import (
  "encoding/json"
	"net/http"
  "strings"
  //"bytes"
  "strconv"
)

var DBp = ProjectsDB{}											// Stores countries


func replyWithAlls(w http.ResponseWriter, DB projectStorage, limit string){

  limitINT, err := strconv.Atoi(limit)
	if err == nil {
	}

  for i := 1; i <= limitINT; i++ {
    if i == 0 {
      continue
    }

    url := "https://git.gvk.idi.ntnu.no/api/v4/projects/" + strconv.Itoa(i) + "?private_token=tAqvSkyKfx7YVw9zouUz"
  	resp, err := http.Get(url)								// GETs url
  	if err != nil {														// If it doesnt work, return error
  		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      ST.TestApi("Gitlab")
  	}


    /*
  	buf := new(bytes.Buffer)
  	buf.ReadFrom(resp.Body)
  	respByte := buf.Bytes()										// Copy COuntry information
    if err := json.Unmarshal(respByte, &tempProject); err != nil {
  		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
  	}*/

    defer resp.Body.Close()
  	var tempProject Project
    json.NewDecoder(resp.Body).Decode(&tempProject)

    url = "https://git.gvk.idi.ntnu.no/api/v4/projects/" + strconv.Itoa(i) + "/repository/commits?per_page=900&private_token=tAqvSkyKfx7YVw9zouUz"
    resp, err = http.Get(url)								// GETs url
    if err != nil {														// If it doesnt work, return error
      http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      ST.TestApi("Gitlab")
    }
    defer resp.Body.Close()
    var tempCommit[] Commit
    json.NewDecoder(resp.Body).Decode(&tempCommit)

    tempProject.Commits = len(tempCommit)


    if(tempProject.Repository != "") { DB.Add(tempProject) }
  }

	a := make([]Project, 0, DB.Count())		// make map variable for printing
	for _, s := range DB.GetAll() {				// For each country in DB
		a = append(a, s)										// Copy them to a
	}
	json.NewEncoder(w).Encode(a)					// Display as JSON on browser
}








func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
  var limit string = r.URL.Query().Get("limit")
  if limit == ""{
    limit = "5"
  }

	if len(parts) == 6 || parts[1] == "conservation"{
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	replyWithAlls(w, &DBp, limit)
}
