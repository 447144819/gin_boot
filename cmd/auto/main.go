// cmd/auto/main.go
package main

import (
	"fmt"
	"gin_boot/internal/utils/gfile"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func main() {
	// 1. 检查命令行参数：期待传入模块名，如 "menu"
	if len(os.Args) < 2 {
		fmt.Println("请指定模块名称，例如：go run cmd/auto/main.go menu")
		return
	}
	moduleName := os.Args[1] // 比如 "menu"

	// 2. 确保模块名首字母大写（用于生成结构体名）
	moduleNameUpper := capitalizeFirst(moduleName)

	// 3. 定义要生成的文件列表与对应模板
	files := []struct {
		path     string
		filename string
		tpl      *template.Template
		data     interface{}
	}{
		{
			path:     filepath.Join("internal", "controller"),
			filename: fmt.Sprintf("%s.go", moduleName),
			tpl:      controllerTemplate(),
			data:     map[string]string{"ModuleName": moduleNameUpper, "ModelTitle": moduleName},
		},
		{
			path:     filepath.Join("internal", "dao"),
			filename: fmt.Sprintf("%s.go", moduleName),
			tpl:      daoTemplate(),
			data:     map[string]string{"ModuleName": moduleNameUpper, "ModelTitle": moduleName},
		},
		{
			path:     filepath.Join("internal", "service"),
			filename: fmt.Sprintf("%s.go", moduleName),
			tpl:      serviceTemplate(),
			data:     map[string]string{"ModuleName": moduleNameUpper, "ModelTitle": moduleName},
		},
		{
			path:     filepath.Join("internal", "dto"),
			filename: fmt.Sprintf("%s.go", moduleName),
			tpl:      dtoTemplate(),
			data:     map[string]string{"ModuleName": moduleNameUpper, "ModelTitle": moduleName},
		},
		{
			path:     filepath.Join("internal", "vo"),
			filename: fmt.Sprintf("%s.go", moduleName),
			tpl:      voTemplate(),
			data:     map[string]string{"ModuleName": moduleNameUpper, "ModelTitle": moduleName},
		},
		{
			path:     filepath.Join("internal", "model"),
			filename: fmt.Sprintf("%s.go", moduleName),
			tpl:      modelTemplate(),
			data:     map[string]string{"ModuleName": moduleNameUpper, "ModelTitle": moduleName},
		},
	}

	// 4. 遍历并生成每一个文件
	for _, f := range files {
		// 确保目录存在
		err := os.MkdirAll(f.path, 0755)
		if err != nil {
			fmt.Printf("❌ 创建目录失败 %s: %v\n", f.path, err)
			continue
		}

		// 文件完整路径
		filePath := filepath.Join(f.path, f.filename)

		// 如果文件不存在旧创建
		if f.path == "internal\\model" && gfile.FileExists(filePath) {
			fmt.Println("文件已存在")
			continue
		}

		// 创建文件
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("❌ 创建文件失败 %s: %v\n", filePath, err)
			continue
		}

		// 渲染模板并写入文件
		err = f.tpl.Execute(file, f.data)
		file.Close()
		if err != nil {
			fmt.Printf("❌ 渲染模板失败 %s: %v\n", filePath, err)
			continue
		}

		fmt.Printf("✅ 成功生成文件: %s\n", filePath)
	}

	fmt.Println("🎉 代码生成完成！")

	// 更新wire
	runWire()
}

func runWire() {
	// 获取项目根目录
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取工作目录:", err)
		os.Exit(1)
	}

	// 运行 go run .\cmd\runwire.go
	wireSetCmd := exec.Command("go", "run", filepath.Join("cmd", "runwire.go"))
	wireSetCmd.Stdout = os.Stdout
	wireSetCmd.Stderr = os.Stderr
	wireSetCmd.Dir = rootDir // 设置工作目录为项目根目录
	fmt.Println("生成wire: go run ./cmd/runwire.go")
	if err := wireSetCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "运行失败 runwire.go: %v\n", err)
		os.Exit(1)
	}
}

// 工具函数：首字母大写
func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-'a'+'A') + s[1:]
}

// ========== 以下是各个文件的模板定义 ==========

func controllerTemplate() *template.Template {
	return template.Must(template.New("controller").Parse(`package controller

import (
	"gin_boot/internal/dto"
	"gin_boot/internal/router/routers"
	"gin_boot/internal/service"
	"gin_boot/internal/utils/converter"
	"gin_boot/pkg/response"
	"github.com/gin-gonic/gin"
)

type {{.ModuleName}}Controller struct {
	svc service.{{.ModuleName}}Service
}

func New{{.ModuleName}}Controller(svc service.{{.ModuleName}}Service) *{{.ModuleName}}Controller {
	return &{{.ModuleName}}Controller{
		svc: svc,
	}
}

func (h *{{.ModuleName}}Controller) RegisterRoutes(server *gin.Engine) {
	apiv1 := server.Group(routers.RouterBase.APIV1)
	api := apiv1.Group("/system/{{.ModelTitle}}")
	api.POST("", h.Create)
	api.PUT("", h.Edit)
	api.DELETE("", h.Delete)
	api.GET("/:id", h.Detail)
	api.GET("/page", h.List)
}

func (h *{{.ModuleName}}Controller) Create(ctx *gin.Context) {
	var req dto.{{.ModuleName}}CreateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}
	if err := h.svc.Create(ctx, req); err != nil {
		response.Error(ctx, err, response.AddError)
	}

	response.Success(ctx, response.AddSuccess)
}

func (h *{{.ModuleName}}Controller) Edit(ctx *gin.Context) {
	var req dto.{{.ModuleName}}EditDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}

	if err := h.svc.Edit(ctx, req); err != nil {
		response.Error(ctx, err, response.EditError)
		return
	}
	response.Success(ctx, response.EditSuccess)
}

func (h *{{.ModuleName}}Controller) Delete(ctx *gin.Context) {
	var ids []uint64
	if err := ctx.ShouldBindJSON(&ids); err != nil {
		response.Error(ctx, nil, response.ParamsError)
		return
	}

	ginCtx := ctx.Request.Context() // 转为 context.Context
	if err := h.svc.Delete(ginCtx, ids); err != nil {
		response.Error(ctx, err, response.DeleteError)
		return
	}
	response.Success(ctx, response.DeleteSuccess)
}

func (h *{{.ModuleName}}Controller) Detail(ctx *gin.Context) {
	id, _ := converter.StringToUint64(ctx.Param("id"))
	if id < 1 {
		response.Error(ctx, nil, response.ParamsError)
		return
	}
	info, err := h.svc.Detail(ctx.Request.Context(), id)
	if err != nil {
		response.Error(ctx, err, response.DoError)
		return
	}
	response.SuccessData(ctx, info)
}

func (h *{{.ModuleName}}Controller) List(ctx *gin.Context) {
	var req dto.{{.ModuleName}}ListDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, err, response.ParamsError)
		return
	}

	datas, total, err := h.svc.List(ctx.Request.Context(), req)
	if err != nil {
		response.Error(ctx, err, response.DoError)
		return
	}
	response.PageSuccess(ctx, datas, total, req.Page, req.Limit)
}
`))
}

func daoTemplate() *template.Template {
	return template.Must(template.New("dao").Parse(`package dao

import (
	"gin_boot/internal/dao/basedao"
	"gin_boot/internal/model"
	"gorm.io/gorm"
)

// {{.ModuleName}}Dao
type {{.ModuleName}}Dao struct {
	*basedao.BaseDao[model.{{.ModuleName}}, uint64]
}

// New{{.ModuleName}}Dao 是构造函数，返回接口类型
func New{{.ModuleName}}Dao(db *gorm.DB) *{{.ModuleName}}Dao {
	// 自动创建表
	db.AutoMigrate(&model.{{.ModuleName}}{})
	return &{{.ModuleName}}Dao{
		basedao.NewBaseDao[model.{{.ModuleName}}, uint64](db),
	}
}
`))
}

func serviceTemplate() *template.Template {
	return template.Must(template.New("service").Parse(`package service

import (
	"context"
	"errors"
	"gin_boot/internal/dao"
	"gin_boot/internal/dto"
	"gin_boot/internal/model"
	"gin_boot/internal/utils/times"
	"gin_boot/internal/vo"
	"gin_boot/pkg/response"
)

// {{.ModuleName}}Service 定义服务行为（接口）
type {{.ModuleName}}Service interface {
	Create(ctx context.Context, req dto.{{.ModuleName}}CreateDTO) error
	Edit(ctx context.Context, req dto.{{.ModuleName}}EditDTO) error
	Delete(ctx context.Context, ids []uint64) error
	Detail(ctx context.Context, id uint64) (vo.{{.ModuleName}}InfoVO, error)
	List(ctx context.Context, req dto.{{.ModuleName}}ListDTO) ([]vo.{{.ModuleName}}InfoVO, int64, error)
}

// {{.ModelTitle}}ServiceImpl 是接口的实际实现（包内实现，不对外暴露）
type {{.ModelTitle}}ServiceImpl struct {
	dao *dao.{{.ModuleName}}Dao
}

func New{{.ModuleName}}Service(dao *dao.{{.ModuleName}}Dao) {{.ModuleName}}Service {
	return &{{.ModelTitle}}ServiceImpl{
		dao: dao,
	}
}

func (s *{{.ModelTitle}}ServiceImpl) ModelToVo(info model.{{.ModuleName}}) vo.{{.ModuleName}}InfoVO {
	return vo.{{.ModuleName}}InfoVO{
		Id:     info.Id,
		// Name:   info.Name,
		CreateTime: times.TimestampToDateTime(info.CreateTime),
		UpdateTime: times.TimestampToDateTime(info.UpdateTime),
	}
}

func (s *{{.ModelTitle}}ServiceImpl) Create(ctx context.Context, req dto.{{.ModuleName}}CreateDTO) error {
	return s.dao.Create(ctx, &model.{{.ModuleName}}{
		Code: req.Code,
		Name: req.Name,
	})
}

func (s *{{.ModelTitle}}ServiceImpl) Edit(ctx context.Context, req dto.{{.ModuleName}}EditDTO) error {
	info, err := s.dao.FindById(ctx, req.Id)
	if info.Id < 1 {
		return errors.New(response.NoExists)
	}
	if err != nil {
		return err
	}
	// info.Name = req.Name
	// info.Code = req.Code
	return s.dao.Update(ctx, &info)
}

func (s *{{.ModelTitle}}ServiceImpl) Delete(ctx context.Context, ids []uint64) error {
	infos, err := s.dao.FindByIds(ctx, "id", ids)
	for _, info := range infos {
		if info.Id < 1 {
			return errors.New(response.NoExists)
		}
		if err != nil {
			return err
		}
		err = s.dao.Delete(ctx, info.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *{{.ModelTitle}}ServiceImpl) Detail(ctx context.Context, id uint64) (vo.{{.ModuleName}}InfoVO, error) {
	info, err := s.dao.FindById(ctx, id)
	if info.Id < 1 {
		return vo.{{.ModuleName}}InfoVO{}, errors.New(response.NoExists)
	}
	return s.ModelToVo(info), err
}

func (s *{{.ModelTitle}}ServiceImpl) List(ctx context.Context, req dto.{{.ModuleName}}ListDTO) ([]vo.{{.ModuleName}}InfoVO, int64, error) {
	var {{.ModelTitle}}sVo []vo.{{.ModuleName}}InfoVO
	where := map[string]interface{}{
		"code":           req.Code,
		"name like ?": "%" + req.Name + "%",
	}
	infos, total, err := s.dao.PageQuery(ctx, req.Page, req.Limit, where, "id desc", []string{})
	if err != nil || total < 1 {
		return nil, total, err
	}
	for _, info := range infos {
		{{.ModelTitle}}sVo = append({{.ModelTitle}}sVo, s.ModelToVo(info))
	}
	return {{.ModelTitle}}sVo, total, err
}
`))
}

func dtoTemplate() *template.Template {
	return template.Must(template.New("dto").Parse(`package dto

type {{.ModuleName}}CreateDTO struct {
	Name       string ` + "`" + `json:"name" form:"name" binding:"required,min=1,max=30"` + "`" + `
	Code       string ` + "`" + `json:"code" form:"code" binding:"required,min=1,max=100"` + "`" + `
	Sort       int    ` + "`" + `json:"sort" form:"sort"` + "`" + `
}

type {{.ModuleName}}EditDTO struct {
	Id uint64 ` + "`" + `json:"id" form:"id" binding:"required"` + "`" + `
	{{.ModuleName}}CreateDTO
}

type {{.ModuleName}}ListDTO struct {
	Name string ` + "`" + `form:"name"` + "`" + `
	Code string ` + "`" + `form:"code"` + "`" + `

    Pagination
}
`))
}

func voTemplate() *template.Template {
	return template.Must(template.New("dto").Parse(`package vo

type {{.ModuleName}}InfoVO struct {
	Id     	   uint64 ` + "`" + `json:"id"` + "`" + `
	Name   	   string ` + "`" + `json:"name"` + "`" + `
	Code   	   string ` + "`" + `json:"code"` + "`" + `
	CreateTime string ` + "`" + `json:"create_time"` + "`" + `
	UpdateTime string ` + "`" + `json:"update_time"` + "`" + `
}
`))
}

func modelTemplate() *template.Template {
	return template.Must(template.New("dto").Parse(`package model

type {{.ModuleName}} struct {
	Id   uint64 ` + "`" + `gorm:"primary_key;auto_increment;comment:主键ID"` + "`" + `
	Name string ` + "`" + `gorm:"type:varchar(100);comment:角色名称"` + "`" + `
	Code string ` + "`" + `gorm:"type:varchar(100);comment:角色标识"` + "`" + `
	Comments string ` + "`" + `gorm:"type:varchar(400);comment:备注"` + "`" + `

	CommonModel
}
`))
}
