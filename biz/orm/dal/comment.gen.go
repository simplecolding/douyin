// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/simplecolding/douyin/biz/biz/orm/model"
)

func newComment(db *gorm.DB, opts ...gen.DOOption) comment {
	_comment := comment{}

	_comment.commentDo.UseDB(db, opts...)
	_comment.commentDo.UseModel(&model.Comment{})

	tableName := _comment.commentDo.TableName()
	_comment.ALL = field.NewAsterisk(tableName)
	_comment.Cid = field.NewInt64(tableName, "cid")
	_comment.Vid = field.NewInt64(tableName, "vid")
	_comment.UID = field.NewInt64(tableName, "uid")
	_comment.Content = field.NewString(tableName, "content")
	_comment.CreatedAt = field.NewTime(tableName, "created_at")
	_comment.UpdatedAt = field.NewTime(tableName, "updated_at")

	_comment.fillFieldMap()

	return _comment
}

type comment struct {
	commentDo

	ALL       field.Asterisk
	Cid       field.Int64  // 主键
	Vid       field.Int64  // 视频id
	UID       field.Int64  // 用户id
	Content   field.String // 评论内容，不能超过255字符
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (c comment) Table(newTableName string) *comment {
	c.commentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comment) As(alias string) *comment {
	c.commentDo.DO = *(c.commentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comment) updateTableName(table string) *comment {
	c.ALL = field.NewAsterisk(table)
	c.Cid = field.NewInt64(table, "cid")
	c.Vid = field.NewInt64(table, "vid")
	c.UID = field.NewInt64(table, "uid")
	c.Content = field.NewString(table, "content")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *comment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["cid"] = c.Cid
	c.fieldMap["vid"] = c.Vid
	c.fieldMap["uid"] = c.UID
	c.fieldMap["content"] = c.Content
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c comment) clone(db *gorm.DB) comment {
	c.commentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c comment) replaceDB(db *gorm.DB) comment {
	c.commentDo.ReplaceDB(db)
	return c
}

type commentDo struct{ gen.DO }

type ICommentDo interface {
	gen.SubQuery
	Debug() ICommentDo
	WithContext(ctx context.Context) ICommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentDo
	WriteDB() ICommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentDo
	Not(conds ...gen.Condition) ICommentDo
	Or(conds ...gen.Condition) ICommentDo
	Select(conds ...field.Expr) ICommentDo
	Where(conds ...gen.Condition) ICommentDo
	Order(conds ...field.Expr) ICommentDo
	Distinct(cols ...field.Expr) ICommentDo
	Omit(cols ...field.Expr) ICommentDo
	Join(table schema.Tabler, on ...field.Expr) ICommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentDo
	Group(cols ...field.Expr) ICommentDo
	Having(conds ...gen.Condition) ICommentDo
	Limit(limit int) ICommentDo
	Offset(offset int) ICommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentDo
	Unscoped() ICommentDo
	Create(values ...*model.Comment) error
	CreateInBatches(values []*model.Comment, batchSize int) error
	Save(values ...*model.Comment) error
	First() (*model.Comment, error)
	Take() (*model.Comment, error)
	Last() (*model.Comment, error)
	Find() ([]*model.Comment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Comment, err error)
	FindInBatches(result *[]*model.Comment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Comment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentDo
	Assign(attrs ...field.AssignExpr) ICommentDo
	Joins(fields ...field.RelationField) ICommentDo
	Preload(fields ...field.RelationField) ICommentDo
	FirstOrInit() (*model.Comment, error)
	FirstOrCreate() (*model.Comment, error)
	FindByPage(offset int, limit int) (result []*model.Comment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentDo) Debug() ICommentDo {
	return c.withDO(c.DO.Debug())
}

func (c commentDo) WithContext(ctx context.Context) ICommentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentDo) ReadDB() ICommentDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentDo) WriteDB() ICommentDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentDo) Session(config *gorm.Session) ICommentDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentDo) Clauses(conds ...clause.Expression) ICommentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentDo) Returning(value interface{}, columns ...string) ICommentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentDo) Not(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentDo) Or(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentDo) Select(conds ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentDo) Where(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICommentDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentDo) Order(conds ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentDo) Distinct(cols ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentDo) Omit(cols ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentDo) Join(table schema.Tabler, on ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentDo) Group(cols ...field.Expr) ICommentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentDo) Having(conds ...gen.Condition) ICommentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentDo) Limit(limit int) ICommentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentDo) Offset(offset int) ICommentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentDo) Unscoped() ICommentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentDo) Create(values ...*model.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentDo) CreateInBatches(values []*model.Comment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentDo) Save(values ...*model.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentDo) First() (*model.Comment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Take() (*model.Comment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Last() (*model.Comment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Find() ([]*model.Comment, error) {
	result, err := c.DO.Find()
	return result.([]*model.Comment), err
}

func (c commentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Comment, err error) {
	buf := make([]*model.Comment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentDo) FindInBatches(result *[]*model.Comment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentDo) Attrs(attrs ...field.AssignExpr) ICommentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentDo) Assign(attrs ...field.AssignExpr) ICommentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentDo) Joins(fields ...field.RelationField) ICommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentDo) Preload(fields ...field.RelationField) ICommentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentDo) FirstOrInit() (*model.Comment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) FirstOrCreate() (*model.Comment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) FindByPage(offset int, limit int) (result []*model.Comment, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentDo) Delete(models ...*model.Comment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentDo) withDO(do gen.Dao) *commentDo {
	c.DO = *do.(*gen.DO)
	return c
}
