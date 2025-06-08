package memberships

import (
	"testing"

	"github.com/azybk/music-catalog/internal/configs"
	"github.com/azybk/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		req memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				req: memberships.LoginRequest{
					Email:    "",
					Password: "",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "",
					Username: "",
					Password: "",
				}, nil)
			},
		},
		{
			name: "failed when get user",
			args: args{
				req: memberships.LoginRequest{
					Email:    "",
					Password: "",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				req: memberships.LoginRequest{
					Email:    "",
					Password: "",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.req.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "",
					Username: "",
					Password: "",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "",
					},
				},
				repository: mockRepo,
			}

			got, err := s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
