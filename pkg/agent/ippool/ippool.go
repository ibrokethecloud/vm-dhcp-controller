package ippool

import (
	"fmt"

	"github.com/sirupsen/logrus"

	networkv1 "github.com/harvester/vm-dhcp-controller/pkg/apis/network.harvesterhci.io/v1alpha1"
	"github.com/harvester/vm-dhcp-controller/pkg/util"
)

func (c *Controller) Update(ipPool *networkv1.IPPool) error {
	if ipPool.Status.IPv4 == nil {
		return fmt.Errorf("ippool status has no records")
	}
	allocated := ipPool.Status.IPv4.Allocated
	filterExcluded(allocated)
	return c.updatePoolCacheAndLeaseStore(allocated, ipPool.Spec.IPv4Config)
}

func (c *Controller) updatePoolCacheAndLeaseStore(latest map[string]string, ipv4Config networkv1.IPv4Config) error {
	for ip, mac := range c.poolCache {
		if newMAC, exists := latest[ip]; exists {
			if mac != newMAC {
				logrus.Infof("set %s with new value %s", ip, newMAC)
				// TODO: update lease
				c.poolCache[ip] = newMAC
			}
		} else {
			logrus.Infof("remove %s", ip)
			if err := c.dhcpAllocator.DeleteLease(c.poolCache[ip]); err != nil {
				return err
			}
			delete(c.poolCache, ip)
		}
	}

	for newIP, newMAC := range latest {
		if _, exists := c.poolCache[newIP]; !exists {
			logrus.Infof("add %s with value %s", newIP, newMAC)
			if err := c.dhcpAllocator.AddLease(
				newMAC,
				ipv4Config.ServerIP,
				newIP,
				ipv4Config.CIDR,
				ipv4Config.Router,
				ipv4Config.DNS,
				ipv4Config.DomainName,
				ipv4Config.DomainSearch,
				ipv4Config.NTP,
				ipv4Config.LeaseTime,
			); err != nil {
				return err
			}
			c.poolCache[newIP] = newMAC
		}
	}

	return nil
}

func filterExcluded(allocated map[string]string) {
	for ip, mac := range allocated {
		if mac == util.ExcludedMark {
			delete(allocated, ip)
		}
	}
}
