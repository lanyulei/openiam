package clouds

/*
  @Author : lanyulei
  @Desc :
*/

type HandleType string

const (
	DescribeInstances HandleType = "DescribeInstances"
	DescribeRegions   HandleType = "DescribeRegions"
)

type CloudResourceType string

const (
	CloudResourceHost CloudResourceType = "Host"
)
