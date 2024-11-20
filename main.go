package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/version/acl/:aclId", GetAclById)
	router.GET("/api/version/acls/user/:userId", GetAclsByUserId)
	router.GET("/api/version/acls/guestbook/:gbId", GetAclsByGuestbookId)

	router.POST("/api/version/acls", PostAcl)
	router.PUT("/api/version/acls", PutAcl)
	router.DELETE("/api/version/acl/:aclId", DeleteAcl)

	router.Run(":8080")
}
