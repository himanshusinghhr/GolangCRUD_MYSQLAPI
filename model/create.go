package model

//CreateTodoo to do
func CreateTodoo(Name string, Task string) error {
	// // sqlQ := "insert into todo values (" + "himanshu2," + "task2);"
	insertIQ, err := con.Query("insert into todo values(?,?)", Name, Task)
	// insertIQ, err := con.Query("insert into todo values('Himanshu2','test2')")

	defer insertIQ.Close()
	if err != nil {
		return err
	}
	return nil
}
