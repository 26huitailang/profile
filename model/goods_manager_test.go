package model

import (
	"os"
	"profile/database"
	"testing"
)

const TestDBName = "test.db"

func TestGoodsManger_GetAllGoods(t *testing.T) {

	itemPhone := Goods{Name: "phone", Description: "desc1", Price: 9, Category: uint(ElectronicEquipment)}
	itemTV := Goods{Name: "tv", Description: "desc1", Price: 19, Category: uint(HouseholdAppliances)}

	tests := []struct {
		name string
		want []Goods
	}{
		{"one record", []Goods{
			itemPhone,
		}},
		{"two record", []Goods{
			itemPhone,
			itemTV,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer removeTestDB(t, TestDBName)
			db, closeDB := database.NewDB(TestDBName)
			defer closeDB()
			db.AutoMigrate(&Goods{}, &GoodsImage{})
			m := &GoodsManger{
				db: db,
			}
			for _, item := range tt.want {
				m.InsertOneGoods(&item)
			}
			if got := m.GetAllGoods(); len(got) != len(tt.want) {
				t.Errorf("GoodsManger.GetAllGoods() len %v, want len %v", len(got), len(tt.want))
			}
		})
	}
}

func removeTestDB(t *testing.T, filename string) {
	err := os.Remove(filename)
	if err != nil {
		t.Fatal("[Teardown] db remove failed!")
	}

}
