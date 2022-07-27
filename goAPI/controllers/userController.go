package controllers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goAPI/models"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

var collection *mongo.Collection
var ctx = context.TODO()

func InitCollection(mongoClient *mongo.Client, dbName, colName string) {
	collection = mongoClient.Database(dbName).Collection(colName)
}

// CreateUser : - check if it's an array of users or a single user
//  			- inserts a single user in a database with a hashed password
//			    - creates a unique file storing the user's data field

func CreateUser(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)
	check(err)
	var users []models.User

	if !isArray(jsonData) {
		var user models.User
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			return
		}
		users = append(users, user)
	} else {
		err := json.Unmarshal(jsonData, &users)
		if err != nil {
			return
		}
	}

	creationProcess(c, users)
}

// GetById : - returns the ID and the name of a user.
//			 - check login method to get all the information

func GetById(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, reduceUser(user))
}

//GetUsers : - Does exactly the same as GetById method for all users

func GetUsers(c *gin.Context) {
	filter := bson.D{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		users = append(users, user)
	}
	if err := cursor.Close(ctx); err != nil {
		log.Fatal(err)
	}
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, "Database is empty")
		return
	}

	c.JSON(http.StatusOK, reduceUsers(users))

}

// DeleteUser : - Deletes a user from the database
//   - Deletes the file associated
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	filter := bson.D{{"_id", id}}
	deleteRes, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Could not delete the user")
		return
	}
	deleteUserFile(id)
	c.JSON(http.StatusOK, gin.H{"Deleted user :": deleteRes.DeletedCount})
}

// DeleteAll : - Deletes all the users from the database
//
//	: - Deletes all the files associated
func DeleteAll(c *gin.Context) {
	filter := bson.D{}
	deleteRes, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Could not delete all users")
		return
	}
	deleteAllDataFiles()
	c.JSON(http.StatusOK, gin.H{"Deleted users :": deleteRes.DeletedCount})
}

// LoginUser : - Gives all the information about a user
//
//	: - Only works if the passwords match
func LoginUser(c *gin.Context) {

	var userLogin models.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, "Could not parse JSON")
		return
	}
	var user models.User

	filter := bson.D{{"_id", userLogin.ID}}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Id does not exist": userLogin.ID})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Passwords do not match for the id :": userLogin.ID})
		return
	}
	c.JSON(http.StatusOK, user)
}

//	UpdateUser : - Update the fields of a user
//			     - Change the datafile only if user's data field has changed

func UpdateUser(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, "Could not parse JSON")
		return
	}
	updatedUser.Password = getHashedPassword(updatedUser)

	var user models.User
	filter := bson.D{{"_id", updatedUser.ID}}
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	if updatedUser.Data != user.Data {
		updateUserFile(updatedUser)
	}
	res, err := updateUser(updatedUser)
	if err != nil {
		panic(err)
	}
	if res == 0 {
		c.JSON(http.StatusBadRequest, "Could not modify the user")
		return
	}
	str := "User " + updatedUser.ID + " updated"
	c.JSON(http.StatusOK, str)
}
