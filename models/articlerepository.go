package models

import "github.com/DmitryOdintsov/GinGormAPI/storage"

func GetAllArticles(a *[]Article) error {
	return storage.DB.Find(a).Error
}

func AddNewArticle(a *Article) error {
	return storage.DB.Create(a).Error
}

func GetArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).First(a).Error
}

func UpdateArticleById(a *Article, id string) error {
	return storage.DB.Update(a).Error
}

func DeleteArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).Delete(a).Error
}
