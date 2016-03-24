package cluster

// MarkOffsetPerPartition marks the offset for a specific partition, for the given topic.
func (c *Consumer) MarkOffsetPerPartition(topic string, part int32, offset int64) {
	c.subs.Fetch(topic, part).MarkOffset(offset, "")
}
