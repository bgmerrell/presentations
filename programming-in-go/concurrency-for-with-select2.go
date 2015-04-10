func Run() (stdout string, stderr string, err error) {
	[...]
	c := make(chan error)
	go func() {
		c <- session.Run(command)
	}()
	select {
	case err = <-c:
		if err == nil {
			return "", "", err
		}
	case <-time.After(time.Duration(timeout) * time.Second):
		err = errors.New("timeout")
		defer handleTimeout()
	}
	return stdout_buf.String(), stderr_buf.String(), err
}
