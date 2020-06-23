---
layout: "cohesity"
page_title: "Cohesity: cohesity_restore_vmware"
description: |-
  Restore a VMware VM 
---

# cohesity\_restore\_vmware

Restore a VMware VM 

## Example Usage
```
provider "cohesity" {
        cluster_vip = "10.2.45.143"
        cluster_username = "abcd"
        cluster_domain = "LOCAL"
}


resource "cohesity_restore_vmware" "protect_vcenter" {
	name = "Restore_Job1"
  job_name = "Job1"
	state = "start"
  vm_name = ["vm1", "vm2"]
}
```

### Argument Reference
The following arguments are supported:

- name - (Required, string) Specifies the name of the restore task
- job_name - (Required, string) Specifies the name of the protection job that backed up the objects to be restored
- backup_timestamp - (Optional, string) Specifies the time of the protection job run. Should be in the format YYYY-MM-DD HH:MM Area/Location
- vm_names - (Optional, set of strings) Specifies the names of the virtual machines to restore
- operation_timeout - (Optional, int) Specifies the time to wait in minutes for the protection job run to complete the run or stop the run. The default value is **120**
- state - (Optional, string) Specifies whether to start or stop a protection job run. The default value is **start**. The supported values are **start** and **stop**
- vmware_parameters - (Optional, Map) Specifies vmware parameters for the restore task


#### Attributes Reference
The following attributes are exported:

- id - ID of the VMware protection job