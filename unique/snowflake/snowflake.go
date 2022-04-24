package snowflake

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func New(nodeId int64) (err error) {
	node, err = snowflake.NewNode(nodeId)
	return err
}

func Gen() snowflake.ID {
	if node == nil {
		return 0
	}
	return node.Generate()
}
