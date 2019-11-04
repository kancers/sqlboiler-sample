// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Migrations", testMigrations)
	t.Run("UserProfiles", testUserProfiles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Migrations", testMigrationsDelete)
	t.Run("UserProfiles", testUserProfilesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Migrations", testMigrationsQueryDeleteAll)
	t.Run("UserProfiles", testUserProfilesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Migrations", testMigrationsSliceDeleteAll)
	t.Run("UserProfiles", testUserProfilesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Migrations", testMigrationsExists)
	t.Run("UserProfiles", testUserProfilesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Migrations", testMigrationsFind)
	t.Run("UserProfiles", testUserProfilesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Migrations", testMigrationsBind)
	t.Run("UserProfiles", testUserProfilesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Migrations", testMigrationsOne)
	t.Run("UserProfiles", testUserProfilesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Migrations", testMigrationsAll)
	t.Run("UserProfiles", testUserProfilesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Migrations", testMigrationsCount)
	t.Run("UserProfiles", testUserProfilesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Migrations", testMigrationsHooks)
	t.Run("UserProfiles", testUserProfilesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Migrations", testMigrationsInsert)
	t.Run("Migrations", testMigrationsInsertWhitelist)
	t.Run("UserProfiles", testUserProfilesInsert)
	t.Run("UserProfiles", testUserProfilesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserProfileToUserUsingUser", testUserProfileToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("UserToUserProfileUsingUserProfile", testUserOneToOneUserProfileUsingUserProfile)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserProfileToUserUsingUserProfile", testUserProfileToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("UserToUserProfileUsingUserProfile", testUserOneToOneSetOpUserProfileUsingUserProfile)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Migrations", testMigrationsReload)
	t.Run("UserProfiles", testUserProfilesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Migrations", testMigrationsReloadAll)
	t.Run("UserProfiles", testUserProfilesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Migrations", testMigrationsSelect)
	t.Run("UserProfiles", testUserProfilesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Migrations", testMigrationsUpdate)
	t.Run("UserProfiles", testUserProfilesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Migrations", testMigrationsSliceUpdateAll)
	t.Run("UserProfiles", testUserProfilesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}