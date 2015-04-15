var hosts struct {
	sync.Mutex
	byName map[string][]string
	byAddr map[string][]string
	expire time.Time
	path   string
}

func lookupStaticHost(host string) []string {
	hosts.Lock()
	defer hosts.Unlock()
	readHosts()
	if len(hosts.byName) != 0 {
		if ips, ok := hosts.byName[host]; ok {
			return ips
		}
	}
	return nil
}
