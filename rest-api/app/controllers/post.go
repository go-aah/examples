// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"aahframework.org/aah.v0"
	"aahframework.org/log.v0"
	"aahframework.org/examples/rest-api/app/models"
)

// PostController implements the post APi endpoints.
type PostController struct {
	*aah.Context
}

// List method returns all the posts available.
func (p *PostController) List() {
	posts := models.AllPosts()
	log.Infof("%v posts found", len(posts))
	p.Reply().JSON(aah.Data{
		"posts": posts,
	})
}

// Create method to create a post via JSON.
func (p *PostController) Create(post *models.Post) {
	log.Infof("Post Info: %+v", *post)
	id := models.CreatePost(post)
	p.Reply().Created().JSON(aah.Data{
		"id": id,
	})
}

// Retrieve method retunrs single post details for given post ID.
func (p *PostController) Retrieve(id int64) {
	log.Infof("Retrieving post, ID: %v", id)
	post, err := models.GetPost(id)
	if err != nil {
		log.Errorf("Post ID %v, %v", id, err)
		p.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	p.Reply().JSON(models.ToPostAlias(post))
}

// Update method updates the post with given content.
func (p *PostController) Update(id int64, post *models.Post) {
	log.Infof("Updating post: %v", id)
	if post.ID == 0 {
		post.ID = id
	}

	if err := models.UpdatePost(post); err != nil {
		log.Errorf("Post ID %v, %v", id, err)
		p.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	p.Reply().NoContent()
}

// Delete method deletes the post of given post ID.
func (p *PostController) Delete(id int64) {
	log.Infof("Deleting post, ID: %v", id)
	if err := models.DeletePost(id); err != nil {
		log.Errorf("Post ID %v, %v", id, err)
		p.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	p.Reply().NoContent()
}
