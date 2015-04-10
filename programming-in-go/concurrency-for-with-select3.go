func (s *Server) Serve() {
	for {
		select {
		// Stop and close the connection
		case <-s.StopCh:
			s.conn.Close()
			return
		default:
			buf := make([]byte, defs.DatagramSize)
			t := time.Now()
			s.conn.SetReadDeadline(t.Add(readTimeout))
			n, addr, err := s.conn.ReadFromUDP(buf)
			if err == nil {
				go s.route(buf[:n], addr)
			} else {
				s.handleErr(err, addr)
			}
		}
	}
}
