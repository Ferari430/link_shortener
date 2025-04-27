package link

// CRUD Links for DB
import (
	"log"
	"my_project/db"

	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {

	return &LinkRepository{Database: database}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("URL finded by Hash")
	return &link, nil
}

func (repo *LinkRepository) GetByURL(url string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "url = ?", url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) UpdateById(link *Link) (*Link, error) {
	//Clauses(clause.Returning{}) == UPDATE links SET hash = 'abc123' WHERE id = 1 RETURNING *;
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(&link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil

}

func (repo *LinkRepository) Delete(id uint64) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		log.Println("cant delete")
		return result.Error
	}

	return nil

}

func (repo *LinkRepository) GetById(id uint64) error {
	var link Link
	result := repo.Database.DB.First(&link, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link

	repo.Database.Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)

	return links
}

func (repo *LinkRepository) GetActive() int64 {
	var count int64

	repo.Database.Table("links").
		Count(&count).
		Where("deleted_at is null")

	return count
}
