package project_test

import (
	"errors"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "example-db/project"
	"example-db/helpers"
)

var _ = Describe("ProjectRepository", func() {
	Context("#GetAllProjectsForAccount", func() {
		It("should return a list of projects for given accountID on success", func() {
			db, mockDB, _ := sqlmock.New()
			defer db.Close()

			testAccountID := "accountid"
			expectedProject := Project{
				ProjectID:  1,
				AccountID:  testAccountID,
				Name:       "Golden Gods, Inc.",
				Type:       "Parsing",
				CreatedBy:  "Dennis Reynolds",
				CreatedOn:  "October 12, 2019",
				ModifiedBy: "Dennis Reynolds",
				ModifiedOn: "October 13, 2019",
			}

			columns := []string{"ProjectID", "AccountID", "Name",
				"Type", "CreatedBy", "CreatedDate",
				"ModifiedBy", "ModifiedDate"}

			mockDB.ExpectQuery("SELECT (.+) FROM DataModeler.Project").
				WillReturnRows(sqlmock.NewRows(columns).AddRow(expectedProject.ProjectID, testAccountID, expectedProject.Name, expectedProject.Type, expectedProject.CreatedBy, expectedProject.CreatedOn, expectedProject.ModifiedBy, expectedProject.ModifiedOn))

			projectRepository := ProjectRepository{DB: db}
			projects, err := projectRepository.GetAllProjectsForAccount(testAccountID)
			Expect(err).ToNot(HaveOccurred())

			if err := mockDB.ExpectationsWereMet(); err != nil {
				Fail(err.Error())
			}

			Expect(len(projects)).To(Equal(1))

			project := projects[0]
			Expect(project.ProjectID).To(Equal(expectedProject.ProjectID))
			Expect(project.AccountID).To(Equal(expectedProject.AccountID))
			Expect(project.Name).To(Equal(expectedProject.Name))
			Expect(project.Type).To(Equal(expectedProject.Type))
			Expect(project.CreatedBy).To(Equal(expectedProject.CreatedBy))
			Expect(project.CreatedOn).To(Equal(expectedProject.CreatedOn))
			Expect(project.ModifiedBy).To(Equal(expectedProject.ModifiedBy))
			Expect(project.ModifiedOn).To(Equal(expectedProject.ModifiedOn))
		})

		It("should return an empty list of projects when no projects are found for the given accountID", func() {
			db, mockDB, _ := sqlmock.New()
			defer db.Close()

			testAccountID := "accountid"
			columns := []string{"ProjectID", "AccountID", "Name",
				"Type", "CreatedBy", "CreatedDate",
				"ModifiedBy", "ModifiedDate"}

			mockDB.ExpectQuery("SELECT (.+) FROM DataModeler.Project").
				WillReturnRows(sqlmock.NewRows(columns))

			projectRepository := ProjectRepository{DB: db}
			projects, err := projectRepository.GetAllProjectsForAccount(testAccountID)
			Expect(err).ToNot(HaveOccurred())

			if err := mockDB.ExpectationsWereMet(); err != nil {
				Fail(err.Error())
			}

			Expect(len(projects)).To(Equal(0))
		})

		It("should return a DBError with a 500 code on failure", func() {
			db, mockDB, _ := sqlmock.New()
			defer db.Close()

			testAccountID := "accountid"
			mockDB.ExpectQuery("SELECT (.+) FROM DataModeler.Project").WillReturnError(errors.New("BLAMO"))
			
			projectRepository := ProjectRepository{DB: db}
			_, err := projectRepository.GetAllProjectsForAccount(testAccountID)
			Expect(err).To(HaveOccurred())

			if dbErr, ok := err.(helpers.DBError); ok {
				Expect(dbErr.Code).To(Equal(500))
			} else {
				Fail("wrong type of error returned")
			}

			if err := mockDB.ExpectationsWereMet(); err != nil {
				Fail(err.Error())
			}
		})
	})
})
