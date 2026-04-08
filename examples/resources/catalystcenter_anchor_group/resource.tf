resource "catalystcenter_anchor_group" "example" {
  anchor_group_name = "AG1"
  mobility_anchors = [
    {
      device_name         = "WLC-ANCHOR"
      ip_address          = "10.0.0.1"
      anchor_priority     = "PRIMARY"
      managed_anchor_wlc  = false
      peer_device_type    = "IOS-XE"
      mac_address         = "aa:bb:cc:00:00:01"
      mobility_group_name = "default"
    }
  ]
}
