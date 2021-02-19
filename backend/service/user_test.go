// +build unit

package service

import (
	"errors"
	"testing"

	"github.com/jimlawless/whereami"
	"github.com/nouney/randomstring"

	"github.com/zozoee27/cookbook/backend/entity"
	"github.com/zozoee27/cookbook/backend/interface/user"
	"github.com/zozoee27/cookbook/backend/testutil"
)

func TestCreateUser(t *testing.T) {
	userDb := user.NewFakeUserDb()
	userDb.ClearAllEntries()

	userService := CreateUserService(userDb)

	randomUser := createRandomUser()
	userService.CreateUser(randomUser)
	testutil.CompareInt(t, userDb.Size(), 1, "Could not insert first user", whereami.WhereAmI())

	randomUser2 := createRandomUser()
	userService.CreateUser(randomUser2)
	testutil.CompareInt(t, userDb.Size(), 2, "Could not insert second user", whereami.WhereAmI())
}

func TestCreateUserWithError(t *testing.T) {
	expectedError := errors.New(randomstring.Generate(10))

	userDb := user.NewFakeUserDbWithError(expectedError)
	userDb.ClearAllEntries()

	userService := CreateUserService(userDb)

	randomUser := createRandomUser()
	actualError := userService.CreateUser(randomUser)

	testutil.CompareError(t, actualError, expectedError, "Returned error is different", whereami.WhereAmI())
}

func TestFindValidUserFromCollection(t *testing.T) {
	userDb := user.NewFakeUserDb()
	userDb.ClearAllEntries()

	userService := CreateUserService(userDb)

	randomUserA := createRandomUser()
	randomUserB := createRandomUser()
	userService.CreateUser(randomUserA)
	userService.CreateUser(randomUserB)

	actualUserA, errA := userService.FindUserFromCollection(randomUserA.Username)
	actualUserB, errB := userService.FindUserFromCollection(randomUserB.Username)

	testutil.CompareUserEntity(t, actualUserA, randomUserA, "Could not find UserA", whereami.WhereAmI())
	testutil.CompareUserEntity(t, actualUserB, randomUserB, "Could not find UserB", whereami.WhereAmI())
	testutil.CompareError(t, errA, nil, "Error from retrieving UserA", whereami.WhereAmI())
	testutil.CompareError(t, errB, nil, "Error from retrieving UserB", whereami.WhereAmI())
}

func TestFindNonExistentUserFromCollection(t *testing.T) {
	userDb := user.NewFakeUserDb()
	userDb.ClearAllEntries()

	userService := CreateUserService(userDb)

	randomUserA := createRandomUser()

	actualUserA, errA := userService.FindUserFromCollection(randomUserA.Username)
	testutil.CompareUserEntity(t, actualUserA, nil, "UserA should not exist", whereami.WhereAmI())
	testutil.CompareError(t, errA, nil, "Error from retrieving user", whereami.WhereAmI())
}

func TestFindUserFromCollectionWithError(t *testing.T) {
	expectedError := errors.New(randomstring.Generate(10))
	userDb := user.NewFakeUserDbWithError(expectedError)
	userDb.ClearAllEntries()

	userService := CreateUserService(userDb)

	randomUserA := createRandomUser()
	randomUserB := createRandomUser()
	userService.CreateUser(randomUserA)

	resultA, errA := userService.FindUserFromCollection(randomUserA.Username)
	resultB, errB := userService.FindUserFromCollection(randomUserB.Username)

	testutil.CompareUserEntity(t, resultA, randomUserA, "Returned user is wrong", whereami.WhereAmI())
	testutil.CompareUserEntity(t, resultB, nil, "UserB should not exist", whereami.WhereAmI())

	testutil.CompareError(t, errA, expectedError, "Returned error is different", whereami.WhereAmI())
	testutil.CompareError(t, errB, expectedError, "Returned error is different", whereami.WhereAmI())
}

func TestClearAllEntries(t *testing.T) {
	userDb := user.NewFakeUserDb()
	userDb.ClearAllEntries()

	userService := CreateUserService(userDb)

	userService.CreateUser(createRandomUser())
	userService.CreateUser(createRandomUser())
	testutil.CompareInt(t, userDb.Size(), 2, "DB Size should be 2", whereami.WhereAmI())

	userService.ClearAllEntries()
	testutil.CompareInt(t, userDb.Size(), 0, "DB Should be empty", whereami.WhereAmI())

	userService.CreateUser(createRandomUser())
	testutil.CompareInt(t, userDb.Size(), 1, "DB Size should be 1", whereami.WhereAmI())
}

func TestClearAllEntriesWithError(t *testing.T) {
	expectedError := errors.New(randomstring.Generate(10))
	userDb := user.NewFakeUserDbWithError(expectedError)

	userService := CreateUserService(userDb)
	userService.CreateUser(createRandomUser())

	actualError := userService.ClearAllEntries()
	testutil.CompareError(t, actualError, expectedError, "Errors are different", whereami.WhereAmI())
}

func createRandomUser() *entity.User {
	randomCharacterLength := 6
	return &entity.User{
		Username:  randomstring.Generate(randomCharacterLength),
		Email:     randomstring.Generate(randomCharacterLength),
		FirstName: randomstring.Generate(randomCharacterLength),
		LastName:  randomstring.Generate(randomCharacterLength),
		Password:  randomstring.Generate(randomCharacterLength),
	}
}
