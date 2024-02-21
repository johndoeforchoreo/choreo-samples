/*
 * Copyright (c) 2023, WSO2 LLC. (https://www.wso2.com/) All Rights Reserved.
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package routes

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wso2/choreo-sample-apps/go/rest-api/internal/models"
	// "net/http"
	// "github.com/wso2/choreo-sample-apps/go/rest-api/internal/utils"
)

func registerAuthorRoutes(router fiber.Router) {
	r := router.Group("/authors")
	r.Get("/", ListAuthors)
	r.Post("/", AddAuthor)
	r.Get("/:id", GetAuthor)
	r.Put("/:id", UpdateAuthor)
}

// AddAuthor
//
//	@Summary	Add a new Auth
//	@Tags		authors
//	@Accept		json
//	@Produce	json
//	@Param		request	body	models.Author	true	"New Author details"
//	@Router		/authors [post]
//	@Success	201	{object}	models.Author		"successful operation"
//	@Failure	400	{object}	utils.ErrorResponse	"invalid Author details"
//	@Failure	409	{object}	utils.ErrorResponse	"author already exists"
func AddAuthor(c *fiber.Ctx) error {	
	author := models.Author{}
	return c.Status(fiber.StatusCreated).JSON(author)
}

// GetAuthor
//
//	@Summary	Get Author by id
//
//	@Tags		authors
//
//	@Produce	json
//	@Param		id	path	string	true	"Book ID"
//	@Router		/authors/{id} [get]
//	@Success	200	{object}	models.Author		"successful operation"
//	@Failure	404	{object}	utils.ErrorResponse	"Author not found"
func GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	author := models.Author{}	
	author.Id = id
	return c.Status(fiber.StatusOK).JSON(author)
}

// UpdateAuthor
//
//	@Summary	Update a reading list book by id
//	@Tags		authors
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string			true	"Author ID"
//	@Param		request	body	models.Author	true	"Updated author details"
//	@Router		/authors/{id} [put]
//	@Success	200	{object}	models.Author		"successful operation"
//	@Failure	400	{object}	utils.ErrorResponse	"invalid book details"
//	@Failure	404	{object}	utils.ErrorResponse	"book not found"
func UpdateAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	author := models.Author{}	
	author.Id = id
	return c.Status(fiber.StatusOK).JSON(author)
}

// ListAuthors
//
//	@Summary	Get all Authors associated with books
//	@Tags		authors
//	@Produce	json
//	@Router		/authors [get]
//	@Success	200	{array}	models.Author	"successful operation"
func ListAuthors(c *fiber.Ctx) error {
	author := models.Author{}	
	return c.Status(fiber.StatusOK).JSON(author)
}
