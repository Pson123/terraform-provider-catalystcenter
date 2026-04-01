resource "catalystcenter_power_profile" "example" {
  profile_name = "Low_Power_Profile"
  description  = "Reduce power consumption on APs"
  rules = [
    {
      interface_type  = "RADIO"
      interface_id    = "6GHZ"
      parameter_type  = "STATE"
      parameter_value = "DISABLE"
    }
  ]
}
