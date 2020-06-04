package mysql

import (
	"fxtservice/pkg/config"
	"github.com/kevin-zx/sqlbase"
	"reflect"
	"testing"
)

func Test_fxtStorage_GetUser(t *testing.T) {

	type fields struct {
		s *sqlbase.Storage
	}
	s, err := sqlbase.NewStorage(config.XwDB["ENGINE"].(string), config.XwDB["PORT"].(int), config.XwDB["NAME"].(string), config.XwDB["USER"].(string), config.XwDB["HOST"].(string), config.XwDB["PASSWORD"].(string))
	if err != nil {
		panic(err)
	}
	defer s.Close()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "译达",
			fields: fields{
				s: s,
			},
			args: args{
				name: "译达",
			},
			want:    "合肥译达翻译服务有限公司",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fxtStorage{
				s: tt.fields.s,
			}
			got, err := f.GetUserByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err!=nil {
				return
			}
			if !reflect.DeepEqual(got.UserName, tt.want) {
				t.Errorf("GetUserByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fxtStorage_GetGuanwangKeywordsAndRecentRank(t *testing.T) {
	type fields struct {
		s *sqlbase.Storage
	}
	s, err := sqlbase.NewStorage(config.XwDB["ENGINE"].(string), config.XwDB["PORT"].(int), config.XwDB["NAME"].(string), config.XwDB["USER"].(string), config.XwDB["HOST"].(string), config.XwDB["PASSWORD"].(string))
	if err != nil {
		panic(err)
	}
	defer s.Close()
	type args struct {
		siteID     uint
		recentDays int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "1",
			fields:  fields{s: s},
			args:    args{
				siteID:     2327,
				recentDays: 7,
			},
			want:    8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fxtStorage{
				s: tt.fields.s,
			}
			got, err := f.GetGuanwangKeywordsAndRecentRank(tt.args.siteID, tt.args.recentDays)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGuanwangKeywordsAndRecentRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GetGuanwangKeywordsAndRecentRank() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_fxtStorage_GetGuanbaoProjectBySiteURLAndClientName(t *testing.T) {
	type fields struct {
		s *sqlbase.Storage
	}
	s, err := sqlbase.NewStorage(config.XwDB["ENGINE"].(string), config.XwDB["PORT"].(int), config.XwDB["NAME"].(string), config.XwDB["USER"].(string), config.XwDB["HOST"].(string), config.XwDB["PASSWORD"].(string))
	if err != nil {
		panic(err)
	}
	defer s.Close()
	type args struct {
		siteURL    string
		clientName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "1",
			fields:  fields{
				s: s,
			},
			args:    args{
				siteURL:    "www.ywjzryp.com",
				clientName: "义乌市聚壮日用品有限公司",
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fxtStorage{
				s: tt.fields.s,
			}
			got, err := f.GetGuanbaoProjectBySiteURLAndClientName(tt.args.siteURL, tt.args.clientName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGuanbaoProjectBySiteURLAndClientName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GetGuanbaoProjectBySiteURLAndClientName() got = %v, want %v", got, tt.want)
			}
		})
	}
}