package controllers

import (
    "github.com/revel/revel"
    "github.com/janekolszak/revmgo"
    "github.com/ip4368/colonnade/app/models"
)

func init() {
    revmgo.ControllerInit()
}

type Courses struct {
    *revel.Controller
    revmgo.MongoController
}

func (c Courses) CoursesForUser() revel.Result {
    loginStat := models.LoginStatus(
        models.User_t{
            Email: c.Session["email"],
            Username: c.Session["username"],
            Name: c.Session["name"],
            UserIdHex: c.Session["userId"],
        })
    var result int = 0
    var coordinator, tutor, student []models.Course_db
    if loginStat == 0 {
        result, coordinator, tutor, student = models.CoursesForUser(c.MongoSession, c.Session["userId"])
    } else { result = 1 }

    // start with initialise response interface
    data := make(map[string]interface{})
    data["error"] = result
    switch result {
        case 0 :
            data["message"] = "Sucess"
            data["data"] = make(map[string]interface{})
            data["data"].(map[string]interface{})["asCoordinator"] = coordinator
            data["data"].(map[string]interface{})["asTutor"] = tutor
            data["data"].(map[string]interface{})["asStudent"] = student
        case 1 :
            data["message"] = "User has not logged in"
        case 2 :
            data["message"] = "Invalid User ID"
        case 3 :
            data["message"] = "Unexpected Error in Database"
    }
    return c.RenderJson(data)
}

func (c Courses) Course(Id string) revel.Result {
    loginStat := models.LoginStatus(
        models.User_t{
            Email: c.Session["email"],
            Username: c.Session["username"],
            Name: c.Session["name"],
            UserIdHex: c.Session["userId"],
        })
    var result int = 0
    var courseData models.Course_db
    if loginStat == 0 {
        result, courseData = models.UserCourse(c.MongoSession, c.Session["userId"], Id)
    } else { result = 1 }

    // start with initialise response interface
    data := make(map[string]interface{})
    data["error"] = result
    switch result {
        case 0 :
            data["message"] = "Sucess"
            data["data"]    = make(map[string]interface{})
            data["data"].(map[string]interface{})["course"] = courseData
        case 1 :
            data["message"] = "User has not logged in"
        case 2 :
            data["message"] = "Invalid User ID"
        case 3 :
            data["message"] = "Invalid Course ID"
        case 4 :
            data["message"] = "Unexpected Error in Database"
    }
    return c.RenderJson(data)
}
