// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"errors"
	"sync"
	"time"
)

/*
This is an IN-MEMORY implementation of Post store to demostrate REST
API implementation of aah framework.
*/

var (
	postStore  = make(map[int64]*Post)
	lastPostID int64
	idMx       = &sync.Mutex{}
	postMx     = &sync.RWMutex{}
)

// Post struct holds single post information.
type Post struct {
	ID        int64     `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// PostAlias is used for time format to ISO 8061
type PostAlias Post

// AllPosts returns all the posts.
func AllPosts() []interface{} {
	postMx.RLock()
	defer postMx.RUnlock()
	posts := make([]interface{}, 0)
	for _, post := range postStore {
		posts = append(posts, ToPostAlias(post))
	}
	return posts
}

// CreatePost method creates the new post entry.
func CreatePost(post *Post) int64 {
	postMx.Lock()
	defer postMx.Unlock()
	id := createPostID()
	post.ID = id
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt
	postStore[id] = post
	return id
}

// GetPost method return the post for given ID.
func GetPost(id int64) (*Post, error) {
	postMx.RLock()
	defer postMx.RUnlock()
	if _, found := postStore[id]; !found {
		return nil, errors.New("post not found")
	}

	return postStore[id], nil
}

// UpdatePost method updates the given info with post store.
func UpdatePost(post *Post) error {
	postMx.Lock()
	defer postMx.Unlock()
	if _, found := postStore[post.ID]; !found {
		return errors.New("post not found")
	}

	postStore[post.ID].Title = post.Title
	postStore[post.ID].Body = post.Body
	postStore[post.ID].UpdatedAt = time.Now()
	return nil
}

// DeletePost method deletes the post for given ID.
func DeletePost(id int64) error {
	postMx.Lock()
	defer postMx.Unlock()
	if _, found := postStore[id]; !found {
		return errors.New("post not found")
	}

	delete(postStore, id)
	return nil
}

// ToPostAlias method formats the time to RFC3339 using alias.
func ToPostAlias(post *Post) interface{} {
	return &struct {
		*PostAlias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		PostAlias: (*PostAlias)(post),
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	}
}

func createPostID() int64 {
	idMx.Lock()
	defer idMx.Unlock()
	lastPostID++
	return lastPostID
}
