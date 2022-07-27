package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"goAPI/configs"
	"goAPI/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"sync"
)

// creationProcess : - Uses concurrency to add multiple users in the database
//   			     - Creates a unique ID file for every user
//   			     - Shows to the client how many users were added to the database

func creationProcess(c *gin.Context, users []models.User) {
	var mu sync.Mutex
	mu.Lock()
	var wg sync.WaitGroup
	wg.Add(len(users))

	var added = len(users)

	for _, user := range users {
		u := user
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			u.Password = getHashedPassword(u)
			createFile(u)
			_, err := collection.InsertOne(ctx, &u)
			if err != nil {
				added--
				c.JSON(http.StatusInternalServerError,
					gin.H{
						"User already in the database ": u.ID,
						"name":                          u.Name,
					})
				return
			}
			c.JSON(http.StatusOK, gin.H{u.Name: "has been inserted to the database"})
		}()
	}
	mu.Unlock()
	wg.Wait()
	c.JSON(http.StatusOK, gin.H{"Number of users added :": added})

}

//updateUser : - update every single field of a user

func updateUser(updatedUser models.User) (int64, error) {
	id := updatedUser.ID
	fmt.Println(id)
	update := bson.M{
		"password":   updatedUser.Password,
		"name":       updatedUser.Name,
		"isactive":   updatedUser.Isactive,
		"balance":    updatedUser.Balance,
		"age":        updatedUser.Age,
		"company":    updatedUser.Company,
		"email":      updatedUser.Email,
		"phone":      updatedUser.Phone,
		"address":    updatedUser.Address,
		"about":      updatedUser.About,
		"registered": updatedUser.Registered,
		"latitude":   updatedUser.Latitude,
		"longitude":  updatedUser.Longitude,
		"tags":       updatedUser.Tags,
		"friends":    updatedUser.Friends,
		"data":       updatedUser.Data,
	}
	updated, err := collection.UpdateOne(ctx, bson.D{{"_id", id}}, bson.M{"$set": update})
	fmt.Println(updated == nil)
	return updated.MatchedCount, err

}

func createFile(user models.User) {
	id := user.ID
	fileName := configs.UsersData + "/" + id
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)

	_, err = file.WriteString(user.Data)
	check(err)
}

func updateUserFile(updatedUser models.User) {
	fileName := configs.UsersData + "/" + updatedUser.ID
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)
	_, err = file.WriteString(updatedUser.Data)
	check(err)
}

func deleteUserFile(id string) {
	err := os.Remove(configs.UsersData + "/" + id)
	check(err)
}

// deleteAllDataFiles : - Delete every single data file concurrently using goroutines
func deleteAllDataFiles() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	mutex.Lock()
	files, err := os.ReadDir(configs.UsersData)
	check(err)
	wg.Add(len(files))

	for _, file := range files {
		file := file
		go func() {
			mutex.Lock()
			err := os.Remove(configs.UsersData + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			wg.Done()
		}()
	}

	mutex.Unlock()
	wg.Wait()
}

//	reduceUser : - converts a user into a reduced user to get the minimal information about her/him and returns it

func reduceUser(user models.User) models.ReducedUser {
	return models.ReducedUser{
		ID:   user.ID,
		Name: user.Name,
	}
}

//	reduceUsers : - returns an array of reduced users

func reduceUsers(users []models.User) []models.ReducedUser {
	var reducedUsers []models.ReducedUser
	for _, user := range users {
		reducedUsers = append(reducedUsers, reduceUser(user))
	}
	return reducedUsers
}

// getHashedPassword : Converts user's password into a hashed one
func getHashedPassword(user models.User) string {
	hashedPsd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	check(err)
	return string(hashedPsd)
}

// isArray : checks if the first byte of the json data is a '[', in that case we will know that it's an array
func isArray(jsonData []byte) bool {
	return bytes.HasPrefix(bytes.TrimSpace(jsonData), []byte{'['})
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
