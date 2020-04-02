package postgres

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/techmexdev/competitive_writing_api/pkg/selecting"
)

type selectingStorage struct {
	*gorm.DB
}

type gormSelection struct {
	gorm.Model
	Selector       string
	Passage        gormPassage
	BestWriting    gormWriting
	WritingChoices []gormWriting
	Review         string
}

// NewSelectingStorage stores selections.
func NewSelectingStorage(db *gorm.DB) selecting.Storage {
	db.AutoMigrate(&gormSelection{})
	return &selectingStorage{db}
}

func (db *selectingStorage) CreateSelection(sel selecting.Selection) error {
	s := toGormSel(sel)
	if err := db.Create(&s).Error; err != nil {
		return err
	}

	return nil
}

func (db *selectingStorage) ReadSelection(id string) (selecting.Selection, error) {
	var s gormSelection
	if err := db.First(&s, id).Error; err != nil {
		return selecting.Selection{}, err
	}

	return toDomainSel(s), nil
}

func (db *selectingStorage) ListSelectionsFromUser(usr selecting.User) ([]selecting.Selection, error) {
	var ss []gormSelection
	if err := db.Where("user_id = ( select id from gorm_user where username = ? )", usr.Username).Find(&ss).Error; err != nil {
		return []selecting.Selection{}, err
	}

	return toDomainSels(ss), nil
}

func (db *selectingStorage) ListSelectionsWithWritings(ww []selecting.Writing) ([]selecting.Selection, error) {
	var wwIDs []string
	for _, w := range ww {
		wwIDs = append(wwIDs, w.ID)
	}

	var ss []gormSelection
	if err := db.Where("writing_id in (?)", wwIDs).Find(&ss).Error; err != nil {
		return []selecting.Selection{}, err
	}

	return toDomainSels(ss), nil
}

func toGormSel(sel selecting.Selection) gormSelection {
	pID, _ := strconv.Atoi(sel.Passage.ID)
	bwID, _ := strconv.Atoi(sel.BestWriting.ID)

	var wcs []gormWriting
	for _, swc := range sel.WritingChoices {
		wcID, _ := strconv.Atoi(swc.ID)
		wc := gormWriting{Model: gorm.Model{ID: uint(wcID)}}
		wcs = append(wcs, wc)
	}

	return gormSelection{
		Selector:       sel.Selector,
		Passage:        gormPassage{Model: gorm.Model{ID: uint(pID)}},
		BestWriting:    gormWriting{Model: gorm.Model{ID: uint(bwID)}},
		WritingChoices: wcs,
		Review:         sel.Review,
	}
}

func toDomainSels(sels []gormSelection) []selecting.Selection {
	var ss []selecting.Selection
	for _, sel := range sels {
		ss = append(ss, toDomainSel(sel))
	}

	return ss
}

func toDomainSel(sel gormSelection) selecting.Selection {
	var wcs []selecting.Writing
	for _, swc := range sel.WritingChoices {
		wcID := strconv.Itoa(int(swc.ID))
		wc := selecting.Writing{ID: wcID}
		wcs = append(wcs, wc)
	}

	return selecting.Selection{
		Selector:       sel.Selector,
		Passage:        selecting.Passage{ID: strconv.Itoa(int(sel.Passage.ID))},
		BestWriting:    selecting.Writing{ID: strconv.Itoa(int(sel.BestWriting.ID))},
		WritingChoices: wcs,
		Review:         sel.Review,
	}
}
