package model

func Delete(name string) error {
	_, err := con.Query("delete from todo where name=?", name)
	if err != nil {
		return err
	}
	return nil
}
