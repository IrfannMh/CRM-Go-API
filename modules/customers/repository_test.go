package customers

import (
	"gorm.io/gorm"
	"testing"
)

func Test_repository_Save(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		customer *Customer
	}
	type testCase struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	var tests []testCase

	mockQuery, mockDb := test.NewMockQueryDB(t)

	f := fields{db: mockDb}
	tests = append(tests, testCase{
		name:    "success create customer",
		fields:  f,
		args:    args{},
		wantErr: false,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				db: tt.fields.db,
			}
			if err := r.Save(tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
