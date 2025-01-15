package usecase_test

import (
	"testing"

	"errors"

	"github.com/stretchr/testify/suite"

	"github.com/golang/mock/gomock"
	"golang_template_source/mock/mock_repository"
	"golang_template_source/domain"
	"golang_template_source/usecase"
)

type testUserUseCaseSuite struct {
	suite.Suite
	ctrl          *gomock.Controller
	mockRepo      *mock_repository.MockUserRepository
	userUseCase   usecase.UserUseCase
}
func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, &testUserUseCaseSuite{})
}

func (s *testUserUseCaseSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mockRepo = mock_repository.NewMockUserRepository(s.ctrl)
	s.userUseCase = usecase.NewUserUseCase(s.mockRepo)
}

func (s *testUserUseCaseSuite) TearDownTest() {
	s.ctrl.Finish()
}
func (s *testUserUseCaseSuite) TestGetAllUsers() {
	expectedUsers := []*domain.SysUser{
		{
			ID:    1,
			Name:  "User1",
			Email: "user1@example.com",
		},
		{
			ID:    2,
			Name:  "User2",
			Email: "user2@example.com",
		},
	}

	s.mockRepo.EXPECT().GetAll().Return(expectedUsers, nil)

	users, err := s.userUseCase.GetAllUsers()

	s.NoError(err)
	s.Equal(expectedUsers, users)
}

func (s *testUserUseCaseSuite) TestGetUserByID() {
	tests := []struct {
		name        string
		userID      int
		expectedUser *domain.SysUser
		expectedErr  error
		mockBehavior func(r *mock_repository.MockUserRepository)
	}{
		{
			name:   "User found",
			userID: 1,
			expectedUser: &domain.SysUser{
				ID:    1,
				Name:  "User1",
				Email: "user1@example.com",
			},
			expectedErr: nil,
			mockBehavior: func(r *mock_repository.MockUserRepository) {
				r.EXPECT().GetByID(1).Return(&domain.SysUser{
					ID:    1,
					Name:  "User1",
					Email: "user1@example.com",
				}, nil)
			},
		},
		{
			name:        "User not found",
			userID:      2,
			expectedUser: nil,
			expectedErr:  errors.New("user not found"),
			mockBehavior: func(r *mock_repository.MockUserRepository) {
				r.EXPECT().GetByID(2).Return(nil, errors.New("user not found"))
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			if tt.mockBehavior != nil {
				tt.mockBehavior(s.mockRepo)
			}

			user, err := s.userUseCase.GetUserByID(tt.userID)

			s.Equal(tt.expectedUser, user)
			s.Equal(tt.expectedErr, err)
		})
	}
}
func (s *testUserUseCaseSuite) TestFindByEmail() {
	tests := []struct {
		name         string
		inputEmail   string
		mockBehavior func(r *mock_repository.MockUserRepository)
		expectedUser *domain.SysUser
		expectedErr  error
	}{
		{
			name:       "successful find",
			inputEmail: "test@gmail.com",
			mockBehavior: func(r *mock_repository.MockUserRepository) {
				r.EXPECT().FindByEmail("test@gmail.com").Return(&domain.SysUser{
					ID:       1,
					Name:     "username test",
					Email:    "test@gmail.com",
					Password: "123456",
					Phone:    "123456",
				}, nil)
			},
			expectedUser: &domain.SysUser{
				ID:       1,
				Name:     "username test",
				Email:    "test@gmail.com",
				Password: "123456",
				Phone:    "123456",
			},
			expectedErr: nil,
		},
		{
			name:       "user not found",
			inputEmail: "notfound@gmail.com",
			mockBehavior: func(r *mock_repository.MockUserRepository) {
				r.EXPECT().FindByEmail("notfound@gmail.com").Return(nil, errors.New("user not found"))
			},
			expectedUser: nil,
			expectedErr:  errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// Arrange
			if tt.mockBehavior != nil {
				tt.mockBehavior(s.mockRepo)
			}

			// Act
			user, err := s.userUseCase.FindByEmail(tt.inputEmail)

			// Assert
			s.Equal(tt.expectedUser, user)
			s.Equal(tt.expectedErr, err)
		})
	}
}