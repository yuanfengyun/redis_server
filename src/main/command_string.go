package main

func cmd_get(db *DB, args []string) string {
	if len(args) != 1 {
		return get_error_string("wrong args num!")
	}
	key := args[0]
	o := db.dict.Get(key)
	if o == nil {
		return get_empty_bulk_string()
	}
	if o.t != 1 {
		return get_error_string("not a string")
	}
	return get_bulk_string(o.p.(string))
}

func cmd_set(db *DB, args []string) string {
	if len(args) != 2 {
		return get_error_string("wrong args num!")
	}
	key := args[0]
	value := args[1]
	db.dict.Set(key, &Object{
		t: 1,
		p: value,
	})
	return get_status_string("OK")
}

func register_string_cmds() {
	cmds["get"] = cmd_get
	cmds["set"] = cmd_set
}