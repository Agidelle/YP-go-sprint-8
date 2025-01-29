package main

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)
	clientID := 1
	client, err := selectClient(db, clientID)
	require.NoError(t, err)
	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)
	clientID := -1
	client, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, client.ID)
	assert.Empty(t, client.FIO)
	assert.Empty(t, client.Login)
	assert.Empty(t, client.Birthday)
	assert.Empty(t, client.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	client, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, cl, client)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	db, err := sql.Open("sqlite", "demo.db")
	defer db.Close()
	require.NoError(t, err)
	cl.ID, err = insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)
	_, err = selectClient(db, cl.ID)
	require.NoError(t, err)
	err = deleteClient(db, cl.ID)
	require.NoError(t, err)
	_, err = selectClient(db, cl.ID)
	require.Equal(t, sql.ErrNoRows, err)
}
