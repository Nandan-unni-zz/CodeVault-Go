package handlers

import (
	"net/http"

	"../utils"
)

// FeedHandler : Handler
func FeedHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		print("GET: /\n")
		utils.RenderTemplate(res, "feed.html", nil)
	}
}

// VoteHandler : Handler
func VoteHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		print("GET: /vote/\n")
		utils.RenderTemplate(res, "vote.html", nil)
	}
}

// ResultHandler : Handler
func ResultHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		print("GET: /result/\n")
		utils.RenderTemplate(res, "result.html", nil)
	}
}
