package v1

// todo: unittest for Goods API

//func TestViewHandler_FindGoods(t *testing.T) {
//	type fields struct {
//		store Store
//	}
//	type args struct {
//		c echo.Context
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := &ViewHandler{
//				store: tt.fields.store,
//			}
//			if err := h.FindGoods(tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("ViewHandler.FindGoods() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestViewHandler_CreateGoods(t *testing.T) {
//	type fields struct {
//		store Store
//	}
//	type args struct {
//		c echo.Context
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := &ViewHandler{
//				store: tt.fields.store,
//			}
//			if err := h.CreateGoods(tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("ViewHandler.CreateGoods() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestViewHandler_EditGoods(t *testing.T) {
//	type fields struct {
//		store Store
//	}
//	type args struct {
//		c echo.Context
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			h := &ViewHandler{
//				store: tt.fields.store,
//			}
//			if err := h.EditGoods(tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("ViewHandler.EditGoods() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//type StubGoodsManager struct{
//	Goods []model.Goods
//}
//
//func (s *StubGoodsManager) InsertOneGoods(item *model.Goods) (*model.Goods, error) {
//	s.Goods = append(s.Goods, *item)
//	return item, nil
//}
//
//func (s *StubGoodsManager) GetAllGoods() []model.Goods {
//	return s.Goods
//}
//
//func (s *StubGoodsManager) UpdateOneGoods(item *model.Goods) (*model.Goods, error) {
//	v := reflect.ValueOf(*item)
//	count := v.NumField()
//	for i := 0; i < count; i++ {
//		f := v.Field(i)
//		switch f.FieldByName() {
//		case reflect.String:
//			fmt.Println(f.String())
//		case reflect.Int:
//			fmt.Println(f.Int())
//		}
//	}
//}
