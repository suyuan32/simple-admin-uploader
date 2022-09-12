// Code generated by goctl. DO NOT EDIT.
package types

// The response data when upload finished | 上传完成返回的数据
// swagger:response UploadResp
type UploadResp struct {
	// Message | 信息
	Msg string `json:"msg"`
	// File name | 文件名称
	Name string `json:"name"`
	// File path | 文件路径
	Path string `json:"path"`
}

// Update file information params | 更新文件信息参数
// swagger:model UpdateFileReq
type UpdateFileReq struct {
	// ID
	// Required : true
	ID int64 `json:"id"`
	// File name | 文件名
	// Required : true
	Name string `json:"name"`
}

// Get file list params | 获取文件列表参数
// swagger:model FileListReq
type FileListReq struct {
	PageInfo
	// File type | 文件类型
	// Required : true
	FileType string `json:"fileType"`
	// File name | 文件名
	// Required : true
	FileName string `json:"fileName"`
	// Create date period | 创建日期时间段
	// Required : true
	Period []string `json:"period"`
}

// The response data of file information | 文件信息数据
// swagger:response FileInfo
type FileInfo struct {
	// ID
	ID int64 `json:"id"`
	// UUID
	UUID string `json:"UUID"`
	// User's UUID | 用户的UUID
	UserUUID string `json:"userUUID"`
	// File name | 文件名
	Name string `json:"name"`
	// File type | 文件类型
	FileType string `json:"fileType"`
	// File size | 文件大小
	Size int64 `json:"size"`
	// File path | 文件路径
	Path string `json:"path"`
	// File public staus | 文件公开状态
	// false private true public | false 私人, true公开
	Status bool `json:"status"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt"`
}

// The response data of file information list | 文件信息列表数据
// swagger:response FileListResp
type FileListResp struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// The file list data | 文件信息列表数据
	// in: body
	Data []FileInfo `json:"data"`
}

// The request params of setting file status | 设置文件状态参数
// swagger:model ChangeStatusReq
type ChangeStatusReq struct {
	// ID
	// Required : true
	ID uint64 `json:"id"`
	// File public staus | 文件公开状态
	// false private true public | false 私人, true公开
	// Required : true
	Status bool `json:"status"`
}

// The request params of download file | 下载文件参数
// swagger:parameters DownloadReq
type DownloadReq struct {
	// ID
	// Required : true
	Id int64 `path:"id"`
}

// The basic response with data | 基础带数据信息
// swagger:response BaseMsg
type BaseMsg struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// The basic response without data | 基础不带数据信息
// swagger:response BaseMsg
type BaseResp struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

// The simplest message | 最简单的信息
// swagger:response SimpleMsg
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// swagger:model PageInfo
// The page request parameters | 列表请求参数
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page"`
	// Page size | 单页数据行数
	// Required: true
	PageSize uint64 `json:"pageSize"`
}

// The page response data model | 列表返回信息
// swagger:response PageList
type PageList struct {
	// Total number | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data []string `json:"data"`
}

// Basic id request | 基础id参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	ID uint `json:"id"`
}

// Basic id request | 基础id参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	ID uint `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// UUID
	// Required: true
	UUID string `json:"UUID" path:"UUID"`
}

// The base response data | 基础信息
// swagger:response BaseInfo
type BaseInfo struct {
	// ID
	ID uint `json:"id"`
	// Create date | 创建日期
	CreatedAt int64 `json:"createdAt"`
	// Update date | 更新日期
	UpdatedAt int64 `json:"updatedAt"`
	// Delete date | 删除日期
	DeletedAt int64 `json:"deletedAt"`
}
