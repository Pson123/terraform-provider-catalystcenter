resource "catalystcenter_power_profile" "example" {
  profile_name = "Low_Power_Profile"
  description  = "Reduce power consumption on APs"
  rules = [
    {
      interface_type  = ""
      interface_id    = ""
      parameter_type  = ""
      parameter_value = ""
    }
  ]
}
