it := db.NewIterator(ro)
defer it.Close()
it.Seek(mykey)
for it = it; it.Valid(); it.Next() {
	doSomethingAwesome(it.Key(), it.Value())
}
