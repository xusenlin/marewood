package controller

import (
	"MareWood/service/serviceRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

//const commitFormat string = `--pretty=format:{
//  "commit": "%H",
//  "parent": "%P",
//  "refs": "%D",
//  "subject": "%s",
//  "author": { "name": "%aN", "email": "%aE", "date": "%ad" },
//  "committer": { "name": "%cN", "email": "%cE", "date": "%cd" }
// },`

//type GitPerson struct {
//	Name  string     `json:"name"`
//	Email string     `json:"email"`
//	Date  *time.Time `json:"date"`
//}
//type GitCommit struct {
//	Commit    string    `json:"commit"`
//	Parent    string    `json:"parent"`
//	Refs      string    `json:"refs"`
//	Subject   string    `json:"subject"`
//	Author    GitPerson `json:"author"`
//	Committer GitPerson `json:"committer"`
//}

func CommitRecord(c *gin.Context) {
	id := c.Query("id")

	args := []string{
		"log",
		"-4",
		//"--date=iso-strict",
		"--first-parent",
	}

	logOut , err := serviceRepository.RunCmdOnRepositoryDir(id,"git",args...)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}

	//logOut = logOut[:len(logOut)-1] // Remove the last ","
	//logOut = fmt.Sprintf("[%s]", logOut) // Add []

	//var gitCommitList []GitCommit
	//err = json.Unmarshal([]byte(logOut), &gitCommitList)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": false,
	//		"data":   "",
	//		"msg":    err.Error(),
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   logOut,
		"msg":    "search successful",
	})
}
