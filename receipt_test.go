package addons_receipt

import (
	"github.com/emdeweier/addons-receipt/assets/utils"
	"github.com/jung-kurt/gofpdf"
	"reflect"
	"testing"
)

func TestReceipt(t *testing.T) {
	type args struct {
		transactionId        string
		transactionName      string
		isFirstHeaderVisible bool
		listModel            []utils.ListModelData
		paperSize            string
	}
	tests := []struct {
		name    string
		args    args
		want    gofpdf.Pdf
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Receipt(tt.args.transactionId, tt.args.transactionName, tt.args.isFirstHeaderVisible, tt.args.listModel, tt.args.paperSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("Receipt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Receipt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
