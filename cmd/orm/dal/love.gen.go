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

	"github.com/simplecolding/douyin/cmd/orm/model"
)

func newLove(db *gorm.DB, opts ...gen.DOOption) love {
	_love := love{}

	_love.loveDo.UseDB(db, opts...)
	_love.loveDo.UseModel(&model.Love{})

	tableName := _love.loveDo.TableName()
	_love.ALL = field.NewAsterisk(tableName)
	_love.Lid = field.NewInt64(tableName, "lid")
	_love.Vid = field.NewInt64(tableName, "vid")
	_love.UID = field.NewInt64(tableName, "uid")
	_love.Status = field.NewBool(tableName, "status")
	_love.CreatedAt = field.NewTime(tableName, "created_at")
	_love.UpdatedAt = field.NewTime(tableName, "updated_at")

	_love.fillFieldMap()

	return _love
}

type love struct {
	loveDo

	ALL       field.Asterisk
	Lid       field.Int64 // 主键
	Vid       field.Int64 // 视频id
	UID       field.Int64 // 用户id
	Status    field.Bool  // 是否删除 1:是  0:否
	CreatedAt field.Time  // 创建时间
	UpdatedAt field.Time  // 更新时间

	fieldMap map[string]field.Expr
}

func (l love) Table(newTableName string) *love {
	l.loveDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l love) As(alias string) *love {
	l.loveDo.DO = *(l.loveDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *love) updateTableName(table string) *love {
	l.ALL = field.NewAsterisk(table)
	l.Lid = field.NewInt64(table, "lid")
	l.Vid = field.NewInt64(table, "vid")
	l.UID = field.NewInt64(table, "uid")
	l.Status = field.NewBool(table, "status")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")

	l.fillFieldMap()

	return l
}

func (l *love) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *love) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["lid"] = l.Lid
	l.fieldMap["vid"] = l.Vid
	l.fieldMap["uid"] = l.UID
	l.fieldMap["status"] = l.Status
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
}

func (l love) clone(db *gorm.DB) love {
	l.loveDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l love) replaceDB(db *gorm.DB) love {
	l.loveDo.ReplaceDB(db)
	return l
}

type loveDo struct{ gen.DO }

type ILoveDo interface {
	gen.SubQuery
	Debug() ILoveDo
	WithContext(ctx context.Context) ILoveDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILoveDo
	WriteDB() ILoveDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILoveDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILoveDo
	Not(conds ...gen.Condition) ILoveDo
	Or(conds ...gen.Condition) ILoveDo
	Select(conds ...field.Expr) ILoveDo
	Where(conds ...gen.Condition) ILoveDo
	Order(conds ...field.Expr) ILoveDo
	Distinct(cols ...field.Expr) ILoveDo
	Omit(cols ...field.Expr) ILoveDo
	Join(table schema.Tabler, on ...field.Expr) ILoveDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILoveDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILoveDo
	Group(cols ...field.Expr) ILoveDo
	Having(conds ...gen.Condition) ILoveDo
	Limit(limit int) ILoveDo
	Offset(offset int) ILoveDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILoveDo
	Unscoped() ILoveDo
	Create(values ...*model.Love) error
	CreateInBatches(values []*model.Love, batchSize int) error
	Save(values ...*model.Love) error
	First() (*model.Love, error)
	Take() (*model.Love, error)
	Last() (*model.Love, error)
	Find() ([]*model.Love, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Love, err error)
	FindInBatches(result *[]*model.Love, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Love) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILoveDo
	Assign(attrs ...field.AssignExpr) ILoveDo
	Joins(fields ...field.RelationField) ILoveDo
	Preload(fields ...field.RelationField) ILoveDo
	FirstOrInit() (*model.Love, error)
	FirstOrCreate() (*model.Love, error)
	FindByPage(offset int, limit int) (result []*model.Love, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILoveDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l loveDo) Debug() ILoveDo {
	return l.withDO(l.DO.Debug())
}

func (l loveDo) WithContext(ctx context.Context) ILoveDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l loveDo) ReadDB() ILoveDo {
	return l.Clauses(dbresolver.Read)
}

func (l loveDo) WriteDB() ILoveDo {
	return l.Clauses(dbresolver.Write)
}

func (l loveDo) Session(config *gorm.Session) ILoveDo {
	return l.withDO(l.DO.Session(config))
}

func (l loveDo) Clauses(conds ...clause.Expression) ILoveDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l loveDo) Returning(value interface{}, columns ...string) ILoveDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l loveDo) Not(conds ...gen.Condition) ILoveDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l loveDo) Or(conds ...gen.Condition) ILoveDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l loveDo) Select(conds ...field.Expr) ILoveDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l loveDo) Where(conds ...gen.Condition) ILoveDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l loveDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILoveDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l loveDo) Order(conds ...field.Expr) ILoveDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l loveDo) Distinct(cols ...field.Expr) ILoveDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l loveDo) Omit(cols ...field.Expr) ILoveDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l loveDo) Join(table schema.Tabler, on ...field.Expr) ILoveDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l loveDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILoveDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l loveDo) RightJoin(table schema.Tabler, on ...field.Expr) ILoveDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l loveDo) Group(cols ...field.Expr) ILoveDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l loveDo) Having(conds ...gen.Condition) ILoveDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l loveDo) Limit(limit int) ILoveDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l loveDo) Offset(offset int) ILoveDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l loveDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILoveDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l loveDo) Unscoped() ILoveDo {
	return l.withDO(l.DO.Unscoped())
}

func (l loveDo) Create(values ...*model.Love) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l loveDo) CreateInBatches(values []*model.Love, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l loveDo) Save(values ...*model.Love) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l loveDo) First() (*model.Love, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Love), nil
	}
}

func (l loveDo) Take() (*model.Love, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Love), nil
	}
}

func (l loveDo) Last() (*model.Love, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Love), nil
	}
}

func (l loveDo) Find() ([]*model.Love, error) {
	result, err := l.DO.Find()
	return result.([]*model.Love), err
}

func (l loveDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Love, err error) {
	buf := make([]*model.Love, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l loveDo) FindInBatches(result *[]*model.Love, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l loveDo) Attrs(attrs ...field.AssignExpr) ILoveDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l loveDo) Assign(attrs ...field.AssignExpr) ILoveDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l loveDo) Joins(fields ...field.RelationField) ILoveDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l loveDo) Preload(fields ...field.RelationField) ILoveDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l loveDo) FirstOrInit() (*model.Love, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Love), nil
	}
}

func (l loveDo) FirstOrCreate() (*model.Love, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Love), nil
	}
}

func (l loveDo) FindByPage(offset int, limit int) (result []*model.Love, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l loveDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l loveDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l loveDo) Delete(models ...*model.Love) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *loveDo) withDO(do gen.Dao) *loveDo {
	l.DO = *do.(*gen.DO)
	return l
}
