package impl_repository

import (
	"context"
	"errors"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func (r *UserRepositoryImpl) CreateUserRepository(ctx context.Context, user *entity.User) *entity.User {
	err := r.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (r *UserRepositoryImpl) FindByIdUserRepo(ctx context.Context, id int64) (*entity.User, error) {
	var user *entity.User
	result := r.DB.WithContext(ctx).Where("id_user = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return &entity.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindAllUserRepo(ctx context.Context) []*entity.User {
	var users []*entity.User
	r.DB.WithContext(ctx).Find(&users)
	return users
}

func (r *UserRepositoryImpl) FindByUsernamePhoneAndEmail(ctx context.Context, username, phone, email string) (*entity.User, error) {
	var user entity.User // Struct untuk menampung hasil query

	// Query dengan kondisi AND
	result := r.DB.WithContext(ctx).Where("username = ? OR nomor_telepon = ? OR email = ?", username, phone, email).First(&user)

	if result.RowsAffected != 0 {
		return nil, errors.New("ada 3 field yang unique, tebak sendiri request mana yang kagak unique, saya malas dan saya bangga") // Kembalikan error jika tidak ada data yang cocok
	}

	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUserRepo(ctx context.Context, user *entity.User) *entity.User {
	err := r.DB.WithContext(ctx).Where("id_user = ?", user.ID).Updates(user).Error
	exception.PanicLogging(err)
	return user
}

func (r *UserRepositoryImpl) DeleteUserRepo(ctx context.Context, user *entity.User) {
	err := r.DB.WithContext(ctx).Delete(&user).Error
	exception.PanicLogging(err)
}

func (r *UserRepositoryImpl) RegisterUserRepo(ctx context.Context, user *entity.User) *entity.User {
	err := r.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (r *UserRepositoryImpl) MarkUserEmailVerified(ctx context.Context, userID int64) error {
	err := r.DB.WithContext(ctx).Model(&entity.User{}).Where("id_user = ?", userID).Update("is_email_verified", true).Error
	if err != nil {
		return err
	}
	return nil
}
