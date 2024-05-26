package complaint

import (
	"e-complaint-api/constants"
	"e-complaint-api/entities"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ComplaintRepo struct {
	DB *gorm.DB
}

func NewComplaintRepo(db *gorm.DB) *ComplaintRepo {
	return &ComplaintRepo{DB: db}
}

func (r *ComplaintRepo) GetPaginated(limit int, page int, search string, filter map[string]interface{}, sortBy string, sortType string) ([]entities.Complaint, error) {
	var complaints []entities.Complaint
	query := r.DB

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("description LIKE ? OR address LIKE ? OR id LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	query = query.Order(sortBy + " " + sortType)

	if err := query.Limit(limit).Offset((page - 1) * limit).Preload("User").Preload("Regency").Preload("Category").Preload("Files").Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func (r *ComplaintRepo) GetMetaData(limit int, page int, search string, filter map[string]interface{}) (entities.Metadata, error) {
	var totalData int64

	query := r.DB.Model(&entities.Complaint{})

	if filter != nil {
		query = query.Where(filter)
	}

	if search != "" {
		query = query.Where("description LIKE ? OR address LIKE ? OR id LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&totalData).Error; err != nil {
		return entities.Metadata{}, err
	}

	metadata := entities.Metadata{
		TotalData: int(totalData),
	}

	return metadata, nil
}

func (r *ComplaintRepo) GetByID(id string) (entities.Complaint, error) {
	var complaint entities.Complaint

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("id = ?", id).First(&complaint).Error; err != nil {
		return entities.Complaint{}, constants.ErrReportNotFound
	}

	return complaint, nil
}

func (r *ComplaintRepo) Create(complaint *entities.Complaint) error {
	if err := r.DB.Create(complaint).Error; err != nil {
		return err
	}

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("id = ?", complaint.ID).First(complaint).Error; err != nil {
		return err
	}

	return nil
}

func (r *ComplaintRepo) Delete(id string, userId int) error {
	var complaint entities.Complaint

	if err := r.DB.Where("id = ?", id).First(&complaint).Error; err != nil {
		return constants.ErrReportNotFound
	}

	if complaint.UserID != userId {
		return constants.ErrUnauthorized
	}

	complaint.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := r.DB.Save(&complaint).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *ComplaintRepo) AdminDelete(id string) error {
	var complaint entities.Complaint

	if err := r.DB.Where("id = ?", id).First(&complaint).Error; err != nil {
		return constants.ErrReportNotFound
	}

	complaint.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := r.DB.Save(&complaint).Error; err != nil {
		return constants.ErrInternalServerError
	}

	return nil
}

func (r *ComplaintRepo) Update(complaint entities.Complaint) (entities.Complaint, error) {
	var oldComplaint entities.Complaint

	if err := r.DB.Where("id = ?", complaint.ID).First(&oldComplaint).Error; err != nil {
		return entities.Complaint{}, constants.ErrReportNotFound
	}

	if oldComplaint.UserID != complaint.UserID {
		return entities.Complaint{}, constants.ErrUnauthorized
	}

	oldComplaint.Description = complaint.Description
	oldComplaint.Type = complaint.Type
	oldComplaint.CategoryID = complaint.CategoryID
	oldComplaint.RegencyID = complaint.RegencyID
	oldComplaint.Address = complaint.Address

	fmt.Println(oldComplaint.CategoryID)
	fmt.Println(complaint.CategoryID)

	if err := r.DB.Save(&oldComplaint).Error; err != nil {
		return entities.Complaint{}, err
	}

	if err := r.DB.Preload("User").Preload("Regency").Preload("Category").Preload("Files").Where("id = ?", oldComplaint.ID).First(&complaint).Error; err != nil {
		return entities.Complaint{}, constants.ErrInternalServerError
	}

	return complaint, nil
}
