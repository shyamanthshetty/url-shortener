package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}
func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertAndRetrieve(t *testing.T) {
	initialLink := "www.abcd.xpense.com/abcd/dwed"
	shortLink := "shortLink"
	userId := "e0dba740-fc4b-4977-872c-d360239e6b1a"

	SaveUrlMapping(shortLink, initialLink, userId)
	retrievedUrl := GetUrlMapping(shortLink)
	assert.Equal(t, initialLink, retrievedUrl)
}
