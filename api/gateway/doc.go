package gateway

import (
	"doc-api/api/gateway/model"
	"doc-api/api/usecase"
	"errors"

	"gorm.io/gorm"
)

type (
	docRepository struct {
		db *gorm.DB
	}
)

func NewDocRepository(db *gorm.DB) usecase.DocRepository {
	return &docRepository{db: db}
}

// userが所有しているdocumentを列挙します
//
// return:
//   - 成功事: err == nil
//   - 失敗時: err != nil
func (r *docRepository) GetOwnDocs(userId string) ([]usecase.DocHeader, error) {
	user := model.User{ManualIDModel: model.ManualIDModel{ID: userId}}
	docs := []model.Document{}
	if err := r.db.Model(&user).Association("Documents").Find(&docs); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	ret := []usecase.DocHeader{}
	for _, doc := range docs {
		ret = append(ret, usecase.DocHeader{
			Id:    doc.ID,
			Title: doc.Title,
		})
	}
	return ret, nil
}

func (r *docRepository) GetOwnDoc(userId string, documentId string) (*usecase.Doc, error) {
	user := model.User{ManualIDModel: model.ManualIDModel{ID: userId}}
	doc := model.Document{AutoUUIDModel: model.AutoUUIDModel{ID: documentId}}

	if err := r.db.Model(&user).Association("Documents").Find(&doc); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	ret := usecase.Doc{
		Id:    doc.ID,
		Title: doc.Title,
		Body:  doc.Body,
	}

	return &ret, nil
}

// doc.idは使用しない
func (r *docRepository) AppendOwnDoc(userId string, doc usecase.Doc) (documentId string, err error) {
	user := model.User{ManualIDModel: model.ManualIDModel{ID: userId}}

	document := model.Document{Body: doc.Body, Title: doc.Title}
	if err := r.db.Model(&user).Association("Documents").Append(&document); err != nil {
		return "", err
	}

	return document.ID, nil
}

func (r *docRepository) UpdateOwnDoc(userId string, doc usecase.Doc) error {
	document := model.Document{AutoUUIDModel: model.AutoUUIDModel{ID: doc.Id}}
	if err := r.db.Model(&document).Updates(model.Document{Title: doc.Title, Body: doc.Body}).Error; err != nil {
		return err
	}

	return nil
}

func (r *docRepository) DeleteOwnDoc(userId string, docId string) error {
	document := model.Document{AutoUUIDModel: model.AutoUUIDModel{ID: docId}}
	if err := r.db.Delete(&document).Error; err != nil {
		return err
	}
	return nil
}
