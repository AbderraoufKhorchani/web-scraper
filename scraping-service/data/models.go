package data

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type baseQuote struct {
	QuoteText string   `json:"text"`
	Author    string   `json:"author"`
	Tags      []string `json:"tags"`
}

type Quote struct {
	gorm.Model
	QuoteText string `json:"text"`
	Author    string `json:"author"`
	Tags      []Tag  `gorm:"many2many:quote_tags;"`
}

type Tag struct {
	gorm.Model
	Name   string  `json:"name"`
	Quotes []Quote `gorm:"many2many:quote_tags;"`
}

type Models struct {
	Quote Quote
}

var db *gorm.DB

func New(dbPool *gorm.DB) Models {
	db = dbPool
	err := db.AutoMigrate(&Quote{}, &Tag{})
	if err != nil {
		fmt.Print(err)
	}
	return Models{
		Quote: Quote{},
	}
}

func (q *Quote) GetByAuthor(author string) ([]baseQuote, error) {

	var quotes []Quote
	result := db.Preload("Tags").Where("author = ?", author).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}

	baseQuotes := q.basingQuote(quotes)
	return baseQuotes, nil
}

func (q *Quote) GetByTag(tag string) ([]baseQuote, error) {
	var quotes []Quote
	result := db.Preload("Tags").Joins("JOIN quote_tags ON quotes.id = quote_tags.quote_id").
		Joins("JOIN tags ON tags.id = quote_tags.tag_id").
		Where("tags.name = ?", tag).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}

	baseQuotes := q.basingQuote(quotes)
	return baseQuotes, nil
}

func (q Quote) GetAll() ([]baseQuote, error) {

	var quotes []Quote
	result := db.Preload("Tags").Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}

	baseQuotes := q.basingQuote(quotes)
	return baseQuotes, nil
}

func (q *Quote) DatabaseIsEmpty() (bool, error) {
	var count int64
	result := db.Model(&Quote{}).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	return count == 0, nil
}

func (q *Quote) AddQuoteWithTags(quote Quote, tagNames []string) error {
	tx := db.Begin()

	if err := tx.Create(&quote).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, name := range tagNames {
		var tag Tag
		if err := tx.Where("name = ?", name).First(&tag).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				tx.Rollback()
				return err
			}
			tag.Name = name
			if err := tx.Create(&tag).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
		if err := tx.Model(&quote).Association("Tags").Append(&tag); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (q *Quote) basingQuote(quotes []Quote) []baseQuote {

	var baseQuotes []baseQuote

	for _, quote := range quotes {

		var baseTags []string
		for _, tag := range quote.Tags {
			baseTags = append(baseTags, tag.Name)
		}
		baseQuote := baseQuote{
			QuoteText: quote.QuoteText,
			Author:    quote.Author,
			Tags:      baseTags,
		}
		baseQuotes = append(baseQuotes, baseQuote)

	}
	return baseQuotes
}
