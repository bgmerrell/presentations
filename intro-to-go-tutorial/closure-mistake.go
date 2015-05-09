for val := range values {
    go func() {
        fmt.Println(val)
    }()
}
