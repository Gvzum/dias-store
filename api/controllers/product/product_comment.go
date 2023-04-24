package product

import (
	"github.com/Gvzum/dias-store.git/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c Controller) ListComments(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	comments, err := listComments(uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

func (c Controller) CreateComment(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)
	var commentProductSchema CommentProductSchema
	if err := ctx.ShouldBindJSON(&commentProductSchema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}
	commentProductSchema.ProductID = uint(id)
	if _, err := createComment(commentProductSchema, user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment created successfully",
	})
}

func (c Controller) UpdateComment(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)
	var commentProductSchema BaseCommentSchema
	if err := ctx.ShouldBindJSON(&commentProductSchema); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	idStr := ctx.Param("comment_id")
	if commentID, err := strconv.ParseUint(idStr, 10, 0); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	} else {
		_, err := updateComment(uint(commentID), user, commentProductSchema.Message)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

	}
}

func (c Controller) DetailedComment(ctx *gin.Context) {
	idStr := ctx.Param("comment_id")
	if commentID, err := strconv.ParseUint(idStr, 10, 0); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	} else {
		comment, err := detailedComment(uint(commentID))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"comment": comment,
		})
	}
}

func (c Controller) DeleteComment(ctx *gin.Context) {
	user, _ := ctx.Value("user").(*models.User)

	idStr := ctx.Param("comment_id")
	if commentID, err := strconv.ParseUint(idStr, 10, 0); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	} else {
		err := deleteComment(uint(commentID), user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "comment deleted successfully",
		})
	}
}
