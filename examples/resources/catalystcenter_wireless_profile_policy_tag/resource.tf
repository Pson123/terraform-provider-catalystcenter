resource "catalystcenter_wireless_profile_policy_tag" "example" {
  wireless_profile_id = "12345678-1234-1234-1234-123456789012"
  policy_tag_name     = "PolicyTag1"
  site_ids            = ["12345678-1234-1234-1234-123456789000"]
}
