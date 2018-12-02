// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"fmt"

	"aahframe.work"
	"aahframework.org/examples/rest-api/app/models"
)

// PostController implements the post APi endpoints.
type PostController struct {
	*aah.Context
}

// List method returns all the posts available.
func (c *PostController) List() {
	posts := models.AllPosts()
	c.Log().Infof("%v posts found", len(posts))
	c.Reply().JSON(aah.Data{
		"posts": posts,
	})
}

// Create method to create a post via JSON.
func (c *PostController) Create(post *models.Post) {
	c.Log().Infof("Post Info: %+v", *post)
	id := models.CreatePost(post)
	newResourceURL := fmt.Sprintf("%s:%v", c.Req.Scheme, c.RouteURL("retrieve_post", id))
	c.Reply().Created().
		Header("Location", newResourceURL).
		JSON(aah.Data{
			"id": id,
		})
}

// Retrieve method retunrs single post details for given post ID.
func (c *PostController) Retrieve(id int64) {
	c.Log().Infof("Retrieving post, ID: %v", id)
	post, err := models.GetPost(id)
	if err != nil {
		c.Log().Errorf("Post ID %v, %v", id, err)
		c.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	c.Reply().JSON(models.ToPostAlias(post))
}

// Update method updates the post with given content.
func (c *PostController) Update(id int64, post *models.Post) {
	c.Log().Infof("Updating post: %v", id)
	if post.ID == 0 {
		post.ID = id
	}

	if err := models.UpdatePost(post); err != nil {
		c.Log().Errorf("Post ID %v, %v", id, err)
		c.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	c.Reply().NoContent()
}

// Delete method deletes the post of given post ID.
func (c *PostController) Delete(id int64) {
	c.Log().Infof("Deleting post, ID: %v", id)
	if err := models.DeletePost(id); err != nil {
		c.Log().Errorf("Post ID %v, %v", id, err)
		c.Reply().NotFound().JSON(aah.Data{
			"error": err.Error(),
		})
		return
	}

	c.Reply().NoContent()
}
