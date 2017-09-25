// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"net/http"

	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func statusHandler(c *gin.Context) {
	result := util.NewResult()
	data := map[string]interface{}{}
	result.Data = &data

	data["articleCount"] = 1

	c.JSON(http.StatusOK, result)
}
