package secondAssignment

import (
  "encoding/json"
	"net/http"
  "strings"
)



type Issues struct {
//	labels		labelInfo `json:"labels"`
Repository						string `json:"title"`
 AuthName      authName    `json:"author"`
}

type labelInfo struct {
	labels		    string    `json:"labels"`
}

type authName struct {
  id            int       `json:"id"`
	username	  	string    `json:"username"`
}


var issueStructure[] Issues
var DBu = UsersDB{}
var DBl = LabelsDB{}


func replyWithAllu(w http.ResponseWriter, DB userStorage){


  url := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/members/all?private_token=yicA2z2b1baW9bBMeYUz"
	resp, err := http.Get(url)								// GETs url
	if err != nil {														// If it doesnt work, return error
  	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
	}
  defer resp.Body.Close()
	var tempUser[] User
  json.NewDecoder(resp.Body).Decode(&tempUser)


  iurl := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/issues?private_token=yicA2z2b1baW9bBMeYUz"
  resp, err = http.Get(iurl)								  // GETs url
  if err != nil {														// If it doesnt work, return error
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
  }
  defer resp.Body.Close()
  json.NewDecoder(resp.Body).Decode(&issueStructure)



  for idx, x := range tempUser{     // For each occurrence
    if idx == 0 {														// Skip header
      println("")
    }
    println(x.Username)
    DBu.Add(x)
  }

  for idx, y := range issueStructure{
    if idx == 0 {														// Skip header
      println("")
    }

    println(y.AuthName.username)

  }

	a := make([]User, 0, DBu.Count())		  // make map variable for printing
	for _, s := range DB.GetAll() {				// For each country in DB
		a = append(a, s)										// Copy them to a
	}
	json.NewEncoder(w).Encode(a)					// Display as JSON on browser
}




func replyWithAlll(w http.ResponseWriter, DB labelsStorage){

  url := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/labels?private_token=yicA2z2b1baW9bBMeYUz"
	resp, err := http.Get(url)								// GETs url
	if err != nil {														// If it doesnt work, return error
  	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
	}
  defer resp.Body.Close()
	var tempLabel[] Label
  json.NewDecoder(resp.Body).Decode(&tempLabel)


  lurl := "https://git.gvk.idi.ntnu.no/api/v4/projects/1/issues?private_token=yicA2z2b1baW9bBMeYUz"
  resp, err = http.Get(lurl)								  // GETs url
  if err != nil {														// If it doesnt work, return error
    http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
    ST.TestApi("Gitlab")
  }
  defer resp.Body.Close()
  json.NewDecoder(resp.Body).Decode(&issueStructure)



  for idx, x := range tempLabel{     // For each occurrence
    if idx == 0 {														// Skip header
      println("")
    }
    DBl.Add(x)
  }


	a := make([]Label, 0, DBl.Count())		  // make map variable for printing
	for _, s := range DBl.GetAll() {				// For each country in DB
		a = append(a, s)										// Copy them to a
	}
	json.NewEncoder(w).Encode(a)					// Display as JSON on browser

}




func HandlerIssues(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
  var issueType string = r.URL.Query().Get("type")


	if len(parts) == 6 || parts[1] == "conservation"{
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
  if(issueType == "user"){
    replyWithAllu(w, &DBu)
  }
  if(issueType == "labels"){
	replyWithAlll(w, &DBl)
  }
}
