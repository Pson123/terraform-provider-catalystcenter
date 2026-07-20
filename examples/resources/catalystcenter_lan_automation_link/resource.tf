resource "catalystcenter_lan_automation_link" "example" {
  action                               = "ADD_LINK"
  primary_device_management_ip_address = "1.2.3.4"
  primary_device_interface_name        = "FortyGigabitEthernet1/0/1"
  peer_device_management_ip_address    = "5.6.7.8"
  peer_device_interface_name           = "TenGigabitEthernet1/0/5"
  ip_pool_name                         = "SDA_Overlapping_LAN_Auto"
}
