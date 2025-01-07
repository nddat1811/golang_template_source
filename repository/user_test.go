package repository_test

import (
	// "database/sql"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"
	"time"

	"golang_template_source/domain"
	"golang_template_source/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testUserRepositorySuite struct {
	suite.Suite
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, &testUserRepositorySuite{})
}

func (s *testUserRepositorySuite) TestGetAll() {
	db, gdb, mock, err := NewConnManager()
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s.Run("Success", func() {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john@example.com", "123-456-7890", "hashed_password1", time.Now(), time.Now())
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).WillReturnRows(rows)

		userRepository := repository.NewUserRepository(gdb)
		users, err := userRepository.GetAll()
		assert.NoError(s.T(), err)
		assert.Len(s.T(), users, 1)
		assert.Equal(s.T(), "John Doe", users[0].Name)
		assert.NoError(s.T(), mock.ExpectationsWereMet())
	})
}

func (s *testUserRepositorySuite) TestGetByID() {
	db, gdb, mock, err := NewConnManager()
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository := repository.NewUserRepository(gdb)

	s.Run("Success", func() {
		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john@example.com", "123-456-7890", "hashed_password1", time.Now(), time.Now())
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`)).
			WithArgs(1, 1).
			WillReturnRows(rows)
		// userRepository := repository.NewUserRepository(gdb)
		user, err := userRepository.GetByID(1)
		assert.NoError(s.T(), err)
		assert.NotNil(s.T(), user)
		assert.Equal(s.T(), "John Doe", user.Name)
		assert.NoError(s.T(), mock.ExpectationsWereMet())
	})

	s.Run("Failure", func() {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`)).
			WithArgs(999, 1).
			WillReturnError(fmt.Errorf("some error"))

		user, err := userRepository.GetByID(999)
		assert.Error(s.T(), err)
		assert.Nil(s.T(), user)
		assert.NoError(s.T(), mock.ExpectationsWereMet())
	})
}

type AnyTime struct{}

func (a AnyTime) Match(v interface{}) bool {
	switch value := v.(type) {
	case time.Time: // Trường hợp truyền trực tiếp
		return true
	case driver.NamedValue: // Trường hợp bọc trong NamedValue
		_, ok := value.Value.(time.Time)
		return ok
	default:
		return false
	}
}

func (s *testUserRepositorySuite) TestCreate() {
	// Khởi tạo mock database
	db, gdb, mock, err := NewConnManager()
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type testCaseIn struct {
		mocks        func()
		newUserInput domain.SysUser
	}

	type testCaseOut struct {
		err error
		id  int
	}

	cases := []struct {
		name     string
		in       testCaseIn
		expected testCaseOut
	}{
		{
			name: "create new user success",
			in: testCaseIn{
				newUserInput: domain.SysUser{
					Name:     "t1",
					Email:    "t1@gmail.com",
					Phone:    "1",
					Password: "1",
				},
				mocks: func() {
					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(
						`INSERT INTO "users" ("name","email","phone","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
						WithArgs(
							"t1", "t1@gmail.com", "1", "1",
							sqlmock.AnyArg(), sqlmock.AnyArg(),
						).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
					mock.ExpectCommit()
				},
			},
			expected: testCaseOut{
				err: nil,
				id:  1,
			},
		},
		{
			name: "create new user failed",
			in: testCaseIn{
				newUserInput: domain.SysUser{
					Name:     "t1",
					Email:    "t1@gmail.com",
					Phone:    "1",
					Password: "1",
				},
				mocks: func() {
					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(
						`INSERT INTO "users" ("name","email","phone","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
						WithArgs(
							"t1", "t1@gmail.com", "1", "1",
							sqlmock.AnyArg(), sqlmock.AnyArg(),
						).WillReturnError(fmt.Errorf("some error"))

					mock.ExpectRollback()
				},
			},
			expected: testCaseOut{
				err: fmt.Errorf("some error"),
				id:  0,
			},
		},
	}

	for _, c := range cases {
		s.T().Run(c.name, func(t *testing.T) {
			c.in.mocks()
			userRepository := repository.NewUserRepository(gdb)

			ret, err := userRepository.Create(&c.in.newUserInput)

			assert.Equal(t, c.expected.id, ret)
			assert.Equal(t, c.expected.err, err)
		})
	}
}

// // Tạo instance của repository
// userRepository := repository.NewUserRepository(gdb)

// s.Run("Success", func() {
// 	// Expect bắt đầu transaction
// 	mock.ExpectBegin()

// 	// Expect câu lệnh INSERT với RETURNING "id"
// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO "users" ("name","email","phone","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
// 		WithArgs(
// 			"t1", "t1@example.com", "1", "1",
// 			AnyTime{}, AnyTime{}, // Matcher AnyTime cho thời gian
// 		).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // Giả lập trả về ID = 1

// 	mock.ExpectCommit()

// 	// Tạo dữ liệu test
// 	now := time.Now().Truncate(time.Millisecond)
// 	user := domain.User{
// 		Name:      "t1",
// 		Email:     "t1@example.com",
// 		Phone:     "1",
// 		Password:  "1",
// 		CreatedAt: now,
// 		UpdatedAt: now,
// 	}

// 	// Gọi hàm Create từ repository
// 	idReturn, err := userRepository.Create(&user)

// 	// So sánh kết quả mong đợi
// 	assert.NoError(s.T(), err)
// 	assert.Equal(s.T(), 1, idReturn)

// 	// Kiểm tra các expectation đã được thực thi
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		s.T().Errorf("unfulfilled expectations: %s", err)
// 	}
// })
