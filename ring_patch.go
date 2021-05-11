package redis

// 这个文件是对Ring的补丁文件

func (c *Ring) EachShard(f func (*Client)) {
	clients := []*Client{}
	c.mu.Lock()
	for _, shard := range c.shards {
		if shard == nil || shard.Client == nil {
			continue
		}
		clients = append(clients, shard.Client)
	}
	c.mu.Unlock()
	for _, client := range clients {
		f(client)
	}
}
