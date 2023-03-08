package main

type Transaction struct {
	ID   uint
	Name string
}

// TestTransaction 事务
func TestTransaction() {
	//CreateTableByModel(&Transaction{})
	//DB.Transaction(func(tx *gorm.DB) error {
	//	DB.Create(&Transaction{
	//		Name: "Test1",
	//	})
	//	DB.Create(&Transaction{
	//		Name: "Test2",
	//	})
	//	tx.Transaction(func(tx *gorm.DB) error {
	//		tx.Create(&Transaction{
	//			Name: "Test3",
	//		})
	//		return errors.New("嵌套事务发生错误")
	//	})
	//	return nil
	//})
	begin := DB.Begin()
	begin.Create(&Transaction{Name: "事务开始3"})
	begin.Create(&Transaction{Name: "事务开始4"})
	begin.SavePoint("p1")
	begin.Create(&Transaction{Name: "事务开始3"})
	//////if true {
	//////	begin.Rollback()
	//////}
	begin.RollbackTo("p1")
	begin.Commit()
}
