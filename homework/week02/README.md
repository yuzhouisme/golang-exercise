第二周作业。



##### 问题：

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


##### 思考：

1. 个人认为应该Wrap，因为需要把错误栈上抛，但若存在敏感信息应考虑屏蔽。
2. 然后个人觉得在上层应该降级处理吧，转成业务自己的ErrorCode。



```go
...

func getClientById(id string) string {
	// mock: call DAO func
	user, err := getOne(id)
	if err != nil {
		// write into log
		fmt.Printf("%+v\n", err)
		//
		if errors.Is(err, sql.ErrNoRows) {
			return "No User"
		} else {
			return err.Error();
		}
	}

	return fmt.Sprintf("SUCCEED!!! User = %v \n", user)
}

func getOne(id string) (*User, error) {
	user := User{ID: id}
	row := db.QueryRow("SELECT * FROM clients WHERE id=?", id)
	if err := row.Scan(&user.ID, &user.BusinessName, &user.Address, &user.Phone, &user.RoleType, &user.Status); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("UserService: getOne id = %v", id))
	}
	return &user, nil
}

```

