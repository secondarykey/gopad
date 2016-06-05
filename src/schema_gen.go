// generated by argen; DO NOT EDIT
package main

import (
	"fmt"

	"github.com/monochromegane/argen"
)

type MemoRelation struct {
	src *Memo
	*ar.Relation
}

func (m *Memo) newRelation() *MemoRelation {
	r := &MemoRelation{
		m,
		ar.NewRelation(db, logger).Table("memos"),
	}
	r.Select(
		"id",
		"title",
		"content",
	)

	return r
}

func (m Memo) Select(columns ...string) *MemoRelation {
	return m.newRelation().Select(columns...)
}

func (r *MemoRelation) Select(columns ...string) *MemoRelation {
	cs := []string{}
	for _, c := range columns {
		if r.src.isColumnName(c) {
			cs = append(cs, fmt.Sprintf("memos.%s", c))
		} else {
			cs = append(cs, c)
		}
	}
	r.Relation.Columns(cs...)
	return r
}

func (m Memo) Find(id int) (*Memo, error) {
	return m.newRelation().Find(id)
}

func (r *MemoRelation) Find(id int) (*Memo, error) {
	return r.FindBy("id", id)
}

func (m Memo) FindBy(cond string, args ...interface{}) (*Memo, error) {
	return m.newRelation().FindBy(cond, args...)
}

func (r *MemoRelation) FindBy(cond string, args ...interface{}) (*Memo, error) {
	return r.Where(cond, args...).Limit(1).QueryRow()
}

func (m Memo) First() (*Memo, error) {
	return m.newRelation().First()
}

func (r *MemoRelation) First() (*Memo, error) {
	return r.Order("id", "ASC").Limit(1).QueryRow()
}

func (m Memo) Last() (*Memo, error) {
	return m.newRelation().Last()
}

func (r *MemoRelation) Last() (*Memo, error) {
	return r.Order("id", "DESC").Limit(1).QueryRow()
}

func (m Memo) Where(cond string, args ...interface{}) *MemoRelation {
	return m.newRelation().Where(cond, args...)
}

func (r *MemoRelation) Where(cond string, args ...interface{}) *MemoRelation {
	r.Relation.Where(cond, args...)
	return r
}

func (r *MemoRelation) And(cond string, args ...interface{}) *MemoRelation {
	r.Relation.And(cond, args...)
	return r
}

func (m Memo) Order(column, order string) *MemoRelation {
	return m.newRelation().Order(column, order)
}

func (r *MemoRelation) Order(column, order string) *MemoRelation {
	r.Relation.OrderBy(column, order)
	return r
}

func (m Memo) Limit(limit int) *MemoRelation {
	return m.newRelation().Limit(limit)
}

func (r *MemoRelation) Limit(limit int) *MemoRelation {
	r.Relation.Limit(limit)
	return r
}

func (m Memo) Offset(offset int) *MemoRelation {
	return m.newRelation().Offset(offset)
}

func (r *MemoRelation) Offset(offset int) *MemoRelation {
	r.Relation.Offset(offset)
	return r
}

func (m Memo) Group(group string, groups ...string) *MemoRelation {
	return m.newRelation().Group(group, groups...)
}

func (r *MemoRelation) Group(group string, groups ...string) *MemoRelation {
	r.Relation.GroupBy(group, groups...)
	return r
}

func (r *MemoRelation) Having(cond string, args ...interface{}) *MemoRelation {
	r.Relation.Having(cond, args...)
	return r
}

func (m Memo) IsValid() (bool, *ar.Errors) {
	result := true
	errors := &ar.Errors{}
	var on ar.On
	if m.IsNewRecord() {
		on = ar.OnCreate()
	} else {
		on = ar.OnUpdate()
	}
	rules := map[string]*ar.Validation{}
	for name, rule := range rules {
		if ok, errs := ar.NewValidator(rule).On(on).IsValid(m.fieldValueByName(name)); !ok {
			result = false
			errors.SetErrors(name, errs)
		}
	}
	customs := []*ar.Validation{}
	for _, rule := range customs {
		custom := ar.NewValidator(rule).On(on).Custom()
		custom(errors)
	}
	if len(errors.Messages) > 0 {
		result = false
	}
	return result, errors
}

type MemoParams Memo

func (m Memo) Build(p MemoParams) *Memo {
	return &Memo{
		Id:      p.Id,
		Title:   p.Title,
		Content: p.Content,
	}
}

func (m Memo) Create(p MemoParams) (*Memo, *ar.Errors) {
	n := m.Build(p)
	_, errs := n.Save()
	return n, errs
}

func (m *Memo) IsNewRecord() bool {
	return ar.IsZero(m.Id)
}

func (m *Memo) IsPersistent() bool {
	return !m.IsNewRecord()
}

func (m *Memo) Save(validate ...bool) (bool, *ar.Errors) {
	if len(validate) == 0 || len(validate) > 0 && validate[0] {
		if ok, errs := m.IsValid(); !ok {
			return false, errs
		}
	}
	errs := &ar.Errors{}
	if m.IsNewRecord() {
		ins := ar.NewInsert(db, logger).Table("memos").Params(map[string]interface{}{
			"title":   m.Title,
			"content": m.Content,
		})

		if result, err := ins.Exec(); err != nil {
			errs.AddError("base", err)
			return false, errs
		} else {
			if lastId, err := result.LastInsertId(); err == nil {
				m.Id = int(lastId)
			}
		}
		return true, nil
	} else {
		upd := ar.NewUpdate(db, logger).Table("memos").Params(map[string]interface{}{
			"id":      m.Id,
			"title":   m.Title,
			"content": m.Content,
		}).Where("id", m.Id)

		if _, err := upd.Exec(); err != nil {
			errs.AddError("base", err)
			return false, errs
		}
		return true, nil
	}
}

func (m *Memo) Update(p MemoParams) (bool, *ar.Errors) {

	if !ar.IsZero(p.Id) {
		m.Id = p.Id
	}
	if !ar.IsZero(p.Title) {
		m.Title = p.Title
	}
	if !ar.IsZero(p.Content) {
		m.Content = p.Content
	}
	return m.Save()
}

func (m *Memo) UpdateColumns(p MemoParams) (bool, *ar.Errors) {

	if !ar.IsZero(p.Id) {
		m.Id = p.Id
	}
	if !ar.IsZero(p.Title) {
		m.Title = p.Title
	}
	if !ar.IsZero(p.Content) {
		m.Content = p.Content
	}
	return m.Save(false)
}

func (m *Memo) Destroy() (bool, *ar.Errors) {
	return m.Delete()
}

func (m *Memo) Delete() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	if _, err := ar.NewDelete(db, logger).Table("memos").Where("id", m.Id).Exec(); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (m Memo) DeleteAll() (bool, *ar.Errors) {
	errs := &ar.Errors{}
	if _, err := ar.NewDelete(db, logger).Table("memos").Exec(); err != nil {
		errs.AddError("base", err)
		return false, errs
	}
	return true, nil
}

func (r *MemoRelation) Query() ([]*Memo, error) {
	rows, err := r.Relation.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []*Memo{}
	for rows.Next() {
		row := &Memo{}
		err := rows.Scan(row.fieldPtrsByName(r.Relation.GetColumns())...)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (r *MemoRelation) QueryRow() (*Memo, error) {
	row := &Memo{}
	err := r.Relation.QueryRow(row.fieldPtrsByName(r.Relation.GetColumns())...)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (m Memo) Exists() bool {
	return m.newRelation().Exists()
}

func (m Memo) Count(column ...string) int {
	return m.newRelation().Count(column...)
}

func (m Memo) All() *MemoRelation {
	return m.newRelation().All()
}

func (r *MemoRelation) All() *MemoRelation {
	return r
}

func (m *Memo) fieldValueByName(name string) interface{} {
	switch name {
	case "id", "memos.id":
		return m.Id
	case "title", "memos.title":
		return m.Title
	case "content", "memos.content":
		return m.Content
	default:
		return ""
	}
}

func (m *Memo) fieldPtrByName(name string) interface{} {
	switch name {
	case "id", "memos.id":
		return &m.Id
	case "title", "memos.title":
		return &m.Title
	case "content", "memos.content":
		return &m.Content
	default:
		return nil
	}
}

func (m *Memo) fieldPtrsByName(names []string) []interface{} {
	fields := []interface{}{}
	for _, n := range names {
		f := m.fieldPtrByName(n)
		fields = append(fields, f)
	}
	return fields
}

func (m *Memo) isColumnName(name string) bool {
	for _, c := range m.columnNames() {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Memo) columnNames() []string {
	return []string{
		"id",
		"title",
		"content",
	}
}
