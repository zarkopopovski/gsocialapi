package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type UsersController struct {
	dbConnection *DBConnection
}

func (uController *UsersController) setUserRelations(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")
	userIDTarget := r.FormValue("user_id_target")
	relationType := r.FormValue("relation_type")

	redisDB := uController.dbConnection.cache

	if relationType == "FOLLOW" {
		redisDB.SAdd("usr:"+userIDSource+":following", userIDTarget)
		redisDB.SAdd("usr:"+userIDTarget+":followed", userIDSource)

		redisDB.Incr("usr:" + userIDSource + ":followingNo")
		redisDB.Incr("usr:" + userIDTarget + ":followedNo")
	} else if relationType == "UNFOLLOW" {
		redisDB.SRem("usr:"+userIDSource+":following", userIDTarget)
		redisDB.SRem("usr:"+userIDTarget+":followed", userIDSource)

		redisDB.Decr("usr:" + userIDSource + ":followingNo")
		redisDB.Decr("usr:" + userIDTarget + ":followedNo")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (uController *UsersController) listAllFollowingConnections(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")

	redisDB := uController.dbConnection.cache

	connectionsList := redisDB.SMembers("usr:" + userIDSource + ":following")

	if connectionsList != nil {
		usersArray := uController.listUserConnectionsData(connectionsList.Val(), userIDSource)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(usersArray); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func (uController *UsersController) listAllFollowedConnections(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")

	redisDB := uController.dbConnection.cache

	connectionsList := redisDB.SMembers("usr:" + userIDSource + ":followed")

	if connectionsList != nil {
		usersArray := uController.listUserConnectionsData(connectionsList.Val(), userIDSource)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(usersArray); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func (uController *UsersController) listConnectionsStats(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")

	redisDB := uController.dbConnection.cache

	connectionsStatsList := redisDB.MGet("usr:"+userIDSource+":followingNo", "usr:"+userIDSource+":followedNo")

	if connectionsStatsList != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(connectionsStatsList.Val()); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (uController *UsersController) addToFavorites(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")
	favoriteID := r.FormValue("favorite_id")

	redisDB := uController.dbConnection.cache

	redisDB.SAdd("usr:"+userIDSource+":favorites", favoriteID)
	redisDB.Incr("usr:" + userIDSource + ":favoritesNo")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (uController *UsersController) removeFromFavorites(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")
	favoriteID := r.FormValue("favorite_id")

	redisDB := uController.dbConnection.cache

	redisDB.SRem("usr:"+userIDSource+":favorites", favoriteID)
	redisDB.Decr("usr:" + userIDSource + ":favoritesNo")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (uController *UsersController) listFavorites(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userIDSource := r.FormValue("user_id_source")

	redisDB := uController.dbConnection.cache

	favoritesList := redisDB.SMembers("usr:" + userIDSource + ":favorites")

	if favoritesList != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(favoritesList); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (uController *UsersController) listUserConnectionsData(userList []string, user_id string) []*User {

	joinedString := ""

	for i := 0; i < len(userList); i++ {
		joinedString += "\"" + userList[i] + "\","
	}

	joinedString = joinedString[0 : len(joinedString)-1]

	fmt.Print(joinedString)

	query := "SELECT cp.credentials_id AS id,cp.full_name AS fullname FROM customer_profile cp WHERE cp.credentials_id IN (" + joinedString + ")"

	rows, err := uController.dbConnection.db.Query(query)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {

		newUser := new(User)

		err := rows.Scan(&newUser.Id, &newUser.FullName)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		users = append(users, newUser)

	}

	return users

}
