package blb

import (
	"fmt"
	"testing"

	"github.com/drinktee/bce-sdk-go/util"
)

var expectCreateBLB = &CreateLoadBalancerArgs{
	Name: "blb-for-test",
}

func TestCreateLoadBalance(t *testing.T) {
	blb, err := blbClient.CreateLoadBalancer(expectCreateBLB)
	if err != nil {
		t.Error(util.FormatTest("TestCreateLoadBalance", err.Error(), "nil"))
	} else {
		if blb.Name != expectCreateBLB.Name {
			t.Error("blb name error")
		}
	}
}

func TestDescribeLoadBalancers(t *testing.T) {
	args := &DescribeLoadBalancersArgs{
		LoadBalancerName: "test",
	}
	list, err := blbClient.DescribeLoadBalancers(args)

	if err != nil {
		fmt.Println(err)
		t.Error(util.FormatTest("TestDescribeLoadBalancers", err.Error(), "nil"))
	}
	if len(list) != 2 {
		t.Error("blb length error")
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	args := &UpdateLoadBalancerArgs{
		LoadBalancerId: "lb-e5b33752",
		Name:           "golang-123",
	}
	err := blbClient.UpdateLoadBalancer(args)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteLoadBalancer(t *testing.T) {
	args := &DeleteLoadBalancerArgs{
		LoadBalancerId: "lb-426fad2b",
	}
	err := blbClient.DeleteLoadBalancer(args)
	if err != nil {
		t.Error(err)
	}
}
