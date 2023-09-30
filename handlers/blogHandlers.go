package handlers

import (
	"errors"
	"sagara-try/config"
	"sagara-try/helpers"
	"sagara-try/models"
)

type Blog models.Blog

func (h *Blog) H_SaveBlog(userId int) (H, error) {
	datum := Blog{}
	datum.UserID = userId
	datum.Title = h.Title
	datum.Type = h.Type
	datum.Content = h.Content
	datum.Status = h.Status
	datum.IsEvent = h.IsEvent

	err := config.GetDB().Debug().Create(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [BlogHandler.Saved] - failed insert: "+err.Error())
		return nil, err
	}

	msg := H{}
	msg["blog_id"] = datum.BlogID
	msg["user_id"] = datum.UserID
	msg["title"] = datum.Title
	msg["type"] = datum.Type
	msg["content"] = datum.Content
	msg["status"] = datum.Status
	msg["is_event"] = datum.IsEvent

	return msg, nil
}

func H_GetExistingBlog(userId int) ([]*Blog, error) {
	datum := make([]*Blog, 0)

	var query error
	query = config.GetDB().Debug().Where("user_id = ?", userId).Order("created DESC").Find(&datum).Error
	err := query

	if err != nil {
		helpers.Logger("error", "In Server: [Handlers.Blog] - failed view all data "+err.Error())
		return nil, err
	}

	return datum, nil
}

func H_GetBlogByUser(userId int) ([]*Blog, error) {
	datum := make([]*Blog, 0)

	var query error
	query = config.GetDB().Debug().Where("user_id = ?", userId).Order("created DESC").Find(&datum).Error
	err := query

	if err != nil {
		helpers.Logger("error", "In Server: [BlogHandler.GetBlogByUser] - failed view all data "+err.Error())
		return nil, err
	}

	return datum, nil
}

func H_GetPaginatedBlog(userId int, page int, limit int) ([]*Blog, int, error) {
	datum := make([]*Blog, 0)
	var totalRows int

	if err := config.GetDB().Model(&Blog{}).Where("user_id = ?", userId).Count(&totalRows).Error; err != nil {
		helpers.Logger("error", "In Server: [BlogHandler.BlogPaginated] - failed total rows "+err.Error())
		return nil, 0, err
	}

	offset := (page - 1) * limit

	var query = config.GetDB().Debug().Where("user_id = ?", userId).Order("created DESC").Offset(offset).Limit(limit).Find(&datum).Error
	if query != nil {
		helpers.Logger("error", "In Server: [BlogHandler.BlogPage] - failed view page data "+query.Error())
		return nil, 0, query
	}

	return datum, totalRows, nil
}

func H_GetOneBlog(blogId int) (*Blog, error) {
	var datum Blog
	err := config.GetDB().Debug().Where("blog_id = ?", blogId).Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [BlogHandler.GetOneBlog] - id is not exist "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func (h *Blog) H_UpdateOneBlog(blogId int, userId int) (*Blog, error) {
	datum := Blog{}
	datum.BlogID = blogId
	h.UserID = userId
	datum.UserID = h.UserID
	datum.Title = h.Title
	datum.Type = h.Type
	datum.Content = h.Content
	datum.Status = true
	datum.IsEvent = h.IsEvent

	err := config.GetDB().Debug().Model(datum).Where("blog_id = ?", blogId).Update(&h).Error
	if err != nil {
		helpers.Logger("error", "In Server: [BlogHandler.Update] - failed updated data : "+err.Error())
		return nil, err
	}

	return &datum, nil
}

func H_DeleteOneBlog(blogId int) (string, error) {
	rowsAffected := config.GetDB().Debug().Model(Blog{}).Where("blog_id = ?", blogId).Delete(Blog{}).RowsAffected
	if rowsAffected == 0 {
		helpers.Logger("error", "In Server: [BlogHandler.Deleted] - id is not exist")
		return "", errors.New("id is not exist")
	}
	return "success to deleted", nil
}
