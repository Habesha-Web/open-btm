
package common

import (
	"context"
	"math"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"details"`
}

type ResponsePagination struct {
	Success bool        `json:"success"`
	Items   interface{} `json:"data"`
	Message string      `json:"details"`
	Total   uint        `json:"total"`
	Page    uint        `json:"page"`
	Size    uint        `json:"size"`
	Pages   uint        `json:"pages"`
}
func Pagination(db *gorm.DB, queryModel interface{}, responseObjectModel interface{}, page uint, size uint, tracer context.Context) (ResponsePagination, interface{}, error) {
	//  protection against requesting large amount of data
	//  set to 50
	var update_size uint
	if size > 50 {
		size = 50
	}

	count_channel := make(chan int64)
	str_chann := make(chan string)
	defer func() {
		close(str_chann)
		close(count_channel)
	}()
	var offset int64 = int64(page-1) * int64(update_size)
	//finding count value
	go func(comm <-chan int64) {
		var local_counter int64
		if tracer != nil {
			db.WithContext(tracer).Select("*").Model(&queryModel).Count(&local_counter)
		} else {
			db.WithContext(tracer).Select("*").Model(&queryModel).Count(&local_counter)
		}
		count_channel <- local_counter

	}(count_channel)
	//  set offset value for page One
	var response_page int64
	go func(comm <-chan string) {
		if page == 1 {
			if tracer != nil {
				db.WithContext(tracer).Order("id asc").Limit(int(size)).Offset(0).Preload(clause.Associations).Find(&responseObjectModel)

			} else {
				db.WithContext(tracer).Order("id asc").Limit(int(size)).Offset(0).Preload(clause.Associations).Find(&responseObjectModel)

			}

			response_page = 1
		} else {
			if tracer != nil {
				db.WithContext(tracer).WithContext(tracer).Order("id asc").Limit(int(size)).Offset(int(offset)).Preload(clause.Associations).Find(&responseObjectModel)

			} else {
				db.WithContext(tracer).Order("id asc").Limit(int(size)).Offset(int(offset)).Preload(clause.Associations).Find(&responseObjectModel)
			}
			// response_channel <- loc_resp
			response_page = int64(page)
		}
		str_chann <- "completed"
	}(str_chann)
	count := <-count_channel
	response_obj := <-str_chann
	pages := math.Ceil(float64(count) / float64(size))

	result := ResponsePagination{
		Success: true,
		Items:   responseObjectModel,
		Message: response_obj,
		Total:   uint(count),
		Page:    uint(response_page),
		Size:    uint(size),
		Pages:   uint(pages),
	}
	return result, responseObjectModel, nil
}

func PaginationPureModel(db *gorm.DB, queryModel interface{}, responseObjectModel interface{}, page uint, size uint, tracer context.Context) (ResponsePagination, interface{}, error) {
	if size > 50 {
		size = 50
	}
	count_channel := make(chan int64)
	str_chann := make(chan string)
	defer func() {
		close(str_chann)
		close(count_channel)
	}()

	var offset int64 = int64(page-1) * int64(size)
	//finding count value
	go func(comm <-chan int64) {
		var local_counter int64
		if tracer != nil {
			db.WithContext(tracer).Select("*").Model(&queryModel).Count(&local_counter)
		} else {
			db.WithContext(tracer).Select("*").Model(&queryModel).Count(&local_counter)
		}
		count_channel <- local_counter

	}(count_channel)
	//  set offset value for page One
	var response_page int64
	go func(comm <-chan string) {
		if page == 1 {
			if tracer != nil {
				db.WithContext(tracer).Model(&queryModel).Order("id asc").Limit(int(size)).Offset(0).Find(&responseObjectModel)
			} else {
				db.WithContext(tracer).Model(&queryModel).Order("id asc").Limit(int(size)).Offset(0).Find(&responseObjectModel)
			}
			response_page = 1
		} else {
			if tracer != nil {
				db.WithContext(tracer).Model(&queryModel).Order("id asc").Limit(int(size)).Offset(int(offset)).Find(&responseObjectModel)
			} else {
				db.WithContext(tracer).Model(&queryModel).Order("id asc").Limit(int(size)).Offset(int(offset)).Find(&responseObjectModel)
			}
			// response_channel <- loc_resp
			response_page = int64(page)
		}
		str_chann <- "completed"
	}(str_chann)

	count := <-count_channel
	response_obj := <-str_chann
	pages := math.Ceil(float64(count) / float64(size))
	result := ResponsePagination{
		Success: true,
		Items:   responseObjectModel,
		Message: response_obj,
		Total:   uint(count),
		Page:    uint(response_page),
		Size:    uint(size),
		Pages:   uint(pages),
	}
	return result, responseObjectModel, nil
}

func PaginationPureModelFilterOneToMany(db *gorm.DB, queryModel interface{}, responseObjectModel interface{}, otm_string string, parentID uint,
	page uint, size uint, tracer context.Context) (ResponsePagination, interface{}, error) {
	if size > 50 {
		size = 50
	}
	count_channel := make(chan int64)
	str_chann := make(chan string)
	defer func() {
		close(str_chann)
		close(count_channel)
	}()

	var offset int64 = int64(page-1) * int64(size)
	//finding count value
	go func(comm <-chan int64) {
		var local_counter int64
		if tracer != nil {
			db.WithContext(tracer).Select("*").Model(&queryModel).Where(otm_string, parentID).Count(&local_counter)
		} else {
			db.WithContext(tracer).Select("*").Model(&queryModel).Where(otm_string, parentID).Count(&local_counter)
		}
		count_channel <- local_counter

	}(count_channel)
	//  set offset value for page One
	var response_page int64
	go func(comm <-chan string) {
		if page == 1 {
			if tracer != nil {
				db.WithContext(tracer).Model(&queryModel).Where(otm_string, parentID).Order("id asc").Limit(int(size)).Offset(0).Find(&responseObjectModel)
			} else {
				db.WithContext(tracer).Model(&queryModel).Where(otm_string, parentID).Order("id asc").Limit(int(size)).Offset(0).Find(&responseObjectModel)
			}
			response_page = 1
		} else {
			if tracer != nil {
				db.WithContext(tracer).Model(&queryModel).Where(otm_string, parentID).Order("id asc").Limit(int(size)).Offset(int(offset)).Find(&responseObjectModel)
			} else {
				db.WithContext(tracer).Model(&queryModel).Where(otm_string, parentID).Order("id asc").Limit(int(size)).Offset(int(offset)).Find(&responseObjectModel)
			}
			// response_channel <- loc_resp
			response_page = int64(page)
		}
		str_chann <- "completed"
	}(str_chann)

	count := <-count_channel
	response_obj := <-str_chann
	pages := math.Ceil(float64(count) / float64(size))
	result := ResponsePagination{
		Success: true,
		Items:   responseObjectModel,
		Message: response_obj,
		Total:   uint(count),
		Page:    uint(response_page),
		Size:    uint(size),
		Pages:   uint(pages),
	}
	return result, responseObjectModel, nil
}

