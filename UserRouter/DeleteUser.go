package UserRouter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
)

func (suite *UserTestSuite) TestDeleteSuperAdminUser() {
	createUserDto, _ := CreateUserRequestPayload("SuperAdmin", suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the Delete User API")
	responseOfDeleteUserApi := HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), true, responseOfDeleteUserApi.Result)
	log.Println("Hitting the get user by id after deleting the user")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), 404, responseOfGetUserById.Code)
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", responseOfGetUserById.Errors[0].InternalMessage)

}

func (suite *UserTestSuite) TestDeleteUserWithValidGroupsAndRoleFilters() {
	createUserDto, roleGroupId := CreateUserRequestPayload("GroupsAndRoleFilter", suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the Delete User API")
	responseOfDeleteUserApi := HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), true, responseOfDeleteUserApi.Result)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
	log.Println("Hitting the get user by id after deleting the user")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), 404, responseOfGetUserById.Code)
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", responseOfGetUserById.Errors[0].InternalMessage)
}

func (suite *UserTestSuite) TestDeleteUserWithValidGroupsOnly() {
	createUserDto, roleGroupId := CreateUserRequestPayload(GroupsOnly, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the Delete User API")
	responseOfDeleteUserApi := HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), true, responseOfDeleteUserApi.Result)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
	log.Println("Hitting the get user by id after deleting the user")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), 404, responseOfGetUserById.Code)
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", responseOfGetUserById.Errors[0].InternalMessage)
}

func (suite *UserTestSuite) TestDeleteUserWithValidFiltersOnly() {
	createUserDto, roleGroupId := CreateUserRequestPayload(RoleFilterOnly, suite.authToken)
	byteValueOfStruct, _ := json.Marshal(createUserDto)
	responseOfCreateUserApi := HitCreateUserApi(byteValueOfStruct, suite.authToken)

	log.Println("Hitting the Delete User API")
	responseOfDeleteUserApi := HitDeleteUserApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), true, responseOfDeleteUserApi.Result)
	HitDeleteRoleGroupByIdApi(strconv.Itoa(roleGroupId), suite.authToken)
	log.Println("Hitting the get user by id after deleting the user")
	responseOfGetUserById := HitGetUserByIdApi(strconv.Itoa(responseOfCreateUserApi.Result[0].Id), suite.authToken)
	assert.Equal(suite.T(), 404, responseOfGetUserById.Code)
	assert.Equal(suite.T(), "[{pg: no rows in result set}]", responseOfGetUserById.Errors[0].InternalMessage)
}
