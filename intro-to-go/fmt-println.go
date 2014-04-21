func Println(a ...interface{}) (n int, err error) {
        return Fprintln(os.Stdout, a...)
}
