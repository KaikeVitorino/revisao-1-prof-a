package q2

import (
	"github.com/revisao-1/utils"
	"testing"
)

func TestAverageLettersPerWord(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    float64
		wantErr bool
	}{
		{
			name:    "returns error if text is empty",
			text:    "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "returns error if text is only spaces",
			text:    " ",
			want:    0,
			wantErr: true,
		},
		{
			name:    "returns error if text is only numbers",
			text:    "123",
			want:    0,
			wantErr: true,
		},
		{
			name:    "returns average letters per word",
			text:    "Lorem ipsum 123 dolor sit amet.456",
			want:    4.4,
			wantErr: false,
		},
		{
			name:    "returns average letters per word",
			text:    "O rato roeu a roupa do rei de Roma",
			want:    2.88,
			wantErr: false,
		},
		{
			name:    "returns average letters per word",
			text:    "Lorem ipsum dolor sit amet.",
			want:    4.4,
			wantErr: false,
		},
		{
			name:    "returns average letters per word",
			text:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			want:    5.34,
			wantErr: false,
		},
		{
			name:    "returns average letters per word",
			text:    "Lorem ipsum dolor sit amet,   consectetur adipiscing elit, sed do    eiusmod tempor    incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur    . Excepteur     sint occaecat cupidatat non proident, sunt in culpa qui    officia deserunt mollit anim    id est laborum.",
			want:    5.34,
			wantErr: false,
		},
		{
			name:    "returns average letters per word",
			text:    " aa dss ew adsd ef,poa sdmkfsmd wqoka sdamowjr grkjmsiddfmqwe f sjkdnfosifgbn rgdjb sdkfjgvnsdoifvmsderftkl ertnirms dffgjnsdo ",
			want:    7.33,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AverageLettersPerWord(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("AverageLettersPerWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !utils.AssertFloatWithPrecision(tt.want, got, 1e-2) {
				t.Errorf("AverageLettersPerWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
