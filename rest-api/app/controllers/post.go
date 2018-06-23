// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"fmt"

	"aahframework.org/aah.v0"
	"aahframework.org/examples/rest-api/app/models"
)

// PostController implements the post APi endpoints.
type PostController struct {
	*aah.Context
}

// List method returns all the posts available.
func (p *PostController) List() {
	posts := models.AllPosts()
	p.Log().Infof("%v posts found", len(posts))
	p.Reply().JSON(aah.Data{
		"posts": posts,
	})
}

// Create method to create a post via JSON.
func (p *PostController) Create(post *models.Post) {
	p.Log().Infof("Post Info: %+v", *post)
	id := models.CreatePost(post)
	newResourceURL := fmt.Sprintf("%s:%v", p.Req.Scheme, p.RouteURL("retrieve_post", id))
	p.Reply().Created().
		Header("Location", newResourceURL).
		JSON(aah.Data{
			"id": id,
		})
}

// Retrieve method retunrs single post details for given post ID.
func (p *PostController) Retrieve(id int64) {
	p.Log().Infof("Retrieving post, ID: %v", id)
	post, err := models.GetPost(id)
	if err != nil {
		p.Log().Errorf("Post ID %v, %v", id, err)
		p.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	p.Reply().JSON(models.ToPostAlias(post))
}

// Update method updates the post with given content.
func (p *PostController) Update(id int64, post *models.Post) {
	p.Log().Infof("Updating post: %v", id)
	if post.ID == 0 {
		post.ID = id
	}

	if err := models.UpdatePost(post); err != nil {
		p.Log().Errorf("Post ID %v, %v", id, err)
		p.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	p.Reply().NoContent()
}

// Delete method deletes the post of given post ID.
func (p *PostController) Delete(id int64) {
	p.Log().Infof("Deleting post, ID: %v", id)
	if err := models.DeletePost(id); err != nil {
		p.Log().Errorf("Post ID %v, %v", id, err)
		p.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	p.Reply().NoContent()
}
