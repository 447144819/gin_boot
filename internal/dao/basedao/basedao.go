package basedao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log/slog"
)

// BaseDao 是一个泛型 DAO，M=模型，T=模型主键类型
type BaseDao[M, T any] struct {
	DB *gorm.DB
}

// NewBaseDao 创建一个泛型 DAO 实例
func NewBaseDao[M, T any](db *gorm.DB) *BaseDao[M, T] {
	if db == nil {
		panic("BaseDao: db 数据库为空")
	}

	return &BaseDao[M, T]{DB: db}
}

func (d *BaseDao[M, T]) FindById(ctx context.Context, id T) (M, error) {
	var entity M
	err := d.DB.WithContext(ctx).First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity, nil // 返回 nil 表示未找到记录
	}
	return entity, err
}

func (d *BaseDao[M, T]) FindByIds(ctx context.Context, field string, ids []T) ([]M, error) {
	var entity []M
	err := d.DB.WithContext(ctx).Model(new(M)).Where(field+" in(?)", ids).Find(&entity).Error
	return entity, err
}

// Create 创建单个记录
func (d *BaseDao[M, T]) Create(ctx context.Context, entity *M) error {
	if entity == nil {
		return errors.New("要创建的数据不存在")
	}
	return d.DB.WithContext(ctx).Create(entity).Error
}

// BatchCreate 批量创建记录，支持事务
func (d *BaseDao[M, T]) BatchCreate(ctx context.Context, entities []M) error {
	if len(entities) == 0 {
		return errors.New("要更新的数据不存在")
	}
	return d.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for i := range entities {
			if err := tx.Create(&entities[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// Update 更新记录
func (d *BaseDao[M, T]) Update(ctx context.Context, entity *M) error {
	if entity == nil {
		return errors.New("要更新的数据不存在")
	}
	return d.DB.WithContext(ctx).Save(entity).Error
}

// 更新状态
func (d *BaseDao[M, T]) UpdateStatus(ctx context.Context, id T, field string, status int) error {
	return d.DB.WithContext(ctx).Model(new(M)).Where("id =?", id).Update(field, status).Error
}

// Delete 删除记录
func (d *BaseDao[M, T]) Delete(ctx context.Context, id T) error {
	return d.DB.WithContext(ctx).Delete(new(M), id).Error
}

// BatchDelete 批量删除记录，支持事务
func (d *BaseDao[M, T]) BatchDelete(ctx context.Context, ids []T) error {
	if len(ids) == 0 {
		return errors.New("删除的ids为空")
	}
	return d.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var entity M
		return tx.Delete(&entity, ids).Error
	})
}

// FindAll 查询所有记录
func (d *BaseDao[M, T]) FindAll(ctx context.Context, where map[string]interface{}, orderBy ...string) ([]M, error) {
	var entities []M
	query := d.DB.WithContext(ctx).Model(new(M))

	// 应用查询条件
	for k, v := range where {
		if v != nil && v != "" {
			query = query.Where(k, v)
		}
	}

	// 应用排序
	if len(orderBy) > 0 {
		for _, order := range orderBy {
			query = query.Order(order)
		}
	} else {
		query = query.Order("id DESC")
	}

	err := query.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

// PageQuery 分页查询，支持灵活的查询条件、排序和字段选择
func (d *BaseDao[M, T]) PageQuery(ctx context.Context, page, pageSize int, where map[string]interface{}, orderBy string, selectFields []string) ([]M, int64, error) {
	var entities []M
	var total int64
	query := d.DB.WithContext(ctx).Model(new(M))

	// 应用查询条件
	for k, v := range where {
		if v != nil && v != "" {
			query = query.Where(k, v)
		}
	}

	// 应用字段选择
	if len(selectFields) > 0 {
		query = query.Select(selectFields)
	}

	// 查询总记录数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 如果没有记录，直接返回空列表
	if total == 0 {
		return entities, 0, nil
	}

	// 应用分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 应用排序
	if orderBy != "" {
		query = query.Order(orderBy)
	}

	// 执行分页查询
	err := query.Offset(offset).Limit(pageSize).Find(&entities).Error
	if err != nil {
		return nil, 0, err
	}

	slog.Debug("BaseDao: PageQuery executed",
		"page", page,
		"pageSize", pageSize,
		"total", total,
		"resultCount", len(entities))
	return entities, total, nil
}

// WithTransaction 执行事务操作
func (d *BaseDao[M, T]) WithTransaction(ctx context.Context, fn func(txDao *BaseDao[M, T]) error) error {
	return d.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txDao := &BaseDao[M, T]{DB: tx}
		return fn(txDao)
	})
}
